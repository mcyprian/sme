package storage

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Airport struct {
	gorm.Model
	Name    string
	Address string
	Offers  []Offer `gorm:"foreignkey:AirportID"`
}

type Helicopter struct {
	gorm.Model
	Manufacturer string
	Type         string
	Offered      []Offer `gorm:"foreignkey:HelicopterID"`
}

type Client struct {
	gorm.Model
	Name   string
	Email  string
	Phone  string
	Orders []Order
}

type Offer struct {
	gorm.Model
	AirportID    uint
	HelicopterID uint
	Orders       []Order
}

type Order struct {
	gorm.Model
	Time     time.Time
	ClientID uint
	OfferID  uint
}

var Db *gorm.DB = nil

// open connection to the database
func Init() {
	if Db == nil {
		var err error
		Db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=postgres password=postgres sslmode=disable")
		if err != nil {
			panic(err)
		}
	}
}
