package main

import (
	"fmt"
	"time"
	"ubiwhere/cmd"
	"ubiwhere/controller"
	"ubiwhere/model"
)

func main() {
	// CLI start
	cmd.Execute()

	// Create CPU and RAM channels
	cpuc := make(chan float64)
	ramc := make(chan float64)
	simu := make(chan model.SimuData)

	go controller.GetInfo(cpuc, ramc)
	go controller.Simulator(simu)
	go controller.CollectDataSample(simu)

	for {
		time.Sleep(1 * time.Second)

		fmt.Printf("CPU: %.2f%% | RAM: %.2fGb\n", <-cpuc, <-ramc)
	}
}
