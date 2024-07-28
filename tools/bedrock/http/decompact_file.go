package http

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)


func UnzipFile(src, dest string) error {
	// Abre o arquivo ZIP
	r, err := zip.OpenReader(src)
	if err != nil {
		return err 
	}
	defer r.Close() 

	// Itera sobre cada arquivo dentro do arquivo ZIP
	for _, fileItem := range r.File {
		// Concatena o caminho de destino com o nome do arquivo
		fpath := filepath.Join(dest, fileItem.Name)

		// Verifica se o arquivo é um diretório
		if fileItem.FileInfo().IsDir() {
			// Cria o diretório
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Cria os diretórios necessários para os arquivos
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err // Retorna o erro se não conseguir criar os diretórios
		}

		// Abre o arquivo de saída para escrita
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileItem.Mode())
		if err != nil {
			return err // Retorna o erro se não conseguir abrir o arquivo de saída
		}

		// Abre o arquivo ZIP para leitura
		rc, err := fileItem.Open()
		if err != nil {
			return err // Retorna o erro se não conseguir abrir o arquivo ZIP para leitura
		}

		// Copia o conteúdo do arquivo ZIP para o arquivo de saída
		_, err = io.Copy(outFile, rc)

		// Fecha o arquivo de saída e o arquivo ZIP após a cópia
		outFile.Close()
		rc.Close()

		if err != nil {
			return err // Retorna o erro se não conseguir copiar o conteúdo
		}
	}
	return nil // Retorna nil se tudo correr bem
}