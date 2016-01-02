package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"
)

// UpdatePetWithFormHandlerFunc turns a function with the right signature into a update pet with form handler
type UpdatePetWithFormHandlerFunc func(UpdatePetWithFormParams, interface{}) httpkit.Responder

// Handle executing the request and returning a response
func (fn UpdatePetWithFormHandlerFunc) Handle(params UpdatePetWithFormParams, principal interface{}) httpkit.Responder {
	return fn(params, principal)
}

// UpdatePetWithFormHandler interface for that can handle valid update pet with form params
type UpdatePetWithFormHandler interface {
	Handle(UpdatePetWithFormParams, interface{}) httpkit.Responder
}

// NewUpdatePetWithForm creates a new http.Handler for the update pet with form operation
func NewUpdatePetWithForm(ctx *httpkit.Context, handler UpdatePetWithFormHandler) *UpdatePetWithForm {
	return &UpdatePetWithForm{Context: ctx, Handler: handler}
}

/*UpdatePetWithForm swagger:route POST /pet/{petId} pet updatePetWithForm

Updates a pet in the store with form data

*/
type UpdatePetWithForm struct {
	Context *httpkit.Context
	Params  UpdatePetWithFormParams
	Handler UpdatePetWithFormHandler
}

func (o *UpdatePetWithForm) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewUpdatePetWithFormParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
