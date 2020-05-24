package github.com/paulogil93/dbconnector

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	//Import to help with postgres driver
	"github.com/jackc/pgx/v4"
)

func dbConn() (db *pgx.Conn) {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := pgx.Connect(context.Background(), dbURL)

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
	Category int    `json:"category"`
	NMec     int    `json:"nmec"`
	Email    string `json:"email"`
	Area     int    `json:"area"`
}

// AddUser function: returns true if POST succeded, else false
func AddUser(id int, email string, name string, initials string, category int, nmec int, area int) bool {
	_, err := dbConn().Query(context.Background(), GetAddUserSQLCmd(id, email, name, initials, category, nmec, area))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't connect to database: %v\n", err)
		return false
	}
	return true
}

// ShowUsers bla bla
func ShowUsers() []byte {
	var user User
	var arr []User

	rows, _ := dbConn().Query(context.Background(), "SELECT * FROM pei.showUsers();")

	for rows.Next() {
		rows.Scan(&user.ID, &user.Email, &user.Name, &user.Initials, &user.Category, &user.NMec, &user.Area)
		arr = append(arr, user)
	}

	j, _ := json.Marshal(arr)
	return j
}

// ShowUserByID function
func ShowUserByID(id int) User {
	var user User

	row := dbConn().QueryRow(context.Background(), "SELECT * FROM pei.showUserByID("+strconv.Itoa(id)+");")
	row.Scan(&user.ID, &user.Email, &user.Name, &user.Initials, &user.Category, &user.NMec, &user.Area)

	return user
}

// CreateCategory function
func CreateCategory(name string) bool {
	_, err := dbConn().Query(context.Background(), "CALL pei.AddCategory("+name+");")
	if err != nil {
		return false
	}
	return true
}

// AddNotification function
func AddNotification(sender int, receiver int, title string, body string) bool {
	fmt.Println(GetAddNotificationSQLCmd(sender, receiver, title, body))
	_, err := dbConn().Query(context.Background(), GetAddNotificationSQLCmd(sender, receiver, title, body))
	if err != nil {
		return false
	}
	return true
}

// GetAddUserSQLCmd returns the correct SQL command for the AddUSer Stored Procedure
func GetAddUserSQLCmd(id int, email string, name string, initials string, category int, nmec int, area int) string {
	return "CALL pei.addUser(" + strconv.Itoa(id) + "," + "'" + email + "'" + "," + "'" + name + "'" +
		"," + "'" + initials + "'" + "," + strconv.Itoa(category) + "," + strconv.Itoa(nmec) + "," + strconv.Itoa(area) + ")"
}

// GetAddNotificationSQLCmd returns the correct SQL command for the AddNotification Stored Procedure
func GetAddNotificationSQLCmd(sender int, receiver int, title string, body string) string {
	return "CALL pei.addNotification(" + strconv.Itoa(sender) + "," + strconv.Itoa(receiver) + "," + "'" + title + "'" +
		"," + "'" + body + "'" + ")"
}
