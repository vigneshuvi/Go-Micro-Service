/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Main.
 */
 package main

 import "testing"
 import services "./services"

const (
	dev = "DEV"
	qa = "QA"
	prod = "PROD"
)

func TestHTTPServerStart(t *testing.T) {
	// test stuff here...
	t.Log("MAIN: Test HTTP Server Start")
	start(true)
	profile := services.GetProfileService()
	if profile.SeverStatus != true {
		t.Error("MAIN ERROR: HTTP server is not yet started")
	}

	if profile.ServiceCount != 7 {
		t.Errorf("MAIN ERROR: Register service is expected %d", profile.ServiceCount)
	}
}


 func TestEnvArguments(t *testing.T) {
	t.Log("MAIN: Testing the environment aruguments")

	if len(defaultEnv) == 0 {
		t.Error("MAIN ERROR: Default Environment Name: DEV")
	}

	envName := getArgument(dev)
	if envName != dev {
		t.Errorf("MAIN ERROR: Environment Name: %s", envName)
	}

	envName = getArgument(qa)
	if envName != qa {
		t.Errorf("MAIN ERROR: Environment Name: %s", envName)
	}

	envName = getArgument(prod)
	if envName != prod {
		t.Errorf("MAIN ERROR: Environment Name: %s", envName)
	}
 }

 func TestEnvironment(t *testing.T) {
	envConfig := getEnvironmentConfig(dev)
	t.Log("MAIN: Port number is expected : 3000")

	if envConfig.HTTP.Enable {
		t.Log("MAIN: HTTP Config is enabled.")
		if envConfig.HTTP.Port != 3000 {
			t.Errorf("MAIN ERROR: Configuration of port: %d", envConfig.HTTP.Port)
		}
	}
	
	if envConfig.HTTPS.Enable {
		t.Log("MAIN: HTTPs Config is enabled.")
		if envConfig.HTTPS.Port != 3043 {
			t.Errorf("MAIN ERROR: Configuration of port: %d", envConfig.HTTP.Port)
		}

		if len(envConfig.HTTPS.Option.Key) == 0  {
			t.Error("MAIN ERROR: HTTPS is key is empty")
		}

		if len(envConfig.HTTPS.Option.Cert) == 0  {
			t.Error("MAIN ERROR: HTTPS is Cert is empty")
		}
	}

	if envConfig.HTTP.Port != 3000 {
		t.Errorf("MAIN ERROR: Configuration of port: %d", envConfig.HTTP.Port)
	}

	t.Log("MAIN: API Version is expected : v1")

	if envConfig.Version != "v1" {
	   t.Errorf("MAIN ERROR: API Version is %s", envConfig.Version)
	}
 }


func TestHTTPServerStatus(t *testing.T) {
	// test stuff here...
	t.Log("MAIN: Test HTTP Server Status")
	envConfig := getEnvironmentConfig(dev)
	startHTTPServer(true, envConfig.HTTP.Port)

	profile := services.GetProfileService()
	if profile.SeverStatus != true {
		t.Error("MAIN ERROR: HTTP server is not yet started")
	 }

	//startHTTPServer(false, 3000)
	if profile.SeverStatus != true {
		t.Error("MAIN ERROR: HTTP server is not yet started")
	}

}


