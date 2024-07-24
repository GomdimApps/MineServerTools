package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetSystemInfo() {  
	cpuPercent, err := cpu.Percent(time.Second, false)

	if err != nil {
    fmt.Printf("Error getting CPU percent: %v\n", err)
    return
  }

	memory, err := mem.VirtualMemory()

	if err!= nil {
    fmt.Printf("Error getting memory usage: %v\n", err)
    return
  }

	d, _ := disk.Usage("/")

	// convertendo os valores de bytes para gigabytes para uma melhor legibilidade
	totalGB := d.Total / 1024 / 1024 / 1024
	freeGB := d.Free / 1024 / 1024 / 1024
	usedGB := d.Used / 1024 / 1024 / 1024

	fmt.Printf("-------Gerenciador do sistema--------\n\n")

  fmt.Printf("Memory Total: %v MB\n",memory.Total / 1024 / 1024)
	fmt.Printf("Memory Free: %v MB\n", memory.Free / 1024 / 1024)
	if memory.UsedPercent >= 85 {
		fmt.Println("##### WARNING: memory usage is over than 85%")
	}
	fmt.Printf("Memory Usage: %.2f%%\n\n", memory.UsedPercent)

	fmt.Printf("CPU Usage: %.2f%%\n\n", cpuPercent[0])

	fmt.Printf("Total de Disco: %v GB\n", totalGB)
	fmt.Printf("Disco Livre: %v GB\n", freeGB)
	fmt.Printf("Disco Usado: %v GB (%.2f%%)\n\n\n", usedGB, d.UsedPercent)
}