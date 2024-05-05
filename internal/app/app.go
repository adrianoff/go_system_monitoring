package app

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/adrianoff/go-system-monitoring/internal/logger"
)

type App struct {
	logger logger.Logger
}

type AppInterface interface {
	GetMainChannel() <-chan float32
	startMainLoop(ch chan float32)
}

func New(logger logger.Logger) AppInterface {
	return &App{
		logger: logger,
	}
}

func (app *App) GetMainChannel() <-chan float32 {
	ch := make(chan float32)

	go app.startMainLoop(ch)

	return ch
}

func (app *App) startMainLoop(ch chan float32) {
	app.logger.Info("Main Loop Starts")
	var N, M time.Duration
	M = 15
	N = 5

	averageChan := make(chan float32)
	diskChan := make(chan float32)

	go app.collectLoadAverage(int(M), int(N), averageChan)
	go app.collectDiskInfo(int(M), int(N), diskChan)

	for {
		app.logger.Info("Main Loop Read averageChan")

		averageVal := <-averageChan
		fmt.Println(averageVal)
		averageVal = averageVal + <-diskChan
		fmt.Println(averageVal)

		app.logger.Info("Main Loop Main Chan Released")
		ch <- averageVal
	}
}

func (app *App) collectLoadAverage(warmUpSeconds, seconds int, ch chan float32) {
	cmd := exec.Command("/usr/bin/top", "n1")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	str := out.String()
	fmt.Println(str)
	rows := strings.SplitN(str, "\n", 4)

	fmt.Println(rows)

	data := make([]int, warmUpSeconds)
	for i := 0; i < warmUpSeconds; i++ {
		time.Sleep(1 * time.Second)
		data[i] = 5
	}
	ch <- calculateAverage(data)

	for {
		app.logger.Info("collectLoadAverage Starts")
		data = data[seconds:]
		for i := 0; i < seconds; i++ {
			data = append(data, 7)
			time.Sleep(1 * time.Second)
		}
		app.logger.Info("collectLoadAverage Chan Released")
		fmt.Println(data)
		ch <- calculateAverage(data)
	}
}

func (app *App) collectDiskInfo(warmUpSeconds, seconds int, ch chan float32) {
	data := make([]int, warmUpSeconds)
	for i := 0; i < warmUpSeconds; i++ {
		time.Sleep(1 * time.Second)
		data[i] = 10
	}
	ch <- calculateAverage(data)

	for {
		app.logger.Info("collectLoadAverage Starts")
		data = data[seconds:]
		for i := 0; i < seconds; i++ {
			data = append(data, 15)
			time.Sleep(1 * time.Second)
		}
		app.logger.Info("collectLoadAverage Chan Released")
		fmt.Println(data)
		ch <- calculateAverage(data)
	}
}

func calculateAverage(values []int) float32 {
	var total int = 0
	for _, value := range values {
		total += value
	}

	return float32(total) / float32(len(values))
}
