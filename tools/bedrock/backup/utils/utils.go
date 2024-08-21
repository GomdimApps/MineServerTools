package utils

import (
	"archive/tar"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/GomdimApps/MineServerTools/tools/bedrock/backup/config"
	"github.com/GomdimApps/MineServerTools/tools/bedrock/backup/logger"
)

const BackupDir = "/var/mine-backups/backup-server-bedrock/"

type BackupInfo struct {
	Name         string `json:"Name"`
	Size         int64  `json:"Size"`
	CreationDate string `json:"CreationDate"`
}

func Backup() {
	serverDir := config.GetServerDir()
	if serverDir == "" {
		logger.LogError("Diretório do servidor não encontrado.")
		os.Exit(1)
	}

	dateFile := time.Now().Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("server-mine-bedrock-%s.tar.zst", dateFile)
	tempTar := filepath.Join(os.TempDir(), "server-mine-bedrock-temp.tar.zst")

	if err := os.Chdir(BackupDir); err != nil {
		logger.LogError("Erro ao acessar o diretório de backup.")
		os.Exit(1)
	}

	removeOldBackups()

	fmt.Println("Iniciando backup...")

	if err := createTar(serverDir, tempTar); err != nil {
		logger.LogError("Erro ao criar o arquivo tar.")
		os.Exit(1)
	}

	if err := os.Rename(tempTar, filepath.Join(BackupDir, fileName)); err != nil {
		logger.LogError("Erro ao mover o arquivo de backup.")
		os.Exit(1)
	}

	logger.LogSuccess(fmt.Sprintf("Backup concluído com sucesso: %s", fileName))
	fmt.Printf("Backup concluído com sucesso: %s\n", fileName)
}

func ViewBackup() {
	if err := os.Chdir(BackupDir); err != nil {
		logger.LogError("Erro ao acessar os backups.")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(BackupDir)
	if err != nil {
		logger.LogError("Erro ao ler o diretório de backups.")
		os.Exit(1)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".tar.zst") {
			fmt.Printf("Nome: %s, Tamanho: %d bytes, Data de Criação: %s\n", file.Name(), file.Size(), file.ModTime().Format("2006-01-02 15:04:05"))
		}
	}
}

func ViewBackupJson() (string, error) {
	if err := os.Chdir(BackupDir); err != nil {
		logger.LogError("Erro ao acessar os backups.")
		return "", err
	}

	files, err := ioutil.ReadDir(BackupDir)
	if err != nil {
		logger.LogError("Erro ao ler o diretório de backups.")
		return "", err
	}

	var backups []BackupInfo
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".tar.zst") {
			backup := BackupInfo{
				Name:         file.Name(),
				Size:         file.Size(),
				CreationDate: file.ModTime().Format("2006-01-02 15:04:05"),
			}
			backups = append(backups, backup)
		}
	}

	// Converte a lista de backups para JSON
	jsonData, err := json.MarshalIndent(backups, "", "    ")
	if err != nil {
		logger.LogError("Erro ao converter os dados para JSON.")
		return "", err
	}

	return string(jsonData), nil
}

func removeOldBackups() {
	files, err := ioutil.ReadDir(BackupDir)
	if err != nil {
		logger.LogError("Erro ao ler o diretório de backups.")
		os.Exit(1)
	}

	var backups []os.FileInfo
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".tar.zst") {
			backups = append(backups, file)
		}
	}

	if len(backups) > 5 {
		for _, file := range backups[5:] {
			os.Remove(filepath.Join(BackupDir, file.Name()))
		}
	}
}

func createTar(sourceDir, tarPath string) error {
	file, err := os.Create(tarPath)
	if err != nil {
		return err
	}
	defer file.Close()

	zstdWriter := zlib.NewWriter(file)
	defer zstdWriter.Close()

	tarWriter := tar.NewWriter(zstdWriter)
	defer tarWriter.Close()

	return filepath.Walk(sourceDir, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(file, sourceDir)

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if fi.Mode().IsRegular() {
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := io.Copy(tarWriter, f); err != nil {
				return err
			}
		}
		return nil
	})
}
