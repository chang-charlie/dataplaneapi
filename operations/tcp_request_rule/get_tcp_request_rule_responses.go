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

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// GetTCPRequestRuleOKCode is the HTTP code returned for type GetTCPRequestRuleOK
const GetTCPRequestRuleOKCode int = 200

/*GetTCPRequestRuleOK Successful operation

swagger:response getTcpRequestRuleOK
*/
type GetTCPRequestRuleOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetTCPRequestRuleOKBody `json:"body,omitempty"`
}

// NewGetTCPRequestRuleOK creates GetTCPRequestRuleOK with default headers values
func NewGetTCPRequestRuleOK() *GetTCPRequestRuleOK {

	return &GetTCPRequestRuleOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Tcp request rule o k response
func (o *GetTCPRequestRuleOK) WithConfigurationVersion(configurationVersion int64) *GetTCPRequestRuleOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Tcp request rule o k response
func (o *GetTCPRequestRuleOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Tcp request rule o k response
func (o *GetTCPRequestRuleOK) WithPayload(payload *GetTCPRequestRuleOKBody) *GetTCPRequestRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Tcp request rule o k response
func (o *GetTCPRequestRuleOK) SetPayload(payload *GetTCPRequestRuleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPRequestRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetTCPRequestRuleNotFoundCode is the HTTP code returned for type GetTCPRequestRuleNotFound
const GetTCPRequestRuleNotFoundCode int = 404

/*GetTCPRequestRuleNotFound The specified resource was not found

swagger:response getTcpRequestRuleNotFound
*/
type GetTCPRequestRuleNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTCPRequestRuleNotFound creates GetTCPRequestRuleNotFound with default headers values
func NewGetTCPRequestRuleNotFound() *GetTCPRequestRuleNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetTCPRequestRuleNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get Tcp request rule not found response
func (o *GetTCPRequestRuleNotFound) WithConfigurationVersion(configurationVersion int64) *GetTCPRequestRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Tcp request rule not found response
func (o *GetTCPRequestRuleNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Tcp request rule not found response
func (o *GetTCPRequestRuleNotFound) WithPayload(payload *models.Error) *GetTCPRequestRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Tcp request rule not found response
func (o *GetTCPRequestRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPRequestRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetTCPRequestRuleDefault General Error

swagger:response getTcpRequestRuleDefault
*/
type GetTCPRequestRuleDefault struct {
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

// NewGetTCPRequestRuleDefault creates GetTCPRequestRuleDefault with default headers values
func NewGetTCPRequestRuleDefault(code int) *GetTCPRequestRuleDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetTCPRequestRuleDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get TCP request rule default response
func (o *GetTCPRequestRuleDefault) WithStatusCode(code int) *GetTCPRequestRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get TCP request rule default response
func (o *GetTCPRequestRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get TCP request rule default response
func (o *GetTCPRequestRuleDefault) WithConfigurationVersion(configurationVersion int64) *GetTCPRequestRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get TCP request rule default response
func (o *GetTCPRequestRuleDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get TCP request rule default response
func (o *GetTCPRequestRuleDefault) WithPayload(payload *models.Error) *GetTCPRequestRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get TCP request rule default response
func (o *GetTCPRequestRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPRequestRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
