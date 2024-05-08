package xenapi

import (
	"fmt"
)

type CrashdumpRecord struct {
	// Unique identifier/object reference
	UUID string
	// the virtual machine
	VM VMRef
	// the virtual disk
	VDI VDIRef
	// additional configuration
	OtherConfig map[string]string
}

type CrashdumpRef string

// A VM crashdump
type crashdump struct{}

var Crashdump crashdump

// GetRecord: Get a record containing the current state of the given crashdump.
func (crashdump) GetRecord(session *Session, self CrashdumpRef) (retval CrashdumpRecord, err error) {
	method := "crashdump.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeCrashdumpRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the crashdump instance with the specified UUID.
func (crashdump) GetByUUID(session *Session, uUID string) (retval CrashdumpRef, err error) {
	method := "crashdump.get_by_uuid"
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
	retval, err = deserializeCrashdumpRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given crashdump.
func (crashdump) GetUUID(session *Session, self CrashdumpRef) (retval string, err error) {
	method := "crashdump.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVM: Get the VM field of the given crashdump.
func (crashdump) GetVM(session *Session, self CrashdumpRef) (retval VMRef, err error) {
	method := "crashdump.get_VM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// GetVDI: Get the VDI field of the given crashdump.
func (crashdump) GetVDI(session *Session, self CrashdumpRef) (retval VDIRef, err error) {
	method := "crashdump.get_VDI"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given crashdump.
func (crashdump) GetOtherConfig(session *Session, self CrashdumpRef) (retval map[string]string, err error) {
	method := "crashdump.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given crashdump.
func (crashdump) SetOtherConfig(session *Session, self CrashdumpRef, value map[string]string) (err error) {
	method := "crashdump.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given crashdump.
func (crashdump) AddToOtherConfig(session *Session, self CrashdumpRef, key string, value string) (err error) {
	method := "crashdump.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given crashdump.  If the key is not in that Map, then do nothing.
func (crashdump) RemoveFromOtherConfig(session *Session, self CrashdumpRef, key string) (err error) {
	method := "crashdump.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy: Destroy the specified crashdump
func (crashdump) Destroy(session *Session, self CrashdumpRef) (err error) {
	method := "crashdump.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified crashdump
func (crashdump) AsyncDestroy(session *Session, self CrashdumpRef) (retval TaskRef, err error) {
	method := "Async.crashdump.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the crashdumps known to the system.
func (crashdump) GetAll(session *Session) (retval []CrashdumpRef, err error) {
	method := "crashdump.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeCrashdumpRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of crashdump references to crashdump records for all crashdumps known to the system.
func (crashdump) GetAllRecords(session *Session) (retval map[CrashdumpRef]CrashdumpRecord, err error) {
	method := "crashdump.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeCrashdumpRefToCrashdumpRecordMap(method+" -> ", result)
	return
}

