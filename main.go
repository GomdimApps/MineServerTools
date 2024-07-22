package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	file, err := DownloadServer("https://minecraft.azureedge.net/bin-linux/bedrock-server-1.21.3.01.zip")
	if err != nil {
		fmt.Printf("Erro ao fazer download do server: %v\n", err)
    return
	}

	SaveMineServerZip("downloads/bedrock_server-1.21.3.01.zip", file)
}

func DownloadServer(url string) (string, error) {
	fmt.Println("Inciando download do server...")
	resp, err := http.Get(url)

	fmt.Println("Fazendo download do server, por favor aguarde")

	if err != nil {
		return "", err
	}
	
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	fmt.Println("Download finalizado com sucesso!")

	return string(body), nil
}

func SaveMineServerZip(filename string, archive string) {
	fmt.Printf("Salvando o arquivo .zip do server em %s\n\n", filename)

	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		fmt.Printf("Erro ao salvar o arquivo: %v\n", err)
    return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Erro ao salvar o arquivo: %v\n", err)
    return
	}

	defer file.Close()

	_, err = file.WriteString(archive)
	if err != nil {
    fmt.Printf("Erro ao escrever no arquivo: %v\n", err)
    return
  }

	fmt.Println("Arquivo salvo com sucesso!")
}