package main

import "fmt"

func status(board [][]byte, word string, i, j int) bool {
	if word == "" {
		return true
	}

	if board[i][j] != word[0] {
		return false
	}

	if len(word) == 1 {
		return true
	}

	board[i][j] = '#'
	if i > 0 && status(board, word[1:], i-1, j) {
		return true
	}
	if j > 0 && status(board, word[1:], i, j-1) {
		return true
	}
	if i < len(board)-1 && status(board, word[1:], i+1, j) {
		return true
	}
	if j < len(board[0])-1 && status(board, word[1:], i, j+1) {
		return true
	}
	board[i][j] = word[0]
	return false
}
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return word == ""
	}
	if word == "" {
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] != board[i][j] {
				continue
			}
			if status(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func main() {
	var word string
	word = "ABCDE"
	board := [][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'S', 'F', 'C', 'S'}}

	fmt.Println(exist(board, word))

}
