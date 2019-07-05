/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Main.
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

import user "./services/v1/user"
import middleware "./middleware"
import config "./config"

// Project Main function and selected environment
func main() {

	log.Println("1. Application started..")

	argsWithProg := os.Args
	var envName = "DEV"

	if len(argsWithProg) > 1 {
		envName = os.Args[1]
	}

	log.Println("2. Set up the cofiguration.")
	envConfig := config.SetConfig(envName)
	log.Println("3.", envConfig.Env.Title)

	registerWebService()
	log.Println("5. Application running on port ", envConfig.Env.HTTP.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", envConfig.Env.HTTP.Port), nil))
}

func handler(req http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(req, "Welcome, %s!", res.URL.Path[1:])
}

// Register Web services.
func registerWebService() {
	log.Println("4. Register Web Services!")

	// Router Bypass Authentication
	http.Handle("/", middleware.Bypass(http.HandlerFunc(handler)))
	http.Handle("/user/", middleware.Authentication(http.HandlerFunc(user.Create)))

	// Router Authentication
	http.Handle("/user/{username}/", middleware.Authentication(http.HandlerFunc(user.GetByUsername)))
}
