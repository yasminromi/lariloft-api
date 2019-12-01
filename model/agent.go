package model

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Agent struct {
	ID     int64  `gorm:"primary_key" json:"id,omitempty"`
	Name   string `gorm:"column:username" json:"username,omitempty"`
	Msisdn string `gorm:"column:msisdn" json:"username,omitempty"`
}

func (dsd *LariLoftDB) CreateAgent(agent *Agent) error {
	result := dsd.Db.Table("public.agents").Create(agent)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dsd *LariLoftDB) GetAgent(id int) (*Agent, error) {
	agent := Agent{}

	result := dsd.Db.Table("public.agents").First(&agent, "id = ?", id)

	if result.Error != nil && !result.RecordNotFound() {
		log.Println("error on get data from agent", result.Error)
		return nil, result.Error
	}
	return &agent, nil
}

func (dsd *LariLoftDB) GetAllAgents() (*[]Agent, error) {
	agents := []Agent{}
	result := dsd.Db.Table("public.agents").Find(&agents)
	if result.Error != nil {
		log.Println("error on get data from agent", result.Error)
		return nil, result.Error
	}
	return &agents, nil
}
