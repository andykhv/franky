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
	writeHeader(writer)
	json.NewEncoder(writer).Encode(dao.GetUser())
}

/*
Optional Query Parameters: song, artist, album, playlist, category, range
*/
func GetRecords(writer http.ResponseWriter, request *http.Request) {
	writeHeader(writer)
	json.NewEncoder(writer).Encode(dao.GetRecords())
}

func writeHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
