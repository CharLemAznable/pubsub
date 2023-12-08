package pubsub_test

import (
	"github.com/CharLemAznable/pubsub"
	"testing"
)

func TestPubSub(t *testing.T) {
	ach := make(chan *MsgA, 1)
	bch := make(chan *MsgB, 1)

	aSub := pubsub.SubscribeFunc[*MsgA](func(msg *MsgA) {
		ach <- msg
	})
	bSub := pubsub.SubscribeChan[*MsgB](bch)
	subs := pubsub.Subscribers{aSub, bSub}

	pubsub.Subscribe("test_topic", subs)

	go func() {
		pubsub.Publish("test_topic", &MsgA{Content: "MSG_A"})
		pubsub.Publish("test_topic", &MsgB{Message: "MSG_B"})
	}()

	select {
	case msgA := <-ach:
		if msgA.Content != "MSG_A" {
			t.Errorf("Expected msgA.Content is 'MSG_A', but got '%s'", msgA.Content)
		}
	}
	select {
	case msgB := <-bch:
		if msgB.Message != "MSG_B" {
			t.Errorf("Expected msgB.Message is 'MSG_B', but got '%s'", msgB.Message)
		}
	}

	pubsub.Unsubscribe("test_topic", subs)
	pubsub.Unsubscribe("test_topic", bSub)
	pubsub.Unsubscribe("test_topic", aSub)
}

type MsgA struct {
	Content string
}

type MsgB struct {
	Message string
}
