package main

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

var ready bool

func init() {
	setup()
}

func setup() {
	ready = true
}

func TestSetup(t *testing.T) {
	setup()
	if ready != true {
		t.Error("Expected ready to be true, but it was false")
	}
}
