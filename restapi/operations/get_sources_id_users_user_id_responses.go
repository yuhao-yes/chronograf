package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*GetSourcesIDUsersUserIDOK Information relating to the user

swagger:response getSourcesIdUsersUserIdOK
*/
type GetSourcesIDUsersUserIDOK struct {

	// In: body
	Payload *models.User `json:"body,omitempty"`
}

// NewGetSourcesIDUsersUserIDOK creates GetSourcesIDUsersUserIDOK with default headers values
func NewGetSourcesIDUsersUserIDOK() *GetSourcesIDUsersUserIDOK {
	return &GetSourcesIDUsersUserIDOK{}
}

// WithPayload adds the payload to the get sources Id users user Id o k response
func (o *GetSourcesIDUsersUserIDOK) WithPayload(payload *models.User) *GetSourcesIDUsersUserIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources Id users user Id o k response
func (o *GetSourcesIDUsersUserIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDUsersUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSourcesIDUsersUserIDNotFound Unknown source id

swagger:response getSourcesIdUsersUserIdNotFound
*/
type GetSourcesIDUsersUserIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSourcesIDUsersUserIDNotFound creates GetSourcesIDUsersUserIDNotFound with default headers values
func NewGetSourcesIDUsersUserIDNotFound() *GetSourcesIDUsersUserIDNotFound {
	return &GetSourcesIDUsersUserIDNotFound{}
}

// WithPayload adds the payload to the get sources Id users user Id not found response
func (o *GetSourcesIDUsersUserIDNotFound) WithPayload(payload *models.Error) *GetSourcesIDUsersUserIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources Id users user Id not found response
func (o *GetSourcesIDUsersUserIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDUsersUserIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSourcesIDUsersUserIDDefault Unexpected internal service error

swagger:response getSourcesIdUsersUserIdDefault
*/
type GetSourcesIDUsersUserIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSourcesIDUsersUserIDDefault creates GetSourcesIDUsersUserIDDefault with default headers values
func NewGetSourcesIDUsersUserIDDefault(code int) *GetSourcesIDUsersUserIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSourcesIDUsersUserIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get sources ID users user ID default response
func (o *GetSourcesIDUsersUserIDDefault) WithStatusCode(code int) *GetSourcesIDUsersUserIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get sources ID users user ID default response
func (o *GetSourcesIDUsersUserIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get sources ID users user ID default response
func (o *GetSourcesIDUsersUserIDDefault) WithPayload(payload *models.Error) *GetSourcesIDUsersUserIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources ID users user ID default response
func (o *GetSourcesIDUsersUserIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDUsersUserIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
