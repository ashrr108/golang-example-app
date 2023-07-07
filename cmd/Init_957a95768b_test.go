package mycmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/aristat/golang-example-app/pkg/jwt"

	"github.com/aristat/golang-example-app/pkg/migrate"

	health_check_service "github.com/aristat/golang-example-app/pkg/health-check-service"
	product_service "github.com/aristat/golang-example-app/pkg/product-service"

	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/logger"

	"go.uber.org/automaxprocs/maxprocs"

	"github.com/aristat/golang-example-app/pkg/daemon"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ... rest of the code ...
