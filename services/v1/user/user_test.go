/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the services.
 */
package user

import "testing"

import (
	"net/http"
	"net/http/httptest"
)

func TestCreateUserHandler(t *testing.T) {
	t.Log("SERVICES: Testing the services for '/user/' URL path ")
	req, err := http.NewRequest("POST", "/user", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(UserHandleRequest)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"status":true,"data":{"firstname":"vignesh","lastname":"kumar","username":"vigneshuvi","roletype":1,"rolename":"admin"},"message":"successfully hit the CREATE user webservice."}`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestGetByUsernameHandler(t *testing.T) {
	t.Log("SERVICES: Testing the services for '/GetByUsername/' URL path ")
	req, err := http.NewRequest("GET", "/user?id=vigneshuvi", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(UserHandleRequest)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"status":true,"data":{"firstname":"vignesh","lastname":"kumar","username":"vigneshuvi","roletype":2,"rolename":"manager"},"message":"successfully hit the GetByUsername webservice."}`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestGetAllHandler(t *testing.T) {
	t.Log("SERVICES: Testing the services for '/user' URL path ")
	req, err := http.NewRequest("GET", "/user", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(UserHandleRequest)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"status":true,"data":[{"firstname":"Vignesh","lastname":"kumar","username":"vigneshuvi","roletype":2,"rolename":"manager"},{"firstname":"Vinoth","lastname":"kumar","username":"vinothkumar","roletype":2,"rolename":"manager"}],"message":"successfully hit the Get All webservice."}`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestPutByUsernameHandler(t *testing.T) {
	t.Log("SERVICES: Testing the services for '/User' Put URL path ")
	req, err := http.NewRequest("PUT", "/user?id=vigneshuvi", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(UserHandleRequest)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"status":true,"data":{"firstname":"vignesh","lastname":"kumar","username":"vigneshuvi","roletype":1,"rolename":"admin"},"message":"successfully hit the Update user webservice."}`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestDeleteByUsernameHandler(t *testing.T) {
	t.Log("SERVICES: Testing the services for '/User' DELETE URL path ")
	req, err := http.NewRequest("DELETE", "/user?id=vigneshuvi", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(UserHandleRequest)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"status":true,"data":{"firstname":"vignesh","lastname":"kumar","username":"vigneshuvi","roletype":2,"rolename":"manager"},"message":"successfully hit the Delete webservice."}`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}