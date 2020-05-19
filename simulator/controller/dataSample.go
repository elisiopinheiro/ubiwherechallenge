package controller

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"ubiwhere/model"
)

/*
function: Generates 4 random variables as a data sample every second (simulating an external device)
 		  and adds it to the channel
param: simu - channel for the simulator data samples
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

/*
function: collects the data samples every second form the channel
params: channel to get the data samples from
*/
func CollectDataSample(simu chan model.SimuData) {
	for {
		// Collect data sample from the channel
		dataSample := <-simu

		// Insert data sample into the database
		Db.Create(&dataSample)

		time.Sleep(1 * time.Second)
	}
}

func GetNMetrics(c *gin.Context) {

	// Get the number of metrics to read
	n := c.Param("n")
	num, err := strconv.Atoi(n)
	if err != nil || num < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid number in url!"})
		return
	}

	if num > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Max 10 metrics!"})
		return
	}

	// Get metrics from database and store them in the array
	var metrics []model.SimuData
	Db.Order("id desc").Limit(num).Find(&metrics)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": metrics})
}
