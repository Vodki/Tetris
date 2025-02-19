package game

import (
	"github.com/thoas/go-funk"
)

type Board struct {
	Grid [][]int
}

func NewBoard() *Board {
	grid := make([][]int, 20)
	for i := range grid {
		grid[i] = make([]int, 10)
	}
	return &Board{Grid: grid}
}

func (b *Board) GridCopy() [][]int {
	copy := make([][]int, len(b.Grid))

	for i := range b.Grid {
		copy[i] = make([]int, len(b.Grid[i]))
		for j := range b.Grid[i] {
			copy[i][j] = b.Grid[i][j]
		}
	}

	return copy
}

func (b *Board) fillBoard(X, Y int) {
	if X >= 0 && Y >= 0 {
		b.Grid[Y][X] += 1
	}
}

func (b *Board) lineIsFull(Y int) bool {
	for _, n := range b.Grid[Y] {
		if n != 1 {
			return false
		}
	}
	return true
}

func (b *Board) clearFullLines() int {
	total := 0
	filledLines := []int{}
	for i := len(b.Grid) - 1; i >= 0; i-- {
		if b.lineIsFull(i) {
			filledLines = append(filledLines, i)
			total++
		} else if len(filledLines) != 0 {
			b.clearLines(filledLines)
			i = len(b.Grid) - 1
			filledLines = []int{}
		}
	}
	return total
}

func (b *Board) clearLines(lines []int) {
	lenLines := len(lines)
	max := funk.MaxInt(lines)
	for i := max; i > 0; i-- {
		if (i - lenLines) < 0 {
			copy(b.Grid[i], make([]int, len(b.Grid)))
		} else {
			copy(b.Grid[i], b.Grid[i-lenLines])
		}
	}
}

func (b *Board) isInvalid() bool {
	for i := range b.Grid {
		for _, n := range b.Grid[i] {
			if n != 0 && n != 1 {
				return true
			}
		}
	}
	return false
}
