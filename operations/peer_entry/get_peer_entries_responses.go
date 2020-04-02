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

// GetPeerEntriesOKCode is the HTTP code returned for type GetPeerEntriesOK
const GetPeerEntriesOKCode int = 200

/*GetPeerEntriesOK Successful operation

swagger:response getPeerEntriesOK
*/
type GetPeerEntriesOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetPeerEntriesOKBody `json:"body,omitempty"`
}

// NewGetPeerEntriesOK creates GetPeerEntriesOK with default headers values
func NewGetPeerEntriesOK() *GetPeerEntriesOK {

	return &GetPeerEntriesOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get peer entries o k response
func (o *GetPeerEntriesOK) WithConfigurationVersion(configurationVersion int64) *GetPeerEntriesOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get peer entries o k response
func (o *GetPeerEntriesOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get peer entries o k response
func (o *GetPeerEntriesOK) WithPayload(payload *GetPeerEntriesOKBody) *GetPeerEntriesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peer entries o k response
func (o *GetPeerEntriesOK) SetPayload(payload *GetPeerEntriesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeerEntriesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetPeerEntriesDefault General Error

swagger:response getPeerEntriesDefault
*/
type GetPeerEntriesDefault struct {
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

// NewGetPeerEntriesDefault creates GetPeerEntriesDefault with default headers values
func NewGetPeerEntriesDefault(code int) *GetPeerEntriesDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetPeerEntriesDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get peer entries default response
func (o *GetPeerEntriesDefault) WithStatusCode(code int) *GetPeerEntriesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get peer entries default response
func (o *GetPeerEntriesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get peer entries default response
func (o *GetPeerEntriesDefault) WithConfigurationVersion(configurationVersion int64) *GetPeerEntriesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get peer entries default response
func (o *GetPeerEntriesDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get peer entries default response
func (o *GetPeerEntriesDefault) WithPayload(payload *models.Error) *GetPeerEntriesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peer entries default response
func (o *GetPeerEntriesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeerEntriesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
