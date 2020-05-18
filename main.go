package main

import (
	"time"
	"ubiwhere/controller"
	"ubiwhere/model"
)
func init() {
	controller.OpenDatabase()
}
func main() {
	// CLI start
	//cmd.Execute()

	// Create external device simulator channel (capacity: 3)
	simu := make(chan model.SimuData, 3)

	go controller.GetCpuAndRamInfo()
	go controller.Simulator(simu)
	go controller.CollectDataSample(simu)

	for {
		time.Sleep(5 * time.Second)

		//fmt.Printf("CPU: %.2f%% | RAM: %.2fGb\n", <-cpuc, <-ramc)
	}
}
