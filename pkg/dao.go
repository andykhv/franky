package franky

type FrankyDAO interface {
	GetUser() *User
	AddUser(user *User)
	DeleteUser(userId string)
	GetRecords() []Record
	AddRecord(userId string, record *Record)
}

type testDAO struct {
}

func TestDAO() FrankyDAO {
	return &testDAO{}
}

func (dao *testDAO) GetUser() *User {
	user := User{"id", "email", "password", "apiKey", "creationDate"}
	return &user
}

func (dao *testDAO) AddUser(user *User) {
}

func (dao *testDAO) DeleteUser(userId string) {
}

func (dao *testDAO) GetRecords() []Record {
	record := Record{"song", "artist", "album", "playlist", 180, 1000, "rap"}
	return []Record{record, record}
}

func (dao *testDAO) AddRecord(userId string, record *Record) {
}
