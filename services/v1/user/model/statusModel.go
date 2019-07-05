/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Status object model.
 */

package model

// Status Object model
type StatusModel struct {
    Status int `json:"status"`
    Data UserModel `json:"data"`
    Message string `json:"message"`
}