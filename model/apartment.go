package model

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Apartment struct {
	ID               int64   `gorm:"primary_key" json:"id,omitempty"`
	Lat              float64 `gorm:"column:lat" json:"lat,omitempty"`
	Long             float64 `gorm:"column:long" json:"long,omitempty"`
	NumberOfRooms    int64   `gorm:"column:number_of_rooms" json:"numberOfRooms,omitempty"`
	MaxEstimate      float64 `gorm:"column:max_estimate" json:"maxEstimate,omitempty"`
	MinEstimate      float64 `gorm:"column:min_estimate" json:"minEstimate,omitempty"`
	KitchenSize      float64 `gorm:"column:kitchen_size" json:"kitchenSize,omitempty"`
	LivingRoomSize   float64 `gorm:"column:living_room_size" json:"livingRoomSize,omitempty"`
	HasOffice        bool    `gorm:"column:has_office" json:"hasOffice,omitempty"`
	HasPool          bool    `gorm:"column:has_pool" json:"hasPool,omitempty"`
	AcceptPet        bool    `gorm:"column:accept_pet" json:"acceptPet,omitempty"`
	PriceCondominium float64 `gorm:"column:price_condominium" json:"priceCondominium,omitempty"`
}

//CreateUser: criar um apartamento
func (dsd *LariLoftDB) CreateApartment(apartment *Apartment) error {
	result := dsd.Db.Table("public.apartments").Create(apartment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetUsers: retorna um apartamento
func (dsd *LariLoftDB) GetApartment(id int) (*Apartment, error) {
	apartment := Apartment{}

	result := dsd.Db.Table("public.apartments").First(&apartment, "id = ?", id)

	if result.Error != nil && !result.RecordNotFound() {
		log.Println("error on get data from apartment", result.Error)
		return nil, result.Error
	}
	return &apartment, nil
}

//GetUsers: retorna todos os apartamentos
func (dsd *LariLoftDB) GetAllApartments() (*[]Apartment, error) {
	apartments := []Apartment{}
	result := dsd.Db.Table("public.apartments").Find(&apartments)
	if result.Error != nil {
		log.Println("error on get data from apartment", result.Error)
		return nil, result.Error
	}
	return &apartments, nil
}
