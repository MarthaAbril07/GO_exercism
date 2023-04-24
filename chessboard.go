package main

import (
	"fmt"
)

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard = map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	cont := 0
	ocupados := cb[file]
	for _, celda := range ocupados {
		if celda {
			cont++
		}
	}
	return cont
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	cont := 0
	for _, filas := range cb {
		for numFila, ocupada := range filas {
			if (numFila+1) == rank && ocupada {
				cont++
			}
		}
	}
	return cont
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * 8
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	cont := 0
	for _, filas := range cb {
		for _, ocupada := range filas {
			if ocupada {
				cont++
			}
		}
	}
	return cont
}

func main() {
	board := Chessboard{
		"A": {true, false, true, false, false, false, false, true},
		"B": {false, false, false, false, true, false, false, false},
		"C": {false, false, true, false, false, false, false, false},
		"D": {false, false, false, false, false, false, false, false},
		"E": {false, false, false, false, false, true, false, true},
		"F": {false, false, false, false, false, false, false, false},
		"G": {false, false, false, true, false, false, false, false},
		"H": {true, true, true, true, true, true, false, true},
	}
	fmt.Println(CountInFile(board, "A"))
	fmt.Println(CountInRank(board, 2))
	fmt.Println(CountAll(board))
	fmt.Println(CountOccupied(board))

}
