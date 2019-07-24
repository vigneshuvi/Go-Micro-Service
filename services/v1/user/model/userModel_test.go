/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Main.
 */
 package model

 import "testing"
 
 func TestUserModel(t *testing.T) {
	 userObject := CreateUser("vignesh", "kumar", "vigneshuvi", 1)
	 t.Log("USER MODEL: Testing the UserModel create and validate parameters")

	 if len(userObject.Firstname) > 0 && userObject.Firstname != "vignesh" {
		t.Errorf("User First name: %s", userObject.Firstname)
	 }

	 if len(userObject.Lastname) > 0 && userObject.Lastname != "kumar" {
		t.Errorf("User First name: %s", userObject.Lastname)
	 }

	 if len(userObject.Username) > 0 && userObject.Username != "vigneshuvi" {
		t.Errorf("Username: %s", userObject.Username)
	 }

	 if len(userObject.RoleName) > 0 && userObject.RoleName != "admin" {
		t.Errorf("User Role name: %s", userObject.Username)
	 }

	 if userObject.RoleType < 0 && userObject.RoleType > 2 {
		t.Errorf("User Role type: %d", userObject.RoleType)
	 }

	 if userObject.RoleType != 1 {
		t.Error("User Role type should be 1")
	 }

	 if userObject != (UserModel{Firstname:"vignesh", Lastname: "kumar", Username: "vigneshuvi", RoleType: 1, RoleName: "admin"}) {
		t.Error("User object is same!")
	 }
 }

