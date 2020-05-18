package controller

import (
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"os"
	"time"
	"ubiwhere/model"
)

/*
function: Gets CPU & RAM info every second and adds them to the Database
*/
func GetCpuAndRamInfo() {

	// Get CPU & RAM info every second and add to database
	for {
		var info model.CpuAndRam
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

		// Insert CPU info into the struct
		info.CPU = 100 - float64(after.Idle-before.Idle)/total*100

		// Get RAM info
		mem, err := memory.Get()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}

		// Insert RAM info into the struct
		info.RAM =float64(mem.Used) / 1000000000

		// Insert CPU & RAM info to the DB
		Db.Create(&info)
	}
}