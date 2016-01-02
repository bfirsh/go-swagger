package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/toolkit"

	"github.com/go-swagger/go-swagger/examples/generated/models"
)

/*UploadFileOK successful operation

swagger:response uploadFileOK
*/
type UploadFileOK struct {

	// In: body
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUploadFileOK creates UploadFileOK with default headers values
func NewUploadFileOK() UploadFileOK {
	return UploadFileOK{}
}

// WithPayload adds the payload to the upload file o k response
func (o *UploadFileOK) WithPayload(payload *models.APIResponse) *UploadFileOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *UploadFileOK) WriteResponse(rw http.ResponseWriter, producer toolkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
