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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionHistoryGetOKCode is the HTTP code returned for type WeaviateActionHistoryGetOK
const WeaviateActionHistoryGetOKCode int = 200

/*WeaviateActionHistoryGetOK Successful response.

swagger:response weaviateActionHistoryGetOK
*/
type WeaviateActionHistoryGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.ActionGetHistoryResponse `json:"body,omitempty"`
}

// NewWeaviateActionHistoryGetOK creates WeaviateActionHistoryGetOK with default headers values
func NewWeaviateActionHistoryGetOK() *WeaviateActionHistoryGetOK {

	return &WeaviateActionHistoryGetOK{}
}

// WithPayload adds the payload to the weaviate action history get o k response
func (o *WeaviateActionHistoryGetOK) WithPayload(payload *models.ActionGetHistoryResponse) *WeaviateActionHistoryGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate action history get o k response
func (o *WeaviateActionHistoryGetOK) SetPayload(payload *models.ActionGetHistoryResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionHistoryGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionHistoryGetUnauthorizedCode is the HTTP code returned for type WeaviateActionHistoryGetUnauthorized
const WeaviateActionHistoryGetUnauthorizedCode int = 401

/*WeaviateActionHistoryGetUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionHistoryGetUnauthorized
*/
type WeaviateActionHistoryGetUnauthorized struct {
}

// NewWeaviateActionHistoryGetUnauthorized creates WeaviateActionHistoryGetUnauthorized with default headers values
func NewWeaviateActionHistoryGetUnauthorized() *WeaviateActionHistoryGetUnauthorized {

	return &WeaviateActionHistoryGetUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionHistoryGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionHistoryGetForbiddenCode is the HTTP code returned for type WeaviateActionHistoryGetForbidden
const WeaviateActionHistoryGetForbiddenCode int = 403

/*WeaviateActionHistoryGetForbidden The used API-key has insufficient permissions.

swagger:response weaviateActionHistoryGetForbidden
*/
type WeaviateActionHistoryGetForbidden struct {
}

// NewWeaviateActionHistoryGetForbidden creates WeaviateActionHistoryGetForbidden with default headers values
func NewWeaviateActionHistoryGetForbidden() *WeaviateActionHistoryGetForbidden {

	return &WeaviateActionHistoryGetForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionHistoryGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateActionHistoryGetNotFoundCode is the HTTP code returned for type WeaviateActionHistoryGetNotFound
const WeaviateActionHistoryGetNotFoundCode int = 404

/*WeaviateActionHistoryGetNotFound Successful query result but no resource was found.

swagger:response weaviateActionHistoryGetNotFound
*/
type WeaviateActionHistoryGetNotFound struct {
}

// NewWeaviateActionHistoryGetNotFound creates WeaviateActionHistoryGetNotFound with default headers values
func NewWeaviateActionHistoryGetNotFound() *WeaviateActionHistoryGetNotFound {

	return &WeaviateActionHistoryGetNotFound{}
}

// WriteResponse to the client
func (o *WeaviateActionHistoryGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// WeaviateActionHistoryGetInternalServerErrorCode is the HTTP code returned for type WeaviateActionHistoryGetInternalServerError
const WeaviateActionHistoryGetInternalServerErrorCode int = 500

/*WeaviateActionHistoryGetInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateActionHistoryGetInternalServerError
*/
type WeaviateActionHistoryGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionHistoryGetInternalServerError creates WeaviateActionHistoryGetInternalServerError with default headers values
func NewWeaviateActionHistoryGetInternalServerError() *WeaviateActionHistoryGetInternalServerError {

	return &WeaviateActionHistoryGetInternalServerError{}
}

// WithPayload adds the payload to the weaviate action history get internal server error response
func (o *WeaviateActionHistoryGetInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateActionHistoryGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate action history get internal server error response
func (o *WeaviateActionHistoryGetInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionHistoryGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionHistoryGetNotImplementedCode is the HTTP code returned for type WeaviateActionHistoryGetNotImplemented
const WeaviateActionHistoryGetNotImplementedCode int = 501

/*WeaviateActionHistoryGetNotImplemented Not (yet) implemented.

swagger:response weaviateActionHistoryGetNotImplemented
*/
type WeaviateActionHistoryGetNotImplemented struct {
}

// NewWeaviateActionHistoryGetNotImplemented creates WeaviateActionHistoryGetNotImplemented with default headers values
func NewWeaviateActionHistoryGetNotImplemented() *WeaviateActionHistoryGetNotImplemented {

	return &WeaviateActionHistoryGetNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateActionHistoryGetNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(501)
}
