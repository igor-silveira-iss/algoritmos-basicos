def maior_soma_subarray(lista_numeros: list[int]) -> tuple[int, int, int]:
    if not lista_numeros:
        raise ValueError("a lista não pode ser vazia")

    melhor_soma = atual_soma = lista_numeros[0]
    melhor_inicio = melhor_fim = 0
    atual_inicio = 0

    for i in range(1, len(lista_numeros)):
        x = lista_numeros[i]
        if atual_soma < 0:
            atual_soma = x
            atual_inicio = i
        else:
            atual_soma += x

        if atual_soma > melhor_soma:
            melhor_soma = atual_soma
            melhor_inicio = atual_inicio
            melhor_fim = i

    return melhor_soma, melhor_inicio, melhor_fim

if __name__ == "__main__":
    numeros = [10, -45, 3, -9, 12]
    print(maior_soma_subarray(numeros))
