package backup

import (
	"fmt"
	"os/exec"
	"strings"
)

func ScheduleTasks() {
	fmt.Println("Agendando tarefas automáticas...")
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.Output()
	if err != nil {
		LogError("Erro ao listar crontab.")
		return
	}

	newCron := string(output) + "\n0 3 * * * /usr/bin/bedrock-tools/tools/backup-server --backup\n"
	cmd = exec.Command("crontab", "-")
	cmd.Stdin = strings.NewReader(newCron)
	if err := cmd.Run(); err != nil {
		LogError("Erro ao agendar tarefas.")
		return
	}

	LogSuccess("Tarefas agendadas.")
}
