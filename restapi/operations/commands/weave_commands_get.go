/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package commands




import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveCommandsGetHandlerFunc turns a function with the right signature into a weave commands get handler
type WeaveCommandsGetHandlerFunc func(WeaveCommandsGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveCommandsGetHandlerFunc) Handle(params WeaveCommandsGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveCommandsGetHandler interface for that can handle valid weave commands get params
type WeaveCommandsGetHandler interface {
	Handle(WeaveCommandsGetParams, interface{}) middleware.Responder
}

// NewWeaveCommandsGet creates a new http.Handler for the weave commands get operation
func NewWeaveCommandsGet(ctx *middleware.Context, handler WeaveCommandsGetHandler) *WeaveCommandsGet {
	return &WeaveCommandsGet{Context: ctx, Handler: handler}
}

/*WeaveCommandsGet swagger:route GET /commands/{commandId} commands weaveCommandsGet

Returns a particular command.

*/
type WeaveCommandsGet struct {
	Context *middleware.Context
	Handler WeaveCommandsGetHandler
}

func (o *WeaveCommandsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveCommandsGetParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}