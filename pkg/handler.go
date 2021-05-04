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
	userId := mux.Vars(request)["id"]
	user, err := dao.GetUser(userId)

	if err != nil {
		writeErrorHeader(writer, err)
	} else {
		writeOkHeaderWithJson(writer)
		json.NewEncoder(writer).Encode(user)
	}
}

func PostUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	user, _ := dao.GetUser(userId)
	err := dao.AddUser(user)

	if err != nil {
		writeErrorHeader(writer, err)
	}

	writer.WriteHeader(http.StatusOK)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	err := dao.DeleteUser(userId)

	if err != nil {
		writeErrorHeader(writer, err)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

/*
Optional Query Parameters: song, artist, album, playlist, category, range
*/
func GetRecords(writer http.ResponseWriter, request *http.Request) {
	records, err := dao.GetRecords()

	if err != nil {
		writeErrorHeader(writer, err)
	} else {
		writeOkHeaderWithJson(writer)
		json.NewEncoder(writer).Encode(records)
	}
}

func PostRecord(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	err := dao.AddRecord(userId, nil)

	if err != nil {
		writeErrorHeader(writer, err)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

func writeOkHeaderWithJson(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func writeErrorHeader(writer http.ResponseWriter, err *HttpError) {
	writer.WriteHeader(err.StatusCode)
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte(err.Err.Error()))
}
