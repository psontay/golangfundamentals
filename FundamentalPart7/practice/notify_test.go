package practice

import "testing"

type MockNotifier struct {
	SentMessage string
}

func (MockNotifier *MockNotifier) Send(message string) error {
	MockNotifier.SentMessage = message
	return nil
}

func TestNotifyUser(t *testing.T) {
	mock := &MockNotifier{}
	err := NotifyUser(mock, "Hello World")
	if mock.SentMessage != "Hello World" || err != nil {
		t.Errorf("Sent message should be \"Hello World\"")
	}
}
