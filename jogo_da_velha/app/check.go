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

	if board[2][2] == "o" &&
		board[1][1] == "o" &&
		board[0][0] == "o" {
		return "o", true
	}
	if board[2][2] == "x" &&
		board[1][1] == "x" &&
		board[0][0] == "x" {
		return "x", true
	}

	if board[2][0] == "o" &&
		board[1][1] == "o" &&
		board[0][2] == "o" {
		return "o", true
	}
	if board[2][0] == "x" &&
		board[1][1] == "x" &&
		board[0][2] == "x" {
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
