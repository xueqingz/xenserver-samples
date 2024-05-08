package xenapi

import (
	"fmt"
)

type PVSCacheStorageRecord struct {
	// Unique identifier/object reference
	UUID string
	// The host on which this object defines PVS cache storage
	Host HostRef
	// SR providing storage for the PVS cache
	SR SRRef
	// The PVS_site for which this object defines the storage
	Site PVSSiteRef
	// The size of the cache VDI (in bytes)
	Size int
	// The VDI used for caching
	VDI VDIRef
}

type PVSCacheStorageRef string

// Describes the storage that is available to a PVS site for caching purposes
type pVSCacheStorage struct{}

var PVSCacheStorage pVSCacheStorage

// GetRecord: Get a record containing the current state of the given PVS_cache_storage.
func (pVSCacheStorage) GetRecord(session *Session, self PVSCacheStorageRef) (retval PVSCacheStorageRecord, err error) {
	method := "PVS_cache_storage.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSCacheStorageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSCacheStorageRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the PVS_cache_storage instance with the specified UUID.
func (pVSCacheStorage) GetByUUID(session *Session, uUID string) (retval PVSCacheStorageRef, err error) {
	method := "PVS_cache_storage.get_by_uuid"
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
	retval, err = deserializePVSCacheStorageRef(method+" -> ", result)
	return
}

// Create: Create a new PVS_cache_storage instance, and return its handle. The constructor args are: host, SR, site, size (* = non-optional).
func (pVSCacheStorage) Create(session *Session, args PVSCacheStorageRecord) (retval PVSCacheStorageRef, err error) {
	method := "PVS_cache_storage.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializePVSCacheStorageRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSCacheStorageRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new PVS_cache_storage instance, and return its handle. The constructor args are: host, SR, site, size (* = non-optional).
func (pVSCacheStorage) AsyncCreate(session *Session, args PVSCacheStorageRecord) (retval TaskRef, err error) {
	method := "Async.PVS_cache_storage.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializePVSCacheStorageRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Destroy the specified PVS_cache_storage instance.
func (pVSCacheStorage) Destroy(session *Session, self PVSCacheStorageRef) (err error) {
	method := "PVS_cache_storage.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSCacheStorageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified PVS_cache_storage instance.
func (pVSCacheStorage) AsyncDestroy(session *Session, self PVSCacheStorageRef) (retval TaskRef, err error) {
	method := "Async.PVS_cache_storage.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSCacheStorageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given PVS_cache_storage.
func (pVSCacheStorage) GetUUID(session *Session, self PVSCacheStorageRef) (retval string, err error) {
	method := "PVS_cache_storage.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSCacheStorageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost: Get the host field of the given PVS_cache_storage.
func (pVSCacheStorage) GetHost(session *Session, self PVSCacheStorageRef) (retval HostRef, err error) {
	method := "PVS_cache_storage.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSCacheStorageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRef(method+" -> ", result)
	return
}

// GetSR: Get the SR field of the given PVS_cache_storage.
func (pVSCacheStorage) GetSR(session *Session, self PVSCacheStorageRef) (retval SRRef, err error) {
	method := "PVS_cache_storage.get_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSCacheStorageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRef(method+" -> ", result)
	return
}

// GetSite: Get the site field of the given PVS_cache_storage.
func (pVSCacheStorage) GetSite(session *Session, self PVSCacheStorageRef) (retval PVSSiteRef, err error) {
	method := "PVS_cache_storage.get_site"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSCacheStorageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSize: Get the size field of the given PVS_cache_storage.
func (pVSCacheStorage) GetSize(session *Session, self PVSCacheStorageRef) (retval int, err error) {
	method := "PVS_cache_storage.get_size"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSCacheStorageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVDI: Get the VDI field of the given PVS_cache_storage.
func (pVSCacheStorage) GetVDI(session *Session, self PVSCacheStorageRef) (retval VDIRef, err error) {
	method := "PVS_cache_storage.get_VDI"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSCacheStorageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// GetAll: Return a list of all the PVS_cache_storages known to the system.
func (pVSCacheStorage) GetAll(session *Session) (retval []PVSCacheStorageRef, err error) {
	method := "PVS_cache_storage.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSCacheStorageRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of PVS_cache_storage references to PVS_cache_storage records for all PVS_cache_storages known to the system.
func (pVSCacheStorage) GetAllRecords(session *Session) (retval map[PVSCacheStorageRef]PVSCacheStorageRecord, err error) {
	method := "PVS_cache_storage.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSCacheStorageRefToPVSCacheStorageRecordMap(method+" -> ", result)
	return
}

