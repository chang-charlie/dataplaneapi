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

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// DeleteServerAcceptedCode is the HTTP code returned for type DeleteServerAccepted
const DeleteServerAcceptedCode int = 202

/*DeleteServerAccepted Configuration change accepted and reload requested

swagger:response deleteServerAccepted
*/
type DeleteServerAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteServerAccepted creates DeleteServerAccepted with default headers values
func NewDeleteServerAccepted() *DeleteServerAccepted {

	return &DeleteServerAccepted{}
}

// WithReloadID adds the reloadId to the delete server accepted response
func (o *DeleteServerAccepted) WithReloadID(reloadID string) *DeleteServerAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete server accepted response
func (o *DeleteServerAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteServerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteServerNoContentCode is the HTTP code returned for type DeleteServerNoContent
const DeleteServerNoContentCode int = 204

/*DeleteServerNoContent Server deleted

swagger:response deleteServerNoContent
*/
type DeleteServerNoContent struct {
}

// NewDeleteServerNoContent creates DeleteServerNoContent with default headers values
func NewDeleteServerNoContent() *DeleteServerNoContent {

	return &DeleteServerNoContent{}
}

// WriteResponse to the client
func (o *DeleteServerNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteServerNotFoundCode is the HTTP code returned for type DeleteServerNotFound
const DeleteServerNotFoundCode int = 404

/*DeleteServerNotFound The specified resource was not found

swagger:response deleteServerNotFound
*/
type DeleteServerNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteServerNotFound creates DeleteServerNotFound with default headers values
func NewDeleteServerNotFound() *DeleteServerNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &DeleteServerNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the delete server not found response
func (o *DeleteServerNotFound) WithConfigurationVersion(configurationVersion int64) *DeleteServerNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete server not found response
func (o *DeleteServerNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete server not found response
func (o *DeleteServerNotFound) WithPayload(payload *models.Error) *DeleteServerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete server not found response
func (o *DeleteServerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteServerDefault General Error

swagger:response deleteServerDefault
*/
type DeleteServerDefault struct {
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

// NewDeleteServerDefault creates DeleteServerDefault with default headers values
func NewDeleteServerDefault(code int) *DeleteServerDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &DeleteServerDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the delete server default response
func (o *DeleteServerDefault) WithStatusCode(code int) *DeleteServerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete server default response
func (o *DeleteServerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete server default response
func (o *DeleteServerDefault) WithConfigurationVersion(configurationVersion int64) *DeleteServerDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete server default response
func (o *DeleteServerDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete server default response
func (o *DeleteServerDefault) WithPayload(payload *models.Error) *DeleteServerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete server default response
func (o *DeleteServerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
