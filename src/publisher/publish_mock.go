package publisher

type mockPublisher struct{}

func (p *mockPublisher) Publish(message string) error {
	return nil
}
