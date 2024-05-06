package cpu

import (
	"bytes"
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func CollectData(warmUpSeconds, seconds int, ch chan CPU) {

	data := make([]CPU, warmUpSeconds)

	for i := 0; i < warmUpSeconds; i++ {
		time.Sleep(1 * time.Second)
		data[i] = parse(run())
	}
	ch <- calculateAverage(data)

	for {
		data = data[seconds:]
		for i := 0; i < seconds; i++ {
			data = append(data, parse(run()))
			time.Sleep(1 * time.Second)
		}
		fmt.Println(data)
		ch <- calculateAverage(data)
	}
}

func parse(str string) CPU {
	rows := strings.SplitN(str, "\n", 4)
	firstLine := strings.Split(rows[0], " ")
	min, _ := strconv.ParseFloat(strings.TrimRight(firstLine[12], ","), 64)
	five, _ := strconv.ParseFloat(strings.TrimRight(firstLine[13], ","), 64)
	fifteen, _ := strconv.ParseFloat(strings.TrimRight(firstLine[14], ","), 64)

	return CPU{
		Min:     float64(min),
		Five:    float64(five),
		Fifteen: float64(fifteen),
	}
}

func run() string {
	cmd := exec.Command("/usr/bin/top", "-n1", "-b")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return out.String()
}

func calculateAverage(cpus []CPU) CPU {
	var minSum float64
	var fiveSum float64
	var fifteenSum float64
	for _, cpu := range cpus {
		minSum += cpu.Min
		fiveSum += cpu.Five
		fifteenSum += cpu.Fifteen
	}

	dataLen := float64(len(cpus))

	return CPU{
		Min:     math.Round((minSum/dataLen)*100) / 100,
		Five:    math.Round((fiveSum/dataLen)*100) / 100,
		Fifteen: math.Round((fifteenSum/dataLen)*100) / 100,
	}
}
