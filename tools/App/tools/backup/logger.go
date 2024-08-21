package backup

import (
	"fmt"
	"log"
	"os"
	"time"
)

const LogFilePath = "/var/log/bedrock-backup.log"

func LogError(message string) {
	logToFile(fmt.Sprintf("[%s] ERRO: %s", time.Now().Format(time.RFC3339), message))
}

func LogSuccess(message string) {
	logToFile(fmt.Sprintf("[%s] SUCESSO: %s", time.Now().Format(time.RFC3339), message))
}

func logToFile(message string) {
	f, err := os.OpenFile(LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(message + "\n"); err != nil {
		log.Fatal(err)
	}
}
