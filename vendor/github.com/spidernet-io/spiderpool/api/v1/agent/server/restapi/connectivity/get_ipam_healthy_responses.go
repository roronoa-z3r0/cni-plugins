// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package connectivity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetIpamHealthyOKCode is the HTTP code returned for type GetIpamHealthyOK
const GetIpamHealthyOKCode int = 200

/*
GetIpamHealthyOK Success

swagger:response getIpamHealthyOK
*/
type GetIpamHealthyOK struct {
}

// NewGetIpamHealthyOK creates GetIpamHealthyOK with default headers values
func NewGetIpamHealthyOK() *GetIpamHealthyOK {

	return &GetIpamHealthyOK{}
}

// WriteResponse to the client
func (o *GetIpamHealthyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetIpamHealthyInternalServerErrorCode is the HTTP code returned for type GetIpamHealthyInternalServerError
const GetIpamHealthyInternalServerErrorCode int = 500

/*
GetIpamHealthyInternalServerError Failed

swagger:response getIpamHealthyInternalServerError
*/
type GetIpamHealthyInternalServerError struct {
}

// NewGetIpamHealthyInternalServerError creates GetIpamHealthyInternalServerError with default headers values
func NewGetIpamHealthyInternalServerError() *GetIpamHealthyInternalServerError {

	return &GetIpamHealthyInternalServerError{}
}

// WriteResponse to the client
func (o *GetIpamHealthyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
