package postgre

import "database/sql"

func GetUserId(db *sql.DB, login, pass_hash string) (id int, err error) {
	result := db.QueryRow("SELECT id FROM users WHERE login = ? AND password_hash = ?", login, pass_hash)
	err = result.Scan(&id)
	return id, err
}
