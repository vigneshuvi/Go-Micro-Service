/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Response object model.
 */

package model


import (
    "encoding/json"
    "net/http"
    "io"
    "fmt"
)

// Service ResponseWriter Object model
type Response struct {
    Writer	http.ResponseWriter
}

func NewResponse(res http.ResponseWriter) Response {
    resp := Response{Writer: res}
    return resp
}

func (r *Response) Json(code int, body []byte, messsage string) {
	r.Writer.Header().Set("Content-Type", "application/json")
    
    status := false
    if code == 200 {
        status = true
    }

    resJSON := StatusModel{Status: status, Data:json.RawMessage(fmt.Sprintf("%s", body)), Message: messsage}
	responseObj, err := json.Marshal(resJSON)
	if err != nil {
        fmt.Println(err)
        r.Writer.WriteHeader(http.StatusBadRequest)
		return
    }
    r.Writer.WriteHeader(code)
    r.Writer.Write(responseObj)
}

func (r *Response) Text(code int, body string) {
    r.Writer.Header().Set("Content-Type", "text/plain")
    r.Writer.WriteHeader(code)
    io.WriteString(r.Writer, fmt.Sprintf("%s", body))
}