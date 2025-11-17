package menu

import (
	"Main.go/monitor"
	"fmt"
	"os"
)

func ExibirMenuSwitch() {

	for {

		fmt.Println("1 - iniciar monitoramento")
		fmt.Println("2 - exibir logs")
		fmt.Println("0 - sair do programa")

		comandoSwitch := comandoSwitch()

		switch comandoSwitch {
		case 1:
			monitor.IniciarMonitoramento()
		case 2:
			fmt.Println("exibindo logs...")
			monitor.ExibirLogs()
		case 0:
			fmt.Println("saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("comando invalido no switch")
			os.Exit(-1)
		}
	}

}

func comandoSwitch() int {
	var comandoSwitch int
	fmt.Println("digite um comando para usar o Switch: ")

	_, err := fmt.Scan(&comandoSwitch)

	if err != nil {
		return -1
	}

	fmt.Println("o comando escolhido foi", comandoSwitch)

	return comandoSwitch
}
