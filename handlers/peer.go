package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	client_native "github.com/haproxytech/client-native/v2"
	"github.com/haproxytech/dataplaneapi/haproxy"
	"github.com/haproxytech/dataplaneapi/misc"
	"github.com/haproxytech/dataplaneapi/operations/peer"
	"github.com/haproxytech/models/v2"
)

//CreatePeerHandlerImpl implementation of the CreatePeerHandler interface using client-native client
type CreatePeerHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

//DeletePeerHandlerImpl implementation of the DeletePeerHandler interface using client-native client
type DeletePeerHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

//GetPeerHandlerImpl implementation of the GetPeerHandler interface using client-native client
type GetPeerHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//GetPeersHandlerImpl implementation of the GetPeersHandler interface using client-native client
type GetPeersHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//ReplacePeerHandlerImpl implementation of the ReplacePeerHandler interface using client-native client
type ReplacePeerHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

//Handle executing the request and returning a response
func (h *CreatePeerHandlerImpl) Handle(params peer.CreatePeerParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return peer.NewCreatePeerDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.CreatePeerSection(params.Data, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return peer.NewCreatePeerDefault(int(*e.Code)).WithPayload(e)
	}

	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return peer.NewCreatePeerDefault(int(*e.Code)).WithPayload(e)
			}
			return peer.NewCreatePeerCreated().WithPayload(params.Data)
		}
		rID := h.ReloadAgent.Reload()
		return peer.NewCreatePeerAccepted().WithReloadID(rID).WithPayload(params.Data)
	}
	return peer.NewCreatePeerAccepted().WithPayload(params.Data)
}

//Handle executing the request and returning a response
func (h *DeletePeerHandlerImpl) Handle(params peer.DeletePeerParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return peer.NewDeletePeerDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.DeletePeerSection(params.Name, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return peer.NewDeletePeerDefault(int(*e.Code)).WithPayload(e)
	}
	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return peer.NewDeletePeerDefault(int(*e.Code)).WithPayload(e)
			}
			return peer.NewDeletePeerNoContent()
		}
		rID := h.ReloadAgent.Reload()
		return peer.NewDeletePeerAccepted().WithReloadID(rID)
	}
	return peer.NewDeletePeerAccepted()
}

//Handle executing the request and returning a response
func (h *GetPeerHandlerImpl) Handle(params peer.GetPeerSectionParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	v, p, err := h.Client.Configuration.GetPeerSection(params.Name, t)
	if err != nil {
		e := misc.HandleError(err)
		return peer.NewGetPeerSectionDefault(int(*e.Code)).WithPayload(e).WithConfigurationVersion(v)
	}
	return peer.NewGetPeerSectionOK().WithPayload(&peer.GetPeerSectionOKBody{Version: v, Data: p}).WithConfigurationVersion(v)
}

//Handle executing the request and returning a response
func (h *GetPeersHandlerImpl) Handle(params peer.GetPeerSectionsParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	v, ps, err := h.Client.Configuration.GetPeerSections(t)
	if err != nil {
		e := misc.HandleError(err)
		return peer.NewGetPeerSectionsDefault(int(*e.Code)).WithPayload(e).WithConfigurationVersion(v)
	}
	return peer.NewGetPeerSectionsOK().WithPayload(&peer.GetPeerSectionsOKBody{Version: v, Data: ps}).WithConfigurationVersion(v)
}
