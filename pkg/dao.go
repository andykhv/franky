package franky

type FrankyDAO interface {
	GetUser() *User
	GetRecords() []Record
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

func (dao *testDAO) GetRecords() []Record {
	record := Record{"song", "artist", "album", "playlist", 180, 1000, "rap"}
	return []Record{record}
}
