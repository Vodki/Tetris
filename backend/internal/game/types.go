package game

type GameMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type GameClient interface {
	SendUpdate(grid [][]int)
	SendGameOver()
}
