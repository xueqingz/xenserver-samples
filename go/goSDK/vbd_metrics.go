package xenapi

import (
	"fmt"
	"time"
)

type VBDMetricsRecord struct {
	// Unique identifier/object reference
	UUID string
	// Read bandwidth (KiB/s)
	IoReadKbs float64
	// Write bandwidth (KiB/s)
	IoWriteKbs float64
	// Time at which this information was last updated
	LastUpdated time.Time
	// additional configuration
	OtherConfig map[string]string
}

type VBDMetricsRef string

// The metrics associated with a virtual block device
type vBDMetrics struct{}

var VBDMetrics vBDMetrics

// GetRecord: Get a record containing the current state of the given VBD_metrics.
func (vBDMetrics) GetRecord(session *Session, self VBDMetricsRef) (retval VBDMetricsRecord, err error) {
	method := "VBD_metrics.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDMetricsRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VBD_metrics instance with the specified UUID.
func (vBDMetrics) GetByUUID(session *Session, uUID string) (retval VBDMetricsRef, err error) {
	method := "VBD_metrics.get_by_uuid"
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
	retval, err = deserializeVBDMetricsRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VBD_metrics.
func (vBDMetrics) GetUUID(session *Session, self VBDMetricsRef) (retval string, err error) {
	method := "VBD_metrics.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIoReadKbs: Get the io/read_kbs field of the given VBD_metrics.
func (vBDMetrics) GetIoReadKbs(session *Session, self VBDMetricsRef) (retval float64, err error) {
	method := "VBD_metrics.get_io_read_kbs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeFloat(method+" -> ", result)
	return
}

// GetIoWriteKbs: Get the io/write_kbs field of the given VBD_metrics.
func (vBDMetrics) GetIoWriteKbs(session *Session, self VBDMetricsRef) (retval float64, err error) {
	method := "VBD_metrics.get_io_write_kbs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeFloat(method+" -> ", result)
	return
}

// GetLastUpdated: Get the last_updated field of the given VBD_metrics.
func (vBDMetrics) GetLastUpdated(session *Session, self VBDMetricsRef) (retval time.Time, err error) {
	method := "VBD_metrics.get_last_updated"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given VBD_metrics.
func (vBDMetrics) GetOtherConfig(session *Session, self VBDMetricsRef) (retval map[string]string, err error) {
	method := "VBD_metrics.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given VBD_metrics.
func (vBDMetrics) SetOtherConfig(session *Session, self VBDMetricsRef, value map[string]string) (err error) {
	method := "VBD_metrics.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VBD_metrics.
func (vBDMetrics) AddToOtherConfig(session *Session, self VBDMetricsRef, key string, value string) (err error) {
	method := "VBD_metrics.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VBD_metrics.  If the key is not in that Map, then do nothing.
func (vBDMetrics) RemoveFromOtherConfig(session *Session, self VBDMetricsRef, key string) (err error) {
	method := "VBD_metrics.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the VBD_metrics instances known to the system.
func (vBDMetrics) GetAll(session *Session) (retval []VBDMetricsRef, err error) {
	method := "VBD_metrics.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDMetricsRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VBD_metrics references to VBD_metrics records for all VBD_metrics instances known to the system.
func (vBDMetrics) GetAllRecords(session *Session) (retval map[VBDMetricsRef]VBDMetricsRecord, err error) {
	method := "VBD_metrics.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDMetricsRefToVBDMetricsRecordMap(method+" -> ", result)
	return
}
