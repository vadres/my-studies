import advinhacao
import forca

print("Jogos")
print("1 - Advinhação  2 - Forca")
game = int(input("Escolha o jogo: "))
if game == 1:
    advinhacao.jogar()
elif game == 2:
    forca.jogar()
