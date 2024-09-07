package main

import (
	"fmt"

	"github.com/GomdimApps/MineServerTools/App/bin/api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	port, err := utils.GetPortApi()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	SetupRoutes(r)

	r.Run(fmt.Sprintf(":%d", port))
}
