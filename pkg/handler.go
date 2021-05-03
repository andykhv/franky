package franky

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var dao = TestDAO()

func DefaultHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Welcome To Franky!\n"))
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	writeHeader(writer)
	json.NewEncoder(writer).Encode(dao.GetUser())
}

func PostUser(writer http.ResponseWriter, request *http.Request) {
	dao.AddUser(dao.GetUser())
	writer.WriteHeader(http.StatusOK)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	dao.DeleteUser(userId)
	writer.WriteHeader(http.StatusOK)
}

/*
Optional Query Parameters: song, artist, album, playlist, category, range
*/
func GetRecords(writer http.ResponseWriter, request *http.Request) {
	writeHeader(writer)
	json.NewEncoder(writer).Encode(dao.GetRecords())
}

func PostRecord(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	dao.AddRecord(userId, nil)
	writer.WriteHeader(http.StatusOK)
}

func writeHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
