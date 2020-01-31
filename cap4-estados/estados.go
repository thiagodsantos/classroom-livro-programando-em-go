package main

import "fmt"

type Estado struct {
	nome      string
	populacao int
	capital   string
}

func main() {
	// string é a chave e Estado é o tipo do dado
	estados := make(map[string]Estado, 2)

	estados["GO"] = Estado{"Goiás", 648264, "Goiânia"}
	estados["PB"] = Estado{"Paraíba", 739572, "João Pessoa"}
	estados["SP"] = Estado{"São Paulo", 2395272, "São Paulo"}

	fmt.Println(estados)

	// Retornando como array associativo
	saoPaulo, encontrado := estados["SP"]
	if encontrado {
		fmt.Println(saoPaulo)
	}

	// Somente para verificar se existe
	_, goEncontrado := estados["GO"]
	if goEncontrado {
		fmt.Println("GO encontrado!")
	}

	// Removendo
	delete(estados, "PB")
	fmt.Println(estados)

	// Iterando
	for sigla, estado := range estados {
		fmt.Printf("%s, (%s) possui %d habitantes.\n", estado.nome, sigla, estado.populacao)
	}
}
