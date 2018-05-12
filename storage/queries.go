package storage

import (
	"fmt"
)

func GenerateExampleData() {
	Init()

	Db.DropTable(&Airport{})
	Db.DropTable(&Helicopter{})
	Db.DropTable(&Offer{})

	Db.AutoMigrate(&Airport{}, &Helicopter{}, &Client{}, &Offer{}, &Order{})

	medlanky := Airport{Name: "Letisko Medlánky", Address: "Turistická 67, Brno"}
	turany := Airport{Name: "Letište Tuřany", Address: "Tuřany 904/1, Brno"}

	ranger := Helicopter{Manufacturer: "Bell", Type: "505 Jet Ranger X"}
	sikorsky := Helicopter{Manufacturer: "Sikorsky", Type: "S-76C"}
	eurocopter := Helicopter{Manufacturer: "Eurocopter", Type: "AS350"}

	cheap := Offer{}
	sport := Offer{}

	Db.Create(&medlanky)
	Db.Create(&turany)
	Db.Create(&ranger)
	Db.Create(&sikorsky)
	Db.Create(&eurocopter)

	Db.Model(&medlanky).Association("Offers").Append(cheap)
	Db.Model(&turany).Association("Offers").Append(sport)

	Db.Model(&ranger).Association("Offered").Append(cheap)
	Db.Model(&eurocopter).Association("Offered").Append(sport)
}

func QueryExampleData() {
	Init()
	var airport Airport
	var offers []Offer

	Db.Last(&airport)
	fmt.Println(airport)

	Db.Find(&offers)
	// Db.Model(&airport).Related(&offers)
	fmt.Println(offers)
}
