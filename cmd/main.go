package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"limble/handlers"
)

func main() {
	db := setUpDb()
	defer db.Close()

	http.HandleFunc("/cost/by-worker", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleWorkerCost(w, r, db)
	})
	http.HandleFunc("/cost/by-location", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleLocationCost(w, r, db)
	})

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setUpDb() *sql.DB {
	// username and password both root for now
	// TODO: use encrypted credentials from k8s or AWS
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	// health check
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	} else {
		fmt.Println("Successfully connected to the database!")
	}

	return db
}