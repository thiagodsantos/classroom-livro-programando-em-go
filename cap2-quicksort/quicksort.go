package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {
		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s não é um número válido!\n", n)
			os.Exit(1)
		}
		numeros[i] = numero
	}

	//fmt.Println(quicksort(numeros))
}
