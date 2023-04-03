package funcoes

import (
	"fmt"
	"io/ioutil"
)

func ImprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(arquivo))
	fmt.Println()
	fmt.Println("--------------------- FIM DOS LOGS-----------------------------------------------")
}
