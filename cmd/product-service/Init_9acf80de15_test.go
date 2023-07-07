package main

import (
    _ "github.com/aristat/golang-example-app/app/common"
    _ "github.com/aristat/golang-example-app/app/config"
    _ "github.com/aristat/golang-example-app/app/logger"
    _ "github.com/aristat/golang-example-app/generated/resources/proto/health_checks"
    _ "github.com/aristat/golang-example-app/generated/resources/proto/products"
    _ "github.com/nats-io/nats.go"
    _ "github.com/nats-io/stan.go"
    _ "github.com/spf13/cobra"
    _ "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
    _ "go.opentelemetry.io/otel"
    _ "go.opentelemetry.io/otel/propagation"
    _ "google.golang.org/grpc"
    _ "google.golang.org/grpc/credentials/insecure"
    _ "google.golang.org/protobuf/types/known/emptypb"
)

func main() {
    // your code goes here
}
