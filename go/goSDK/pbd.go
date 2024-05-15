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

// GetAllRecords: Return a map of PBD references to PBD records for all PBDs known to the system.
// Version: rio
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

// GetAllRecords1: Return a map of PBD references to PBD records for all PBDs known to the system.
// Version: rio
func (pBD) GetAllRecords1(session *Session) (retval map[PBDRef]PBDRecord, err error) {
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

// GetAll: Return a list of all the PBDs known to the system.
// Version: rio
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

// GetAll1: Return a list of all the PBDs known to the system.
// Version: rio
func (pBD) GetAll1(session *Session) (retval []PBDRef, err error) {
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

// SetDeviceConfig: Sets the PBD&apos;s device_config field
// Version: miami
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
// Version: miami
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

// SetDeviceConfig3: Sets the PBD&apos;s device_config field
// Version: miami
func (pBD) SetDeviceConfig3(session *Session, self PBDRef, value map[string]string) (err error) {
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

// AsyncSetDeviceConfig3: Sets the PBD&apos;s device_config field
// Version: miami
func (pBD) AsyncSetDeviceConfig3(session *Session, self PBDRef, value map[string]string) (retval TaskRef, err error) {
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

// Unplug: Deactivate the specified PBD, causing the referenced SR to be detached and nolonger scanned
// Version: rio
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
// Version: rio
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

// Unplug2: Deactivate the specified PBD, causing the referenced SR to be detached and nolonger scanned
// Version: rio
func (pBD) Unplug2(session *Session, self PBDRef) (err error) {
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

// AsyncUnplug2: Deactivate the specified PBD, causing the referenced SR to be detached and nolonger scanned
// Version: rio
func (pBD) AsyncUnplug2(session *Session, self PBDRef) (retval TaskRef, err error) {
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

// Plug: Activate the specified PBD, causing the referenced SR to be attached and scanned
// Version: rio
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
// Version: rio
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

// Plug2: Activate the specified PBD, causing the referenced SR to be attached and scanned
// Version: rio
//
// Errors:
// SR_UNKNOWN_DRIVER - The SR could not be connected because the driver was not recognised.
func (pBD) Plug2(session *Session, self PBDRef) (err error) {
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

// AsyncPlug2: Activate the specified PBD, causing the referenced SR to be attached and scanned
// Version: rio
//
// Errors:
// SR_UNKNOWN_DRIVER - The SR could not be connected because the driver was not recognised.
func (pBD) AsyncPlug2(session *Session, self PBDRef) (retval TaskRef, err error) {
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PBD.  If the key is not in that Map, then do nothing.
// Version: miami
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

// RemoveFromOtherConfig3: Remove the given key and its corresponding value from the other_config field of the given PBD.  If the key is not in that Map, then do nothing.
// Version: miami
func (pBD) RemoveFromOtherConfig3(session *Session, self PBDRef, key string) (err error) {
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

// RemoveFromOtherConfig2: Remove the given key and its corresponding value from the other_config field of the given PBD.  If the key is not in that Map, then do nothing.
// Version: rio
func (pBD) RemoveFromOtherConfig2(session *Session, self PBDRef) (err error) {
	method := "PBD.remove_from_other_config"
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given PBD.
// Version: miami
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

// AddToOtherConfig4: Add the given key-value pair to the other_config field of the given PBD.
// Version: miami
func (pBD) AddToOtherConfig4(session *Session, self PBDRef, key string, value string) (err error) {
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

// AddToOtherConfig2: Add the given key-value pair to the other_config field of the given PBD.
// Version: rio
func (pBD) AddToOtherConfig2(session *Session, self PBDRef) (err error) {
	method := "PBD.add_to_other_config"
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

// SetOtherConfig: Set the other_config field of the given PBD.
// Version: miami
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

// SetOtherConfig3: Set the other_config field of the given PBD.
// Version: miami
func (pBD) SetOtherConfig3(session *Session, self PBDRef, value map[string]string) (err error) {
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

// SetOtherConfig2: Set the other_config field of the given PBD.
// Version: rio
func (pBD) SetOtherConfig2(session *Session, self PBDRef) (err error) {
	method := "PBD.set_other_config"
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

// GetOtherConfig: Get the other_config field of the given PBD.
// Version: rio
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

// GetOtherConfig2: Get the other_config field of the given PBD.
// Version: rio
func (pBD) GetOtherConfig2(session *Session, self PBDRef) (retval map[string]string, err error) {
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

// GetCurrentlyAttached: Get the currently_attached field of the given PBD.
// Version: rio
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

// GetCurrentlyAttached2: Get the currently_attached field of the given PBD.
// Version: rio
func (pBD) GetCurrentlyAttached2(session *Session, self PBDRef) (retval bool, err error) {
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

// GetDeviceConfig: Get the device_config field of the given PBD.
// Version: rio
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

// GetDeviceConfig2: Get the device_config field of the given PBD.
// Version: rio
func (pBD) GetDeviceConfig2(session *Session, self PBDRef) (retval map[string]string, err error) {
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

// GetSR: Get the SR field of the given PBD.
// Version: rio
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

// GetSR2: Get the SR field of the given PBD.
// Version: rio
func (pBD) GetSR2(session *Session, self PBDRef) (retval SRRef, err error) {
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

// GetHost: Get the host field of the given PBD.
// Version: rio
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

// GetHost2: Get the host field of the given PBD.
// Version: rio
func (pBD) GetHost2(session *Session, self PBDRef) (retval HostRef, err error) {
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

// GetUUID: Get the uuid field of the given PBD.
// Version: rio
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

// GetUUID2: Get the uuid field of the given PBD.
// Version: rio
func (pBD) GetUUID2(session *Session, self PBDRef) (retval string, err error) {
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

// Destroy: Destroy the specified PBD instance.
// Version: rio
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
// Version: rio
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

// Destroy2: Destroy the specified PBD instance.
// Version: rio
func (pBD) Destroy2(session *Session, self PBDRef) (err error) {
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

// AsyncDestroy2: Destroy the specified PBD instance.
// Version: rio
func (pBD) AsyncDestroy2(session *Session, self PBDRef) (retval TaskRef, err error) {
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

// Create: Create a new PBD instance, and return its handle. The constructor args are: host*, SR*, device_config*, other_config (* = non-optional).
// Version: rio
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
// Version: rio
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

// Create2: Create a new PBD instance, and return its handle. The constructor args are: host*, SR*, device_config*, other_config (* = non-optional).
// Version: rio
func (pBD) Create2(session *Session, args PBDRecord) (retval PBDRef, err error) {
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

// AsyncCreate2: Create a new PBD instance, and return its handle. The constructor args are: host*, SR*, device_config*, other_config (* = non-optional).
// Version: rio
func (pBD) AsyncCreate2(session *Session, args PBDRecord) (retval TaskRef, err error) {
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

// GetByUUID: Get a reference to the PBD instance with the specified UUID.
// Version: rio
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

// GetByUUID2: Get a reference to the PBD instance with the specified UUID.
// Version: rio
func (pBD) GetByUUID2(session *Session, uUID string) (retval PBDRef, err error) {
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

// GetRecord: Get a record containing the current state of the given PBD.
// Version: rio
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

// GetRecord2: Get a record containing the current state of the given PBD.
// Version: rio
func (pBD) GetRecord2(session *Session, self PBDRef) (retval PBDRecord, err error) {
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
