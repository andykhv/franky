package franky

type FrankyDAO interface {
	GetUser(userId string) (*User, *HttpError)
	AddUser(user *User) *HttpError
	DeleteUser(userId string) *HttpError
	GetRecords() ([]Record, *HttpError)
	AddRecord(userId string, record *Record) *HttpError
}

type testDAO struct {
}

func TestDAO() FrankyDAO {
	return &testDAO{}
}

func (dao *testDAO) GetUser(userId string) (*User, *HttpError) {
	user := User{"id", "email", "password", "apiKey", "creationDate"}
	return &user, nil
}

func (dao *testDAO) AddUser(user *User) *HttpError {
	return nil
}

func (dao *testDAO) DeleteUser(userId string) *HttpError {
	return nil
}

func (dao *testDAO) GetRecords() ([]Record, *HttpError) {
	record := Record{"song", "artist", "album", "playlist", 180, 1000, "rap"}
	return []Record{record, record}, nil
}

func (dao *testDAO) AddRecord(userId string, record *Record) *HttpError {
	return nil
}
