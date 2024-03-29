package model

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// Modelo de acesso ao banco
type LariLoftDB struct {
	Db *gorm.DB
}

//Inicialização do repository
func (dsd *LariLoftDB) MustInit() {
	var err error
	dsd.Db, err = gorm.Open("postgres", Connection())
	dsd.Db.LogMode(true)
	if err != nil {
		log.Println("error on connect database")
		return
	}
}

//Helper para formatar string de conexão
func Connection() string {
	return os.Getenv("DATABASE_URL")
}

func ElasticSearch() string {
	return os.Getenv("BONSAI_URL")
}
