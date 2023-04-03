package funcoes

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			fmt.Println(sites)
			break
		}
	}
	arquivo.Close()
	return sites

}
