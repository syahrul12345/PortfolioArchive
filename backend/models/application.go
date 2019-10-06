package models

import (
	"backend/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Application struct represents one application object
type Application struct {
	gorm.Model
	Name       string
	NameHash   string
	Blurb      string
	Category   string
	MyFileName string
	AvatarName string
	Website    string
	Github     string
	Platform   string
	Tools      []string
	Languages  []string
}

//Validate if the Application payload is valid
func (app *Application) Validate() (map[string]interface{}, bool) {
	if len(app.Name) == 0 {
		return utils.Message(false, "Application must have a name"), false
	}
	if len(app.Blurb) == 0 {
		return utils.Message(false, "Application must have a blurb"), false
	}
	if len(app.Category) == 0 {
		return utils.Message(false, "Application must have a Category"), false
	}
	fmt.Println(app)
	return utils.Message(true, "Application Requirement passed"), true
	//Check if it already exists

}

//Create a new application in the db
func (app *Application) Create() map[string]interface{} {
	resp, ok := app.Validate()
	if !ok {
		return resp
	}
	return resp
}
