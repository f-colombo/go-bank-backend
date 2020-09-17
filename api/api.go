package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"duomly.com/go-bank-backend/helpers"
	"duomly.com/go-bank-backend/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type ErrResponse struct {
	Message string
}

func Users_Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// login := users.StrLogin(
	login := users.Login(
		formattedBody.Username,
		formattedBody.Password)

	if login["message"] == "All is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		// resp := ErrResponse{Message: login["message"].(string)}
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func Users_Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	// register := users.StrRegister(
	register := users.Register(
		formattedBody.Username,
		formattedBody.Email,
		formattedBody.Password)

	if register["message"] == "All is fine" {
		resp := register
		json.NewEncoder(w).Encode(resp)
	} else {
		// resp := ErrResponse{Message: register["message"].(string)}
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi() {
	router := mux.NewRouter()

	router.HandleFunc("/users/login", Users_Login).Methods("POST")
	router.HandleFunc("/users/register", Users_Register).Methods("POST")

	fmt.Println("App is working on port: 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
