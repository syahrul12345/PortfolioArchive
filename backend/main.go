package main

import (
	"backend/app"
	"backend/controller"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sevlyar/go-daemon"
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}

	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")
	serve()
}

func serve() {
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
		log.Print("web page loaded")
		const staticPath = "../dist"
		const indexPath = "index.html"
		fileServer := http.FileServer(http.Dir(staticPath))
		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			log.Print(err)
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
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileServer.ServeHTTP(w, r)
	}))
	port := "5556"
	log.Print("The website can be served at http://localhost:5556")
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
