package controller

import (
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"os"
	"time"
)

/*
param: c1 - channel for the cpu info
param: c2 - channel for the ram info
function: this function gets CPU & RAM info every second and adds them to each channel to be collected in the future
*/
func GetInfo(c1 chan float64, c2 chan float64) {
	// Get CPU & RAM info every second and add them to each channel
	for {
		time.Sleep(1 * time.Second)

		// Get CPU info
		before, err := cpu.Get()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		time.Sleep(time.Duration(1) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		total := float64(after.Total - before.Total)
		c1 <- 100 - float64(after.Idle-before.Idle)/total*100

		// Get RAM info
		mem, err := memory.Get()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		c2 <- float64(mem.Used) / 1000000000
	}
}
