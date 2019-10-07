package main

import (
	"backend/app"
	"backend/controller"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const staticPath = "../dist"
		const indexPath = "index.html"
		fileServer := http.FileServer(http.Dir(staticPath))
		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		path = filepath.Join(staticPath, path)
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			// file does not exist, serve index.html
			http.ServeFile(w, r, filepath.Join(staticPath, indexPath))
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileServer.ServeHTTP(w, r)
	}))
	port := "5556"
	fmt.Println("The website can be served at http://localhost:5556")
	//For Testing
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Println(err)
	}
}
