/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Response object model.
 */

package model


import "testing"

import (
	"net/http"
	"net/http/httptest"
)


func TestJSONBadRequestResponse(t *testing.T) { 
    t.Log("SERVICES: Testing the JSON Response Model")

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    resp := NewResponse(rr)
    data := []byte(`{ "votes": { "option_A": "3" } }`)
	resp.Json(http.StatusBadRequest, data, "successfully tested the response model.") 

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"status":false,"data":{"votes":{"option_A":"3"}},"message":"successfully tested the response model."}`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestJSONSuccessResponse(t *testing.T) { 
    t.Log("SERVICES: Testing the JSON Response Model")

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    resp := NewResponse(rr)
    data := []byte(`{ "votes": { "option_A": "3" } }`)
	resp.Json(http.StatusOK, data, "successfully tested the response model.") 

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"status":true,"data":{"votes":{"option_A":"3"}},"message":"successfully tested the response model."}`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}


func TestJSONFailResponse(t *testing.T) { 
    t.Log("SERVICES: Testing the JSON Response Model")

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    resp := NewResponse(rr)
	resp.Json(http.StatusOK, nil, "successfully tested the response model.") 

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }
}

func TestTextResponse(t *testing.T) { 
    t.Log("SERVICES: Testing the Text Response Model")

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    resp := NewResponse(rr)
    if resp.Writer == nil  {
		t.Error("Register Service Profile Count: application/json")
    }
    
    resp.Text(http.StatusOK,  "successfully tested the response model.") 
    
    if resp.Writer.Header().Get("Content-Type") == "application/json"  {
		t.Error("Register Service Profile Count: application/json")
	}

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("SERVICES ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `successfully tested the response model.`
    if rr.Body.String() != expected {
        t.Errorf("SERVICES ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}