package franky

import (
	"github.com/gorilla/mux"
)

var queryParameters = []string{
	"song", "{[\\w]+}",
	"artist", "{[\\w]+}",
	"album", "{[\\w]+}",
	"playlist", "{[\\w]+}",
	"category", "{[\\w]+}",
	"range", "{[0-9]+-[0-9]+}",
}

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", DefaultHandler)

	router.HandleFunc("/users/{id:[0-9]+}", GetUser).
		Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", DefaultHandler).
		Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", DefaultHandler).
		Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", DefaultHandler).
		Methods("DELETE")

	router.HandleFunc("/users/{id:[0-9]+}/records", DefaultHandler).
		Methods("GET").
		Queries(queryParameters...)
	router.HandleFunc("/users/{id:[0-9]+}/records", DefaultHandler).
		Methods("POST")

	return router
}
