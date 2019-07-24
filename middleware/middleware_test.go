/*
Package middleware provides help to check and bypass the micro service requests.

 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
*/
package middleware

import (
	"net/http"
    "net/http/httptest"
	"testing"
)

import services "../services"


// Test Authentication function contains the basic validation for Auth token and header format.
func TestAuthentication(t *testing.T) {
	t.Log("MIDDLEWARE: Testing the Authenticate Middleware")

	// test stuff here...
	req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
	//req.Header.Set("xdate", "19/07/2019")

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
	handler := authentication(http.HandlerFunc(services.Handler))

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusUnauthorized {
        t.Errorf("MIDDLEWARE ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusUnauthorized)
    }

    // Check the response body is what we expect.
    expected := `Unauthorized`
    if rr.Body.String() != expected {
        t.Errorf("MIDDLEWARE ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
	}
}

// Test Authentication function contains the basic validation for Auth token and header format.
func TestAuthenticationHeader(t *testing.T) {
	t.Log("MIDDLEWARE: Testing the Authenticate Middleware")

	// test stuff here...
	req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
	req.Header.Set("xdate", "19/07/2019")

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
	handler := authentication(http.HandlerFunc(services.Handler))

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("MIDDLEWARE ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `Welcome to Go Micro Services!`
    if rr.Body.String() != expected {
        t.Errorf("MIDDLEWARE ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
	}
}

// Test ByPass function contains the basic validation for Auth token and header format.
func TestByPass(t *testing.T) {
	t.Log("MIDDLEWARE: Testing the ByPass Middleware")

	// test stuff here...
	req, err := http.NewRequest("GET", "/vignesh", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
	handler := byPass(http.HandlerFunc(services.Handler))

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusNotFound {
        t.Errorf("MIDDLEWARE ERROR: Handler returned wrong status code: got %v want %v",
            status, http.StatusNotFound)
    }

    // Check the response body is what we expect.
    expected := `Not Found!`
    if rr.Body.String() != expected {
        t.Errorf("MIDDLEWARE ERROR: Handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
	}
}

// Test ByPass function contains the basic validation for Auth token and header format.
func TestRegisterWSWithMiddleware(t *testing.T) {
    t.Log("MIDDLEWARE: Testing the RegisterWSWithMiddleware")
    RegisterWSWithMiddleWare("", false, nil)
    RegisterWSWithMiddleWare("/test1", true, nil)
    RegisterWSWithMiddleWare("/test2", false, services.Handler)
    RegisterWSWithMiddleWare("/test3", true, services.Handler)

    profile := services.GetProfileService()

    if profile.ServiceCount != 2 {
		t.Error("MAIN ERROR: Expected environment is `1`")
	}

}

// Test RegisterWSWithMiddleWareHandler 
func TestRegisterWSWithMiddleWareHandler(t *testing.T) {
    t.Log("MIDDLEWARE: Testing the RegisterWSWithMiddleWareHandler")
    RegisterWSWithMiddleWareHandler("", false, nil)
    RegisterWSWithMiddleWareHandler("/test1", true, nil)
    // rh := http.RedirectHandler("/", 307)
    // RegisterWSWithMiddleWareHandler("/test2", false, rh)
    // RegisterWSWithMiddleWareHandler("/test3", true, rh)

    profile := services.GetProfileService()

    if profile.ServiceCount != 2 {
		t.Error("MAIN ERROR: Expected environment is `1`")
	}

}