import time
import random

class No:
    def __init__(self, valor):
        self.valor = valor
        self.esquerda = None
        self.direita = None

def inserir(raiz, valor):
    if raiz is None:
        return No(valor)
    if valor < raiz.valor:
        raiz.esquerda = inserir(raiz.esquerda, valor)
    elif valor > raiz.valor:
        raiz.direita = inserir(raiz.direita, valor)
    return raiz

def buscar(raiz, valor):
    while raiz is not None:
        if valor == raiz.valor:
            return True
        raiz = raiz.esquerda if valor < raiz.valor else raiz.direita
    return False

def em_ordem(raiz):
    if raiz is None:
        return []
    return em_ordem(raiz.esquerda) + [raiz.valor] + em_ordem(raiz.direita)

if __name__ == "__main__":
    r = None
    for v in [8, 3, 10, 1, 6, 14]:
        r = inserir(r, v)

    print(em_ordem(r))
    print(buscar(r, 6))
    print(buscar(r, 7))
