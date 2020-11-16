package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var board = [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}

var cod_player = [2]string{"x", "o"}

func main() {
	player := 0
	for i := 0; i < 9; {

		if hasWinner, winner := check(); hasWinner {
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
				var gameCond, errorMessage = play(player, l, c)
				if gameCond {
					i++
					//fmt.Println(text)
					if player == 0 {
						player = 1
					} else {
						player = 0
					}
				} else {
					fmt.Println(errorMessage)
				}

				continue
			}
		}

		fmt.Println("Use o padrão de coordenada: num,num")
	}
}

func play(player, l, c int) (bool, string) {
	l--
	c--
	switch {
	case l < 0 || c < 0 || l > 2 || c > 2:
		return false, "Fora dos limites do tabuleiro!"
	case board[l][c] != "_":
		return false, "Campo preenchido!"
	}

	board[l][c] = cod_player[player]

	print()
	return true, ""
}

func print() {
	fmt.Println("---------------------------")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println("---------------------------")
}

func check() (hasWinner bool, winner string) {
	op, xp := 0, 0
	for i := 0; i < 3; i++ {
		for u := 0; u < 3; u++ {
			if board[i][u] == "o" {
				op++
				xp--
			} else if board[i][u] == "x" {
				xp++
				op--
			}
		}
		if op >= 3 {
			return true, "o"
		}

		if xp >= 3 {
			return true, "x"
		}
		xp, op = 0, 0
	}

	for u := 0; u < 3; u++ {
		for i := 0; i < 3; i++ {
			if board[i][u] == "o" {
				op++
				xp--
			} else if board[i][u] == "x" {
				xp++
				op--
			}
		}

		if op >= 3 {
			return true, "o"
		}

		if xp >= 3 {
			return true, "x"
		}
		xp, op = 0, 0
	}

	for u := 0; u < 3; u++ {
		if board[u][u] == "o" {
			op++
			xp--
		} else if board[u][u] == "x" {
			xp++
			op--
		}
	}

	if op >= 3 {
		return true, "o"
	}

	if xp >= 3 {
		return true, "x"
	}

	if board[2][0] == "o" &&
		board[1][1] == "o" &&
		board[0][2] == "o" {
		return true, "o"
	}
	if board[2][0] == "x" &&
		board[1][1] == "x" &&
		board[0][2] == "x" {
		return true, "x"
	}

	return false, ""
}
