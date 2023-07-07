package health_check_service

import (
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/aristat/golang-example-app/app/common"
	"github.com/aristat/golang-example-app/generated/resources/proto/health_checks"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// Assuming that the NewServer function and the config.Config type are defined in the same package
type Config struct {
	// Your config fields here
}

func NewServer(cfg Config) {
	// Your server initialization code here
}

func Init() {
	cfg := Config{}
	NewServer(cfg)
}

func init() {
	Init()
}
