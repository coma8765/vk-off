package messenger

type HandleMessageFunc func(message Message) (err error)

type Messenger interface {
	NewMessenger() (*Message, error)
	SendMessage(ref MessageRef)
	SetCallback(messageFunc HandleMessageFunc)
}
