package funcoes

import (
	"net/http"
	"strconv"
	"time"
)

func testaSite(site string) {

	client := &http.Client{Timeout: 10 * time.Second}
	start := time.Now()
	response, err := client.Get(site)

	if err != nil {
		const status = "offline"
		s := err.Error()
		erro := "Site não respondeu a requisição: " + s
		registraLogErro(site, status, erro)
		return
	}
	defer response.Body.Close()

	responseTime := time.Since(start).Seconds()
	statusCode := response.StatusCode
	responseSize := response.ContentLength

	const status = "online"
	size := strconv.Itoa(int(responseSize))
	stcode := strconv.Itoa(statusCode)
	respTime := strconv.FormatFloat(responseTime, 'f', 6, 64)

	sliceDados := []string{site, status, size, stcode, respTime}

	registraLogSucesso(sliceDados)
	return

}
