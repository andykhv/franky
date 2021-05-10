package test

import franky "github.com/andykhv/franky/pkg"

type testDAO struct {
}

func NewTestDAO() franky.FrankyDAO {
	return &testDAO{}
}

func (dao *testDAO) GetUser(userId string) (*franky.User, *franky.HttpError) {
	user := franky.User{userId, "email", "password", "apiKey", "creationDate"}
	return &user, nil
}

func (dao *testDAO) AddUser(user *franky.User) *franky.HttpError {
	return nil
}

func (dao *testDAO) DeleteUser(userId string) *franky.HttpError {
	return nil
}

func (dao *testDAO) GetRecords() ([]franky.Record, *franky.HttpError) {
	record := franky.Record{"song", "artist", "album", "playlist", 180, 1000, "rap"}
	return []franky.Record{record, record}, nil
}

func (dao *testDAO) AddRecord(userId string, record *franky.Record) *franky.HttpError {
	return nil
}
