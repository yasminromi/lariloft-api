package model

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Visit struct {
	ID          int64   `gorm:"primary_key" json:"id,omitempty"`
	UserID      int64   `gorm:"user_id" json:"userId,omitempty"`
	ApartmentID int64   `gorm:"apartment_id" json:"apartmentId,omitempty"`
	AgentID     int64   `gorm:"agent_id" json:"agentId,omitempty"`
	Rate        float64 `gorm:"rate" json:"rate,omitempty"`
	ViewRate    float64 `gorm:"view_rate" json:"viewRate,omitempty"`
	LightRate   float64 `gorm:"light_rate" json:"lightRate,omitempty"`
	Opinion     string  `gorm:"opinion" json:"opinion,omitempty"`
}

func (dsd *LariLoftDB) CreateVisit(visit *Visit) error {
	result := dsd.Db.Table("public.visits").Create(visit)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dsd *LariLoftDB) GetVisit(id int) (*Visit, error) {
	visit := Visit{}
	result := dsd.Db.Table("public.visits").Where("id = ?", id).First(&visit)
	if result.Error != nil {
		log.Println("error on get data from visits", result.Error)
		return nil, result.Error
	}
	return &visit, nil
}

func (dsd *LariLoftDB) GetAllVisitApartmentsByUser(userID int) (*[]Apartment, error) {
	apartments := []Apartment{}
	result := dsd.Db.Table("public.visits").Where("user_id = ?", userID).Find(&apartments)
	if result.Error != nil {
		log.Println("error on get data from visit", result.Error)
		return nil, result.Error
	}
	return &apartments, nil
}
