package franky

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type FrankyHandler struct {
	dao *FrankyDAO
}

func NewFrankyHandler(dao *FrankyDAO) *FrankyHandler {
	handler := FrankyHandler{dao}

	return &handler
}

func (handler *FrankyHandler) defaultHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Welcome To Franky!\n"))
}

func (handler *FrankyHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]
	user, err := (*handler.dao).GetUser(userId)

	if err != nil {
		writeErrorHeader(writer, err)
	} else {
		writeOkHeaderWithJson(writer)
		json.NewEncoder(writer).Encode(user)
	}
}

func (handler *FrankyHandler) PostUser(writer http.ResponseWriter, request *http.Request) {
	bodyReadable := request.Body

	body, err := ioutil.ReadAll(bodyReadable)
	if err != nil {
		httpError := &HttpError{http.StatusBadRequest, err}
		writeErrorHeader(writer, httpError)
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		httpError := &HttpError{http.StatusBadRequest, err}
		writeErrorHeader(writer, httpError)
		return
	}

	httpError := (*handler.dao).AddUser(&user)
	if err != nil {
		writeErrorHeader(writer, httpError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *FrankyHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
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
func (handler *FrankyHandler) GetRecords(writer http.ResponseWriter, request *http.Request) {
	records, err := (*handler.dao).GetRecords()

	if err != nil {
		writeErrorHeader(writer, err)
	} else {
		writeOkHeaderWithJson(writer)
		json.NewEncoder(writer).Encode(records)
	}
}

func (handler *FrankyHandler) PostRecord(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["id"]

	bodyReader := request.Body

	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		httpError := &HttpError{http.StatusBadRequest, err}
		writeErrorHeader(writer, httpError)
		return
	}

	var records []Record
	err = json.Unmarshal(body, &records)
	if err != nil {
		httpError := &HttpError{http.StatusBadRequest, err}
		writeErrorHeader(writer, httpError)
		return
	}

	httpError := (*handler.dao).AddRecord(userId, &records[0])
	if httpError != nil {
		writeErrorHeader(writer, httpError)
		return
	}

	writer.WriteHeader(http.StatusOK)
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
