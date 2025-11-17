package monitor

import (
	"fmt"
	"net/http"
)

func IniciarMonitoramento() {
	fmt.Println("iniciando monitoramento...")
	site := "https://idealctvm.com.br/pt"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
