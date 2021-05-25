package franky

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (handler *FrankyHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		httpError := &HttpError{http.StatusBadRequest, err}
		writeErrorHeader(writer, httpError)
		return
	}

	var token map[string]string
	err = json.Unmarshal(body, &token)
	if err != nil {
		httpError := &HttpError{http.StatusBadRequest, err}
		writeErrorHeader(writer, httpError)
		return
	}

	userId := mux.Vars(request)["id"]
	user, httpErr := (*handler.dao).GetUser(userId)
	if httpErr != nil {
		writeErrorHeader(writer, httpErr)
		return
	}
	if user.ApiKey != token["token"] {
		httpError := &HttpError{http.StatusUnauthorized, errors.New("invalid token")}
		writeErrorHeader(writer, httpError)
		return
	}

	writeOkHeaderWithJson(writer)
	json.NewEncoder(writer).Encode(user)
}

func (handler *FrankyHandler) PostUser(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
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

	user.setCreationDate()
	user.generateId()
	user.generateApiKey()
	err = user.saltAndHashPassword()
	if err != nil {
		httpError := &HttpError{http.StatusInternalServerError, err}
		writeErrorHeader(writer, httpError)
		return
	}

	httpError := (*handler.dao).AddUser(&user)
	if httpError != nil {
		writeErrorHeader(writer, httpError)
		return
	}

	writeOkHeaderWithJson(writer)
	json.NewEncoder(writer).Encode(fmt.Sprintf(`{"token":"%s", "id":"%s"}`, user.ApiKey, user.Id))
}

func (handler *FrankyHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		httpError := &HttpError{http.StatusBadRequest, err}
		writeErrorHeader(writer, httpError)
		return
	}

	var token map[string]string
	err = json.Unmarshal(body, &token)
	if err != nil {
		httpError := &HttpError{http.StatusBadRequest, err}
		writeErrorHeader(writer, httpError)
		return
	}

	userId := mux.Vars(request)["id"]
	user, httpErr := (*handler.dao).GetUser(userId)
	if httpErr != nil {
		writeErrorHeader(writer, httpErr)
		return
	}
	if user.ApiKey != token["token"] {
		httpError := &HttpError{http.StatusUnauthorized, errors.New("invalid token")}
		writeErrorHeader(writer, httpError)
		return
	}

	httpErr = (*handler.dao).DeleteUser(userId)
	if err != nil {
		writeErrorHeader(writer, httpErr)
		return
	}

	writer.WriteHeader(http.StatusOK)
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
