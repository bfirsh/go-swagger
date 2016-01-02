package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"
)

// UploadFileHandlerFunc turns a function with the right signature into a upload file handler
type UploadFileHandlerFunc func(UploadFileParams, interface{}) httpkit.Responder

// Handle executing the request and returning a response
func (fn UploadFileHandlerFunc) Handle(params UploadFileParams, principal interface{}) httpkit.Responder {
	return fn(params, principal)
}

// UploadFileHandler interface for that can handle valid upload file params
type UploadFileHandler interface {
	Handle(UploadFileParams, interface{}) httpkit.Responder
}

// NewUploadFile creates a new http.Handler for the upload file operation
func NewUploadFile(ctx *httpkit.Context, handler UploadFileHandler) *UploadFile {
	return &UploadFile{Context: ctx, Handler: handler}
}

/*UploadFile swagger:route POST /pet/{petId}/uploadImage pet uploadFile

uploads an image

*/
type UploadFile struct {
	Context *httpkit.Context
	Params  UploadFileParams
	Handler UploadFileHandler
}

func (o *UploadFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewUploadFileParams()

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
