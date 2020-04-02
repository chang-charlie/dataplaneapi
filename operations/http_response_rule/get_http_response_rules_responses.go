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

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// GetHTTPResponseRulesOKCode is the HTTP code returned for type GetHTTPResponseRulesOK
const GetHTTPResponseRulesOKCode int = 200

/*GetHTTPResponseRulesOK Successful operation

swagger:response getHttpResponseRulesOK
*/
type GetHTTPResponseRulesOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetHTTPResponseRulesOKBody `json:"body,omitempty"`
}

// NewGetHTTPResponseRulesOK creates GetHTTPResponseRulesOK with default headers values
func NewGetHTTPResponseRulesOK() *GetHTTPResponseRulesOK {

	return &GetHTTPResponseRulesOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http response rules o k response
func (o *GetHTTPResponseRulesOK) WithConfigurationVersion(configurationVersion int64) *GetHTTPResponseRulesOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http response rules o k response
func (o *GetHTTPResponseRulesOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http response rules o k response
func (o *GetHTTPResponseRulesOK) WithPayload(payload *GetHTTPResponseRulesOKBody) *GetHTTPResponseRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http response rules o k response
func (o *GetHTTPResponseRulesOK) SetPayload(payload *GetHTTPResponseRulesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPResponseRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetHTTPResponseRulesDefault General Error

swagger:response getHttpResponseRulesDefault
*/
type GetHTTPResponseRulesDefault struct {
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

// NewGetHTTPResponseRulesDefault creates GetHTTPResponseRulesDefault with default headers values
func NewGetHTTPResponseRulesDefault(code int) *GetHTTPResponseRulesDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetHTTPResponseRulesDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) WithStatusCode(code int) *GetHTTPResponseRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) WithConfigurationVersion(configurationVersion int64) *GetHTTPResponseRulesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) WithPayload(payload *models.Error) *GetHTTPResponseRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPResponseRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
