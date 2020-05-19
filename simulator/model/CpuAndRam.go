package model

import "github.com/jinzhu/gorm"

// CPU and RAM info struct
type CpuAndRam struct {
	gorm.Model
	CPU float64
	RAM float64
}
