/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the User object model.
 */

package model

// User Object model
type ConfigModel struct {
    Title string `json:"title"`
    Version string `json:"version"`
    Auth string `json:"auth"`
    HTTP HTTPModel `json:"http"`
    HTTPS HTTPSModel `json:"https"`
}