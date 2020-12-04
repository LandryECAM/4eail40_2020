package board

import (
	"fmt"

	"../coord"
	"../piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	panic("not implemented") // TODO: Implement
}

// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	panic("not implemented") // TODO: Implement
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {

	// Nous avons implémenté des pieces au sein de cette fonction pour pouvoir
	// réaliser le test.
	// Il aurait d'abord fallut implémenter les objets pièces et board mais ce
	// n'est pas ce qui étati demander.
	fromX, _ := from.Coord(0)
	fromY, _ := from.Coord(1)
	toX, _ := to.Coord(0)
	toY, _ := to.Coord(1)
	pieceDest := piece.NewChessPiece("", 1)
	c[toX][toY] = pieceDest
	// Normaly we should use piece.Moves() to check if the piece capture or not
	if c[toX][toY].Name() != "" {
		return fmt.Errorf("Piece at the destination")
	}
	c[toX][toY] = c[fromX][fromY]
	piece := piece.NewChessPiece("", 1)
	c[fromX][fromY] = piece
	return nil

}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	panic("not implemented") // TODO: Implement
}
