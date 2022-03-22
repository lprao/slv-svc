package messaging

type Client interface {
	NewProducer([]string) (Producer, error)
	NewConsumerGroup() (ConsumerGroup, error)
}

type Producer interface {
	Publish()
}

type ConsumerGroup interface {
}
