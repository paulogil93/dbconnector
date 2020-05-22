package dbconnector

import (
	"database/sql"
	"encoding/json"
	"os"
)

// User struct to handle teachers data
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Initials string `json:"initials"`
	Category string `json:"category"`
	NMec     int    `json:"nmec"`
	Email    string `json:"email"`
	Area     string `json:"area"`
}

func dbConn() (db *sql.DB) {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		panic(err.Error())
	}

	return db
}

// // AddUser bla bla
// func AddUser(id int, name string, initials string, category string) {
// 	statement, _ := dbConn().Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, initials TEXT, category TEXT)")
// 	statement.Exec()
// 	statement, _ = dbConn().Prepare("INSERT INTO users(id, name, initials, category) VALUES (?, ?, ?, ?)")
// 	statement.Exec(id, name, initials, category)
// }

//AddUser function
func AddUser(id int, email string, name string, initials string, category int, nmec int, area int) {
	stmt, _ := dbConn().Prepare("CALL pei.addUser(?, ?, ?, ?, ?, ?, ?)")
	stmt.Exec(id, email, name, initials, category, nmec, area)
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
