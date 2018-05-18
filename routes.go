package main

import (
	"fmt"
	"net/http"
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
	generateHTML(writer, offers, "airport", "planeThumbnail", "base", "navbar", "index")
}

// GET /order
// Show the order page
func order(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "base", "navbar", "order")
}

// POST /order_flight
// Create client if not present and his new order
func orderFlight(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	order := storage.Order{
		StartTime: time.Now().Local(),
		Name:      request.PostFormValue("name"),
		Email:     request.PostFormValue("email"),
		Phone:     request.PostFormValue("phone"),
		// TODO add offer id
	}
	fmt.Println(order)
	storage.Db.Create(&order)
	http.Redirect(writer, request, "/", 302)
}

func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	generateHTML(writer, vals.Get("msg"), "base", "navbar", "error")
}
