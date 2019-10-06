package main

import (
	"backend/app"
	"backend/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	//Create an Account
	router.HandleFunc("/api/user/new", controller.CreateAccount).Methods("POST")
	//Login to account
	router.HandleFunc("/api/user/login", controller.Login).Methods("POST")
	//Change Password
	router.HandleFunc("/api/user/changePassword", controller.ChangePassword).Methods("POST")
	// Adds in a new application
	router.HandleFunc("/api/user/createApp", controller.CreateApplication).Methods("POST")
	//Get All Applications
	router.HandleFunc("/api/user/getAllApps", controller.GetAllApps).Methods("GET")
	//Serve the compiled vueJS frontend
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../dist/")))
	port := "5556"
	fmt.Println("The website can be served at http://localhost:5556")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
