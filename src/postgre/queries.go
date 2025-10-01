package postgre

import (
	"database/sql"
)

func GetUserId(db *sql.DB, login string) (id int, passwordHash string, err error) {
	result := db.QueryRow("SELECT id, password_hash FROM users WHERE login = $1", login)
	err = result.Scan(&id, &passwordHash)
	return id, passwordHash, err
}

func AddUser(db *sql.DB, login, pass_hash, totp_secret string) (err error) {
	_, err = db.Exec(
		"INSERT INTO users (login, password_hash, totp_secret) VALUES ($1, $2, $3);",
		login,
		pass_hash,
		totp_secret,
	)
	return err
}
