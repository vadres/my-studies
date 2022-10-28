import random


def jogar():
    print("--------------------------------")
    print("Bem vindo ao jogo de Adivinhação")
    print("--------------------------------")

    numero_secreto = random.randrange(1, 101)
    print(numero_secreto)

    chute = int(input("Digite o numero: "))

    print(f"Você digitou {chute}")

    if numero_secreto == chute:
        print("Você acertou")
    else:
        print("Você errou")

    print("--------------------------------\n", "Fim de Jogo!")


if __name__ == "__main__":
    jogar()
