/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsCreateHandlerFunc turns a function with the right signature into a weaviate things create handler
type WeaviateThingsCreateHandlerFunc func(WeaviateThingsCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateThingsCreateHandlerFunc) Handle(params WeaviateThingsCreateParams) middleware.Responder {
	return fn(params)
}

// WeaviateThingsCreateHandler interface for that can handle valid weaviate things create params
type WeaviateThingsCreateHandler interface {
	Handle(WeaviateThingsCreateParams) middleware.Responder
}

// NewWeaviateThingsCreate creates a new http.Handler for the weaviate things create operation
func NewWeaviateThingsCreate(ctx *middleware.Context, handler WeaviateThingsCreateHandler) *WeaviateThingsCreate {
	return &WeaviateThingsCreate{Context: ctx, Handler: handler}
}

/*WeaviateThingsCreate swagger:route POST /things things weaviateThingsCreate

Create a new Thing based on a Thing template related to this key.

Registers a new Thing. Given meta-data and schema values are validated.

*/
type WeaviateThingsCreate struct {
	Context *middleware.Context
	Handler WeaviateThingsCreateHandler
}

func (o *WeaviateThingsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateThingsCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WeaviateThingsCreateBody weaviate things create body
// swagger:model WeaviateThingsCreateBody
type WeaviateThingsCreateBody struct {

	// If `async` is true, return a 202 with the new ID of the Thing. You will receive this response before the data is made persistent. If `async` is false, you will receive confirmation after the value is made persistent. The value of `async` defaults to false.
	Async bool `json:"async,omitempty"`

	// thing
	Thing *models.ThingCreate `json:"thing,omitempty"`
}

// Validate validates this weaviate things create body
func (o *WeaviateThingsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateThing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateThingsCreateBody) validateThing(formats strfmt.Registry) error {

	if swag.IsZero(o.Thing) { // not required
		return nil
	}

	if o.Thing != nil {
		if err := o.Thing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "thing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateThingsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateThingsCreateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateThingsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
