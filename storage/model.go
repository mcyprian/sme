package storage

import (
	"fmt"
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
	Price        float64
	Orders       []Order
}

type Order struct {
	gorm.Model
	StartTime time.Time
	EndTime   time.Time
	ClientID  uint
	OfferID   uint
}

var Db *gorm.DB = nil

// open connection to the database
func Init(host string, password string) {
	if Db == nil {
		var err error
		connectionStr := fmt.Sprintf("host=%s user=postgres dbname=postgres password=%s sslmode=disable", host, password)
		Db, err = gorm.Open("postgres", connectionStr)
		if err != nil {
			panic(err)
		}
	}
}
