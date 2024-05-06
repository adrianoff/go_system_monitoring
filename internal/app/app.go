package app

import (
	"time"

	"github.com/adrianoff/go-system-monitoring/internal/app/cpu"
	"github.com/adrianoff/go-system-monitoring/internal/logger"
)

type App struct {
	logger logger.Logger
}

type AppInterface interface {
	GetMainChannel() <-chan MonitoringInfo
	startMainLoop(ch chan MonitoringInfo)
}

func New(logger logger.Logger) AppInterface {
	return &App{
		logger: logger,
	}
}

func (app *App) GetMainChannel() <-chan MonitoringInfo {
	ch := make(chan MonitoringInfo)

	go app.startMainLoop(ch)

	return ch
}

func (app *App) startMainLoop(ch chan MonitoringInfo) {
	app.logger.Info("Main Loop Starts")
	var N, M time.Duration
	M = 15
	N = 5

	cpuChan := make(chan cpu.CPU)

	go cpu.CollectData(int(M), int(N), cpuChan)

	for {
		app.logger.Info("Main Loop Read averageChan")

		monitoringInfo := MonitoringInfo{}
		monitoringInfo.Cpu = <-cpuChan

		app.logger.Info("Main Loop Main Chan Released")
		ch <- monitoringInfo
	}
}
