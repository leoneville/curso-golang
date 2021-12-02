package main

import "fmt"

func media(numeros ...float64) (total float64) {

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f\n", media(7.7, 8.1, 5.9, 9.9))
}
