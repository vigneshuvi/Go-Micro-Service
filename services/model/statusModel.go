/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Status object model.
 */

package model
import (
	"encoding/json"
)

// Status Object model
type StatusModel struct {
    Status bool `json:"status"`
    Data json.RawMessage `json:"data"`
    Message string `json:"message"`
}

func CreateStatus(status bool, data json.RawMessage, message string) StatusModel {
    stausbject := StatusModel{Status: status, Data: data, Message: message}
    return stausbject
}