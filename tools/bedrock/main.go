package main

import (
	"fmt"
	"os"

	"github.com/GomdimApps/MineServerTools/tools/bedrock/http"
	"github.com/GomdimApps/MineServerTools/tools/bedrock/system"
)

func main() {
	fmt.Printf("--------Welcome to MineServerTools!--------\n\n")
	fmt.Printf("What do you want to do?\n")
	fmt.Printf("1. Download and save the latest server version\n2. See your current system info\n\n: ")
	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		http.RunDownloadProcess()
	case "2":
		system.GetSystemInfo()
	default:
		fmt.Println("Invalid option.")
		os.Exit(1)
	}
}
