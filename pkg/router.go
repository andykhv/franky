package franky

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)

	return router
}

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("franky!\n"))
}
