package util

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
)

// Add adds two number
func Add(a int, b int) int {
	grpclog.Info(fmt.Sprintf("grpclog.Info a: %d", a))
	grpclog.Error(fmt.Sprintf("grpclog.Info b: %d", b))
	return a + b
}
