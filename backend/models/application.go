package models

import (
	"backend/utils"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//ApplicationRequest represents the incoming application request
type ApplicationRequest struct {
	Name       string
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

//ApplicationResponse represents the outgoing application payload
type ApplicationResponse struct {
	Name       string
	Blurb      string
	Category   string
	MyFileName string
	AvatarName string
	Website    string
	Github     string
	Platform   string
	Tools      []Tool
	Languages  []Language
}

//Application struct represents one application object that will be saved
type Application struct {
	gorm.Model
	Name       string
	Search     string
	Blurb      string
	Category   string
	MyFileName string
	AvatarName string
	Website    string
	Github     string
	Platform   string
	Tools      pq.ByteaArray `gorm:"type:varchar(1000)[]"`
	Languages  pq.ByteaArray `gorm:"type:varchar(1000)[]"`
}

//Tool struct
type Tool struct {
	Name string
}

//Language struct
type Language struct {
	Name string
}

//Validate if the Application payload is valid
func (app *ApplicationRequest) Validate() (map[string]interface{}, bool) {
	if len(app.Name) == 0 {
		return utils.Message(false, "Application must have a name"), false
	}
	if len(app.Blurb) == 0 {
		return utils.Message(false, "Application must have a blurb"), false
	}
	if len(app.Category) == 0 {
		return utils.Message(false, "Application must have a Category"), false
	}
	return utils.Message(true, "Application Requirement passed"), true

}

//Create an Application Object which can be saved from the db.
func (app *ApplicationRequest) Create() map[string]interface{} {
	validateResp, ok := app.Validate()
	if !ok {
		return validateResp
	}
	nameHash := hash(app.Name)
	application := &Application{}
	//Check if the application already exists or not
	err := GetDB().Table("applications").Where("search = ?", nameHash).First(application).Error
	if err == nil {
		//Application already exists
		return utils.Message(false, "Application with same name already exists")
	}
	//We'll need to convert ApplicatonPayload to an Appllication struct
	application = &Application{
		Name:       app.Name,
		Search:     nameHash,
		Blurb:      app.Blurb,
		Category:   app.Category,
		MyFileName: app.MyFileName,
		AvatarName: app.AvatarName,
		Website:    app.Website,
		Github:     app.Github,
		Platform:   app.Platform,
	}

	for i := range app.Tools {
		tool := &Tool{Name: app.Tools[i]}
		rawData, _ := json.Marshal(tool)
		application.Tools = append(application.Tools, rawData)
	}
	for k := range app.Languages {
		language := &Language{Name: app.Languages[k]}
		rawData, _ := json.Marshal(language)
		application.Languages = append(application.Languages, rawData)
	}
	GetDB().Create(application)
	resp := utils.Message(true, "Succesfully Created new Application and saved into DB")
	return resp
}

//GetApplications returns the Apllication struct when we put a post to it
func GetApplications() []ApplicationResponse {
	//Get All Records
	applicationList := &[]Application{}
	GetDB().Find(applicationList)
	//dereference it to get the array itself
	list := *applicationList
	applicationPayloadList := []ApplicationResponse{}
	for i := range list {
		application := list[i]
		applicationPayload := GetPayload(application)
		applicationPayloadList = append(applicationPayloadList, applicationPayload)
	}
	return applicationPayloadList
}

//GetPayload returns the ApplicationPayload when the Application struct is provided
func GetPayload(app Application) ApplicationResponse {
	appPayload := &ApplicationResponse{
		Name:       app.Name,
		Blurb:      app.Blurb,
		Category:   app.Category,
		MyFileName: app.MyFileName,
		AvatarName: app.AvatarName,
		Website:    app.Website,
		Github:     app.Github,
		Platform:   app.Platform,
		Tools:      getToolArray(app.Tools),
		Languages:  getLanguageArray(app.Languages),
	}
	//I know that byteArrays only have 1 byte
	return *appPayload
}

func hash(name string) string {
	h := sha256.New()
	h.Write([]byte(name))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash
}
func toByte(byteArray pq.ByteaArray) []byte {
	return byteArray[0]
}

func getToolArray(toolArray pq.ByteaArray) []Tool {
	tools := []Tool{}
	for _, item := range toolArray {
		tool := &Tool{}
		json.Unmarshal(item, tool)
		tools = append(tools, *tool)
	}
	return tools
}
func getLanguageArray(languageArray pq.ByteaArray) []Language {
	languages := []Language{}
	for _, item := range languageArray {
		language := &Language{}
		json.Unmarshal(item, language)
		languages = append(languages, *language)
	}
	return languages
}
