package cpu

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func collectData(warmUpSeconds, seconds int, ch chan float32) {
	data := make([]int, warmUpSeconds)
	for i := 0; i < warmUpSeconds; i++ {
		time.Sleep(1 * time.Second)
		data[i] = 5
		cmd := exec.Command("/usr/bin/top")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		out.String()
	}
	ch <- calculateAverage(data)

	for {
		data = data[seconds:]
		for i := 0; i < seconds; i++ {
			data = append(data, 7)
			time.Sleep(1 * time.Second)
		}
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
