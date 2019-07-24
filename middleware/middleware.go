/*
Package middleware provides help to check and bypass the micro service requests.

 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
*/
package middleware

import (
	"log"
	"net/http"
)

import services "../services"


//Middleware is a middleware handler that does request logging
type Middleware struct {
	handler http.Handler
}

// Authentication function contains the basic validation for Auth token and header format.
func authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(req http.ResponseWriter, res *http.Request) {
		xdate := res.Header.Get("xdate")
		if len(xdate) == 0 {
			log.Println("Unauthorized Request")
			req.WriteHeader(http.StatusUnauthorized)
			req.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(req, res)
		}
	})
}

// Bypass function helps to by-pass the basic authentication.
func byPass(next http.Handler) http.Handler {
	return http.HandlerFunc(func(req http.ResponseWriter, res *http.Request) {
		next.ServeHTTP(req, res)
	})
}

func RegisterWSWithMiddleWareHandler(path string, auth bool, fn http.Handler) {
	if len(path) == 0 || fn == nil {
		return
	}
	services.IncreseServiceCount()

	if auth {
		http.Handle(path, authentication(fn))
	} else {
		http.Handle(path, byPass(fn))
	}
}

func RegisterWSWithMiddleWare(path string, auth bool, fn func(res http.ResponseWriter, req *http.Request)) {
	if len(path) == 0 || fn == nil {
		return
	}
	services.IncreseServiceCount()

	if auth {
		http.Handle(path, authentication(http.HandlerFunc(fn)))
	} else {
		http.Handle(path, byPass(http.HandlerFunc(fn)))
	}
}
