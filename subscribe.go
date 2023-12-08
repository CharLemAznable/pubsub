package pubsub

type Subscriber interface {
	Subscribe(any)
}

type SubscribeFunc[T any] func(T)

func (f SubscribeFunc[T]) Subscribe(msg any) {
	if message, ok := msg.(T); ok {
		f(message)
	}
}

type SubscribeChan[T any] chan T

func (ch SubscribeChan[T]) Subscribe(msg any) {
	if message, ok := msg.(T); ok {
		ch <- message
	}
}

type Subscribers []Subscriber

func (s Subscribers) Subscribe(msg any) {
	for _, sub := range s {
		go sub.Subscribe(msg)
	}
}
