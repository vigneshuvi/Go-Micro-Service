/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the services.
 */
package services

import (
	"fmt"
	"encoding/json"
	"net/http"
	"sync"
)

import model "./model"


var serviceProfile *model.ServiceProfile
var once sync.Once

// Help to set Config.
func getServiceProfile() *model.ServiceProfile {
	once.Do(func() {
		serviceProfile = &model.ServiceProfile{Env: "DEV", SeverStatus: false, ServiceCount: 0}
	})
	return serviceProfile
}

func SetServiceEnv(env string) {
	serviceProf := getServiceProfile()
	serviceProf.Env = env
}

func SetProfileStatus(status bool) {
	serviceProf := getServiceProfile()
	serviceProf.SeverStatus = status
}

func SetServiceCount(count int) {
	serviceProf := getServiceProfile()
	serviceProf.ServiceCount = count
}

func IncreseServiceCount() {
	serviceProf := getServiceProfile()
	serviceProf.ServiceCount = serviceProf.ServiceCount  + 1
}

func ResetServiceProfile(count int) {
	serviceProf := getServiceProfile()
	serviceProf.SeverStatus = false
	serviceProf.ServiceCount = 0
}

func GetProfileService() *model.ServiceProfile {
	return getServiceProfile();
}

// Handler for / Web services.
func Handler(res http.ResponseWriter, req *http.Request) {
	resp := model.NewResponse(res)
	if len(req.URL.Path[1:]) > 0 {
		resp.Text(http.StatusNotFound,  "Not Found!") 
		return
	} 
	resp.Text(http.StatusOK,  "Welcome to Go Micro Services!") 
}

// Handler for Service Profile Web services.
func GetProfileHandler(res http.ResponseWriter, req *http.Request) {
	responseObj, err := json.Marshal(GetProfileService())
	if err != nil {
		fmt.Println(err)
		resp := model.NewResponse(res)
		resp.Json(http.StatusNotFound, responseObj, "successfully hit the GetProfileHandler webservice.")
		return
	}
	resp := model.NewResponse(res)
	resp.Json(http.StatusOK, responseObj, "successfully hit the GetProfileHandler webservice.")
}

