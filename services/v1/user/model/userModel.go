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
}