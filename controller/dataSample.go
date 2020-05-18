package controller

import (
	"math/rand"
	"time"
	"ubiwhere/model"
)

/*
param: simu - channel for the simulator data samples
function: Generates 4 random variables as a data sample every second (simulating an external device)
 		  and adds it to the channel
*/
func Simulator(simu chan model.SimuData) {
	var data model.SimuData

	for {
		data.V1 = rand.Intn(100)
		data.V2 = rand.Intn(100)
		data.V3 = rand.Intn(100)
		data.V4 = rand.Intn(100)

		// Add data sample to the channel
		simu <- data

		// Sleep for 1s
		time.Sleep(1 * time.Second)
	}
}

func CollectDataSample(simu chan model.SimuData) {
	for {
		// Collect data sample from the channel
		dataSample := <-simu

		// Insert data sample into the database
		Db.Create(&dataSample)

		time.Sleep(1 * time.Second)
	}
}
