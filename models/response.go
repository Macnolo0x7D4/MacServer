package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}

func (this *Response) NotFound() {
	this.Status = http.StatusNotFound
	this.Message = "Resource Not Found."
}

func (this *Response) Forbidden(){
	this.Status = http.StatusForbidden
	this.Message = "You do not have access to this part of the website."
}

func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

func (this *Response) Send() {
	this.writer.Header().Set("Content-Type", this.contentType)
	this.writer.WriteHeader(this.Status)

	output, _ := json.Marshal(&this)
	fmt.Fprintf(this.writer, string(output))
}

func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

func SendForbidden(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.Forbidden()
	response.Send()
}

func SendBadRequest(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.BadRequest()
	response.Send()
}

func (this *Response) NoContent() {
	this.Status = http.StatusNoContent
	this.Message = "No content"
}

func (this *Response) UnprocessableEntity() {
	this.Status = http.StatusUnprocessableEntity
	this.Message = "Resource not adapted"
}

func (this *Response) BadRequest(){
	this.Status = http.StatusBadRequest
	this.Message = "Bad Request. Try to log in again and refresh the website."
}
