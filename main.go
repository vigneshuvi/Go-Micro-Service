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
	"strings"

)

import user "./services/v1/user"
import services "./services"
import middleware "./middleware"
import config "./config"
import model "./config/model"

var defaultEnv = "DEV"


// Project Main function and selected environment
func main() {
	start(false)
}

// Get the argument from run & test commend
func getArgument(env string) string {
	argsWithProg := os.Args
	var envName = env

	if len(argsWithProg) > 1 {
		envName = os.Args[1]
		if strings.Contains(envName, "test") {
			envName = env
		}
	}
	return envName
}

// Get the Environment Config based on go run arguments
func getEnvironmentConfig(env string) model.ConfigModel {
	envName := getArgument(env)
	services.SetServiceEnv(envName)

	log.Println("2.", envName, "- Set up the cofiguration.",)
	config := config.SetConfig(envName)
	log.Println("3.", config.Env.Title)

	return config.Env
}

// Start the HTTP server and register service
func start(debug bool) {
	log.Println("1. Application started..")
	services.SetProfileStatus(false)

	envConfig := getEnvironmentConfig(defaultEnv)
	log.Println(envConfig)

	registerWebService(envConfig)
	startHTTPServer(debug, envConfig.HTTP.Port)
}


// Start the HtTP Server
func startHTTPServer(debug bool, port int) {
	log.Println("5. Application running on port ", port)
	if !debug {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %s\n", err.Error())
		}
	}
	services.SetProfileStatus(true)
}

// Register Web services.
func registerWebService(envConfig model.ConfigModel) {
	log.Println("4. Register Web Services!")

	version :=  "/"+envConfig.Version

	// Router Bypass Authentication
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	rh := http.RedirectHandler("/", 307)
	middleware.RegisterWSWithMiddleWareHandler("/home",false, rh)
  
	//middleware.RegisterWSWithMiddleWare("/", false, services.Handler)
	middleware.RegisterWSWithMiddleWare(version+"/u1v2i3C4o5n6f7i8g9", false, services.Handler)
	middleware.RegisterWSWithMiddleWare(version+"/test", false, services.Handler)


	middleware.RegisterWSWithMiddleWare(version+"/getProfile", false, services.GetProfileHandler)
	middleware.RegisterWSWithMiddleWare(version+"/user", false, user.UserHandleRequest)
	middleware.RegisterWSWithMiddleWare(version+"/user?", false, user.UserHandleRequest)


	// Router Authentication
	middleware.RegisterWSWithMiddleWare(version+"/AuthUser", true, user.Create)
}