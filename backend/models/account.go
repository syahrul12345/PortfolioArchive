package models

import (
	"backend/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Token Comment
type Token struct {
	UserID   uint
	UserName string
	jwt.StandardClaims
}

//Account Comment
type Account struct {
	gorm.Model
	Email    string
	Password string
	Token    string
}

//NewAccount when creating
type NewAccount struct {
	Email       string
	Password    string
	NewPassword string
}

//Validate comment
func (acc *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(acc.Email, "@") {
		return utils.Message(false, "Email Address Required"), false
	}
	if len(acc.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}
	//Email must be unique
	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", acc.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection Error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email already in use"), false
	}
	return utils.Message(false, "Requirement passed"), true

}

//Validate the new account object
func (newAcc *NewAccount) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(newAcc.Email, "@") {
		return utils.Message(false, "Email Address Required"), false
	}
	if len(newAcc.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}
	if len(newAcc.NewPassword) < 6 {
		return utils.Message(false, "New Password is required"), false
	}
	//Email must be unique
	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", newAcc.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection Error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email Exists"), true
	}

	return utils.Message(false, "No such email"), false
}

//Create account
func (acc *Account) Create() map[string]interface{} {
	response, ok := acc.Validate()
	if !ok {
		return response
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	acc.Password = string(hashedPassword)
	//stores the account into the database
	GetDB().Create(acc)

	if acc.ID <= 0 {
		return utils.Message(false, "Failed to create account")
	}
	//create a new JWT token
	tk := &Token{UserID: acc.ID, UserName: acc.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	acc.Token = tokenString
	acc.Password = ""

	response = utils.Message(true, "Account has been created")
	response["account"] = acc
	return response
}

//Save will save the new account
func (newAcc *NewAccount) Save() map[string]interface{} {
	response, ok := newAcc.Validate()
	if !ok {
		return response
	}
	account := &Account{}
	err := GetDB().Table("accounts").Where("email =?", newAcc.Email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email")
		}
		return utils.Message(false, "Connection error. Please retry")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(newAcc.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Invalid Login Credentials")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newAcc.NewPassword), bcrypt.DefaultCost)
	GetDB().Model(account).Update("Password", string(hashedPassword))
	account.Password = ""

	//create JWT TOKEN
	tk := &Token{UserID: account.ID, UserName: account.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	resp := utils.Message(true, "Succesfully logged in")
	resp["account"] = account
	return resp
}

//Login comment
func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email")
		}
		return utils.Message(false, "Connection error. Please retry")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Invalid Login Credentials")
	}
	//worked, logged in
	account.Password = ""
	//create JWT token
	tk := &Token{UserID: account.ID, UserName: account.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	resp := utils.Message(true, "Succesfully logged in")
	resp["account"] = account
	return resp
}

//GetUser comment
func GetUser(u uint) *Account {
	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}
	acc.Password = ""
	return acc
}
