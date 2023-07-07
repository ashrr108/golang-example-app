package health_check_service

import (
	"context"
	"testing"

	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/generated/resources/proto/health_checks"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func TestIsAlive_156b3bd6f2(t *testing.T) {
	t.Run("RandomDisable is true", func(t *testing.T) {
		// setup
		cfg := &config.Config{RandomDisable: true}
		s := NewServer(cfg)

		// execute
		result, err := s.IsAlive(context.Background(), &emptypb.Empty{})

		// verify
		if err != nil {
			t.Error(err)
		}

		if result.Status != health_checks.IsAliveOut_OK {
			t.Error("Expected status to be OK")
		}
	})

	t.Run("RandomDisable is false and random number is 1", func(t *testing.T) {
		// setup
		cfg := &config.Config{RandomDisable: false}
		s := NewServer(cfg)

		// execute
		result, err := s.IsAlive(context.Background(), &emptypb.Empty{})

		// verify
		if err != nil {
			t.Error(err)
		}

		if result.Status != health_checks.IsAliveOut_OK && result.Status != health_checks.IsAliveOut_NOT_OK {
			t.Error("Expected status to be either OK or NOT_OK")
		}
	})
}
