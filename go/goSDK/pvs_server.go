package xenapi

import (
	"fmt"
)

type PVSServerRecord struct {
	// Unique identifier/object reference
	UUID string
	// IPv4 addresses of this server
	Addresses []string
	// First UDP port accepted by this server
	FirstPort int
	// Last UDP port accepted by this server
	LastPort int
	// PVS site this server is part of
	Site PVSSiteRef
}

type PVSServerRef string

// individual machine serving provisioning (block) data
type pVSServer struct{}

var PVSServer pVSServer

// GetAllRecords: Return a map of PVS_server references to PVS_server records for all PVS_servers known to the system.
// Version: ely
func (pVSServer) GetAllRecords(session *Session) (retval map[PVSServerRef]PVSServerRecord, err error) {
	method := "PVS_server.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRefToPVSServerRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of PVS_server references to PVS_server records for all PVS_servers known to the system.
// Version: ely
func (pVSServer) GetAllRecords1(session *Session) (retval map[PVSServerRef]PVSServerRecord, err error) {
	method := "PVS_server.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRefToPVSServerRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the PVS_servers known to the system.
// Version: ely
func (pVSServer) GetAll(session *Session) (retval []PVSServerRef, err error) {
	method := "PVS_server.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the PVS_servers known to the system.
// Version: ely
func (pVSServer) GetAll1(session *Session) (retval []PVSServerRef, err error) {
	method := "PVS_server.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRefSet(method+" -> ", result)
	return
}

// Forget: forget a PVS server
// Version: ely
func (pVSServer) Forget(session *Session, self PVSServerRef) (err error) {
	method := "PVS_server.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncForget: forget a PVS server
// Version: ely
func (pVSServer) AsyncForget(session *Session, self PVSServerRef) (retval TaskRef, err error) {
	method := "Async.PVS_server.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Forget2: forget a PVS server
// Version: ely
func (pVSServer) Forget2(session *Session, self PVSServerRef) (err error) {
	method := "PVS_server.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncForget2: forget a PVS server
// Version: ely
func (pVSServer) AsyncForget2(session *Session, self PVSServerRef) (retval TaskRef, err error) {
	method := "Async.PVS_server.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Introduce: introduce new PVS server
// Version: ely
func (pVSServer) Introduce(session *Session, addresses []string, firstPort int, lastPort int, site PVSSiteRef) (retval PVSServerRef, err error) {
	method := "PVS_server.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	addressesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "addresses"), addresses)
	if err != nil {
		return
	}
	firstPortArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "first_port"), firstPort)
	if err != nil {
		return
	}
	lastPortArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "last_port"), lastPort)
	if err != nil {
		return
	}
	siteArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "site"), site)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, addressesArg, firstPortArg, lastPortArg, siteArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRef(method+" -> ", result)
	return
}

// AsyncIntroduce: introduce new PVS server
// Version: ely
func (pVSServer) AsyncIntroduce(session *Session, addresses []string, firstPort int, lastPort int, site PVSSiteRef) (retval TaskRef, err error) {
	method := "Async.PVS_server.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	addressesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "addresses"), addresses)
	if err != nil {
		return
	}
	firstPortArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "first_port"), firstPort)
	if err != nil {
		return
	}
	lastPortArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "last_port"), lastPort)
	if err != nil {
		return
	}
	siteArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "site"), site)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, addressesArg, firstPortArg, lastPortArg, siteArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Introduce5: introduce new PVS server
// Version: ely
func (pVSServer) Introduce5(session *Session, addresses []string, firstPort int, lastPort int, site PVSSiteRef) (retval PVSServerRef, err error) {
	method := "PVS_server.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	addressesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "addresses"), addresses)
	if err != nil {
		return
	}
	firstPortArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "first_port"), firstPort)
	if err != nil {
		return
	}
	lastPortArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "last_port"), lastPort)
	if err != nil {
		return
	}
	siteArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "site"), site)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, addressesArg, firstPortArg, lastPortArg, siteArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRef(method+" -> ", result)
	return
}

// AsyncIntroduce5: introduce new PVS server
// Version: ely
func (pVSServer) AsyncIntroduce5(session *Session, addresses []string, firstPort int, lastPort int, site PVSSiteRef) (retval TaskRef, err error) {
	method := "Async.PVS_server.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	addressesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "addresses"), addresses)
	if err != nil {
		return
	}
	firstPortArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "first_port"), firstPort)
	if err != nil {
		return
	}
	lastPortArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "last_port"), lastPort)
	if err != nil {
		return
	}
	siteArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "site"), site)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, addressesArg, firstPortArg, lastPortArg, siteArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetSite: Get the site field of the given PVS_server.
// Version: ely
func (pVSServer) GetSite(session *Session, self PVSServerRef) (retval PVSSiteRef, err error) {
	method := "PVS_server.get_site"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSite2: Get the site field of the given PVS_server.
// Version: ely
func (pVSServer) GetSite2(session *Session, self PVSServerRef) (retval PVSSiteRef, err error) {
	method := "PVS_server.get_site"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastPort: Get the last_port field of the given PVS_server.
// Version: ely
func (pVSServer) GetLastPort(session *Session, self PVSServerRef) (retval int, err error) {
	method := "PVS_server.get_last_port"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// GetLastPort2: Get the last_port field of the given PVS_server.
// Version: ely
func (pVSServer) GetLastPort2(session *Session, self PVSServerRef) (retval int, err error) {
	method := "PVS_server.get_last_port"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// GetFirstPort: Get the first_port field of the given PVS_server.
// Version: ely
func (pVSServer) GetFirstPort(session *Session, self PVSServerRef) (retval int, err error) {
	method := "PVS_server.get_first_port"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// GetFirstPort2: Get the first_port field of the given PVS_server.
// Version: ely
func (pVSServer) GetFirstPort2(session *Session, self PVSServerRef) (retval int, err error) {
	method := "PVS_server.get_first_port"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// GetAddresses: Get the addresses field of the given PVS_server.
// Version: ely
func (pVSServer) GetAddresses(session *Session, self PVSServerRef) (retval []string, err error) {
	method := "PVS_server.get_addresses"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// GetAddresses2: Get the addresses field of the given PVS_server.
// Version: ely
func (pVSServer) GetAddresses2(session *Session, self PVSServerRef) (retval []string, err error) {
	method := "PVS_server.get_addresses"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given PVS_server.
// Version: ely
func (pVSServer) GetUUID(session *Session, self PVSServerRef) (retval string, err error) {
	method := "PVS_server.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given PVS_server.
// Version: ely
func (pVSServer) GetUUID2(session *Session, self PVSServerRef) (retval string, err error) {
	method := "PVS_server.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByUUID: Get a reference to the PVS_server instance with the specified UUID.
// Version: ely
func (pVSServer) GetByUUID(session *Session, uUID string) (retval PVSServerRef, err error) {
	method := "PVS_server.get_by_uuid"
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
	retval, err = deserializePVSServerRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the PVS_server instance with the specified UUID.
// Version: ely
func (pVSServer) GetByUUID2(session *Session, uUID string) (retval PVSServerRef, err error) {
	method := "PVS_server.get_by_uuid"
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
	retval, err = deserializePVSServerRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given PVS_server.
// Version: ely
func (pVSServer) GetRecord(session *Session, self PVSServerRef) (retval PVSServerRecord, err error) {
	method := "PVS_server.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given PVS_server.
// Version: ely
func (pVSServer) GetRecord2(session *Session, self PVSServerRef) (retval PVSServerRecord, err error) {
	method := "PVS_server.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSServerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRecord(method+" -> ", result)
	return
}
