package main

import (
	"fmt"
)

type No struct {
	Valor    int
	Esquerda *No
	Direita  *No
}

func inserirValor(raiz *No, valor int) *No {
	if raiz == nil {
		return &No{Valor: valor}
	}
	if valor < raiz.Valor {
		raiz.Esquerda = inserirValor(raiz.Esquerda, valor)
	} else if valor > raiz.Valor {
		raiz.Direita = inserirValor(raiz.Direita, valor)
	}
	return raiz
}

func buscarValor(raiz *No, valor int) bool {
	for raiz != nil {
		if valor == raiz.Valor {
			return true
		}
		if valor < raiz.Valor {
			raiz = raiz.Esquerda
		} else {
			raiz = raiz.Direita
		}
	}
	return false
}

func ordenarValores(raiz *No, saida []int) []int {
	if raiz == nil {
		return saida
	}
	saida = ordenarValores(raiz.Esquerda, saida)
	saida = append(saida, raiz.Valor)
	return ordenarValores(raiz.Direita, saida)
}

func main() {
	var raiz *No
	for _, v := range []int{8, 3, 10, 1, 6, 14} {
		raiz = inserirValor(raiz, v)
	}
	fmt.Println(ordenarValores(raiz, nil))
	fmt.Println(buscarValor(raiz, 6))
	fmt.Println(buscarValor(raiz, 7))
}
