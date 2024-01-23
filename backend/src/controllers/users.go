package controllers

import (
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user!"))

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
