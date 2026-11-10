package main

import (
	"errors"
	"fmt"
)

func maiorSomaSubarray(listaNumeros []int) ([]int, error) {
	if len(listaNumeros) == 0 {
		return nil, errors.New("a lista não pode ser vazia")
	}

	atualSoma := listaNumeros[0]
	melhorSoma := atualSoma
	melhorFim := 0
	melhorInicio := melhorFim
	atualInicio := 0

	for i := 0; i < len(listaNumeros); i++ {
		valorAtual := listaNumeros[i]
		if atualSoma < 0 {
			atualSoma = valorAtual
			atualInicio = i
		} else {
			atualSoma += valorAtual
		}

		if atualSoma > melhorSoma {
			melhorSoma = atualSoma
			melhorInicio = atualInicio
			melhorFim = i
		}
	}

	return []int{melhorSoma, melhorInicio, melhorFim}, nil
}

func main() {
	numeros := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maiorSomaSubarray(numeros))
}
