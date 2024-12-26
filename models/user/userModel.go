package user_model

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func BodyToUser(body []byte) *User {
	if len(body) == 0 {
		return nil
	}

	var user User
	err := json.Unmarshal(body, &user)
	if err != nil {
		return nil
	}

	return &user
}

func GetAllUsers() []User {
	// mock data
	var users []User
	users = append(users, User{
		Id:       1,
		Email:    "dimdim@as.das",
		Password: "123",
	})
	users = append(users, User{
		Id:       2,
		Email:    "dimdim2@as.das",
		Password: "1234",
	})
	return users
}

const DB_CREATE_USER_STATEMENT = "INSERT INTO users (email, password) VALUES (?, ?)"

func CreateUser(
	userInput *http.Request,
	database *sql.DB,
) bool {
	body, err := io.ReadAll(userInput.Body)
	if err != nil {
		return false
	}
	defer userInput.Body.Close()
	user := BodyToUser(body)
	if user == nil {
		return false
	}
	_, err = database.Exec(
		DB_CREATE_USER_STATEMENT,
		user.Email,
		user.Password,
	)
	if err != nil {
		return false
	}

	return true
}

func GetUserByEmail(
	userEmail string,
	database *sql.DB,
) User {
	var user User
	decodedValue, err := url.QueryUnescape(userEmail)
	if err != nil {
		panic(err.Error())
	}
	prepared, err := database.Prepare("select id, email, password from users where email = ?")
	if err != nil {
		panic(err.Error())
	}
	defer prepared.Close()
	rows, err := prepared.Query(decodedValue)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.Id, &user.Email, &user.Password)
	}
	return user
}
