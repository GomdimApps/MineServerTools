package scheduler

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/GomdimApps/MineServerTools/tools/bedrock/backup/logger"
)

func ScheduleTasks() {
	fmt.Println("Agendando tarefas automáticas...")
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.Output()
	if err != nil {
		logger.LogError("Erro ao listar crontab.")
		return
	}

	newCron := string(output) + "\n0 3 * * * /usr/bin/bedrock-tools/tools/backup-server --backup\n"
	cmd = exec.Command("crontab", "-")
	cmd.Stdin = strings.NewReader(newCron)
	if err := cmd.Run(); err != nil {
		logger.LogError("Erro ao agendar tarefas.")
		return
	}

	logger.LogSuccess("Tarefas agendadas.")
}
