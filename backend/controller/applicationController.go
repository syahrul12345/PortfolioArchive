package controller

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//CreateApplication will create a new application for the Dapp
var CreateApplication = func(writer http.ResponseWriter, request *http.Request) {
	application := &models.ApplicationPayload{}
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
	applications := models.GetApplications()
	resp := utils.Message(false, "Get all application data")
	resp["applications"] = applications
	utils.Respond(writer, resp)
}
