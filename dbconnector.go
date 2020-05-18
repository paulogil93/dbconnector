package dbconnector

import (
	"database/sql"
	"encoding/json"

	//_ "github.com/mattn/go-sqlite3"
)

func dbConn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		panic(err.Error())
	}

	return db
}

// User struct to handle teachers data
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Initials string `json:"initials"`
	Category string `json:"category"`
}

// AddUser bla bla
func AddUser(id int, name string, initials string, category string) {
	statement, _ := dbConn().Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, initials TEXT, category TEXT)")
	statement.Exec()
	statement, _ = dbConn().Prepare("INSERT INTO users(id, name, initials, category) VALUES (?, ?, ?, ?)")
	statement.Exec(id, name, initials, category)
}

// ShowUsers bla bla
func ShowUsers() []byte {
	var user User
	var arr []User

	rows, _ := dbConn().Query("SELECT * FROM users;")

	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Initials, &user.Category)
		arr = append(arr, user)
	}

	j, _ := json.Marshal(arr)
	return j

}
