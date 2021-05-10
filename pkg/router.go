package franky

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router(handler *FrankyHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.defaultHandler)
	router.HandleFunc("/users/{id:[0-9]+}", handler.GetUser).
		Methods(http.MethodGet)
	router.HandleFunc("/users", handler.PostUser).
		Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/users/{id:[0-9]+}", handler.DeleteUser).
		Methods(http.MethodDelete)

	router.HandleFunc("/users/{id:[0-9]+}/records", handler.GetRecords).
		Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}/records", handler.PostRecord).
		Methods(http.MethodPost)

	return router
}
