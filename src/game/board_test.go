package game

import (
	"math/rand"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	// board is the right size
	if len(board.Grid) != BoardSize {
		t.Fatalf("Expected board size %d, got %d", BoardSize, len(board.Grid))
	}

	for _, row := range board.Grid {
		if len(row) != BoardSize {
			t.Fatalf("Expected row size %d, got %d", BoardSize, len(row))
		}
	}

	// Ensure all cells init with EmptyCell
	for i := range board.Grid {
		for j := range board.Grid[i] {
			if board.Grid[i][j] != EmptyCell {
				t.Fatalf("Expected EmptyCell at (%d, %d), got %c", i, j, board.Grid[i][j])
			}
		}
	}
}

func TestPlaceShip(t *testing.T) {
	board := NewBoard()
	shipSize := 3

	randomiser := rand.New(rand.NewSource(42))

	success := board.PlaceShip(shipSize, randomiser)
	if !success {
		t.Fatal("Failed to place ship")
	}

	ShipCellMarkers := []struct{ row, col int }{}
	for r, row := range board.Grid {
		for c, cell := range row {
			if cell == ShipCellMarker {
				ShipCellMarkers = append(ShipCellMarkers, struct{ row, col int }{r, c})
			}
		}
	}

	if len(ShipCellMarkers) != shipSize {
		t.Fatalf("Expected %d ship cells, got %d", shipSize, len(ShipCellMarkers))
	}

	for i, pos := range ShipCellMarkers {
		r, c := pos.row, pos.col

		if i == len(ShipCellMarkers)-1 {
			if board.Grid[r][c] != ShipCellMarker {
				t.Fatalf("Last ship cell at row %d, col %d is not correctly placed", r, c)
			}
			continue
		}

		if r+1 < BoardSize && board.Grid[r+1][c] == ShipCellMarker {
			continue
		}
		if c+1 < BoardSize && board.Grid[r][c+1] == ShipCellMarker {
			continue
		}

		t.Fatalf("Ship cells are not consecutive at row %d, col %d", r, c)
	}
}
