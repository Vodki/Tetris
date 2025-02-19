package game

import (
	"Tetris/internal/msg"
	"encoding/json"
	"log"
	"time"
)

const ROWS = 20
const COLS = 10

const (
	Empty int = iota
	// You could have more values or even a struct if you need color information.
)

type Engine struct {
	Board        *Board
	Current      *Tetromino
	Next         *Tetromino
	Ticker       *time.Ticker
	GameOver     bool
	Score        int
	level        int
	clearedLines int
	clientChan   chan msg.WSMessage
	CommandChan  chan msg.WSMessage
}

func NewEngine(clientChan chan msg.WSMessage) *Engine {
	return &Engine{
		Board:        NewBoard(),
		Current:      NewRandomTetromino(),
		Next:         NewRandomTetromino(),
		GameOver:     false,
		Score:        0,
		level:        1,
		clearedLines: 0,
		clientChan:   clientChan,
		CommandChan:  make(chan msg.WSMessage),
	}
}

func (g *Engine) CalculateScore(lines int) int {
	switch lines {
	case 1:
		return 100 * g.level
	case 2:
		return 300 * g.level
	case 3:
		return 500 * g.level
	case 4:
		return 800 * g.level
	default:
		return 0
	}
}

func (g *Engine) reset() {
	g.Current = NewRandomTetromino()
	g.Next = NewRandomTetromino()
	g.Board = NewBoard()
	g.GameOver = false
	g.Score = 0
}

func (g *Engine) Start() {
	for {
		start := false
		for !start {
			msg := <-g.CommandChan
			if msg.Data == "start" {
				start = true
			}
		}
		g.Ticker = time.NewTicker(500 * time.Millisecond)
		for !g.GameOver {
			select {
			case <-g.Ticker.C:
				g.tick()
			case cmd := <-g.CommandChan:
				g.handleCommand(cmd)
			}
		}
		g.Ticker.Stop()
		g.reset()
		log.Print("GAME OVER")
	}

}

func (g *Engine) sendGrid() {
	tmpGrid := g.Board.GridCopy()
	t := g.Current

	for _, p := range t.GetCurrentShape() {
		x := t.Position.X + p.X
		y := t.Position.Y + p.Y
		if x >= 0 && y >= 0 {
			tmpGrid[y][x] = 1
		}
	}
	gridJson, _ := json.Marshal(tmpGrid)
	msg := msg.NewMessage("GameUpdate", string(gridJson))
	g.clientChan <- *msg
}

func (g *Engine) handleCommand(msg msg.WSMessage) {
	switch msg.Data {
	case "Rotate":
		g.rotateCurrentTetromino()
	case "MoveLeft":
		g.moveLeft()
	case "MoveRight":
		g.moveRight()
	case "MoveDown":
		g.moveDown()
	case "HardDrop":
		g.hardDrop()
	}
}

// tick handles each game update.
func (g *Engine) tick() {
	if g.canMoveDown() {
		g.Current.Position.Y++
	} else {
		g.lockCurrentTetromino()
		n := g.Board.clearFullLines()
		g.Score += g.CalculateScore(n)
		if (n+g.clearedLines)/10 > g.clearedLines/10 {
			g.level++
		}
		g.spawnNewTetromino()
		if !g.isValidPosition(g.Current) || g.Board.isInvalid() {
			g.GameOver = true
		}
	}
	g.sendGrid()
}

func (g *Engine) spawnNewTetromino() {
	g.Current = g.Next
	g.Next = NewRandomTetromino()
}

func (g *Engine) lockCurrentTetromino() {
	t := g.Current
	for _, p := range t.GetCurrentShape() {
		x := t.Position.X + p.X
		y := t.Position.Y + p.Y
		g.Board.fillBoard(x, y)
	}
}

func (g *Engine) isValidPosition(t *Tetromino) bool {
	for _, p := range t.GetCurrentShape() {
		// Calculate the absolute position on the board.
		x := t.Position.X + p.X
		y := t.Position.Y + p.Y

		// Check boundaries.
		if x < 0 || x >= COLS || y < 0 || y >= ROWS {
			return false
		}
		// Check if the cell is already occupied.
		if g.Board.Grid[y][x] != 0 {
			return false
		}
	}
	return true
}

func (g *Engine) canMoveDown() bool {
	// Create a copy of current tetromino with Y+1 and check if it's valid.
	nextPosition := g.Current.Position
	nextPosition.Y++
	temp := *g.Current
	temp.Position = nextPosition
	return g.isValidPosition(&temp)
}

func (g *Engine) rotateCurrentTetromino() {
	// Save the current rotation index.
	originalRotation := g.Current.RotationIndex
	// Rotate.
	g.Current.Rotate()

	// If the new rotation is invalid, revert.
	if !g.isValidPosition(g.Current) {
		g.Current.RotationIndex = originalRotation
	} else {
		g.sendGrid()
	}
}

func (g *Engine) hardDrop() {
	for g.canMoveDown() {
		g.Current.Position.Y++
	}
	g.lockCurrentTetromino()
	n := g.Board.clearFullLines()
	g.Score += g.CalculateScore(n)
	if (n+g.clearedLines)/10 > g.clearedLines/10 {
		g.level++
	}
	g.spawnNewTetromino()
	if !g.isValidPosition(g.Current) || g.Board.isInvalid() {
		g.GameOver = true
	}
	g.sendGrid()
}

func (g *Engine) moveLeft() {
	g.Current.Position.X--
	if !g.isValidPosition(g.Current) {
		g.Current.Position.X++
	} else {
		g.sendGrid()
	}
}

func (g *Engine) moveRight() {
	g.Current.Position.X++
	if !g.isValidPosition(g.Current) {
		g.Current.Position.X--
	} else {
		g.sendGrid()
	}
}

func (g *Engine) moveDown() {
	if g.canMoveDown() {
		g.Current.Position.Y++
	} else {
		g.lockCurrentTetromino()
		n := g.Board.clearFullLines()
		g.Score += g.CalculateScore(n)
		if (n+g.clearedLines)/10 > g.clearedLines/10 {
			g.level++
		}

		g.spawnNewTetromino()
		if !g.isValidPosition(g.Current) || g.Board.isInvalid() {
			g.GameOver = true
		}
	}
	g.sendGrid()
}
