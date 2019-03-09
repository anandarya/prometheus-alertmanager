// Code generated by go-swagger; DO NOT EDIT.

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSilencesParams creates a new GetSilencesParams object
// no default values defined in spec.
func NewGetSilencesParams() GetSilencesParams {

	return GetSilencesParams{}
}

// GetSilencesParams contains all the bound params for the get silences operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSilences
type GetSilencesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A list of matchers to filter silences by
	  In: query
	  Collection Format: multi
	*/
	Filter []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSilencesParams() beforehand.
func (o *GetSilencesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFilter, qhkFilter, _ := qs.GetOK("filter")
	if err := o.bindFilter(qFilter, qhkFilter, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFilter binds and validates array parameter Filter from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetSilencesParams) bindFilter(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	filterIC := rawData

	if len(filterIC) == 0 {
		return nil
	}

	var filterIR []string
	for _, filterIV := range filterIC {
		filterI := filterIV

		filterIR = append(filterIR, filterI)
	}

	o.Filter = filterIR

	return nil
}
