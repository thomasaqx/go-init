package menu

import (
	"fmt"
	"strings"
)

func ExibirMenu() {
	fmt.Print("Digite 'if' para usar o menu com if-else, ou 'swt' para usar o menu com Switch: ")
	var comando string
	fmt.Scan(&comando)

	
	// converte entrada de maiúsculo para minúsculo
	switch strings.ToLower(comando) {
	case "if":
		ExibirMenuIf()
	case "swt":
		ExibirMenuSwitch()
	default:
		fmt.Println("comando invalido")
	}
}
