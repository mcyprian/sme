package storage

import (
	"fmt"
	"time"
)

func GenerateExampleData() {
	Db.DropTable(&Airport{})
	Db.DropTable(&Helicopter{})
	Db.DropTable(&Offer{})
	Db.DropTable(&Order{})

	Db.AutoMigrate(&Airport{}, &Helicopter{}, &Offer{}, &Order{})

	medlanky := Airport{Name: "Letisko Medlánky", Address: "Turistická 67, Brno"}
	turany := Airport{Name: "Letište Tuřany", Address: "Tuřany 904/1, Brno"}

	ranger := Helicopter{Manufacturer: "Bell", Type: "505 Jet Ranger X"}
	sikorsky := Helicopter{Manufacturer: "Sikorsky", Type: "S-76C"}
	eurocopter := Helicopter{Manufacturer: "Eurocopter", Type: "AS350"}
	dynamic := Helicopter{Manufacturer: "Dynamic", Type: "WT9 MUNI-OUU 34"}
	zlin := Helicopter{Manufacturer: "Zlin", Type: "Z-226 MUNI-MGM"}
	piper := Helicopter{Manufacturer: "Piper", Type: "J3C-65 Cub MUNI-ONY"}
	piperPaw := Helicopter{Manufacturer: "Piper", Type: "Pawnee MUNI-MLP"}
	blanik := Helicopter{Manufacturer: "Blanik", Type: "L13 MUNI-1823"}
	blanikS := Helicopter{Manufacturer: "Blanik", Type: "Super L23 MUNI-5550"}
	orlik := Helicopter{Manufacturer: "Orlik", Type: "VT116 MUNI-4321"}
	vosa := Helicopter{Manufacturer: "Vosa", Type: "VSO10 MUNI-1504"}

	Db.Create(&medlanky)
	Db.Create(&turany)
	Db.Create(&ranger)
	Db.Create(&sikorsky)
	Db.Create(&eurocopter)
	Db.Create(&dynamic)
	Db.Create(&zlin)
	Db.Create(&piper)
	Db.Create(&piperPaw)
	Db.Create(&blanik)
	Db.Create(&blanikS)
	Db.Create(&orlik)
	Db.Create(&vosa)

	flight := Order{StartTime: time.Now(), EndTime: time.Now().Add(time.Hour * 3),
		Name: "Ondrej Nečas", Email: "onecas@seznam.cz",
		Phone: "+421 758 633 715"}

	cheap := Offer{HelicopterID: ranger.ID, Price: 100.0}
	sport := Offer{HelicopterID: eurocopter.ID, Price: 180.0}
	zlinoffer := Offer{HelicopterID: zlin.ID, Price: 90.0}

	offer := Offer{HelicopterID: zlin.ID, Price: 150.0}
	Db.Create(&offer)
	Db.Model(&turany).Association("Offers").Append(offer)

	Db.Create(&cheap)
	Db.Create(&sport)

	Db.Model(&medlanky).Association("Offers").Append(cheap)
	Db.Model(&medlanky).Association("Offers").Append(zlinoffer)
	Db.Model(&turany).Association("Offers").Append(sport)
	Db.Model(&cheap).Association("Orders").Append(flight)
}

func QueryExampleData() {
	var airport Airport
	var helicopter Helicopter
	var offers []Offer
	var orders []Order

	Db.Last(&airport)
	Db.First(&helicopter)
	fmt.Println(airport)

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
	OfferID      uint
}

func GetAllOffers() ([]*OffersRow, error) {
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

type OrderRow struct {
	OrderID    int
	Name       string
	ReturnCode string
}

func GetOrder(returnCode string) *OrderRow {
	ords, err := Db.Table("orders").Select(
		"orders.id, orders.name, orders.return_code").Rows()
	if err != nil {
		panic(err)
	}
	orderToReturn := new(OrderRow)

	for ords.Next() {
		err = ords.Scan(&orderToReturn.OrderID, &orderToReturn.Name, &orderToReturn.ReturnCode)
		if orderToReturn.ReturnCode == returnCode {
			return orderToReturn
		}
	}
	// signalize no order with this return code
	orderToReturn.OrderID = -1
	return orderToReturn
}

func GetCurrentOffers() (map[string][]*OffersRow, error) {
	isOrdered := false

	offers := make([]*OffersRow, 0)

	// Get list of currently ordered offers
	ords, er := Db.Table("orders").Select(
		"orders.offer_id").Where("orders.end_time::date <= date '1000-01-01'").Rows()
	if er != nil {
		return nil, er
	}
	rows, err := Db.Table("offers").Select(
		"airports.name, helicopters.manufacturer, helicopters.type, offers.price, offers.id").Joins(
		"join airports on airports.id = offers.airport_id").Joins(
		"join helicopters on helicopters.id = offers.helicopter_id").Rows()

	if err != nil {
		return nil, err
	}
	// For every aircraft chceck if its not already ordered
	for rows.Next() {
		rst := new(OffersRow)
		err := rows.Scan(&rst.Airport, &rst.Manufacturer, &rst.Type, &rst.Price, &rst.OfferID)
		if err != nil {
			panic(err)
		}

		ordid := uint(0)
		for ords.Next() {
			er := ords.Scan(&ordid)
			if er != nil {
				return nil, er
			}
			if ordid == rst.OfferID {
				isOrdered = true
			}
		}
		if !isOrdered {
			offers = append(offers, rst)
		}
		isOrdered = false
		// Get list of orders again
		ords, er = Db.Table("orders").Select(
			"orders.offer_id").Where("orders.end_time::date <= date '0001-01-01'").Rows()
		if er != nil {
			return nil, er
		}
	}
	return GroupByAirport(offers), nil
}

// GroupByAirport creates map which arirports as a key for the front page listing
func GroupByAirport(offers []*OffersRow) map[string][]*OffersRow {
	offersMap := make(map[string][]*OffersRow)
	for _, offer := range offers {
		offersMap[offer.Airport] = append(offersMap[offer.Airport], offer)
	}
	return offersMap
}

func AddEndTimeToOrder(orderID int) {
	Db.Table("orders").Where("id= ?", orderID).Update("end_time", time.Now().Local())
}
