package funcoes

import (
	"fmt"
	"time"
)

const monitoramentos = 5
const intervalo = 1

func IniciarMonitoramento() {
	fmt.Println()
	fmt.Println("-----------------------Monitorando... ---------------------")
	fmt.Println()
	sites := leSitesDoArquivo()
	fmt.Println()
	fmt.Println("--------------------------------------------------------------")
	sites = sites[:len(sites)-1]

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(intervalo * time.Second)
		fmt.Println("")
	}
}
