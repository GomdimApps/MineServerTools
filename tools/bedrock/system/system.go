package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
		fmt.Println("Erro ao ler o arquivo de configuração do Bedrock")
		return ""
	}

	for _, line := range strings.Split(string(content), "\n") {
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
		fmt.Printf("Erro ao obter a porcentagem de uso da CPU: %v\n", err)
		return
	}

	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Erro ao obter o uso da memória: %v\n", err)
		return
	}

	diskUsage, _ := disk.Usage("/")
	totalGB := diskUsage.Total / 1024 / 1024 / 1024
	freeGB := diskUsage.Free / 1024 / 1024 / 1024
	usedGB := diskUsage.Used / 1024 / 1024 / 1024

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
		fmt.Printf("Erro ao obter o tamanho do diretório do servidor: %v\n", err)
		return
	}

	fmt.Printf("\nUtilização da CPU: %.2f%%\n", cpuPercent[0])
	fmt.Printf("Utilização da memória: %.2f%%\n\n", memory.UsedPercent)
	fmt.Printf("Memória Total: %v MB\n", memory.Total/1024/1024)
	fmt.Printf("Memória livre: %v MB\n", memory.Free/1024/1024)
	fmt.Printf("Total de Disco: %v GB\n", totalGB)
	fmt.Printf("Disco Livre: %v GB\n", freeGB)
	fmt.Printf("Disco Usado: %v GB (%.2f%%)\n\n", usedGB, diskUsage.UsedPercent)

	if serverDirSize < 1024*1024*1024 {
		fmt.Printf("Tamanho do servidor: %v MB\n", serverDirSize/1024/1024)
	} else {
		fmt.Printf("Tamanho do servidor: %.2f GB\n", float64(serverDirSize)/(1024*1024*1024))
	}

	cmd := exec.Command("sh", "-c", "top -b -n 1 | grep -q \"bedrock\" && echo \"Operando\" || echo \"Inativo\"")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Erro ao verificar o status do servidor")
		return
	}
	fmt.Printf("Status do servidor: %s", output)

	cmd = exec.Command("sh", "-c", "ip route get 8.8.8.8 | awk '{print $7}'")
	ipOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("Erro ao obter o IP do servidor: %v\n", err)
		return
	}
	fmt.Printf("\nIP do servidor: %s\n", strings.TrimSpace(string(ipOutput)))

	cmd = exec.Command("sh", "-c", `output=$(lsof -i -P -n | grep bedrock | awk '/IPv4/ {print "IPv4: " $9} /IPv6/ {print "IPv6: " $9}'); if [ -z "$output" ]; then echo "Servidor inativo"; else echo "$output"; fi`)
	portsOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("Erro ao obter as portas utilizadas: %v\n", err)
		return
	}
	fmt.Printf("Portas utilizadas:\n%s", string(portsOutput))
}
