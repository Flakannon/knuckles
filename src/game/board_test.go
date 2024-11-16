package game

import (
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

	success := board.PlaceShip(shipSize)
	if !success {
		t.Fatal("Failed to place ship")
	}

	// Count the number of ShipCellMarkers on the board
	ShipCellMarkers := map[int]int{}
	shipCount := 0
	for r, row := range board.Grid {
		for c, cell := range row {
			if cell == ShipCellMarker {
				shipCount++
				ShipCellMarkers[r] = c
			}
		}
	}

	if shipCount != shipSize {
		t.Fatalf("Expected %d ship cells, got %d", shipSize, shipCount)
	}

	// test ship is placed in a valid location (no overlap) and orientation is consecutive
	for r, c := range ShipCellMarkers {
		if r+1 < BoardSize && board.Grid[r+1][c] == ShipCellMarker {
			t.Fatalf("Ship cells are not consecutive")
		}
		if c+1 < BoardSize && board.Grid[r][c+1] == ShipCellMarker {
			t.Fatalf("Ship cells are not consecutive")
		}
	}
}
