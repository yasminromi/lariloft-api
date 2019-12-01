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

//CreateHackathonUser: criar associação de hackathon e hackathoner
func (dsd *WeeHackDB) CreateHackathonUser(hackathonUser *HackathonUser) error {
	result := dsd.Db.Table("public.hackathon_user").Create(hackathonUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//CreateHackathonUser: criar associação de hackathon e hackathoner
func (dsd *WeeHackDB) CreateByUserID(userID int64, hackathonIDs []int64) error {

	for _, hackathonID := range hackathonIDs {

		hackathonUser := &HackathonUser{
			HackathonID: hackathonID,
			UserID:      userID,
		}

		dsd.CreateHackathonUser(hackathonUser)
	}
	return nil
}

//CreateHackathonUser: criar associação de hackathon e hackathoner
func (dsd *WeeHackDB) CreateByHackathonID(hackathonID int64, userIDs []int64) error {

	for _, userID := range userIDs {

		hackathonUser := &HackathonUser{
			HackathonID: hackathonID,
			UserID:      userID,
		}

		dsd.CreateHackathonUser(hackathonUser)
	}

	return nil
}

//GetScouts: retorna uma associação de hackathon e hackathoner
func (dsd *LariLoftDB) GetVisit(id int) (*Visit, error) {
	visit := Visit{}
	result := dsd.Db.Table("public.visits").Where("id = ?", id).First(&visit)
	if result.Error != nil {
		log.Println("error on get data from visits", result.Error)
		return nil, result.Error
	}
	return &visit, nil
}

//GetUsers: retorna todos os hackathon
func (dsd *WeeHackDB) GetAllHackathonUsers() (*[]HackathonUser, error) {
	hackathonUsers := []HackathonUser{}
	result := dsd.Db.Table("public.hackathon").Find(&hackathonUsers)
	if result.Error != nil {
		log.Println("error on get data from hackathonUser", result.Error)
		return nil, result.Error
	}
	return &hackathonUsers, nil
}
