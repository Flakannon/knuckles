package publisher

import "testing"

func TestPublishMessage(t *testing.T) {
	t.Run("should publish message", func(t *testing.T) {
		p := &mockPublisher{}
		err := PublishMessage(p, "test message")
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
	})
}
