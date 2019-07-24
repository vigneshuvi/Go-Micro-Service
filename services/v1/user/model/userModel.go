/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the User object model.
 */

package model

// User Object model
type UserModel struct {
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
    Username string `json:"username"`
    RoleType int `json:"roletype"`
    RoleName string `json:"rolename"`
}

func CreateUser(fname string, lname string, uname string, utype int) UserModel {
    roleName := "user"
    switch utype {
        case 1:
            roleName = "admin"
        case 2:
            roleName = "manager"
    }
    userObject := UserModel{Firstname: fname, Lastname: lname, Username: uname, RoleType: utype, RoleName: roleName}
    return userObject
}