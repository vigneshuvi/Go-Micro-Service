/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the services.
 */
package services

import "testing"

import (
	"net/http"
	"net/http/httptest"
)
import model "./model"


func TestProfile(t *testing.T) { 
    t.Log("SERVICES: Testing the profile singleton object")
    SetServiceEnv("DEV")
    profile := GetProfileService()

    if &(model.ServiceProfile{}) == profile {
		t.Error("MAIN ERROR: Profile is zero value")
    }
    if profile.Env != "DEV" {
		t.Error("MAIN ERROR: Expected environment is `DEV`")
    }

    SetProfileStatus(true)

    if profile.SeverStatus != true {
		t.Error("MAIN ERROR: HTTP server is not yet started")
    }
    
    SetServiceCount(2)

    if profile.ServiceCount != 2 {
		t.Error("MAIN ERROR: Expected environment is `2`")
	}

}


func TestHandler(t *testing.T) {
	t.Log("SERVICES: Testing the services for '/' URL path ")
	req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Handler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `Welcome to Go Micro Services!`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestGetProfileHandler(t *testing.T) {
	t.Log("SERVICES: Testing the services for '/getProfile' URL path ")
	req, err := http.NewRequest("GET", "/getProfile", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetProfileHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"status":true,"data":{"env":"DEV","service_count":2,"server_status":true},"message":"successfully hit the GetProfileHandler webservice."}`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}