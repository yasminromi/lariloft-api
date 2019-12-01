package model

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID              int64   `gorm:"primary_key" json:"id,omitempty"`
	Name            string  `gorm:"name" json:"name,omitempty"`
	Age             int64   `gorm:"age" json:"age,omitempty"`
	Work            string  `gorm:"work" json:"work,omitempty"`
	ActualNeighbor  string  `gorm:"actual_neighbor" json:"actualNeighbor,omitempty"`
	DesiredNeighbor string  `gorm:"desired_neighbor" json:"desiredNeighbor,omitempty"`
	Motive          string  `gorm:"motive" json:"motive,omitempty"`
	Rooms           int64   `gorm:"rooms" json:"rooms,omitempty"`
	CoLivers        string  `gorm:"co_livers" json:"coLivers,omitempty"`
	Kids            int64   `gorm:"kids" json:"kids,omitempty"`
	Pet             bool    `gorm:"pet" json:"pet,omitempty"`
	InterestHome    string  `gorm:"interest_home" json:"interestHome,omitempty"`
	InterestOutside string  `gorm:"interest_outside" json:"interestOutside,omitempty"`
	InterestRooms   string  `gorm:"interest_rooms" json:"interestRooms,omitempty"`
	Budget          float64 `gorm:"budget" json:"budget,omitempty"`
	Msisdn          string  `gorm:"msisdn" json:"msisdn,omitempty"`
}

//CreateUser: criar um usuário
func (dsd *LariLoftDB) CreateUser(user *User) error {
	result := dsd.Db.Table("public.users").Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetUsers: retorna um usuário
func (dsd *LariLoftDB) GetUser(id int) (*User, error) {
	user := User{}

	result := dsd.Db.Table("public.users").First(&user, "id = ?", id)

	if result.Error != nil && !result.RecordNotFound() {
		log.Println("error on get data from user", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

//GetUsers: retorna todos os usuários
func (dsd *LariLoftDB) GetAllUsers() (*[]User, error) {
	users := []User{}
	result := dsd.Db.Table("public.users").Find(&users)
	if result.Error != nil {
		log.Println("error on get data from user", result.Error)
		return nil, result.Error
	}
	return &users, nil
}
