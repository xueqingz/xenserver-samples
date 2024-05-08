package xenapi

import (
	"fmt"
	"time"
)

type VIFMetricsRecord struct {
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

type VIFMetricsRef string

// The metrics associated with a virtual network device
type vIFMetrics struct{}

var VIFMetrics vIFMetrics

// GetRecord: Get a record containing the current state of the given VIF_metrics.
func (vIFMetrics) GetRecord(session *Session, self VIFMetricsRef) (retval VIFMetricsRecord, err error) {
	method := "VIF_metrics.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVIFMetricsRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VIF_metrics instance with the specified UUID.
func (vIFMetrics) GetByUUID(session *Session, uUID string) (retval VIFMetricsRef, err error) {
	method := "VIF_metrics.get_by_uuid"
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
	retval, err = deserializeVIFMetricsRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VIF_metrics.
func (vIFMetrics) GetUUID(session *Session, self VIFMetricsRef) (retval string, err error) {
	method := "VIF_metrics.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIoReadKbs: Get the io/read_kbs field of the given VIF_metrics.
func (vIFMetrics) GetIoReadKbs(session *Session, self VIFMetricsRef) (retval float64, err error) {
	method := "VIF_metrics.get_io_read_kbs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIoWriteKbs: Get the io/write_kbs field of the given VIF_metrics.
func (vIFMetrics) GetIoWriteKbs(session *Session, self VIFMetricsRef) (retval float64, err error) {
	method := "VIF_metrics.get_io_write_kbs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastUpdated: Get the last_updated field of the given VIF_metrics.
func (vIFMetrics) GetLastUpdated(session *Session, self VIFMetricsRef) (retval time.Time, err error) {
	method := "VIF_metrics.get_last_updated"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given VIF_metrics.
func (vIFMetrics) GetOtherConfig(session *Session, self VIFMetricsRef) (retval map[string]string, err error) {
	method := "VIF_metrics.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given VIF_metrics.
func (vIFMetrics) SetOtherConfig(session *Session, self VIFMetricsRef, value map[string]string) (err error) {
	method := "VIF_metrics.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VIF_metrics.
func (vIFMetrics) AddToOtherConfig(session *Session, self VIFMetricsRef, key string, value string) (err error) {
	method := "VIF_metrics.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VIF_metrics.  If the key is not in that Map, then do nothing.
func (vIFMetrics) RemoveFromOtherConfig(session *Session, self VIFMetricsRef, key string) (err error) {
	method := "VIF_metrics.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the VIF_metrics instances known to the system.
func (vIFMetrics) GetAll(session *Session) (retval []VIFMetricsRef, err error) {
	method := "VIF_metrics.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVIFMetricsRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VIF_metrics references to VIF_metrics records for all VIF_metrics instances known to the system.
func (vIFMetrics) GetAllRecords(session *Session) (retval map[VIFMetricsRef]VIFMetricsRecord, err error) {
	method := "VIF_metrics.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVIFMetricsRefToVIFMetricsRecordMap(method+" -> ", result)
	return
}

