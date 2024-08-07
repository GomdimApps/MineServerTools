package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

const ConfigFilePath = "/etc/mineservertools/bedrock-server.conf"

func GetServerDir() string {
	content, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo de configuração: %v\n", err)
		return ""
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "server-dir=") {
			return strings.Trim(strings.Split(line, "=")[1], "\"")
		}
	}
	return ""
}

func DirSize(path string) (uint64, error) {
	var size uint64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += uint64(info.Size())
		}
		return nil
	})
	return size, err
}

func main() {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("Error getting CPU percent: %v\n", err)
		return
	}

	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error getting memory usage: %v\n", err)
		return
	}

	d, _ := disk.Usage("/")
	totalGB := d.Total / 1024 / 1024 / 1024
	freeGB := d.Free / 1024 / 1024 / 1024
	usedGB := d.Used / 1024 / 1024 / 1024

	if memory.UsedPercent >= 85 {
		fmt.Println("##### AVISO: a utilização da memória é superior a 85%")
	}

	serverDir := GetServerDir()
	if serverDir == "" {
		fmt.Println("Erro ao obter o diretório do servidor.")
		return
	}

	serverDirSize, err := DirSize(serverDir)
	if err != nil {
		fmt.Printf("Error getting server directory size: %v\n", err)
		return
	}

	fmt.Printf("\nUtilização da CPU: %.2f%%\n\n", cpuPercent[0])
	fmt.Printf("Utilização da memória: %.2f%%\n\n", memory.UsedPercent)
	fmt.Printf("Memória Total: %v MB\n", memory.Total/1024/1024)
	fmt.Printf("Memória livre: %v MB\n", memory.Free/1024/1024)
	fmt.Printf("Total de Disco: %v GB\n", totalGB)
	fmt.Printf("Disco Livre: %v GB\n", freeGB)
	fmt.Printf("Disco Usado: %v GB (%.2f%%)\n\n", usedGB, d.UsedPercent)

	if serverDirSize < 1024*1024*1024 {
		fmt.Printf("Tamanho total da pasta do servidor: %v MB\n", serverDirSize/1024/1024)
	} else {
		fmt.Printf("Tamanho total da pasta do servidor: %.2f GB\n", float64(serverDirSize)/(1024*1024*1024))
	}
}
