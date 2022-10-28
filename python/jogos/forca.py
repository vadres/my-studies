import random

def jogar():
    print("--------------------------------")
    print("---Bem vindo ao jogo da Forca---")
    print("--------------------------------")

    palavra = get_palavra().upper()
    letras_acertadas = ["_" for letra in palavra]

    enforcou = 5

    print_lista(letras_acertadas)
    while enforcou > 0 and not check_acertou(letras_acertadas):
        letra = input("Qual letra? ")
        letra = letra.strip().upper()

        if not letra_in_palavra(letra, palavra, letras_acertadas):
            enforcou = enforcou - 1
        print_lista(letras_acertadas)

    if check_acertou(letras_acertadas):
        print("Você ganhou!")
    else:
        print("Perdeu cão")


def get_palavra():
    with open("palavras.txt", "r") as arquivo:
        palavras = [linha.strip() for linha in arquivo]
        index = random.randrange(0, len(palavras))
        return palavras[index]


def letra_in_palavra(letra, palavra, letras_acertadas):
    index = 0
    acertou = False
    for l in palavra:
        if l == letra:
            letras_acertadas[index] = letra
            acertou = True
        index = index + 1
    return acertou


def check_acertou(lista):
    return "_" not in lista


def print_lista(lista):
    for l in lista:
        print(l, end=" ")
    print("")

