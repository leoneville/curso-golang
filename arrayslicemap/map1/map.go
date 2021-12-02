package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[43224334243] = "Pedro"
	aprovados[45364656321] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[45364656321])
	delete(aprovados, 45364656321)
	fmt.Println(aprovados[45364656321])
}
