package http

import (
	"fmt"
)

func RunDownloadProcess() error {
	url, err := GetLatestServerVersion()
	if err != nil {
		return fmt.Errorf("erro ao buscar a ultima versão: %v", err)
	}

	fmt.Printf("Baixando URL: %s\n", url)

	file, err := DownloadServer(url)
	if err != nil {
		return fmt.Errorf("erro ao fazer download do server: %v", err)
	}

	filename := "bedrock_server.zip"
	destDir := "/tmp" // Define o diretório de destino aqui

	err = SaveFile(filename, file, destDir)
	if err != nil {
		return fmt.Errorf("erro ao salvar arquivo: %v", err)
	}

	fmt.Println("Download completado com sucesso!")
	return nil
}
