package game

type CommandType int

const (
	MoveLeft CommandType = iota
	MoveRight
	Rotate
	StartGame
)

type Command struct {
	Type CommandType
}
