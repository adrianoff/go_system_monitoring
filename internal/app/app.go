package app

import (
	"time"

	"github.com/adrianoff/go-system-monitoring/internal/logger"
)

type App struct {
	logger logger.Logger
}

type AppInterface interface {
	GetMainChannel() <-chan int
	startMainLoop(ch chan int)
}

func New(logger logger.Logger) AppInterface {
	return &App{
		logger: logger,
	}
}

func (app *App) GetMainChannel() <-chan int {
	ch := make(chan int)

	go app.startMainLoop(ch)

	return ch
}

func (app *App) startMainLoop(ch chan int) {
	for {
		time.Sleep(1 * time.Second)
		ch <- 1
	}
}
