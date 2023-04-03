package funcoes

import (
	"fmt"
	"os"
	"time"
)

func registraLogErro(site string, status string, erro string) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	arquivo.WriteString("Nome do Site: " + site + "\n" +
		"Horário da Consulta: " + time.Now().Format("02/01/2006 15:04:05") + "\n" +
		"Situação do Site: " + status + "\n" +
		"Erro: " + erro + "\n" + "\n")
	arquivo.Close()
}

func registraLogSucesso(sliceDados []string) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	arquivo.WriteString("Nome do Site: " + sliceDados[0] + "\n" +
		"Horário da Consulta: " + time.Now().Format("02/01/2006 15:04:05") + "\n" +
		"Situação do Site: " + sliceDados[1] + "\n" +
		"Tamanho da Resposta: " + sliceDados[2] + "\n" +
		"Tempo de Resposta: " + sliceDados[4] + "\n" +
		"Código de status: " + sliceDados[3] + "\n" + "\n")

	arquivo.Close()
}
