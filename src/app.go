package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)

    //Routes
    router.HandleFunc("/", Index)
    router.HandleFunc("/users.json", UsersIndex).Methods("GET")
    router.HandleFunc("/users.json", UsersAdd).Methods("POST")
    router.HandleFunc("/users/{user_id}.json", UsersView).Methods("GET")
    router.HandleFunc("/users/{user_id}.json", UsersEdit).Methods("PUT", "PATCH")
    router.HandleFunc("/users/{user_id}.json", UsersDelete).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func UsersIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Users Index!")
}

func UsersAdd(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Users Add!")
}

func UsersView(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user_id := vars["user_id"]
    fmt.Fprintln(w, "Show User:", user_id)
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