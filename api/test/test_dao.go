package test

import (
	"fmt"
	"net/http"

	franky "github.com/andykhv/franky/api"
)

type testDAO struct {
}

func NewTestDAO() franky.FrankyDAO {
	return &testDAO{}
}

func (dao *testDAO) GetUser(userId string) (*franky.User, *franky.HttpError) {
	if userId == "123" {
		return user1, nil
	}

	err := fmt.Errorf("userId %s not found", userId)
	httpError := franky.HttpError{StatusCode: http.StatusNotFound, Err: err}

	return nil, &httpError
}

func (dao *testDAO) AddUser(user *franky.User) *franky.HttpError {
	if user.Email == "email" {
		err := fmt.Errorf("email already exists")
		httpError := franky.HttpError{StatusCode: http.StatusNotFound, Err: err}
		return &httpError
	}

	return nil
}

func (dao *testDAO) DeleteUser(userId string) *franky.HttpError {
	if userId == "123" {
		return nil
	}

	err := fmt.Errorf("userId %s not found", userId)
	httpError := franky.HttpError{StatusCode: http.StatusNotFound, Err: err}
	return &httpError
}

func (dao *testDAO) GetRecords(recordRequest *franky.RecordRequest) ([]franky.Record, *franky.HttpError) {
	record := franky.Record{UserId: "123", Song: "song", Artist: "artist", Album: "album", Playlist: "playlist", Duration: 180, Date: 1622143906, Category: "rap"}
	return []franky.Record{record, record}, nil
}

func (dao *testDAO) AddRecord(userId string, record *franky.Record) *franky.HttpError {
	return nil
}
