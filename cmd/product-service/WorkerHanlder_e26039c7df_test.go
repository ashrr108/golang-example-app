package product_service_test

import (
	"testing"

	"github.com/aristat/golang-example-app/app/common"
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/aristat/golang-example-app/generated/resources/proto/health_checks"
	"github.com/aristat/golang-example-app/generated/resources/proto/products"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type natsService struct {
	logger logger.Logger
}

func (s *natsService) workerHandler(m *stan.Msg) {
	s.logger.Info("[NATS] Received a message: %s\n", logger.Args(string(m.Data)))
	m.Ack()
}

type MockMsg struct {
	data []byte
}

func (m *MockMsg) Data() []byte {
	return m.data
}

func (m *MockMsg) Ack() error {
	return nil
}

type MockLogger struct{}

func (l *MockLogger) Info(format string, args ...interface{}) {
}

func TestWorkerHanlder_e26039c7df(t *testing.T) {
	mockLogger := &MockLogger{}
	mockMsg := &MockMsg{data: []byte("Test Message")}

	service := &natsService{logger: mockLogger}

	t.Run("Test Successful Acknowledgement", func(t *testing.T) {
		service.workerHandler(mockMsg)
	})

	t.Run("Test Failed Acknowledgement", func(t *testing.T) {
		service.workerHandler(mockMsg)
	})
}
