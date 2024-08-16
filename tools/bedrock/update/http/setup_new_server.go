package http

import (
	"fmt"
	"os"
	"path/filepath"
)

func SetupNewServer(destDir string) error {

	if destDir == "" {
		return fmt.Errorf("o diretório para instalação do novo servidor não foi fornecido")
	}

	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		return fmt.Errorf("o diretório especificado não existe: %s", destDir)
	}

	url, err := GetLatestServerVersion()
	if err != nil {
		return fmt.Errorf("erro ao buscar a última versão: %v", err)
	}

	file, err := DownloadServer(url)
	if err != nil {
		return fmt.Errorf("erro ao fazer download do server: %v", err)
	}

	filename := "bedrock_server.zip"
	filePath := filepath.Join(destDir, filename)
	err = SaveNewServerFile(filename, file, destDir)
	if err != nil {
		return fmt.Errorf("erro ao salvar arquivo: %v", err)
	}

	err = UnzipFile(filePath, destDir)
	if err != nil {
		return fmt.Errorf("erro ao descompactar o arquivo: %v", err)
	}

	return nil
}

func SaveNewServerFile(filename string, content []byte, destDir string) error {
	filePath := filepath.Join(destDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(content); err != nil {
		return err
	}

	fmt.Printf("Finalizando instalação, tamanho: %d bytes\n", len(content))
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Para iniciar user o comando 'console-bedrock --start -d %s'\n\n", destDir)

	return nil
}
