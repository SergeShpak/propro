package main

import (
	"log"
)

func main() {
	r := newRouter()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
