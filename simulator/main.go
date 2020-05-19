package main

import (
	"ubiwhere/controller"
	"ubiwhere/model"
	"ubiwhere/rest"
)

func init() {
	// Open the database connection
	controller.OpenDatabase()
}

func main() {
	// Create a channel for external device simulator (capacity: 3)
	simu := make(chan model.SimuData, 3)

	// Goroutines
	go controller.GetCpuAndRamInfo()
	go controller.Simulator(simu)
	go controller.CollectDataSample(simu)

	r := rest.SetupRouter()
	_ = r.Run(":8080")

}
