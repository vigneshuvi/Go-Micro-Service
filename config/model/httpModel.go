/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the HTTP object model.
 */

package model

// HTTP Object model
type HTTPModel struct {
    Enable bool `json:"enable"`
    Port int `json:"port"`
}