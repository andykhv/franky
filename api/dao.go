package franky

type FrankyDAO interface {
	GetUser(userId string) (*User, *HttpError)
	AddUser(user *User) *HttpError
	DeleteUser(userId string) *HttpError
	GetRecords(request *RecordRequest) ([]Record, *HttpError)
	AddRecord(userId string, record *Record) *HttpError
}
