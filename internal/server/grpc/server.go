package internalgrpc

import (
	"net"
	"time"

	"github.com/adrianoff/go-system-monitoring/internal/app"
	"github.com/adrianoff/go-system-monitoring/internal/logger"
	"github.com/adrianoff/go-system-monitoring/internal/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedMonitoringServiceServer
	listener net.Listener
	server   *grpc.Server
	app      app.AppInterface
	logger   logger.Logger
	address  string
}

func NewServer(logger logger.Logger, app app.AppInterface, address string) *Server {
	return &Server{
		logger:  logger,
		app:     app,
		address: address,
	}
}

func (s *Server) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", s.address)
	if err != nil {
		s.logger.Error(err)
		return nil
	}
	s.server = grpc.NewServer()
	pb.RegisterMonitoringServiceServer(s.server, s)
	s.logger.Info("starting grpc server on ", s.listener.Addr().String())
	if err := s.server.Serve(s.listener); err != nil {
		s.logger.Error(err)
		s.listener.Close()
	}
	return nil
}

func (s *Server) Stop() {
	s.logger.Info("server grpc is stopping...")
	s.server.Stop()
	s.listener.Close()
}

func (s *Server) StreamSnapshots(request *pb.SnapshotRequest, server pb.MonitoringService_StreamSnapshotsServer) error {

	time.Sleep(2 * time.Second)
	snapshot := pb.Snapshot{
		LoadAverage: &pb.LoadAverage{
			Min:     10,
			Five:    20,
			Fifteen: 30,
		},
	}

	server.Send(&snapshot)

	time.Sleep(2 * time.Second)
	server.Send(&snapshot)

	time.Sleep(2 * time.Second)
	server.Send(&snapshot)

	return nil
}
