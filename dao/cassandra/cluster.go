package cassandra

const (
	usersKeyspace       = "users"
	recordsKeyspace     = "records"
	usersTable          = "users_by_email"
	getUserStatement    = `select * from users_by_email where email=? LIMIT 1`
	addUserStatement    = `insert into users_by_email (email, password, api_key, creation_date, id) values (?, ?, ?, ?, ?)`
	deleteUserStatement = `delete from users_by_email where email=?`
)
