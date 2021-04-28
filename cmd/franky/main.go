package main

import (
	"fmt"
	"log"
	"net/http"

	franky "github.com/andykhv/franky/pkg"
)

func main() {
	router := franky.Router()

	fmt.Println("starting...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
