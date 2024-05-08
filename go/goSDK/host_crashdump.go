package xenapi

import (
	"fmt"
	"time"
)

type HostCrashdumpRecord struct {
	// Unique identifier/object reference
	UUID string
	// Host the crashdump relates to
	Host HostRef
	// Time the crash happened
	Timestamp time.Time
	// Size of the crashdump
	Size int
	// additional configuration
	OtherConfig map[string]string
}

type HostCrashdumpRef string

// Represents a host crash dump
type hostCrashdump struct{}

var HostCrashdump hostCrashdump

// GetRecord: Get a record containing the current state of the given host_crashdump.
func (hostCrashdump) GetRecord(session *Session, self HostCrashdumpRef) (retval HostCrashdumpRecord, err error) {
	method := "host_crashdump.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostCrashdumpRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the host_crashdump instance with the specified UUID.
func (hostCrashdump) GetByUUID(session *Session, uUID string) (retval HostCrashdumpRef, err error) {
	method := "host_crashdump.get_by_uuid"
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
	retval, err = deserializeHostCrashdumpRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given host_crashdump.
func (hostCrashdump) GetUUID(session *Session, self HostCrashdumpRef) (retval string, err error) {
	method := "host_crashdump.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost: Get the host field of the given host_crashdump.
func (hostCrashdump) GetHost(session *Session, self HostCrashdumpRef) (retval HostRef, err error) {
	method := "host_crashdump.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTimestamp: Get the timestamp field of the given host_crashdump.
func (hostCrashdump) GetTimestamp(session *Session, self HostCrashdumpRef) (retval time.Time, err error) {
	method := "host_crashdump.get_timestamp"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTime(method+" -> ", result)
	return
}

// GetSize: Get the size field of the given host_crashdump.
func (hostCrashdump) GetSize(session *Session, self HostCrashdumpRef) (retval int, err error) {
	method := "host_crashdump.get_size"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given host_crashdump.
func (hostCrashdump) GetOtherConfig(session *Session, self HostCrashdumpRef) (retval map[string]string, err error) {
	method := "host_crashdump.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given host_crashdump.
func (hostCrashdump) SetOtherConfig(session *Session, self HostCrashdumpRef, value map[string]string) (err error) {
	method := "host_crashdump.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given host_crashdump.
func (hostCrashdump) AddToOtherConfig(session *Session, self HostCrashdumpRef, key string, value string) (err error) {
	method := "host_crashdump.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given host_crashdump.  If the key is not in that Map, then do nothing.
func (hostCrashdump) RemoveFromOtherConfig(session *Session, self HostCrashdumpRef, key string) (err error) {
	method := "host_crashdump.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy: Destroy specified host crash dump, removing it from the disk.
func (hostCrashdump) Destroy(session *Session, self HostCrashdumpRef) (err error) {
	method := "host_crashdump.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy specified host crash dump, removing it from the disk.
func (hostCrashdump) AsyncDestroy(session *Session, self HostCrashdumpRef) (retval TaskRef, err error) {
	method := "Async.host_crashdump.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Upload: Upload the specified host crash dump to a specified URL
func (hostCrashdump) Upload(session *Session, self HostCrashdumpRef, uRL string, options map[string]string) (err error) {
	method := "host_crashdump.upload"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	uRLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "url"), uRL)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, uRLArg, optionsArg)
	return
}

// AsyncUpload: Upload the specified host crash dump to a specified URL
func (hostCrashdump) AsyncUpload(session *Session, self HostCrashdumpRef, uRL string, options map[string]string) (retval TaskRef, err error) {
	method := "Async.host_crashdump.upload"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCrashdumpRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	uRLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "url"), uRL)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, uRLArg, optionsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetAll: Return a list of all the host_crashdumps known to the system.
func (hostCrashdump) GetAll(session *Session) (retval []HostCrashdumpRef, err error) {
	method := "host_crashdump.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostCrashdumpRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of host_crashdump references to host_crashdump records for all host_crashdumps known to the system.
func (hostCrashdump) GetAllRecords(session *Session) (retval map[HostCrashdumpRef]HostCrashdumpRecord, err error) {
	method := "host_crashdump.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostCrashdumpRefToHostCrashdumpRecordMap(method+" -> ", result)
	return
}

