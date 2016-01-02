package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/toolkit"
)

/*DeletePetBadRequest Invalid pet value

swagger:response deletePetBadRequest
*/
type DeletePetBadRequest struct {
}

// NewDeletePetBadRequest creates DeletePetBadRequest with default headers values
func NewDeletePetBadRequest() DeletePetBadRequest {
	return DeletePetBadRequest{}
}

// WriteResponse to the client
func (o *DeletePetBadRequest) WriteResponse(rw http.ResponseWriter, producer toolkit.Producer) {

	rw.WriteHeader(400)
}
