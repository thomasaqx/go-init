package menu

import (
	"Main.go/monitor"
	"fmt"
	"os"
)

func ExibirMenuIf() {

	for {

		fmt.Println("1 - iniciar monitoramento")
		fmt.Println("2 - exibir logs")
		fmt.Println("0 - sair do programa")

		comandoIf := ComandoIf()

		if comandoIf == 1 {
			monitor.IniciarMonitoramento()
		} else if comandoIf == 2 {
			fmt.Println("exibindo logs...")
			monitor.ExibirLogs()
		} else if comandoIf == 0 {
			fmt.Println("saindo do programa...")
			os.Exit(0)
		} else {
			fmt.Println("comando invalido, digite novamente")
		}
	}
}

func ComandoIf() int {
	var comandoInt int
	fmt.Println("digite um comando para usar o if :")

	_, err := fmt.Scan(&comandoInt)
	if err != nil {
		return -1
	}

	fmt.Println("o comando escolhido foi", comandoInt)

	return comandoInt
}
