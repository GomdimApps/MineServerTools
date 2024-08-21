package network

import (
	"fmt"
	"io"
	"net/http"
)

func DownloadServer(url string) ([]byte, error) {
	fmt.Printf("Iniciando download do URL: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro, statusCode: %d", resp.StatusCode)
	}

	archive, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Download concluído, tamanho do arquivo: %d bytes\n", len(archive))

	return archive, nil
}
