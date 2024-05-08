package xenapi

import (
	"fmt"
)

type PVSProxyRecord struct {
	// Unique identifier/object reference
	UUID string
	// PVS site this proxy is part of
	Site PVSSiteRef
	// VIF of the VM using the proxy
	VIF VIFRef
	// true = VM is currently proxied
	CurrentlyAttached bool
	// The run-time status of the proxy
	Status PvsProxyStatus
}

type PVSProxyRef string

// a proxy connects a VM/VIF with a PVS site
type pVSProxy struct{}

var PVSProxy pVSProxy

// GetRecord: Get a record containing the current state of the given PVS_proxy.
func (pVSProxy) GetRecord(session *Session, self PVSProxyRef) (retval PVSProxyRecord, err error) {
	method := "PVS_proxy.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSProxyRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSProxyRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the PVS_proxy instance with the specified UUID.
func (pVSProxy) GetByUUID(session *Session, uUID string) (retval PVSProxyRef, err error) {
	method := "PVS_proxy.get_by_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	uUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uuid"), uUID)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uUIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSProxyRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given PVS_proxy.
func (pVSProxy) GetUUID(session *Session, self PVSProxyRef) (retval string, err error) {
	method := "PVS_proxy.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSProxyRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// GetSite: Get the site field of the given PVS_proxy.
func (pVSProxy) GetSite(session *Session, self PVSProxyRef) (retval PVSSiteRef, err error) {
	method := "PVS_proxy.get_site"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSProxyRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSSiteRef(method+" -> ", result)
	return
}

// GetVIF: Get the VIF field of the given PVS_proxy.
func (pVSProxy) GetVIF(session *Session, self PVSProxyRef) (retval VIFRef, err error) {
	method := "PVS_proxy.get_VIF"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSProxyRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVIFRef(method+" -> ", result)
	return
}

// GetCurrentlyAttached: Get the currently_attached field of the given PVS_proxy.
func (pVSProxy) GetCurrentlyAttached(session *Session, self PVSProxyRef) (retval bool, err error) {
	method := "PVS_proxy.get_currently_attached"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSProxyRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeBool(method+" -> ", result)
	return
}

// GetStatus: Get the status field of the given PVS_proxy.
func (pVSProxy) GetStatus(session *Session, self PVSProxyRef) (retval PvsProxyStatus, err error) {
	method := "PVS_proxy.get_status"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSProxyRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumPvsProxyStatus(method+" -> ", result)
	return
}

// Create: Configure a VM/VIF to use a PVS proxy
func (pVSProxy) Create(session *Session, site PVSSiteRef, vIF VIFRef) (retval PVSProxyRef, err error) {
	method := "PVS_proxy.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	siteArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "site"), site)
	if err != nil {
		return
	}
	vIFArg, err := serializeVIFRef(fmt.Sprintf("%s(%s)", method, "VIF"), vIF)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, siteArg, vIFArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSProxyRef(method+" -> ", result)
	return
}

// AsyncCreate: Configure a VM/VIF to use a PVS proxy
func (pVSProxy) AsyncCreate(session *Session, site PVSSiteRef, vIF VIFRef) (retval TaskRef, err error) {
	method := "Async.PVS_proxy.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	siteArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "site"), site)
	if err != nil {
		return
	}
	vIFArg, err := serializeVIFRef(fmt.Sprintf("%s(%s)", method, "VIF"), vIF)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, siteArg, vIFArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: remove (or switch off) a PVS proxy for this VM
func (pVSProxy) Destroy(session *Session, self PVSProxyRef) (err error) {
	method := "PVS_proxy.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSProxyRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: remove (or switch off) a PVS proxy for this VM
func (pVSProxy) AsyncDestroy(session *Session, self PVSProxyRef) (retval TaskRef, err error) {
	method := "Async.PVS_proxy.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSProxyRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetAll: Return a list of all the PVS_proxys known to the system.
func (pVSProxy) GetAll(session *Session) (retval []PVSProxyRef, err error) {
	method := "PVS_proxy.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSProxyRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of PVS_proxy references to PVS_proxy records for all PVS_proxys known to the system.
func (pVSProxy) GetAllRecords(session *Session) (retval map[PVSProxyRef]PVSProxyRecord, err error) {
	method := "PVS_proxy.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSProxyRefToPVSProxyRecordMap(method+" -> ", result)
	return
}

