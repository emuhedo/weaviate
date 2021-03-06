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

package knowledge_tools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateToolsMapHandlerFunc turns a function with the right signature into a weaviate tools map handler
type WeaviateToolsMapHandlerFunc func(WeaviateToolsMapParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateToolsMapHandlerFunc) Handle(params WeaviateToolsMapParams) middleware.Responder {
	return fn(params)
}

// WeaviateToolsMapHandler interface for that can handle valid weaviate tools map params
type WeaviateToolsMapHandler interface {
	Handle(WeaviateToolsMapParams) middleware.Responder
}

// NewWeaviateToolsMap creates a new http.Handler for the weaviate tools map operation
func NewWeaviateToolsMap(ctx *middleware.Context, handler WeaviateToolsMapHandler) *WeaviateToolsMap {
	return &WeaviateToolsMap{Context: ctx, Handler: handler}
}

/*WeaviateToolsMap swagger:route POST /tools/map knowledge tools weaviateToolsMap

Tool to render a map of concepts, based on ontologies available over the network.

Tool to render a map of concepts, based on ontologies available over the network.

*/
type WeaviateToolsMap struct {
	Context *middleware.Context
	Handler WeaviateToolsMapHandler
}

func (o *WeaviateToolsMap) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateToolsMapParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
