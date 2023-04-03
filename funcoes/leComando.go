package funcoes

import "fmt"

func LeComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println()
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println()
	return comandoLido
}
