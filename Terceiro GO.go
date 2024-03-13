package main

import (
	"fmt"
)

type salsicha int

var vina salsicha = 10

func main() {
	tiposPróprios := "Dá pra criar os próprios tipos e encontrei uma forma de descobrir o tipo da variável de forma mais fácil!\n"
	fmt.Printf(tiposPróprios, `\n`, "%T")
	x := 10
	fmt.Printf("%v\n", x)
	fmt.Printf("%v", vina)
	vina = x
    //Volto depois para continuar!
}
