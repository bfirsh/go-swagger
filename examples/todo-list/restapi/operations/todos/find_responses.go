package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/toolkit"

	"github.com/go-swagger/go-swagger/examples/todo-list/models"
)

/*FindOK OK

swagger:response findOK
*/
type FindOK struct {

	// In: body
	Payload []*models.Item `json:"body,omitempty"`
}

// NewFindOK creates FindOK with default headers values
func NewFindOK() *FindOK {
	return &FindOK{}
}

// WithPayload adds the payload to the find o k response
func (o *FindOK) WithPayload(payload []*models.Item) *FindOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *FindOK) WriteResponse(rw http.ResponseWriter, producer toolkit.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FindDefault error

swagger:response findDefault
*/
type FindDefault struct {
	_statusCode int `json:"-"`

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindDefault creates FindDefault with default headers values
func NewFindDefault(code int) *FindDefault {
	if code <= 0 {
		code = 500
	}

	return &FindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find default response
func (o *FindDefault) WithStatusCode(code int) *FindDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the find default response
func (o *FindDefault) WithPayload(payload *models.Error) *FindDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *FindDefault) WriteResponse(rw http.ResponseWriter, producer toolkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
