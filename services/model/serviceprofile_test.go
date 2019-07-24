/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Main.
 */
 package model

 import (
	"encoding/json"
	"testing"
)
 
 func TestServiceProfileModel(t *testing.T) {
	 profileObject := CreateServiceProfile("DEV", 1, true)
	 t.Log("SERVICE PROFILE MODEL: Testing the ServiceProfile create and validate parameters")

	 if len(profileObject.Env) > 0 && profileObject.Env != "DEV" {
		t.Errorf("Service Profile Env : %s", profileObject.Env)
	 }

	 if profileObject.SeverStatus != true {
		t.Error("Service Profile status is false")
	 }

	 if profileObject.ServiceCount != 1  {
		t.Errorf("Register Service Profile Count: %d", profileObject.ServiceCount)
	 }

	 if profileObject != (ServiceProfile{Env:"DEV", ServiceCount: 1, SeverStatus: true}) {
		t.Error("Profile object is same!")
	 }

	if &profileObject == nil {
		t.Error("Profile object is nil")
	}

	responseObj, err := json.Marshal(profileObject)
	if err != nil {
		t.Error("Unable to convert to object!")
	}
	if &responseObj == nil {
		t.Error("Profile object is nil")
	}

 }

