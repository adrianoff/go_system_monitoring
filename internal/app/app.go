package app

import (
	"github.com/adrianoff/go-system-monitoring/internal/logger"
)

type App struct {
	logger logger.Logger
}

type AppInterface interface {
	GetAverage()
}

func New(logger logger.Logger) AppInterface {
	return &App{
		logger: logger,
	}
}

func (app *App) GetAverage() {
}
