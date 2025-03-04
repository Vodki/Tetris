package msg

import "fmt"

type WSMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func NewMessage(Type, Data string) *WSMessage {
	return &WSMessage{
		Type: Type,
		Data: Data,
	}
}

func (msg *WSMessage) ToString() string {
	return fmt.Sprintf("{Type:%s Data:%s}", msg.Type, msg.Data)
}

type GameMessage struct {
	Type   string `json:"type"`
	Grid   string `json:"data"`
	Score  int    `json:"score"`
	GameOn bool   `json:"gameOn"`
	Level  int    `json:"level"`
}

func NewGameMessage(Type, Grid string, Score, Level int, GameOn bool) *GameMessage {
	return &GameMessage{
		Type:   Type,
		Grid:   Grid,
		Score:  Score,
		GameOn: GameOn,
		Level:  Level,
	}
}
