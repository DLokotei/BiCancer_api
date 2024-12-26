package main

import (
	"database/sql"
	"devocean/bicancer/db_sql"
	user_model "devocean/bicancer/models/user"
	"devocean/bicancer/router"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const SERVER_PORT string = ":3690"

func main() {
	fmt.Println("Starting of code!")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading the .env file.")
	}

	// Open connection to Database
	database, err := sql.Open("mysql", db_sql.GetDSN())
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()

	// Check success connection to Database
	err = database.Ping()
	if err != nil {
		// problem with configuration or mySql server
		panic(err.Error())
	}

	router.Setup(database)

	// Define handlers to endpoints
	http.HandleFunc("/all_users", func(w http.ResponseWriter, r *http.Request) {
		allUsers := user_model.GetAllUsers()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allUsers)
	})

	println("Server listening on port" + SERVER_PORT + " ...")
	http.ListenAndServe(SERVER_PORT, nil)

}
