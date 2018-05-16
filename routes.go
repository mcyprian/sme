package main

import (
	"fmt"
	"net/http"
	"text/template"

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
	generateHTML(writer, offers, "base", "navbar", "index")
}

func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	generateHTML(writer, vals.Get("msg"), "base", "navbar", "error")
}
