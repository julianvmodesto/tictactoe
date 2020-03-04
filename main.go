package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	t := NewTicTacToe()
	t.Play()
}

type State int

const (
	Pending State = iota + 1
	Won
	Draw
)

type TicTacToe struct {
	board  [][]rune
	player rune
	moves  int
}

func NewTicTacToe() *TicTacToe {
	var board [][]rune
	board = make([][]rune, 3)
	for i := range board {
		board[i] = make([]rune, 3)
	}

	return &TicTacToe{
		board:  board,
		player: 'X',
	}
}

func (t *TicTacToe) GetMove() []int {
	var move []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Player %s's turn. Enter a row 0, 1, or 2:\n", t.Player())
	for scanner.Scan() {
		in := scanner.Text()
		row, err := strconv.Atoi(in)
		if err != nil || row < 0 || row > 2 {
			fmt.Printf("Got unexpected row '%s'\n", in)
			fmt.Printf("Player %s's turn. Enter a row 0, 1, or 2:\n", t.Player())
			continue
		}
		move = append(move, row)
		break
	}

	fmt.Printf("Player %s's turn. Enter a col 0, 1, or 2:\n", t.Player())
	for scanner.Scan() {
		in := scanner.Text()
		col, err := strconv.Atoi(in)
		if err != nil || col < 0 || col > 2 {
			fmt.Printf("Got unexpected col '%s'\n", in)
			fmt.Printf("Player %s's turn. Enter a col 0, 1, or 2:\n", t.Player())
			continue
		}
		move = append(move, col)
		break
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scan err: %v\n", err)
	}
	return move
}

func (t *TicTacToe) Play() {
	var err error
	current := Pending

	for current == Pending {
		t.Print()
		move := t.GetMove()
		current, err = t.PlayMove(move)
		if err != nil {
			fmt.Printf("move error: %v\n", err)
		} else {
			t.moves++
		}
	}
	switch current {
	case Won:
		fmt.Printf("Player %s won!\n", t.Player())
	case Draw:
		fmt.Println("Draw!")
	}
}

func (t *TicTacToe) PlayMove(move []int) (State, error) {
	if err := t.Move(move); err != nil {
		return Pending, err
	}

	if t.Won() {
		return Won, nil
	}
	t.SwitchPlayer()

	if t.moves < 9 {
		return Pending, nil
	}
	return Draw, nil
}

func (t *TicTacToe) Move(move []int) error {
	row, col := move[0], move[1]
	if t.board[row][col] != 0 {
		return fmt.Errorf("expected valid move to empty place, but place already taken by '%s'", string(t.board[row][col]))
	}
	t.board[row][col] = t.player
	return nil
}

func (t *TicTacToe) SwitchPlayer() {
	if t.player == 'X' {
		t.player = 'O'
	} else {
		t.player = 'X'
	}
}

func (t *TicTacToe) Won() bool {
	for row := 0; row < len(t.board); row++ {
		wonRow := true
		for col := 0; col < len(t.board); col++ {
			if t.board[row][col] != t.player {
				wonRow = false
				break
			}
		}
		if wonRow {
			return true
		}
	}

	for col := 0; col < len(t.board); col++ {
		wonCol := true
		for row := 0; row < len(t.board); row++ {
			if t.board[row][col] != t.player {
				wonCol = false
				break
			}
		}
		if wonCol {
			return true
		}
	}

	wonDiagonal := true
	for i := 0; i < len(t.board); i++ {
		if t.board[i][i] != t.player {
			wonDiagonal = false
			break
		}
	}
	if wonDiagonal {
		return true
	}
	wonDiagonal = true
	for i := 0; i < len(t.board); i++ {
		if t.board[i][2-i] != t.player {
			wonDiagonal = false
			break
		}
	}
	if wonDiagonal {
		return true
	}
	return false
}

func (t *TicTacToe) Player() string {
	return string(t.player)
}

func (t *TicTacToe) Print() {
	fmt.Printf("     0   1   2 \n")
	fmt.Printf("    -----------\n")
	for row := range t.board {
		for col := range t.board[row] {
			if col == 0 {
				fmt.Printf(" %d |", row)
				fmt.Printf(" ")
			}
			if t.board[row][col] != 0 {
				fmt.Printf("%s", string(t.board[row][col]))
			} else {
				fmt.Printf(" ")
			}
			if col != 2 {
				fmt.Printf(" | ")
			} else {
				fmt.Println()
			}
		}
		if row != 2 {
			fmt.Printf("    --- --- ---\n")
		} else {
			fmt.Println()
		}
	}
}
