package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Olá mundo")
	fmt.Println("Precisa digitar isso tudo?")
	fmt.Println("E tem que ser em aspas duplas, hmm hmmm")
	texto := "Pra declarar variáveis tem que usar :=, wtf, tambem conhecido como Operador Curto de Declaração"
	fmt.Println(texto)
	números := "Declarar números"
	fmt.Println(números)
	digoNumeros := 1234
	fmt.Println(digoNumeros)
	numerosQuebrados := "Será se tem números flutuantes"
	fmt.Println(numerosQuebrados)
	dePontoFlutuante := 1.234
	fmt.Println(dePontoFlutuante)
	comVirgula := "Não funciona com virgula, a vírgula serve para separar argumentos ou valores"
	fmt.Println(comVirgula)
	um := 1
	valor := "valor"
	paraCada := "para cada"
	variavel := "variável"
	fmt.Println(um, valor, paraCada, variavel)
	fmt.Println("Olha, a concatenação funcionou conforme esperado, que loucura, ou será que não?")
	pergunta := "será se já tem números do tipo Pi nesse módulo ou preciso importa outros módulos para isso?"
	fmt.Println(pergunta)
	Pi := "Tentei, mas ele não reconhece Pi como uma função que retorna o valor de Pi, vou procurar uma lib com isso e já volto"
	fmt.Println(Pi)
	fmt.Println("Deixar aqui o Pi que eu lembro enquanto pesquiso a lib:", 3.1415)
	numeroPi := math.Pi
	fmt.Println(numeroPi)
	fmt.Println("Uuuhh, olha isso, lib math tem a constante de Pi, grandão até, quantos números tem? Tem como contar caracteres em go? seri len?")
	flutuantePstring := fmt.Sprintf("%f", numeroPi)
	tamanhoDoPi := len(flutuantePstring)
	fmt.Println(tamanhoDoPi)
	fmt.Println("Aí já deu um trabalho, tive que converter de Float pra String pra contar quantos caracteres o Pi padrão da Lib math tem, achei interessantemente desnecessário, mas funcionou")
    fmt.Println("Vou encerrar esse algoritmo por aqui, ficou grandinho, vou continuar em O segundo GO.go")
}
