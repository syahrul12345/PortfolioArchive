package controller

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"net/http"
)

//CreateApplication will create a new application for the Dapp
var CreateApplication = func(writer http.ResponseWriter, request *http.Request) {
	application := &models.Application{}
	err := json.NewDecoder(request.Body).Decode(application)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Invalid request"))
	} else {
		resp := application.Create()
		utils.Respond(writer, resp)
	}

}
