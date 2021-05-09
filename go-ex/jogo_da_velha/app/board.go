package app

import (
	"errors"
	"fmt"
	"strings"
)

var board = [][]string{
	{"_", "_", "_"},
	{"_", "_", "_"},
	{"_", "_", "_"},
}

func Board() [][]string {
	return board
}

func Play(simbol string, l, c int) error {
	l--
	c--
	switch {
	case l < 0 || c < 0 || l > 2 || c > 2:
		return errors.New("Fora dos limites do tabuleiro!")
	case board[l][c] != "_":
		return errors.New("Campo preenchido!")
	}

	board[l][c] = simbol

	print()
	return nil
}

func print() {
	fmt.Println("---------------------------")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println("---------------------------")
}
