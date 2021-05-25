package franky

type FrankyDAO interface {
	GetUser(userId string) (*User, *HttpError)
	AddUser(user *User) *HttpError
	DeleteUser(userId string) *HttpError
	GetRecords() ([]Record, *HttpError)
	AddRecord(userId string, record *Record) *HttpError
}
