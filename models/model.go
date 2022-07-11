package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db gorm.DB

type Short struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect"`
	Shorted  string `json:"shorted" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	dsn := "host=172.17.0.2 user= admin password=test dbname=admin port=5432 sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

}

//docker run --name auth-psql -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -d postgres:14
//796987ecd8a9ae25fdb482179c1b48cbf5d9f253dc24313ae0d566b63e29088a
// docker postgres IP = "172.17.0.2"
