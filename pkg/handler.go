package franky

import (
	"encoding/json"
	"net/http"
)

var dao = TestDAO()

func DefaultHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Welcome To Franky!\n"))
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(dao.GetUser())
}
