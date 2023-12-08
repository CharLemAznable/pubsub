package pubsub

import (
	"github.com/CharLemAznable/ge"
	"sync"
)

type PubSub interface {
	Subscribe(string, Subscriber)
	Unsubscribe(string, Subscriber)
	Publish(string, any)
}

func NewPubSub() PubSub {
	return &hub{subscribers: make(map[string][]Subscriber)}
}

type hub struct {
	sync.RWMutex
	subscribers map[string][]Subscriber
}

func (h *hub) Subscribe(topic string, subscriber Subscriber) {
	h.Lock()
	defer h.Unlock()
	h.subscribers[topic] = ge.AppendElementUnique(h.subscribers[topic], subscriber)
}

func (h *hub) Unsubscribe(topic string, subscriber Subscriber) {
	h.Lock()
	defer h.Unlock()
	if subscribers, ok := h.subscribers[topic]; ok {
		h.subscribers[topic] = ge.RemoveElementByValue(subscribers, subscriber)
	}
}

func (h *hub) Publish(topic string, msg any) {
	h.RLock()
	defer h.RUnlock()
	if subscribers, ok := h.subscribers[topic]; ok {
		for _, sub := range subscribers {
			go sub.Subscribe(msg)
		}
	}
}
