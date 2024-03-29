package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/mcyprian/sme/storage"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "base", data)
}

func index(writer http.ResponseWriter, request *http.Request) {
	offers, err := storage.GetCurrentOffers()
	if err != nil {
		panic(err)
	}
	generateHTML(writer, offers, "airport", "planeThumbnail", "base", "navbar", "return", "index")
}

type Offer struct {
	Airport      string
	Manufacturer string
	Type         string
	Price        float64
	OfferID      uint
}

// GET /order
// Show the order page
func order(writer http.ResponseWriter, request *http.Request) {
	airports, ok := request.URL.Query()["airport"]
	if !ok || len(airports) < 1 {
		log.Println("Url Param 'airport' is missing")
		return
	}
	airport := airports[0]

	manufacturers, ok := request.URL.Query()["manufacturer"]
	if !ok || len(manufacturers) < 1 {
		log.Println("Url Param 'manufacturer' is missing")
		return
	}
	manufacturer := manufacturers[0]

	types, ok := request.URL.Query()["type"]
	if !ok || len(types) < 1 {
		log.Println("Url Param 'type' is missing")
		return
	}
	aircraftType := types[0]

	offersID, ok := request.URL.Query()["offerID"]
	if !ok || len(offersID) < 1 {
		log.Println("Url Param 'type' is missing")
		return
	}
	i, err := strconv.ParseInt(offersID[0], 10, 32)
	if err != nil {
		panic(err)
	}
	offerID := uint(i)

	offer := new(Offer)
	offer.Airport = airport
	offer.Manufacturer = manufacturer
	offer.Type = aircraftType
	offer.OfferID = offerID
	generateHTML(writer, offer, "base", "navbar", "return", "order")
}

// POST /order_flight
// Create client if not present and his new order
func orderFlight(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	i, err := strconv.ParseInt(request.PostFormValue("offerID"), 10, 32)
	if err != nil {
		panic(err)
	}
	offerID := uint(i)
	manufacturer := request.PostFormValue("manufacturer")
	aircraftType := request.PostFormValue("type")

	order := storage.Order{
		StartTime:  time.Now().Local(),
		Name:       request.PostFormValue("name"),
		Email:      request.PostFormValue("email"),
		Phone:      request.PostFormValue("phone"),
		OfferID:    offerID,
		ReturnCode: generateID(),
	}
	storage.Db.Create(&order)
	sendOrderMail(order.Email, order.ID, order.StartTime, order.ReturnCode, manufacturer+" - "+aircraftType)
	msg := `Dear %s, <br><br>
	selected aircraft was successfully reserved.
	<br>
	Please wait for confirmation e-mail with your return code.
	<br><br>
	Thank you,
	<br>
	FlyIT team`

	generateHTML(writer, fmt.Sprintf(msg, order.Name), "base", "navbar", "return", "return_confirm")
	// http.Redirect(writer, request, "/", 302)
}

func return_confirm(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	returnCode := request.PostFormValue("code")

	offer, ok := storage.GetOrder(returnCode)
	if ok {
		storage.AddEndTimeToOrder(offer.OrderID)
		generateHTML(writer, "Dear "+offer.Name+",<br><br> your aicraft was successfully returned.<br><br>Thank you,<br> FlyIT team", "base", "navbar", "return", "return_confirm")
	} else {
		generateHTML(writer, "Return denied. WRONG return code.", "base", "navbar", "return", "return_confirm")
	}
}

func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	generateHTML(writer, vals.Get("msg"), "base", "navbar", "return", "error")
}
