package main

import (
    "fmt"
    "log"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "database/sql"
    _ "github.com/lib/pq"
)

//Define User Structure and json Return Keys
type User struct {
	Id int `json:"id"`
	UserName string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}

//Define Users Slice
type Users []User

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func DbConnect() *sql.DB {

	//Connect to DB
	db, err := sql.Open("postgres", "host=postgres-server user=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Find() []User {

	db := DbConnect()

	//Select Query on DB
	rows, err := db.Query("SELECT id, username, first_name, last_name, email FROM users")

	if err != nil {
		log.Fatal(err)
	}

	//Local users Slice
	var users []User

	//Loop through Results, Append to users Slice
	for rows.Next() {
		var user User
        rows.Scan(&user.Id, &user.UserName, &user.FirstName, &user.LastName, &user.Email)
        users = append(users, user)
    }

    return users
}

func FindOne(user_id int) User {

	db := DbConnect()

	//Select Query on DB
	rows, err := db.Query("SELECT id, username, first_name, last_name, email FROM users WHERE id = $1", user_id)

	if err != nil {
		log.Fatal(err)
	}

	//Local users Slice
	var user User

	//Loop through Results, Append to users Slice
	for rows.Next() {
		rows.Scan(&user.Id, &user.UserName, &user.FirstName, &user.LastName, &user.Email)
    }

    return user
}

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	
	//Users
	users := Find()

    //Set Content-Type
    w.Header().Set("Content-Type", "application/json")

    //Encode Results as json
    json.NewEncoder(w).Encode(users)
}

func UsersAdd(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Users Add!")
}

func UsersView(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user_id_string := vars["user_id"]
    user_id, err := strconv.Atoi(user_id_string)
    if err != nil {
		log.Fatal(err)
	}
    user := FindOne(user_id)

    //Set Content-Type
    w.Header().Set("Content-Type", "application/json")

    //Encode Results as json
    json.NewEncoder(w).Encode(user)
}

func UsersEdit(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user_id := vars["user_id"]
    fmt.Fprintln(w, "Edit User:", user_id)
}

func UsersDelete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user_id := vars["user_id"]
    fmt.Fprintln(w, "Delete User:", user_id)
}