package game

import (
	"fmt"
	"math/rand"
)

// Constants for board dimensions and ship representation
const (
	BoardSize      = 10
	EmptyCell      = '.'
	ShipCellMarker = 'S'
)

// Board represents a 2D grid for the game
type Board struct {
	Grid [][]rune
}

// NewBoard initializes an empty game board
func NewBoard() *Board {
	grid := make([][]rune, BoardSize)
	for i := range grid {
		grid[i] = make([]rune, BoardSize)
		for j := range grid[i] {
			grid[i][j] = EmptyCell
		}
	}
	return &Board{Grid: grid}
}

// PlaceShip randomly places a ship of given size on the board
func (b *Board) PlaceShip(shipSize int) bool {
	for {
		row := rand.Intn(BoardSize)
		col := rand.Intn(BoardSize)
		horizontal := rand.Intn(2) == 0 // 0 for horizontal, 1 for vertical for now

		if b.canPlaceShip(row, col, shipSize, horizontal) {
			for i := 0; i < shipSize; i++ {
				if horizontal {
					b.Grid[row][col+i] = ShipCellMarker
				} else {
					b.Grid[row+i][col] = ShipCellMarker
				}
			}
			return true
		}
	}
}

// canPlaceShip checks if a ship can be placed at the specified location
func (b *Board) canPlaceShip(row, col, shipSize int, horizontal bool) bool {
	for i := 0; i < shipSize; i++ {
		r := row
		c := col
		if horizontal {
			c += i
		} else {
			r += i
		}

		if r >= BoardSize || c >= BoardSize || b.Grid[r][c] != EmptyCell {
			return false
		}
	}
	return true
}

// PrintBoard displays the board in the console (for debugging)
func (b *Board) PrintBoard() {
	for _, row := range b.Grid {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}
