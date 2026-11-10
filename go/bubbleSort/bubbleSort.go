package main

import "fmt"

func bubbleSort(listaNumeros []int64) []int64 {
	n := len(listaNumeros)
	for i := 0; i < n; i++ {
		trocou := false
		for j := 0; j < n-i-1; j++ {
			if listaNumeros[j] > listaNumeros[j+1] {
				fmt.Println("Valor de j:", listaNumeros[j], "Valor de j+1:", listaNumeros[j+1])
				listaNumeros[j], listaNumeros[j+1] = listaNumeros[j+1], listaNumeros[j]
				fmt.Println(listaNumeros)
				trocou = true
			}
		}
		if !trocou {
			break
		}
	}
	return listaNumeros
}

func main() {
	listaNumeros := []int64{89, 56, 78, 10, 15}
	bubbleSort(listaNumeros)
}
