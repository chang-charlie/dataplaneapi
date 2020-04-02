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

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// ReplaceTCPResponseRuleOKCode is the HTTP code returned for type ReplaceTCPResponseRuleOK
const ReplaceTCPResponseRuleOKCode int = 200

/*ReplaceTCPResponseRuleOK TCP Response Rule replaced

swagger:response replaceTcpResponseRuleOK
*/
type ReplaceTCPResponseRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.TCPResponseRule `json:"body,omitempty"`
}

// NewReplaceTCPResponseRuleOK creates ReplaceTCPResponseRuleOK with default headers values
func NewReplaceTCPResponseRuleOK() *ReplaceTCPResponseRuleOK {

	return &ReplaceTCPResponseRuleOK{}
}

// WithPayload adds the payload to the replace Tcp response rule o k response
func (o *ReplaceTCPResponseRuleOK) WithPayload(payload *models.TCPResponseRule) *ReplaceTCPResponseRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp response rule o k response
func (o *ReplaceTCPResponseRuleOK) SetPayload(payload *models.TCPResponseRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPResponseRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceTCPResponseRuleAcceptedCode is the HTTP code returned for type ReplaceTCPResponseRuleAccepted
const ReplaceTCPResponseRuleAcceptedCode int = 202

/*ReplaceTCPResponseRuleAccepted Configuration change accepted and reload requested

swagger:response replaceTcpResponseRuleAccepted
*/
type ReplaceTCPResponseRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.TCPResponseRule `json:"body,omitempty"`
}

// NewReplaceTCPResponseRuleAccepted creates ReplaceTCPResponseRuleAccepted with default headers values
func NewReplaceTCPResponseRuleAccepted() *ReplaceTCPResponseRuleAccepted {

	return &ReplaceTCPResponseRuleAccepted{}
}

// WithReloadID adds the reloadId to the replace Tcp response rule accepted response
func (o *ReplaceTCPResponseRuleAccepted) WithReloadID(reloadID string) *ReplaceTCPResponseRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace Tcp response rule accepted response
func (o *ReplaceTCPResponseRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace Tcp response rule accepted response
func (o *ReplaceTCPResponseRuleAccepted) WithPayload(payload *models.TCPResponseRule) *ReplaceTCPResponseRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp response rule accepted response
func (o *ReplaceTCPResponseRuleAccepted) SetPayload(payload *models.TCPResponseRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPResponseRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceTCPResponseRuleBadRequestCode is the HTTP code returned for type ReplaceTCPResponseRuleBadRequest
const ReplaceTCPResponseRuleBadRequestCode int = 400

/*ReplaceTCPResponseRuleBadRequest Bad request

swagger:response replaceTcpResponseRuleBadRequest
*/
type ReplaceTCPResponseRuleBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPResponseRuleBadRequest creates ReplaceTCPResponseRuleBadRequest with default headers values
func NewReplaceTCPResponseRuleBadRequest() *ReplaceTCPResponseRuleBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &ReplaceTCPResponseRuleBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the replace Tcp response rule bad request response
func (o *ReplaceTCPResponseRuleBadRequest) WithConfigurationVersion(configurationVersion int64) *ReplaceTCPResponseRuleBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Tcp response rule bad request response
func (o *ReplaceTCPResponseRuleBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Tcp response rule bad request response
func (o *ReplaceTCPResponseRuleBadRequest) WithPayload(payload *models.Error) *ReplaceTCPResponseRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp response rule bad request response
func (o *ReplaceTCPResponseRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPResponseRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceTCPResponseRuleNotFoundCode is the HTTP code returned for type ReplaceTCPResponseRuleNotFound
const ReplaceTCPResponseRuleNotFoundCode int = 404

/*ReplaceTCPResponseRuleNotFound The specified resource was not found

swagger:response replaceTcpResponseRuleNotFound
*/
type ReplaceTCPResponseRuleNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPResponseRuleNotFound creates ReplaceTCPResponseRuleNotFound with default headers values
func NewReplaceTCPResponseRuleNotFound() *ReplaceTCPResponseRuleNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &ReplaceTCPResponseRuleNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the replace Tcp response rule not found response
func (o *ReplaceTCPResponseRuleNotFound) WithConfigurationVersion(configurationVersion int64) *ReplaceTCPResponseRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Tcp response rule not found response
func (o *ReplaceTCPResponseRuleNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Tcp response rule not found response
func (o *ReplaceTCPResponseRuleNotFound) WithPayload(payload *models.Error) *ReplaceTCPResponseRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp response rule not found response
func (o *ReplaceTCPResponseRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPResponseRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplaceTCPResponseRuleDefault General Error

swagger:response replaceTcpResponseRuleDefault
*/
type ReplaceTCPResponseRuleDefault struct {
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

// NewReplaceTCPResponseRuleDefault creates ReplaceTCPResponseRuleDefault with default headers values
func NewReplaceTCPResponseRuleDefault(code int) *ReplaceTCPResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &ReplaceTCPResponseRuleDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the replace TCP response rule default response
func (o *ReplaceTCPResponseRuleDefault) WithStatusCode(code int) *ReplaceTCPResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace TCP response rule default response
func (o *ReplaceTCPResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace TCP response rule default response
func (o *ReplaceTCPResponseRuleDefault) WithConfigurationVersion(configurationVersion int64) *ReplaceTCPResponseRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace TCP response rule default response
func (o *ReplaceTCPResponseRuleDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace TCP response rule default response
func (o *ReplaceTCPResponseRuleDefault) WithPayload(payload *models.Error) *ReplaceTCPResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace TCP response rule default response
func (o *ReplaceTCPResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
