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
