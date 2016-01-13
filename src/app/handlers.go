package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func UsersIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Users Index Yeah!")
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