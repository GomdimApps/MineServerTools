package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/GomdimApps/MineServerTools/tools/api/utils"

	"github.com/gin-gonic/gin"
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

func main() {
	r := gin.Default()
	port, err := utils.GetPortApi()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	authorized := r.Group("/", authMiddleware())
	{
		authorized.GET("/api/server/backup-bedrock/view/", func(c *gin.Context) {
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
		})

		authorized.GET("/api/server/status/view/", func(c *gin.Context) {
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
		})
	}

	r.Run(fmt.Sprintf(":%d", port))
}
