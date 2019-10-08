package controller

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//CreateApplication will create a new application for the Dapp
var CreateApplication = func(writer http.ResponseWriter, request *http.Request) {
	log.Print("New Application Created")
	application := &models.ApplicationRequest{}
	err := json.NewDecoder(request.Body).Decode(application)
	if err != nil {
		fmt.Println(err)
		utils.Respond(writer, utils.Message(false, "Invalid request"))
	} else {
		resp := application.Create()
		utils.Respond(writer, resp)
	}

}

//GetAllApps returns all the apps saved in the db
var GetAllApps = func(writer http.ResponseWriter, request *http.Request) {
	log.Print("Get all Apps")
	applications := models.GetApplications()
	resp := utils.Message(false, "Get all application data")
	resp["applications"] = applications
	utils.Respond(writer, resp)
}
