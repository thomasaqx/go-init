package monitor

import (
	"fmt"
	"io/ioutil"
)

func ExibirLogs(){
	
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
	
	
}