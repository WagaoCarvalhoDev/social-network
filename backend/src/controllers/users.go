package controllers

import (
	dbmysql "backend/src/db_mysql"
	"backend/src/models"
	"backend/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := dbmysql.ConnectMySql()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	repository.CreateUser(user)

}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding all users!"))
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding user!"))
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Removing user!"))
}
