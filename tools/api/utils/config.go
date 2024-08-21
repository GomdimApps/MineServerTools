package utils

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

const configPath = "/etc/mineservertools/mtools.conf"

func getValueFromConfig(key string) (string, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			return strings.TrimPrefix(line, key+"="), nil
		}
	}

	return "", errors.New("chave não encontrada no arquivo de configuração")
}

func GetPortApi() (int, error) {
	portStr, err := getValueFromConfig("api-port")
	if err != nil {
		return 0, err
	}
	port, err := strconv.Atoi(portStr)
	return port, err
}

func GetTokenApi() (string, error) {
	return getValueFromConfig("token-api")
}
