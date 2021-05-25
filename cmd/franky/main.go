package main

import (
	"fmt"
	"log"
	"net/http"

	franky "github.com/andykhv/franky/api"
	"github.com/andykhv/franky/test"
)

func main() {
	dao := test.NewTestDAO()
	handler := franky.NewFrankyHandler(&dao)
	router := franky.Router(handler)

	fmt.Println("starting...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
