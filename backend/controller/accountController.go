package controller

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"net/http"
)

//CreateAccount will create an account for the CMS
var CreateAccount = func(writer http.ResponseWriter, request *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(request.Body).Decode(account)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Invalid Request"))
	} else {
		resp := account.Create()
		utils.Respond(writer, resp)
	}
}

//Login to the cms
var Login = func(writer http.ResponseWriter, request *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(request.Body).Decode(account)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Invalid Request"))
	} else {
		resp := models.Login(account.Email, account.Password)
		utils.Respond(writer, resp)
	}
}

// ChangePassword of the account
var ChangePassword = func(writer http.ResponseWriter, request *http.Request) {
	newAccount := &models.NewAccount{}
	err := json.NewDecoder(request.Body).Decode(newAccount)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Invalid Request, please provide new password"))
	} else {
		resp := newAccount.Save()
		utils.Respond(writer, resp)
	}
}
