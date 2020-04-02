// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// GetNameserverOKCode is the HTTP code returned for type GetNameserverOK
const GetNameserverOKCode int = 200

/*GetNameserverOK Successful operation

swagger:response getNameserverOK
*/
type GetNameserverOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetNameserverOKBody `json:"body,omitempty"`
}

// NewGetNameserverOK creates GetNameserverOK with default headers values
func NewGetNameserverOK() *GetNameserverOK {

	return &GetNameserverOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get nameserver o k response
func (o *GetNameserverOK) WithConfigurationVersion(configurationVersion int64) *GetNameserverOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get nameserver o k response
func (o *GetNameserverOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get nameserver o k response
func (o *GetNameserverOK) WithPayload(payload *GetNameserverOKBody) *GetNameserverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nameserver o k response
func (o *GetNameserverOK) SetPayload(payload *GetNameserverOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNameserverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNameserverNotFoundCode is the HTTP code returned for type GetNameserverNotFound
const GetNameserverNotFoundCode int = 404

/*GetNameserverNotFound The specified resource was not found

swagger:response getNameserverNotFound
*/
type GetNameserverNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNameserverNotFound creates GetNameserverNotFound with default headers values
func NewGetNameserverNotFound() *GetNameserverNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetNameserverNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get nameserver not found response
func (o *GetNameserverNotFound) WithConfigurationVersion(configurationVersion int64) *GetNameserverNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get nameserver not found response
func (o *GetNameserverNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get nameserver not found response
func (o *GetNameserverNotFound) WithPayload(payload *models.Error) *GetNameserverNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nameserver not found response
func (o *GetNameserverNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNameserverNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetNameserverDefault General Error

swagger:response getNameserverDefault
*/
type GetNameserverDefault struct {
	_statusCode int
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNameserverDefault creates GetNameserverDefault with default headers values
func NewGetNameserverDefault(code int) *GetNameserverDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetNameserverDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get nameserver default response
func (o *GetNameserverDefault) WithStatusCode(code int) *GetNameserverDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get nameserver default response
func (o *GetNameserverDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get nameserver default response
func (o *GetNameserverDefault) WithConfigurationVersion(configurationVersion int64) *GetNameserverDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get nameserver default response
func (o *GetNameserverDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get nameserver default response
func (o *GetNameserverDefault) WithPayload(payload *models.Error) *GetNameserverDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nameserver default response
func (o *GetNameserverDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNameserverDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
