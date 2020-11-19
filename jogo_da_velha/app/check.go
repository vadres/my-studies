package app

func Edge() (string, bool) {
	board := Board()
	op, xp := 0, 0
	for i := 0; i < 3; i++ {
		for u := 0; u < 3; u++ {
			plusDecrease(board[i][u] == "o", &op, &xp)
			plusDecrease(board[i][u] == "x", &xp, &op)
		}
		if op >= 3 {
			return "o", true
		}

		if xp >= 3 {
			return "x", true
		}
		xp, op = 0, 0
	}

	for u := 0; u < 3; u++ {
		for i := 0; i < 3; i++ {
			plusDecrease(board[i][u] == "o", &op, &xp)
			plusDecrease(board[i][u] == "x", &xp, &op)
		}

		if op >= 3 {
			return "o", true
		}

		if xp >= 3 {
			return "x", true
		}
		xp, op = 0, 0
	}

	if diag(board, "0") {
		return "o", true
	}
	if diag(board, "x") {
		return "x", true
	}

	return "", false
}

func plusDecrease(cond bool, toPlus, toDecrease *int) {
	if cond {
		*toPlus++
		*toDecrease--
	}
}

func diag(board [][]string, s string) bool {
	return (board[2][2] == s &&
		board[1][1] == s &&
		board[0][0] == s) ||
		(board[2][0] == s &&
			board[1][1] == s &&
			board[0][2] == s)
}
