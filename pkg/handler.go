package franky

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type frankyHandler struct {
	dao *FrankyDAO
}

func NewFrankyHandler() *frankyHandler {
	dao := TestDAO()
	handler := frankyHandler{&dao}

	return &handler
}

func (handler *frankyHandler) defaultHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Welcome To Franky!\n"))
}

func (handler *frankyHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	user, err := (*handler.dao).GetUser(userId)

	if err != nil {
		writeErrorHeader(writer, err)
	} else {
		writeOkHeaderWithJson(writer)
		json.NewEncoder(writer).Encode(user)
	}
}

func (handler *frankyHandler) PostUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	user, _ := (*handler.dao).GetUser(userId)
	err := (*handler.dao).AddUser(user)

	if err != nil {
		writeErrorHeader(writer, err)
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *frankyHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	err := (*handler.dao).DeleteUser(userId)

	if err != nil {
		writeErrorHeader(writer, err)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

/*
Optional Query Parameters: song, artist, album, playlist, category, range
*/
func (handler *frankyHandler) GetRecords(writer http.ResponseWriter, request *http.Request) {
	records, err := (*handler.dao).GetRecords()

	if err != nil {
		writeErrorHeader(writer, err)
	} else {
		writeOkHeaderWithJson(writer)
		json.NewEncoder(writer).Encode(records)
	}
}

func (handler *frankyHandler) PostRecord(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	err := (*handler.dao).AddRecord(userId, nil)

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
