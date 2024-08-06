package main

import (
	"github.com/GomdimApps/MineServerTools/tools/bedrock/update/http"
)

func main() {
	// Chama a função RunDownloadProcess sem tratar erro, pois a função não retorna erro
	http.RunDownloadProcess()
}
