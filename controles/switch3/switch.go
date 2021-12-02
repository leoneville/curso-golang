package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case func():
		return "Função"
	case bool:
		return "Booleano"
	default:
		return "Não conheço esse tipo"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(true))
	fmt.Println(tipo(time.Now()))
}
