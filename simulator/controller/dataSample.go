package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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

/*
function: gets the last N metrics and sends them through http repsonse
 */
func GetNMetrics(c *gin.Context) {

	// Get the number of metrics to read
	n := c.Param("n")
	num, err := strconv.Atoi(n)
	if err != nil || num < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid number in url!"})
		return
	}

	// Limit to 10 metrics
	if num > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Max 10 metrics!"})
		return
	}

	// Get metrics from database and store them in the array
	var metrics []model.SimuData
	Db.Order("id desc").Limit(num).Find(&metrics)

	// Send response
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": metrics})
}

/*
function: gets last N metrics for passed variables in query string
 */
func GetNMetricsVars(c *gin.Context) {

	// Get the number of metrics to read
	n := c.Param("n")
	num, err := strconv.Atoi(n)
	if err != nil || num < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid number in url!"})
		return
	}

	// Limit to 10 metrics
	if num > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Max 10 metrics!"})
		return
	}

	// Gets passed query string
	params := c.Request.URL.Query()

	// Build an array with the wanted variables (cicle the params to get the valid variables and ignores duplicated)
	var variables []string
	for k, v := range params {
		if k == "var" {
			for _, s := range v {
				if strings.ToLower(s) != "v1" && strings.ToLower(s) != "v2" && strings.ToLower(s) != "v3" && strings.ToLower(s) != "v4" {
					c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "You passed an invalid variable! Use v1, v2, v3 or v4."})
					return
				}
				// append if not exists
				_, found := Find(variables, strings.ToLower(s))
				if !found {
					variables = append(variables, strings.ToLower(s))
				}
			}
		}
	}

	// Collect data from DB
	rows, err := Db.Table("simu_data").Order("id desc").Select(variables).Limit(num).Rows()
	if err != nil {
		fmt.Println("DB Error: ", err.Error())
		return
	}

	// Number of columns
	columns, _ := rows.Columns()
	colNum := len(columns)

	// Prepare a map array with the values
	var results []map[string]interface{}
	for rows.Next() {
		// Prepare to read row using Scan
		r := make([]interface{}, colNum)
		for i := range r {
			r[i] = &r[i]
		}

		// Read rows using Scan
		err = rows.Scan(r...)

		// Create a row map to store row's data
		var row = map[string]interface{}{}
		for i := range r {
			row[columns[i]] = r[i]
		}

		// Append to the final results slice
		results = append(results, row)
	}

	// Send response
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": results})
}

func GetAvgVars(c *gin.Context) {

	params := c.Request.URL.Query()

	// Build an array with the wanted variables (cicle the params to get the valid variables and ignores duplicated)
	var variables []string
	for k, v := range params {
		if k == "var" {
			for _, s := range v {
				if strings.ToLower(s) != "v1" && strings.ToLower(s) != "v2" && strings.ToLower(s) != "v3" && strings.ToLower(s) != "v4" {
					c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "You passed an invalid variable! Use v1, v2, v3 or v4."})
					return
				}
				// append if not exists
				_, found := Find(variables, strings.ToLower(s))
				if !found {
					variables = append(variables, "AVG("+strings.ToLower(s)+")")
				}
			}
		}
	}

	// Collect data from DB
	rows, err := Db.Table("simu_data").Select(variables).Rows()
	if err != nil {
		fmt.Println("DB Error: ", err.Error())
		return
	}

	// Number of columns
	columns, _ := rows.Columns()
	colNum := len(columns)

	// Prepare a map array with the values
	var results []map[string]interface{}
	for rows.Next() {
		// Prepare to read row using Scan
		r := make([]interface{}, colNum)
		for i := range r {
			r[i] = &r[i]
		}

		// Read rows using Scan
		err = rows.Scan(r...)

		// Create a row map to store row's data
		var row = map[string]interface{}{}
		for i := range r {
			row[columns[i]] = r[i]
		}

		// Append to the final results slice
		results = append(results, row)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": results})
}

/*
function: Takes a slice and looks for an element in it. If found it will
   	      return it's key, otherwise it will return -1 and a bool of false.
params: slice - the slice to look into, val - the value to look for
*/
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
