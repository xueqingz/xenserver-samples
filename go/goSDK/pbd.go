package xenapi

import (
	"fmt"
)

type PBDRecord struct {
	// Unique identifier/object reference
	UUID string
	// physical machine on which the pbd is available
	Host HostRef
	// the storage repository that the pbd realises
	SR SRRef
	// a config string to string map that is provided to the host&apos;s SR-backend-driver
	DeviceConfig map[string]string
	// is the SR currently attached on this host?
	CurrentlyAttached bool
	// additional configuration
	OtherConfig map[string]string
}

type PBDRef string

// The physical block devices through which hosts access SRs
type pBD struct{}

var PBD pBD

// GetRecord: Get a record containing the current state of the given PBD.
func (pBD) GetRecord(session *Session, self PBDRef) (retval PBDRecord, err error) {
	method := "PBD.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePBDRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the PBD instance with the specified UUID.
func (pBD) GetByUUID(session *Session, uUID string) (retval PBDRef, err error) {
	method := "PBD.get_by_uuid"
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
	retval, err = deserializePBDRef(method+" -> ", result)
	return
}

// Create: Create a new PBD instance, and return its handle. The constructor args are: host*, SR*, device_config*, other_config (* = non-optional).
func (pBD) Create(session *Session, args PBDRecord) (retval PBDRef, err error) {
	method := "PBD.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializePBDRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializePBDRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new PBD instance, and return its handle. The constructor args are: host*, SR*, device_config*, other_config (* = non-optional).
func (pBD) AsyncCreate(session *Session, args PBDRecord) (retval TaskRef, err error) {
	method := "Async.PBD.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializePBDRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// Destroy: Destroy the specified PBD instance.
func (pBD) Destroy(session *Session, self PBDRef) (err error) {
	method := "PBD.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified PBD instance.
func (pBD) AsyncDestroy(session *Session, self PBDRef) (retval TaskRef, err error) {
	method := "Async.PBD.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given PBD.
func (pBD) GetUUID(session *Session, self PBDRef) (retval string, err error) {
	method := "PBD.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost: Get the host field of the given PBD.
func (pBD) GetHost(session *Session, self PBDRef) (retval HostRef, err error) {
	method := "PBD.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSR: Get the SR field of the given PBD.
func (pBD) GetSR(session *Session, self PBDRef) (retval SRRef, err error) {
	method := "PBD.get_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDeviceConfig: Get the device_config field of the given PBD.
func (pBD) GetDeviceConfig(session *Session, self PBDRef) (retval map[string]string, err error) {
	method := "PBD.get_device_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToStringMap(method+" -> ", result)
	return
}

// GetCurrentlyAttached: Get the currently_attached field of the given PBD.
func (pBD) GetCurrentlyAttached(session *Session, self PBDRef) (retval bool, err error) {
	method := "PBD.get_currently_attached"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given PBD.
func (pBD) GetOtherConfig(session *Session, self PBDRef) (retval map[string]string, err error) {
	method := "PBD.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToStringMap(method+" -> ", result)
	return
}

// SetOtherConfig: Set the other_config field of the given PBD.
func (pBD) SetOtherConfig(session *Session, self PBDRef, value map[string]string) (err error) {
	method := "PBD.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given PBD.
func (pBD) AddToOtherConfig(session *Session, self PBDRef, key string, value string) (err error) {
	method := "PBD.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	keyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "key"), key)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, keyArg, valueArg)
	return
}

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PBD.  If the key is not in that Map, then do nothing.
func (pBD) RemoveFromOtherConfig(session *Session, self PBDRef, key string) (err error) {
	method := "PBD.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	keyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "key"), key)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, keyArg)
	return
}

// Plug: Activate the specified PBD, causing the referenced SR to be attached and scanned
//
// Errors:
// SR_UNKNOWN_DRIVER - The SR could not be connected because the driver was not recognised.
func (pBD) Plug(session *Session, self PBDRef) (err error) {
	method := "PBD.plug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPlug: Activate the specified PBD, causing the referenced SR to be attached and scanned
//
// Errors:
// SR_UNKNOWN_DRIVER - The SR could not be connected because the driver was not recognised.
func (pBD) AsyncPlug(session *Session, self PBDRef) (retval TaskRef, err error) {
	method := "Async.PBD.plug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Unplug: Deactivate the specified PBD, causing the referenced SR to be detached and nolonger scanned
func (pBD) Unplug(session *Session, self PBDRef) (err error) {
	method := "PBD.unplug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncUnplug: Deactivate the specified PBD, causing the referenced SR to be detached and nolonger scanned
func (pBD) AsyncUnplug(session *Session, self PBDRef) (retval TaskRef, err error) {
	method := "Async.PBD.unplug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetDeviceConfig: Sets the PBD&apos;s device_config field
func (pBD) SetDeviceConfig(session *Session, self PBDRef, value map[string]string) (err error) {
	method := "PBD.set_device_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetDeviceConfig: Sets the PBD&apos;s device_config field
func (pBD) AsyncSetDeviceConfig(session *Session, self PBDRef, value map[string]string) (retval TaskRef, err error) {
	method := "Async.PBD.set_device_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetAll: Return a list of all the PBDs known to the system.
func (pBD) GetAll(session *Session) (retval []PBDRef, err error) {
	method := "PBD.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePBDRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of PBD references to PBD records for all PBDs known to the system.
func (pBD) GetAllRecords(session *Session) (retval map[PBDRef]PBDRecord, err error) {
	method := "PBD.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePBDRefToPBDRecordMap(method+" -> ", result)
	return
}
