package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/toolkit"

	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-1/models"
)

/*FindTodosOK list the todo operations

swagger:response findTodosOK
*/
type FindTodosOK struct {

	// In: body
	Payload []*models.Item `json:"body,omitempty"`
}

// NewFindTodosOK creates FindTodosOK with default headers values
func NewFindTodosOK() *FindTodosOK {
	return &FindTodosOK{}
}

// WithPayload adds the payload to the find todos o k response
func (o *FindTodosOK) WithPayload(payload []*models.Item) *FindTodosOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *FindTodosOK) WriteResponse(rw http.ResponseWriter, producer toolkit.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FindTodosDefault generic error response

swagger:response findTodosDefault
*/
type FindTodosDefault struct {
	_statusCode int `json:"-"`

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindTodosDefault creates FindTodosDefault with default headers values
func NewFindTodosDefault(code int) *FindTodosDefault {
	if code <= 0 {
		code = 500
	}

	return &FindTodosDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find todos default response
func (o *FindTodosDefault) WithStatusCode(code int) *FindTodosDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the find todos default response
func (o *FindTodosDefault) WithPayload(payload *models.Error) *FindTodosDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *FindTodosDefault) WriteResponse(rw http.ResponseWriter, producer toolkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
