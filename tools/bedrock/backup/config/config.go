package config

import (
	"io/ioutil"
	"strings"

	"github.com/GomdimApps/MineServerTools/tools/bedrock/backup/logger"
)

const ConfigFilePath = "/etc/mineservertools/bedrock-server.conf"

func GetServerDir() string {
	content, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		logger.LogError("Erro ao ler o arquivo de configuração.")
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
