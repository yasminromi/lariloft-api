package model

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Agent struct {
	ID     int64  `gorm:"primary_key" json:"id,omitempty"`
	Name   string `gorm:"column:name" json:"username,omitempty"`
	Msisdn string `gorm:"column:msisdn" json:"username,omitempty"`
}
