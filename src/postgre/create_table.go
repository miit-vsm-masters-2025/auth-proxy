package postgre

import "database/sql"

func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id TEXT PRIMARY KEY," +
		"login TEXT," +
		"password_hash VARCHAR(256)" +
		"totp_secret TEXT " +
		"created_at TIMESTAMP default current_timestamp);")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE INDEX index_name ON users(login);")
	if err != nil {
		panic(err)
	}
}
