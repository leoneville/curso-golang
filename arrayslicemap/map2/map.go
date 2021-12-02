package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
		"José João":      11325.45,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
