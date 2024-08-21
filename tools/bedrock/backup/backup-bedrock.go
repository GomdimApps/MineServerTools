package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GomdimApps/MineServerTools/tools/bedrock/backup/scheduler"
	"github.com/GomdimApps/MineServerTools/tools/bedrock/backup/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go [--backup|--schedule|--view]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--backup":
		utils.Backup()
	case "--schedule":
		scheduler.ScheduleTasks()
	case "--view":
		utils.ViewBackup()
	case "--view-json":
		jsonOutput, err := utils.ViewBackupJson()
		if err != nil {
			log.Fatalf("Erro: %v", err)
		}
		fmt.Println(jsonOutput)
	default:
		fmt.Println("Uso: go run main.go [--backup|--schedule|--view]")
		os.Exit(1)
	}
}
