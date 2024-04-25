package internalgrpc

import (
	"net"

	"github.com/adrianoff/go-system-monitoring/internal/logger"
	"google.golang.org/grpc"
)

type Server struct {
	listener net.Listener
	server *grpc.Server
	logger logger.Logger
}