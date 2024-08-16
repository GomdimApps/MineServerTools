package main

import (
	"fmt"
	"os"

	"github.com/GomdimApps/MineServerTools/tools/bedrock/update/http"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--new-server":
		if len(os.Args) < 3 {
			fmt.Println("Erro: O diretório para o novo servidor deve ser especificado.")
			printUsage()
			os.Exit(1)
		}
		destDir := os.Args[2]
		err := http.SetupNewServer(destDir)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Novo servidor instalado")

	case "--update":
		err := http.RunDownloadProcess()
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Atualização do servidor completada")

	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Uso: go run main.go [--new-server <diretório>|--update]")
	fmt.Println("  --new-server <diretório>")
	fmt.Println("        Diretório para instalar o novo servidor")
	fmt.Println("  --update")
	fmt.Println("        Atualizar o servidor existente")
}
