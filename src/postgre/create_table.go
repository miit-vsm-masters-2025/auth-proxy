package postgre

import "database/sql"

func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY," +
		"login VARCHAR(512)," +
		"password_hash VARCHAR(256)," +
		"totp_secret VARCHAR(256) ," +
		"created_at TIMESTAMP default current_timestamp);")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE INDEX IF NOT EXISTS index_name ON users(password_hash, login);")
	if err != nil {
		panic(err)
	}
}
