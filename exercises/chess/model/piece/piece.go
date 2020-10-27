package piece

import (
	"fmt"

	"../coord"
	"../player"
)

type ChessPiece struct {
	name   string
	player player.Color
}

func NewChessPiece(name string, player player.Color) ChessPiece {
	return ChessPiece{name, player}
}

func (piece ChessPiece) String() string {
	return fmt.Sprintf(piece.name)
}

func (piece ChessPiece) Color() player.Color {
	// TODO:
	return piece.player
}

func (piece ChessPiece) Moves(isCapture bool) map[coord.ChessCoordinates]bool {
	// TODO:
	return nil
}

func (piece ChessPiece) Name() string {
	return piece.name
}
