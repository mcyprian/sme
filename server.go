package main

import (
	"fmt"
	"net/http"

	"github.com/mcyprian/sme/storage"
)

func main() {
	storage.GenerateExampleData()
	storage.QueryExampleData()

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", files)
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/err", Err)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	fmt.Println("Listening at: http://0.0.0.0:8080")
	server.ListenAndServe()
}
