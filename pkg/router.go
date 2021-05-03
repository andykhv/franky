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
	router.HandleFunc("/users/{id:[0-9]+}", PostUser).
		Methods(http.MethodPost)
	router.HandleFunc("/users/{id:[0-9]+}", PostUser).
		Methods(http.MethodPut)
	router.HandleFunc("/users/{id:[0-9]+}", DeleteUser).
		Methods(http.MethodDelete)

	router.HandleFunc("/users/{id:[0-9]+}/records", GetRecords).
		Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}/records", PostRecord).
		Methods(http.MethodPost)

	return router
}
