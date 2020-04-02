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

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// GetPeerEntryOKCode is the HTTP code returned for type GetPeerEntryOK
const GetPeerEntryOKCode int = 200

/*GetPeerEntryOK Successful operation

swagger:response getPeerEntryOK
*/
type GetPeerEntryOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetPeerEntryOKBody `json:"body,omitempty"`
}

// NewGetPeerEntryOK creates GetPeerEntryOK with default headers values
func NewGetPeerEntryOK() *GetPeerEntryOK {

	return &GetPeerEntryOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get peer entry o k response
func (o *GetPeerEntryOK) WithConfigurationVersion(configurationVersion int64) *GetPeerEntryOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get peer entry o k response
func (o *GetPeerEntryOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get peer entry o k response
func (o *GetPeerEntryOK) WithPayload(payload *GetPeerEntryOKBody) *GetPeerEntryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peer entry o k response
func (o *GetPeerEntryOK) SetPayload(payload *GetPeerEntryOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeerEntryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetPeerEntryNotFoundCode is the HTTP code returned for type GetPeerEntryNotFound
const GetPeerEntryNotFoundCode int = 404

/*GetPeerEntryNotFound The specified resource already exists

swagger:response getPeerEntryNotFound
*/
type GetPeerEntryNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPeerEntryNotFound creates GetPeerEntryNotFound with default headers values
func NewGetPeerEntryNotFound() *GetPeerEntryNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetPeerEntryNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get peer entry not found response
func (o *GetPeerEntryNotFound) WithConfigurationVersion(configurationVersion int64) *GetPeerEntryNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get peer entry not found response
func (o *GetPeerEntryNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get peer entry not found response
func (o *GetPeerEntryNotFound) WithPayload(payload *models.Error) *GetPeerEntryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peer entry not found response
func (o *GetPeerEntryNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeerEntryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetPeerEntryDefault General Error

swagger:response getPeerEntryDefault
*/
type GetPeerEntryDefault struct {
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

// NewGetPeerEntryDefault creates GetPeerEntryDefault with default headers values
func NewGetPeerEntryDefault(code int) *GetPeerEntryDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetPeerEntryDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get peer entry default response
func (o *GetPeerEntryDefault) WithStatusCode(code int) *GetPeerEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get peer entry default response
func (o *GetPeerEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get peer entry default response
func (o *GetPeerEntryDefault) WithConfigurationVersion(configurationVersion int64) *GetPeerEntryDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get peer entry default response
func (o *GetPeerEntryDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get peer entry default response
func (o *GetPeerEntryDefault) WithPayload(payload *models.Error) *GetPeerEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peer entry default response
func (o *GetPeerEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeerEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
