/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Status object model.
 */

 package model

 import "testing"
 
 func TestStatusModel(t *testing.T) {
     data := []byte(`{ "votes": { "option_A": "3" } }`)

     statusObject := CreateStatus(true, data, "Successfully fetch user details")

	 t.Log("STATUS MODEL: Testing the Status Model create and validate parameters")

     if len(statusObject.Message) > 0 && statusObject.Message != "Successfully fetch user details" {
		t.Errorf("Status Message: %s", statusObject.Message)
     }

	 if len(statusObject.Data) > 0 && statusObject.Data == nil {
		t.Errorf("Status Data: %s", statusObject.Data)
	 }
	 
     if statusObject.Status != true {
		t.Error("Status Code: false")
	 }

	 if &statusObject == nil {
		t.Error("Status object is nil")
	 }

 }

