package vksdk

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"time"
	"vk_off/internal/vk_off/adapters/messenger"
)

type VKMessenger struct {
	messenger.Messenger
	client *api.VK
}

func (m *VKMessenger) NewMessenger() VKMessenger {
	token := "vk1.a.3VjHdFwcevSfpshZHGSwMdaCeh4nnlTSywyq_L7PDtINYgwozBuqzFZJL0zxpkTvRZqqcP4Ok6pb0beP7CkZvnsI__Bmv1geLnGhB5ICFw4-m94F1AtiBYhhwgO_n9zc4lE8x4jjCIzY0JMpMZW7XdVq10MwYUO_RKmSaJSemyuNLQwE_rwTM1I2iEPWIy_UCdk8xCKVo3ao5qc-Fw1y0A" // use os.Getenv("TOKEN")
	vk := api.NewVK(token)

	return VKMessenger{
		client: vk,
	}
}

type VKMessageRef struct {
	messenger.MessageRef
	ChatId int
}

func (m *VKMessenger) SendMessage(ref *VKMessageRef) (*messenger.Message, error) {
	b := params.NewMessagesSendBuilder()
	b.Message(ref.Text)
	b.RandomID(0)
	b.PeerID(ref.ChatId)

	message, err := m.client.MessagesSend(b.Params)
	message
	if err != nil {
		return nil, err
	}

	return &messenger.Message{
		MessageRef: messenger.Message{
			MessageRef: messenger.MessageRef{},
			ID:         "",
			IssuedAt:   time.Time{},
			SenderID:   "",
		},
		ID:       "",
		IssuedAt: time.Time{},
		SenderID: "",
	}, nil
}

func usage() {
	vk := api.NewVK(token)
	messenger.Messenger.SendMessage()
}

//func main() {
//	token := "vk1.a.3VjHdFwcevSfpshZHGSwMdaCeh4nnlTSywyq_L7PDtINYgwozBuqzFZJL0zxpkTvRZqqcP4Ok6pb0beP7CkZvnsI__Bmv1geLnGhB5ICFw4-m94F1AtiBYhhwgO_n9zc4lE8x4jjCIzY0JMpMZW7XdVq10MwYUO_RKmSaJSemyuNLQwE_rwTM1I2iEPWIy_UCdk8xCKVo3ao5qc-Fw1y0A" // use os.Getenv("TOKEN")
//	//token := os.Getenv("VK_TOKEN") // use os.Getenv("TOKEN")
//	vk := api.NewVK(token)
//
//	// get information about the group
//	group, err := vk.GroupsGetByID(nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Initializing Long Poll
//	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// New message event
//	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
//		log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)
//
//		if obj.Message.Text == "ping" {
//			b := params.NewMessagesSendBuilder()
//			b.Message("pong")
//			b.RandomID(0)
//			b.PeerID(obj.Message.PeerID)
//
//			_, err := vk.MessagesSend(b.Params)
//			if err != nil {
//				log.Fatal(err)
//			}
//		}
//	})
//
//	// Run Bots Long Poll
//	log.Println("Start Long Poll")
//	if err := lp.Run(); err != nil {
//		log.Fatal(err)
//	}
//
//}
