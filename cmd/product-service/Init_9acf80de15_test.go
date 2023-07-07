package product_service

import (
	"testing"
	"github.com/aristat/golang-example-app/internal/domain/products"
	"github.com/aristat/golang-example-app/internal/config"
	"github.com/aristat/golang-example-app/internal/logger"
	"github.com/nats-io/stan.go"
	"fmt"
)

type mockLogger struct {
	AlertFunc func(string, ...logger.Option)
	InfoFunc  func(string, ...logger.Option)
}

func (m *mockLogger) Alert(msg string, opts ...logger.Option) {
	m.AlertFunc(msg, opts...)
}

func (m *mockLogger) Info(msg string, opts ...logger.Option) {
	m.InfoFunc(msg, opts...)
}

type mockMsg struct {
	Data []byte
}

func (m *mockMsg) GetData() []byte {
	return m.Data
}

func TestListProduct(t *testing.T) {
	emptyProduct := products.Empty
	cfg := config.Config{}
	log := logger.New(cfg)

	mockLog := &mockLogger{
		AlertFunc: func(msg string, opts ...logger.Option) {},
		InfoFunc:  func(msg string, opts ...logger.Option) {},
	}

	mockMessage := &mockMsg{
		Data: []byte("mock message"),
	}

	service := &Service{
		logger: mockLog,
	}
	service.workerHandler(mockMessage)

	fmt.Println("Test passed.")
}
