package model

import "github.com/jinzhu/gorm"

// Data sample struct
type SimuData struct {
	gorm.Model
	V1 int
	V2 int
	V3 int
	V4 int
}
