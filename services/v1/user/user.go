/*
Package user provides user web services for all request type of GET, POST, PUT, DELETE..

 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
*/
package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

import Model "./model"

// Create function will help us to create a new user
func Create(res http.ResponseWriter, req *http.Request) {

	userObject := Model.UserModel{Firstname: "vignesh", Lastname: "kumar", Username: "vigneshuvi"}
	resJSON := Model.StatusModel{Status: 200, Data: userObject, Message: "successfully hit the CREATE user webservice."}
	responseObj, err := json.Marshal(resJSON)
	if err != nil {
		fmt.Println(err)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(responseObj)
}

// GetByUsername function will help us to get the user by name
func GetByUsername(res http.ResponseWriter, req *http.Request) {

	userObject := Model.UserModel{Firstname: "vignesh", Lastname: "kumar", Username: "vigneshuvi"}
	responseObj, err := json.Marshal(userObject)
	if err != nil {
		fmt.Println(err)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(responseObj)
}
