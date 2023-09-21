package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	db, err := sql.Open("sqlite3", "dbapp.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.ListenAndServe(":8000", r)
}

