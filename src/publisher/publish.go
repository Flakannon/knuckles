package publisher

type IPublisher interface {
	Publish(message string) error
}

func PublishMessage(p IPublisher, message string) error {
	return p.Publish(message)
}
