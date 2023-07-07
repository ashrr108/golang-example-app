package product_service

import (
	"fmt"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/nats-io/stan.go"
	"testing"
)

type mockNatsService struct {
	logger logger.Logger
}

func (s *mockNatsService) workerHandler(m *stan.Msg) {
	s.logger.Info("[NATS] Received a message: %s\n", fmt.Sprintf("%s", string(m.Data)))
	m.Ack()
}

// Mocking logger.Logger for testing
type mockLogger struct {
	loggedMessage string
}

func (m *mockLogger) Info(format string, args ...interface{}) {
	m.loggedMessage = format
}

func (m *mockLogger) Alert(format string, args ...interface{}) {
	// Implement Alert method
}

// Mocking stan.Msg for testing
type mockMsg struct {
	data []byte
	ack  bool
}

func (m *mockMsg) Data() []byte {
	return m.data
}

func (m *mockMsg) Ack() error {
	m.ack = true
	return nil
}

func TestWorkerHandler_e26039c7df(t *testing.T) {
	t.Run("successfully logs and acknowledges message", func(t *testing.T) {
		mockLog := &mockLogger{}
		mockMessage := &mockMsg{data: []byte("test message")}
		service := &mockNatsService{logger: mockLog}

		service.workerHandler(mockMessage)

		if mockLog.loggedMessage != "[NATS] Received a message: %s\n" {
			t.Error("Expected log message was not present")
		}

		if !mockMessage.ack {
			t.Error("Message was not acknowledged")
		}
	})

	t.Run("handles empty message", func(t *testing.T) {
		mockLog := &mockLogger{}
		mockMessage := &mockMsg{data: []byte("")}
		service := &mockNatsService{logger: mockLog}

		service.workerHandler(mockMessage)

		if mockLog.loggedMessage != "[NATS] Received a message: %s\n" {
			t.Error("Expected log message was not present")
		}

		if !mockMessage.ack {
			t.Error("Message was not acknowledged")
		}
	})
}
