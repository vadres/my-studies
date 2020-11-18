package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cod_player = [2]string{"x", "o"}

func Game() {
	player := 0
	for i := 0; i < 9; {

		if winner, hasWinner := Edge(); hasWinner {
			fmt.Printf("\n================\n")
			fmt.Printf("%s vencedor", strings.ToUpper(winner))
			fmt.Printf("\n================\n")
			break
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s Jogando - ", strings.ToUpper(cod_player[player]))
		fmt.Print("(lin,col): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		var coord = strings.Split(text, ",")
		if l, errL := strconv.Atoi(coord[0]); errL == nil {
			if c, errC := strconv.Atoi(coord[1]); errC == nil {
				if erro := Play(cod_player[player], l, c); erro == nil {
					i++
					//fmt.Println(text)
					if player == 0 {
						player = 1
					} else {
						player = 0
					}
				} else {
					fmt.Println(erro)
				}

				continue
			}
		}

		fmt.Println("Use o padrão de coordenada: num,num")
	}
}
