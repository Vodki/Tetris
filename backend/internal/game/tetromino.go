package game

import (
	"time"

	"math/rand"
)

type Point struct {
	X, Y int
}

type Tetromino struct {
	Rotations     [][]Point
	RotationIndex int
	Position      Point
	ID            string
	Color         int
}

func (t *Tetromino) GetCurrentShape() []Point {
	return t.Rotations[t.RotationIndex%len(t.Rotations)]
}

func (t *Tetromino) Rotate() {
	t.RotationIndex = (t.RotationIndex + 1) % len(t.Rotations)
}

var localRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// NewRandomTetromino returns a random copy of a tetromino.
func NewRandomTetromino() *Tetromino {
	choice := AllTetrominoes[localRand.Intn(len(AllTetrominoes))]
	t := choice // copy the chosen tetromino
	t.RotationIndex = len(t.Rotations) * 1000
	return &t
}

var AllTetrominoes = []Tetromino{
	TetrominoI,
	TetrominoJ,
	TetrominoL,
	TetrominoO,
	TetrominoS,
	TetrominoT,
	TetrominoZ,
}

var (
	// I Tetromino
	TetrominoI = Tetromino{
		ID: "I",
		Rotations: [][]Point{
			// Vertical
			{{0, -1}, {0, 0}, {0, 1}, {0, 2}},
			// Horizontal
			{{-1, 0}, {0, 0}, {1, 0}, {2, 0}},
		},
		RotationIndex: 0,
		Position:      Point{X: 4, Y: 1},
		Color:         1,
	}

	// J Tetromino
	TetrominoJ = Tetromino{
		ID: "J",
		Rotations: [][]Point{
			// Rotation 0
			{{-1, -1}, {-1, 0}, {0, 0}, {1, 0}},
			// Rotation 1
			{{0, -1}, {0, 0}, {0, 1}, {1, -1}},
			// Rotation 2
			{{-1, 0}, {0, 0}, {1, 0}, {1, 1}},
			// Rotation 3
			{{-1, 1}, {0, -1}, {0, 0}, {0, 1}},
		},
		RotationIndex: 0,
		Position:      Point{X: 4, Y: 1},
		Color:         2,
	}

	// L Tetromino
	TetrominoL = Tetromino{
		ID: "L",
		Rotations: [][]Point{
			// Rotation 0
			{{-1, 0}, {0, 0}, {1, 0}, {1, -1}},
			// Rotation 1
			{{0, -1}, {0, 0}, {0, 1}, {1, 1}},
			// Rotation 2
			{{-1, 1}, {-1, 0}, {0, 0}, {1, 0}},
			// Rotation 3
			{{-1, -1}, {0, -1}, {0, 0}, {0, 1}},
		},
		RotationIndex: 0,
		Position:      Point{X: 4, Y: 1},
		Color:         3,
	}

	// O Tetromino (Square) – Only one rotation needed since it's symmetric.
	TetrominoO = Tetromino{
		ID: "O",
		Rotations: [][]Point{
			{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
		},
		RotationIndex: 0,
		Position:      Point{X: 4, Y: 0},
		Color:         4,
	}

	// S Tetromino
	TetrominoS = Tetromino{
		ID: "S",
		Rotations: [][]Point{
			// Rotation 0
			{{0, 0}, {1, 0}, {-1, 1}, {0, 1}},
			// Rotation 1
			{{0, -1}, {0, 0}, {1, 0}, {1, 1}},
		},
		RotationIndex: 0,
		Position:      Point{X: 4, Y: 0},
		Color:         5,
	}

	// T Tetromino
	TetrominoT = Tetromino{
		ID: "T",
		Rotations: [][]Point{
			// Rotation 0
			{{-1, 0}, {0, 0}, {1, 0}, {0, 1}},
			// Rotation 1
			{{0, -1}, {0, 0}, {1, 0}, {0, 1}},
			// Rotation 2
			{{0, -1}, {-1, 0}, {0, 0}, {1, 0}},
			// Rotation 3
			{{0, -1}, {-1, 0}, {0, 0}, {0, 1}},
		},
		RotationIndex: 0,
		Position:      Point{X: 4, Y: 0},
		Color:         6,
	}

	// Z Tetromino
	TetrominoZ = Tetromino{
		ID: "Z",
		Rotations: [][]Point{
			// Rotation 0
			{{-1, 0}, {0, 0}, {0, 1}, {1, 1}},
			// Rotation 1
			{{1, -1}, {0, 0}, {1, 0}, {0, 1}},
		},
		RotationIndex: 0,
		Position:      Point{X: 4, Y: 0},
		Color:         7,
	}
)
