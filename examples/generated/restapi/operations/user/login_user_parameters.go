package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/toolkit"
)

// NewLoginUserParams creates a new LoginUserParams object
// with the default values initialized.
func NewLoginUserParams() LoginUserParams {
	return LoginUserParams{}
}

// LoginUserParams contains all the bound params for the login user operation
// typically these are obtained from a http.Request
//
// swagger:parameters loginUser
type LoginUserParams struct {
	/*The password for login in clear text
	  Required: true
	  In: query
	*/
	Password string
	/*The user name for login
	  Required: true
	  In: query
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *LoginUserParams) BindRequest(r *http.Request, route *httpkit.MatchedRoute) error {
	var res []error
	qs := toolkit.Values(r.URL.Query())

	qPassword, qhkPassword, _ := qs.GetOK("password")
	if err := o.bindPassword(qPassword, qhkPassword, route.Formats); err != nil {
		res = append(res, err)
	}

	qUsername, qhkUsername, _ := qs.GetOK("username")
	if err := o.bindUsername(qUsername, qhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginUserParams) bindPassword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("password", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("password", "query", raw); err != nil {
		return err
	}

	o.Password = raw

	return nil
}

func (o *LoginUserParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("username", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("username", "query", raw); err != nil {
		return err
	}

	o.Username = raw

	return nil
}
