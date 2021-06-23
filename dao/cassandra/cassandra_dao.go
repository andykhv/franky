package cassandra

// TODO:
// - this DAO implementation is incomplete

import (
	"net/http"
	"time"

	franky "github.com/andykhv/franky/api"
	"github.com/gocql/gocql"
)

type CassandraDAO struct {
	usersSession   *gocql.Session
	recordsSession *gocql.Session
}

func NewCassandraDAO(hosts ...string) (*CassandraDAO, *franky.HttpError) {
	usersSession, err := initUsersKeyspaceSession(hosts...)
	if err != nil {
		return nil, &franky.HttpError{StatusCode: http.StatusBadGateway, Err: err}
	}

	recordsSession, err := initRecordsKeyspaceSession(hosts...)
	if err != nil {
		return nil, &franky.HttpError{StatusCode: http.StatusBadGateway, Err: err}
	}

	return &CassandraDAO{usersSession, recordsSession}, nil
}

func (dao *CassandraDAO) GetUser(email string) (*franky.User, *franky.HttpError) {
	var user map[string]interface{}

	dao.usersSession.Query(getUserStatement, email).MapScan(user)

	frankyUser := franky.User{
		Id:           user["id"].(string),
		Email:        user["email"].(string),
		Password:     user["password"].(string),
		ApiKey:       user["api_key"].(string),
		CreationDate: user["creation_date"].(time.Time).Unix(),
	}

	return &frankyUser, nil
}

func (dao *CassandraDAO) AddUser(user *franky.User) *franky.HttpError {
	err := dao.usersSession.Query(
		addUserStatement,
		user.Email,
		user.Password,
		user.ApiKey,
		user.CreationDate,
		user.Id).Exec()

	if err != nil {
		return &franky.HttpError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	return nil
}

func (dao *CassandraDAO) DeleteUser(email string) *franky.HttpError {
	err := dao.usersSession.Query(deleteUserStatement, email).Exec()

	if err != nil {
		return &franky.HttpError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	return nil
}

func (dao *CassandraDAO) GetRecords() ([]franky.Record, *franky.HttpError)                 {}
func (dao *CassandraDAO) AddRecord(userId string, record *franky.Record) *franky.HttpError {}
func (dao *CassandraDAO) exit() {
	dao.usersSession.Close()
	dao.recordsSession.Close()
}

func initUsersKeyspaceSession(hosts ...string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = usersKeyspace
	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	return session, nil
}

func initRecordsKeyspaceSession(hosts ...string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = recordsKeyspace
	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	return session, nil
}
