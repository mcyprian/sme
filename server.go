package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	fmt.Println("Listening at: http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
