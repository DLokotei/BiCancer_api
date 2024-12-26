package router

import (
	"database/sql"
	user_model "devocean/bicancer/models/user"
	"encoding/json"
	"net/http"
)

type User2 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Setup(db *sql.DB) {
	http.HandleFunc("/test", handleTest)
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handleUser(w, r, db)
	})
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	qer := r.URL.Query()
	id := qer.Get("id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func handleUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodGet {
		userEmail := r.URL.Query().Get("email")
		if len(userEmail) != 0 {
			// requested data of single user
			user := user_model.GetUserByEmail(userEmail, db)
			returnJsonOkResponce(w, user)
			// returnJsonOkResponce(w, userEmail)
			return
		}
		allUsers := user_model.GetAllUsers()
		returnJsonOkResponce(w, allUsers)
		return
	}
	if r.Method == http.MethodPost {
		allUsers := user_model.GetAllUsers()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allUsers)
		return
	}
}

func returnJsonOkResponce(w http.ResponseWriter, jsonData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonData)
}
