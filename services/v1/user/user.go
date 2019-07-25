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
	"log"
	"strings"

)

import Model "./model"
import ServiceModel "../../model/"

func UserHandleRequest(res http.ResponseWriter, req *http.Request) {
	log.Println("Incoming User Request:", req.Method)
	path := strings.ReplaceAll(req.URL.Path[1:],"v1/user","")
	userId := req.URL.Query().Get("id")
	if len(userId) > 0 {
		log.Println("Incoming User Request userId:", userId)
	}

	switch req.Method {
	case http.MethodGet:
		if len(userId) > 0 {
			GetByUsername(userId, res, req)
		} else if path == "user" {
			Get(res, req)
		} else {
			resp := ServiceModel.NewResponse(res)
			resp.Json(http.StatusBadRequest, []byte("{}"), "User id is not found") 
		}
		break
	case http.MethodPost:
		Create(res, req)
		break
	case http.MethodPut:
		if len(userId) > 0 { 
			Update(userId, res, req)
		} else {
			resp := ServiceModel.NewResponse(res)
			resp.Json(http.StatusBadRequest, []byte("{}"), "User id is not found") 
		}
		break
	case http.MethodDelete:
		if  len(userId) > 0 { 
			Delete(userId, res, req)
		} else {
			resp := ServiceModel.NewResponse(res)
			resp.Json(http.StatusBadRequest, []byte("{}"), "User id is not found") 
		}
		break
	default:
		resp := ServiceModel.NewResponse(res)
		resp.Json(http.StatusServiceUnavailable, []byte("{}"), "Method not allowed") 
		break
	}
}


// Create function will help us to create a new user
func Create(res http.ResponseWriter, req *http.Request) {

	userObject := Model.CreateUser("vignesh", "kumar", "vigneshuvi", 1)
	responseObj, err := json.Marshal(userObject)
	if err != nil {
		fmt.Println(err)
		resp := ServiceModel.NewResponse(res)
		resp.Json(http.StatusBadRequest, responseObj, "successfully hit the CREATE user webservice.") 
		return
	}
	resp := ServiceModel.NewResponse(res)
	resp.Json(http.StatusOK, responseObj, "successfully hit the CREATE user webservice.") 
}


// Update function will help us to update a new user
func Update(userId string, res http.ResponseWriter, req *http.Request) {

	userObject := Model.CreateUser("vignesh", "kumar", userId, 1)
	responseObj, err := json.Marshal(userObject)
	if err != nil {
		fmt.Println(err)
		resp := ServiceModel.NewResponse(res)
		resp.Json(http.StatusBadRequest, responseObj, "successfully hit the Update user webservice.") 
		return
	}
	resp := ServiceModel.NewResponse(res)
	resp.Json(http.StatusOK, responseObj, "successfully hit the Update user webservice.") 
}

// Get function will help us to get the user by name
func Get(res http.ResponseWriter, req *http.Request) {

	userListObject := []Model.UserModel {Model.CreateUser("Vignesh", "kumar", "vigneshuvi", 2), Model.CreateUser("Vinoth", "kumar", "vinothkumar", 2),}
	responseObj, err := json.Marshal(userListObject)
	if err != nil {
		fmt.Println(err)
		resp := ServiceModel.NewResponse(res)
		resp.Json(http.StatusBadRequest, responseObj, "successfully hit the Get All user webservice.") 
		return
	}
	resp := ServiceModel.NewResponse(res)
	resp.Json(http.StatusOK, responseObj, "successfully hit the Get All webservice.") 
}

// GetByUsername function will help us to get the user by name
func GetByUsername(userId string,  res http.ResponseWriter, req *http.Request) {

	userObject := Model.CreateUser("vignesh", "kumar", userId, 2)
	responseObj, err := json.Marshal(userObject)
	if err != nil {
		fmt.Println(err)
		resp := ServiceModel.NewResponse(res)
		resp.Json(http.StatusBadRequest, responseObj, "successfully hit the GetByUsername user webservice.") 
		return
	}
	resp := ServiceModel.NewResponse(res)
	resp.Json(http.StatusOK, responseObj, "successfully hit the GetByUsername webservice.") 
}

// Delete function will help us to delete the user by name
func Delete(userId string,  res http.ResponseWriter, req *http.Request) {

	userObject := Model.CreateUser("vignesh", "kumar", userId, 2)
	responseObj, err := json.Marshal(userObject)
	if err != nil {
		fmt.Println(err)
		resp := ServiceModel.NewResponse(res)
		resp.Json(http.StatusBadRequest, responseObj, "successfully hit the Delete user webservice.") 
		return
	}
	resp := ServiceModel.NewResponse(res)
	resp.Json(http.StatusOK, responseObj, "successfully hit the Delete webservice.") 
}
