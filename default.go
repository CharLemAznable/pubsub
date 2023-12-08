package pubsub

var globalPubSub = NewPubSub()

func Subscribe(topic string, subscriber Subscriber) {
	globalPubSub.Subscribe(topic, subscriber)
}

func Unsubscribe(topic string, subscriber Subscriber) {
	globalPubSub.Unsubscribe(topic, subscriber)
}

func Publish(topic string, msg any) {
	globalPubSub.Publish(topic, msg)
}
