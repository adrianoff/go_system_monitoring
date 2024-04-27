package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrianoff/go-system-monitoring/internal/app"
	"github.com/adrianoff/go-system-monitoring/internal/logger"
	internalgrpc "github.com/adrianoff/go-system-monitoring/internal/server/grpc"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "../configs/config.yaml", "Path to configuration file")
}

func main() {
	config := NewConfig(configFile)
	logger := logger.New(config.Logger.Level, "Monitoring: ")
	monitoringApp := app.New(logger)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	serverGrps := internalgrpc.NewServer(logger, monitoringApp, config.GRPCServer.Address)
	go func() {
		<-ctx.Done()
		serverGrps.Stop()
	}()

	if err := serverGrps.Start(); err != nil {
		logger.Error("failed to start grpc server: " + err.Error())
		cancel()
		os.Exit(1)
	}
}
