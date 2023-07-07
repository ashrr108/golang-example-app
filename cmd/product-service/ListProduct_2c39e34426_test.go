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
	go_opentelemetry_io_contrib_instrumentation_google_golang_org_grpc_otelgrpc "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	go_opentelemetry_io_otel "go.opentelemetry.io/otel"
	go_opentelemetry_io_otel_propagation "go.opentelemetry.io/otel/propagation"
	google_golang_org_grpc "google.golang.org/grpc"
	google_golang_org_grpc_credentials_insecure "google.golang.org/grpc/credentials/insecure"
	google_golang_org_protobuf_types_known_emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func TestFetchModules(t *testing.T) {
	// your test logic here
}
