/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the HTTPS object model.
 */

package model

// HTTPs Object model
type HTTPSModel struct {
    Enable bool `json:"enable"`
    Port int `json:"port"`
    Option struct {
        Key     string `json:"key"`
        Cert string `json:"cert"`
    } `json:"option"`
}