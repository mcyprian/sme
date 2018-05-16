package main

import (
	"fmt"
	"net/http"

	"github.com/mcyprian/sme/storage"
)

func main() {
	loadConfig()
	storage.Init(config.DBHost, config.DBPasswd)
	storage.GenerateExampleData()
	storage.QueryExampleData()

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	fmt.Println("Listening at: http://0.0.0.0:8080")
	server.ListenAndServe()
}
