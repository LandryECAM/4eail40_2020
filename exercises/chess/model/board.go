// Package model contains the gameplay logic for the game of chess
package model

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

//TODO Implement type Board

type Board interface {
	Initialize()
}
type Coord interface {
}

type Board8x8 struct {
	board map[Coord2D]Piece
}
type Coord2D struct {
	X int
	Y int
}

func (board8x8 *Board8x8) Initialize() {
	board8x8.board = make(map[Coord2D]Piece)
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			coord := Coord2D{x, y}
			piece := Piece{" ", ""}
			if y == 1 {
				piece = Piece{"P", "blanc"}
			} else if y == 6 {
				piece = Piece{"P", "noir"}
			} else if y == 0 {
				if x == 0 || x == 7 {
					piece = Piece{"T", "blanc"}
				} else if x == 1 || x == 6 {
					piece = Piece{"C", "blanc"}
				} else if x == 2 || x == 5 {
					piece = Piece{"F", "blanc"}
				} else if x == 3 {
					piece = Piece{"R", "blanc"}
				} else if x == 4 {
					piece = Piece{"r", "blanc"}
				}
			} else if y == 7 {
				if x == 0 || x == 7 {
					piece = Piece{"T", "noir"}
				} else if x == 1 || x == 6 {
					piece = Piece{"C", "noir"}
				} else if x == 2 || x == 5 {
					piece = Piece{"F", "noir"}
				} else if x == 3 {
					piece = Piece{"R", "noir"}
				} else if x == 4 {
					piece = Piece{"r", "noir"}
				}
			}
			board8x8.board[coord] = piece
		}
	}
}

func (board8x8 *Board8x8) String() string {
	result := ""
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			coord := Coord2D{x, y}
			a := board8x8.board[coord]
			if a.color == "noir" {
				result += Cyan + a.name + Reset
			} else {
				result += a.name
			}
			result += " | "
		}
		result += "\n"
	}
	return result
}
