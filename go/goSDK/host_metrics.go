package xenapi

import (
	"fmt"
	"time"
)

type HostMetricsRecord struct {
	// Unique identifier/object reference
	UUID string
	// Total host memory (bytes)
	MemoryTotal int
	// Free host memory (bytes)
	MemoryFree int
	// Pool master thinks this host is live
	Live bool
	// Time at which this information was last updated
	LastUpdated time.Time
	// additional configuration
	OtherConfig map[string]string
}

type HostMetricsRef string

// The metrics associated with a host
type hostMetrics struct{}

var HostMetrics hostMetrics

// GetRecord: Get a record containing the current state of the given host_metrics.
func (hostMetrics) GetRecord(session *Session, self HostMetricsRef) (retval HostMetricsRecord, err error) {
	method := "host_metrics.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostMetricsRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the host_metrics instance with the specified UUID.
func (hostMetrics) GetByUUID(session *Session, uUID string) (retval HostMetricsRef, err error) {
	method := "host_metrics.get_by_uuid"
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
	retval, err = deserializeHostMetricsRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given host_metrics.
func (hostMetrics) GetUUID(session *Session, self HostMetricsRef) (retval string, err error) {
	method := "host_metrics.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryTotal: Get the memory/total field of the given host_metrics.
func (hostMetrics) GetMemoryTotal(session *Session, self HostMetricsRef) (retval int, err error) {
	method := "host_metrics.get_memory_total"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryFree: Get the memory/free field of the given host_metrics.
func (hostMetrics) GetMemoryFree(session *Session, self HostMetricsRef) (retval int, err error) {
	method := "host_metrics.get_memory_free"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLive: Get the live field of the given host_metrics.
func (hostMetrics) GetLive(session *Session, self HostMetricsRef) (retval bool, err error) {
	method := "host_metrics.get_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastUpdated: Get the last_updated field of the given host_metrics.
func (hostMetrics) GetLastUpdated(session *Session, self HostMetricsRef) (retval time.Time, err error) {
	method := "host_metrics.get_last_updated"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given host_metrics.
func (hostMetrics) GetOtherConfig(session *Session, self HostMetricsRef) (retval map[string]string, err error) {
	method := "host_metrics.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given host_metrics.
func (hostMetrics) SetOtherConfig(session *Session, self HostMetricsRef, value map[string]string) (err error) {
	method := "host_metrics.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given host_metrics.
func (hostMetrics) AddToOtherConfig(session *Session, self HostMetricsRef, key string, value string) (err error) {
	method := "host_metrics.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given host_metrics.  If the key is not in that Map, then do nothing.
func (hostMetrics) RemoveFromOtherConfig(session *Session, self HostMetricsRef, key string) (err error) {
	method := "host_metrics.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the host_metrics instances known to the system.
func (hostMetrics) GetAll(session *Session) (retval []HostMetricsRef, err error) {
	method := "host_metrics.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostMetricsRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of host_metrics references to host_metrics records for all host_metrics instances known to the system.
func (hostMetrics) GetAllRecords(session *Session) (retval map[HostMetricsRef]HostMetricsRecord, err error) {
	method := "host_metrics.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostMetricsRefToHostMetricsRecordMap(method+" -> ", result)
	return
}

