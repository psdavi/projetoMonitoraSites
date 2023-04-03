package main

import (
	"fmt"
	"os"
	"projetoMonitora/funcoes"
)

func main() {

	funcoes.ExibeIntroducao()

	for {
		funcoes.ExibeMenu()
		comando := funcoes.LeComando()

		switch comando {
		case 1:
			funcoes.IniciarMonitoramento()
		case 2:
			fmt.Println("---------------------------  Exibindo Logs... ----------------------------------------")
			fmt.Println()
			funcoes.ImprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}
}
