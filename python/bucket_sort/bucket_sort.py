def insertion_sort(balde: list[float]) -> None:
    for i in range(1, len(balde)):
        chave = balde[i]
        j = i - 1
        while j >= 0 and balde[j] > chave:
            balde[j + 1] = balde[j]
            j-=1
        balde[j + 1] = chave

def bucket_sort(lista_numeros: list[float]) -> list[float]:
    n = len(lista_numeros)
    if n == 0:
        return list()

    baldes = [[] for _ in range(n)]
    for num in lista_numeros:
        indice = min(int(num * n), n - 1)
        baldes[indice].append(num)

    resultado = []
    for balde in baldes:
        insertion_sort(balde)
        resultado.extend(balde)
    return resultado

if __name__ == "__main__":
    dados = [0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68]
    print(bucket_sort(dados))
