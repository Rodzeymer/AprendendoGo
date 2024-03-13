package main

import (
	"fmt"
	"reflect"
)

func main() {
	primeiroParágrafo := "Essa primeira parte do algoritmo em GO me lembra Java, deve ser por isso que GO é mais leve que Python, Python te coloca muita coisa na mão que talvez você nem use, agora GO não te dá nada, tem que importar tudo."
	segundoParágrafo := "Mas essa parte de importação te deixa com um algoritmo mais leva, mas o programador tem que ficar mais atento ao que precisa pra fazer o programa funcionar, no fim das contas são tradeoff diferentes."
	fmt.Println(primeiroParágrafo)
	fmt.Println(segundoParágrafo)
	tipos := "Como saber qual tipo de dado estamos tratando?"
	fmt.Println(tipos, "Claro que temos mais uma lib para importar, a 'reflect'")
	numeroInteiro := 10
	numeroQuebrado := 13.31
	string := "Stringer Things"
	booleanum := false
	fmt.Println(reflect.TypeOf(numeroInteiro))
	fmt.Println(reflect.TypeOf(numeroQuebrado))
	fmt.Println(reflect.TypeOf(string))
	fmt.Println(reflect.TypeOf(booleanum))
	fmt.Println("Legal, agora podemos descobrir se uma string é uma string e se um número quebrado, decimal ou fracionário ou uma string")
	interpreteA_String := "Coloque em uma nova linha\nAgora coloque uma t\tabulação, eita, calma, que \"apressadinho\""
	fmt.Println(interpreteA_String)
	fmt.Println("Estamos nos entendendo")
	stringLiteral := `"Coloque em uma nova linha\nAgora coloque uma t\tabulação, eita, calma, que \"apressadinho\""`
	fmt.Println(stringLiteral)
	fmt.Println("Deu certo?", "Maravilha, tudo certo.", "Até agora pelo menos.")
	poemaLiteral := `"Rapaz, dá pra brincar bastante,
				subir na estante,
						brincar de escada,
							   a cada degrau 
								que subimos, e
								quanto mais subimos
										nos mostra 
											que dá para subir ainda mais,
														e ampliar nossos horizontes.
	O que me lembra da minha frase favorita
	'Sempre se mantenha curioso, 
	é mais produtivo viver com o peso da busca 
	que com o vazio da ignorância."`
	fmt.Println(poemaLiteral)
	fmt.Println("Gostei, dá pra brincar com a formatação de forma muito intuitiva", "Gostosinho esse tal de GO")
	jáDeu := "Pronto, já deu por este algoritmo no Terceiro Go continuamos a aprender essa delícia"
    fmt.Println(jáDeu)
	
}
