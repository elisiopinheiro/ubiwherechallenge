package controller

import (
	"fmt"
	"math/rand"
	"time"
	"ubiwhere/model"
)

/*
param: simu - channel for the simulator data samples
function: generates 4 random variables as a data sample every second (simulating an external device)
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

		time.Sleep(1 * time.Second)
	}
}

func CollectDataSample(simu chan model.SimuData) {
	for {
		dataSample := <-simu
		fmt.Printf("V1: %d | V2: %d | V3: %d | V4: %d\n\n", dataSample.V1, dataSample.V2, dataSample.V3, dataSample.V4)

		time.Sleep(1 * time.Second)
	}
}
