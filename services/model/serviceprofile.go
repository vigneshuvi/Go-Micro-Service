/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Service Profile object model.
 */

package model

// Service Profile Object model
type ServiceProfile struct {
    Env string `json:"env"`
    ServiceCount int `json:"service_count"`
    SeverStatus bool `json:"server_status"`
}

func CreateServiceProfile(env string, count int, status bool) ServiceProfile {
    profileObject := ServiceProfile{Env: env, SeverStatus: status, ServiceCount: count}
    return profileObject
}