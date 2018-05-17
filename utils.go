package main

import (
	"crypto/rand"
	"io"
)

func generateID() []byte {
	id := make([]byte, 16)
	r := rand.Reader
	if _, err := io.ReadFull(r, id); err != nil {
		panic(err)
	}
	return id
}
