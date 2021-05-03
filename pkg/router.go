package franky

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", DefaultHandler)

	router.HandleFunc("/users/{id:[0-9]+}", GetUser).
		Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}", DefaultHandler).
		Methods(http.MethodPost)
	router.HandleFunc("/users/{id:[0-9]+}", DefaultHandler).
		Methods(http.MethodPut)
	router.HandleFunc("/users/{id:[0-9]+}", DefaultHandler).
		Methods(http.MethodDelete)

	router.HandleFunc("/users/{id:[0-9]+}/records", GetRecords).
		Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}/records", DefaultHandler).
		Methods(http.MethodPost)

	return router
}
