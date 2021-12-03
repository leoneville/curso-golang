package main

import "fmt"

// Uma solução melhor seria ...

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	resultado := factorial(5)
	fmt.Println(resultado)
}
