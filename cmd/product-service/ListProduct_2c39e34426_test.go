package product_service

import (
	"context"
	"fmt"
	"testing"

	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/aristat/golang-example-app/generated/resources/proto/health_checks"
	"github.com/aristat/golang-example-app/generated/resources/proto/products"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type mockHealthChecksClient struct {
	health_checks.UnimplementedHealthChecksServer
}

func (m *mockHealthChecksClient) IsAlive(context.Context, *products.Empty) (*health_checks.IsAliveOut, error) {
	return &health_checks.IsAliveOut{Status: health_checks.IsAliveOut_OK}, nil
}

func TestListProduct_2c39e34426(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	health_checks.RegisterHealthChecksServer(s, &mockHealthChecksClient{})
	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Printf("Server exited with error: %v", err)
		}
	}()

	cfg := &config.Config{HealthCheckUrl: "bufnet"}
	logger := logger.New(cfg)

	server := &server{cfg: cfg, logger: logger}

	ctx := context.Background()
	in := &products.ListProductIn{}

	out, err := server.ListProduct(ctx, in)

	assert.NoError(t, err)
	assert.NotNil(t, out)
	assert.Equal(t, products.ListProductOut_OK, out.Status)
	assert.Len(t, out.Products, 2)
	assert.Equal(t, int32(1), out.Products[0].Id)
	assert.Equal(t, "first_product", out.Products[0].Name)
	assert.Equal(t, int32(2), out.Products[1].Id)
	assert.Equal(t, "second_product", out.Products[1].Name)
}
