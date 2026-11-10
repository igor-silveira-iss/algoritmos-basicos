def bubble_sort(lista_numeros: list[int]) -> list[int]:
    """
    Método bubble_sort efetua a ordenação de uma lista/arryas de números
    """
    n = len(lista_numeros)
    for i in range(n):
        trocou = False
        for j in range(n - i - 1):
            if lista_numeros[j] > lista_numeros[j + 1]:
                print("Valor de j:", lista_numeros[j], "Valor de j+1:", lista_numeros[j + 1])
                lista_numeros[j], lista_numeros[j + 1] = lista_numeros[j + 1], lista_numeros[j]
                print(lista_numeros)
                trocou = True
        if not trocou:
            break
    return lista_numeros

numeros = [105, 110, 78, 99, 10, 3, 8]
bubble_sort(numeros)
