package http

import (
	"fmt"
	"os"
	"path/filepath"
)

func SaveFile(filename string, content []byte) error {
	fmt.Printf("Salvando arquivo em: %s\n", filename)

	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	fmt.Println("Arquivo salvo com sucesso")

	return nil
}
