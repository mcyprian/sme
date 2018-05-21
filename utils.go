package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func generateID() string {
	id := make([]byte, 16)
	r := rand.Reader
	if _, err := io.ReadFull(r, id); err != nil {
		panic(err)
	}
	stringID := (fmt.Sprintf(`%02X`, id))
	return stringID
}
