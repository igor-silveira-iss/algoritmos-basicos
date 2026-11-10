package main

import "fmt"

func insertionSort(balde []float64) {
	for i := 0; i < len(balde); i++ {
		chave := balde[i]
		j := i - 1
		for j >= 0 && balde[j] > chave {
			balde[j+1] = balde[j]
			j -= 1
		}
		balde[j+1] = chave
	}
}

func bucketSort(listaNumeros []float64) []float64 {
	n := len(listaNumeros)
	if n == 0 {
		return []float64{}
	}

	baldes := make([][]float64, n)
	for _, v := range listaNumeros {
		indice := int(v * float64(n))
		if indice >= n {
			indice = n - 1
		}
		baldes[indice] = append(baldes[indice], v)
	}

	resultado := make([]float64, 0, n)
	for _, balde := range baldes {
		insertionSort(balde)
		resultado = append(resultado, balde...)
	}
	return resultado
}

func main() {
	numeros := []float64{0.15, 0.85, 0.21, 0.19, 0.54}
	fmt.Println(bucketSort(numeros))
}
