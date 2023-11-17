package messenger

import "time"

type MessageRef struct {
	ChatId string
	Text   string
}

type Message struct {
	MessageRef

	ID       string
	IssuedAt time.Time
	SenderID string
}
