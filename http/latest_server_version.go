package http

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetLatestServerVersion() (string, error) {
	fmt.Println("Buscando versão mais atualizada do server")
	const baseURL = "https://www.minecraft.net/en-us/download/server/bedrock"

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error, statusCode: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var downloadLink string
	doc.Find("a.downloadlink").Each(func(i int, s *goquery.Selection) {
		if platform, exists := s.Attr("data-platform"); exists && platform == "serverBedrockLinux" {
			downloadLink, _ = s.Attr("href")
		}
	})

	if downloadLink == "" {
		return "", fmt.Errorf("link para download não encontrado")
	}

	fmt.Println("Última versão disponível, iniciando download...")

	return downloadLink, nil
}