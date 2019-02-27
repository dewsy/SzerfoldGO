// Code generated by go-swagger; DO NOT EDIT.

package daily

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/go-swagger/go-swagger/cmd/swagger/models"
)

// GetDailyOKCode is the HTTP code returned for type GetDailyOK
const GetDailyOKCode int = 200

/*GetDailyOK list of daily items since given id

swagger:response getDailyOK
*/
type GetDailyOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Daily `json:"body,omitempty"`
}

// NewGetDailyOK creates GetDailyOK with default headers values
func NewGetDailyOK() *GetDailyOK {

	return &GetDailyOK{}
}

// WithPayload adds the payload to the get daily o k response
func (o *GetDailyOK) WithPayload(payload []*models.Daily) *GetDailyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get daily o k response
func (o *GetDailyOK) SetPayload(payload []*models.Daily) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDailyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Daily, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetDailyDefault generic error response

swagger:response getDailyDefault
*/
type GetDailyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDailyDefault creates GetDailyDefault with default headers values
func NewGetDailyDefault(code int) *GetDailyDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDailyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get daily default response
func (o *GetDailyDefault) WithStatusCode(code int) *GetDailyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get daily default response
func (o *GetDailyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get daily default response
func (o *GetDailyDefault) WithPayload(payload *models.Error) *GetDailyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get daily default response
func (o *GetDailyDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDailyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
