package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/GomdimApps/MineServerTools/tools/api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	rateLimiters  = make(map[string]*rate.Limiter)
	ipBlacklist   = make(map[string]time.Time)
	ipLimit       = 25
	blockDuration = 40 * time.Second
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token_custom, err := utils.GetTokenApi()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		token := c.GetHeader("Authorization")
		expectedToken := fmt.Sprintf("Bearer %s", token_custom)
		if token != expectedToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if _, exists := rateLimiters[ip]; !exists {
			rateLimiters[ip] = rate.NewLimiter(3, 1)
		}

		if !rateLimiters[ip].Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ipBlockMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if blockUntil, blocked := ipBlacklist[ip]; blocked {
			if time.Now().Before(blockUntil) {
				c.JSON(http.StatusForbidden, gin.H{"error": "IP blocked due to excessive requests"})
				c.Abort()
				return
			}
			delete(ipBlacklist, ip)
		}

		if _, exists := rateLimiters[ip]; exists {
			if rateLimiters[ip].Allow() {
				countRequests(ip)
			}
		}

		c.Next()
	}
}

func countRequests(ip string) {
	if _, exists := ipBlacklist[ip]; !exists {
		ipBlacklist[ip] = time.Now().Add(blockDuration)
	}

	if len(ipBlacklist) > ipLimit {
		ipBlacklist[ip] = time.Now().Add(blockDuration)
	}
}

func SetupRoutes(r *gin.Engine) {
	authorized := r.Group("/", authMiddleware(), rateLimitMiddleware(), ipBlockMiddleware())
	{
		authorized.GET("/api/server/backup-bedrock/view/", viewBackup)
		authorized.GET("/api/server/status/view/", viewServerStatus)
	}
}

func viewBackup(c *gin.Context) {
	cmd := exec.Command("backup-bedrock", "--view-json")
	stdout, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := strings.TrimSpace(string(stdout))

	var backups []utils.Backup
	err = json.Unmarshal([]byte(output), &backups)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao parsear o JSON"})
		return
	}

	c.JSON(http.StatusOK, backups)
}

func viewServerStatus(c *gin.Context) {
	cmd := exec.Command("info-bedrock", "--json")
	stdout, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := strings.TrimSpace(string(stdout))

	var serverInfo utils.ServerInfo
	err = json.Unmarshal([]byte(output), &serverInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao parsear o JSON"})
		return
	}

	c.JSON(http.StatusOK, serverInfo)
}
