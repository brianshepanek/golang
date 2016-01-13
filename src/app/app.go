package main

import (
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
