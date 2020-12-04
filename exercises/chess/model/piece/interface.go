// Package piece will handle game pieces.
package piece

import (
	"fmt"

	"../coord"
	"../player"
)

// Piece represents a game piece
type Piece interface {
	fmt.Stringer
	// Color returns the appartenance.
	Color() player.Color
	// Moves returns a set of valid move.
	Moves(isCapture bool) map[coord.ChessCoordinates]bool
	// Returns the name of the piece.
	Name() string
}
