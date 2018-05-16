package storage

import (
	"fmt"
	"time"
)

func GenerateExampleData() {
	Db.DropTable(&Airport{})
	Db.DropTable(&Helicopter{})
	Db.DropTable(&Client{})
	Db.DropTable(&Offer{})
	Db.DropTable(&Order{})

	Db.AutoMigrate(&Airport{}, &Helicopter{}, &Client{}, &Offer{}, &Order{})

	medlanky := Airport{Name: "Letisko Medlánky", Address: "Turistická 67, Brno"}
	turany := Airport{Name: "Letište Tuřany", Address: "Tuřany 904/1, Brno"}

	ranger := Helicopter{Manufacturer: "Bell", Type: "505 Jet Ranger X"}
	sikorsky := Helicopter{Manufacturer: "Sikorsky", Type: "S-76C"}
	eurocopter := Helicopter{Manufacturer: "Eurocopter", Type: "AS350"}

	ondro := Client{Name: "Ondrej Nečas", Email: "onecas@seznam.cz", Phone: "+421 758 633 715"}

	Db.Create(&medlanky)
	Db.Create(&turany)
	Db.Create(&ranger)
	Db.Create(&sikorsky)
	Db.Create(&eurocopter)
	Db.Create(&ondro)
	flight := Order{StartTime: time.Now(), EndTime: time.Now().Add(time.Hour * 3), ClientID: ondro.ID}

	cheap := Offer{HelicopterID: ranger.ID, Price: 100.0}
	sport := Offer{HelicopterID: eurocopter.ID, Price: 180.0}

	Db.Create(&cheap)
	Db.Create(&sport)

	Db.Model(&medlanky).Association("Offers").Append(cheap)
	Db.Model(&turany).Association("Offers").Append(sport)
	Db.Model(&cheap).Association("Orders").Append(flight)

}

func QueryExampleData() {
	var airport Airport
	var helicopter Helicopter
	var client Client
	var offers []Offer
	var orders []Order

	Db.Last(&airport)
	Db.First(&helicopter)
	Db.First(&client)
	fmt.Println(airport)
	fmt.Println(client)

	// Db.Find(&offers)
	// Db.Model(&airport).Related(&offers)
	Db.Model(&helicopter).Related(&offers)
	Db.Model(&offers[0]).Related(&orders)
	Db.Find(&orders)
	fmt.Println(offers)
	fmt.Println(orders)
}

type OffersRow struct {
	Airport      string
	Manufacturer string
	Type         string
	Price        float64
}

func GetCurrentOffers() ([]*OffersRow, error) {
	rows, err := Db.Table("offers").Select(
		"airports.name, helicopters.manufacturer, helicopters.type, offers.price").Joins(
		"left join airports on airports.id = offers.airport_id").Joins(
		"left join helicopters on helicopters.id = offers.helicopter_id").Rows()

	offers := make([]*OffersRow, 0)
	if err != nil {
		return offers, err
	}
	for rows.Next() {
		rst := new(OffersRow)
		err := rows.Scan(&rst.Airport, &rst.Manufacturer, &rst.Type, &rst.Price)
		if err != nil {
			panic(err)
		}

		offers = append(offers, rst)
	}
	return offers, nil
}
