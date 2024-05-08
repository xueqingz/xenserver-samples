package xenapi

import (
	"fmt"
	"time"
)

type PIFMetricsRecord struct {
	// Unique identifier/object reference
	UUID string
	// Read bandwidth (KiB/s)
	IoReadKbs float64
	// Write bandwidth (KiB/s)
	IoWriteKbs float64
	// Report if the PIF got a carrier or not
	Carrier bool
	// Report vendor ID
	VendorID string
	// Report vendor name
	VendorName string
	// Report device ID
	DeviceID string
	// Report device name
	DeviceName string
	// Speed of the link in Mbit/s (if available)
	Speed int
	// Full duplex capability of the link (if available)
	Duplex bool
	// PCI bus path of the pif (if available)
	PciBusPath string
	// Time at which this information was last updated
	LastUpdated time.Time
	// additional configuration
	OtherConfig map[string]string
}

type PIFMetricsRef string

// The metrics associated with a physical network interface
type pIFMetrics struct{}

var PIFMetrics pIFMetrics

// GetRecord: Get a record containing the current state of the given PIF_metrics.
func (pIFMetrics) GetRecord(session *Session, self PIFMetricsRef) (retval PIFMetricsRecord, err error) {
	method := "PIF_metrics.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePIFMetricsRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the PIF_metrics instance with the specified UUID.
func (pIFMetrics) GetByUUID(session *Session, uUID string) (retval PIFMetricsRef, err error) {
	method := "PIF_metrics.get_by_uuid"
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
	retval, err = deserializePIFMetricsRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given PIF_metrics.
func (pIFMetrics) GetUUID(session *Session, self PIFMetricsRef) (retval string, err error) {
	method := "PIF_metrics.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIoReadKbs: Get the io/read_kbs field of the given PIF_metrics.
func (pIFMetrics) GetIoReadKbs(session *Session, self PIFMetricsRef) (retval float64, err error) {
	method := "PIF_metrics.get_io_read_kbs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIoWriteKbs: Get the io/write_kbs field of the given PIF_metrics.
func (pIFMetrics) GetIoWriteKbs(session *Session, self PIFMetricsRef) (retval float64, err error) {
	method := "PIF_metrics.get_io_write_kbs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCarrier: Get the carrier field of the given PIF_metrics.
func (pIFMetrics) GetCarrier(session *Session, self PIFMetricsRef) (retval bool, err error) {
	method := "PIF_metrics.get_carrier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVendorID: Get the vendor_id field of the given PIF_metrics.
func (pIFMetrics) GetVendorID(session *Session, self PIFMetricsRef) (retval string, err error) {
	method := "PIF_metrics.get_vendor_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVendorName: Get the vendor_name field of the given PIF_metrics.
func (pIFMetrics) GetVendorName(session *Session, self PIFMetricsRef) (retval string, err error) {
	method := "PIF_metrics.get_vendor_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDeviceID: Get the device_id field of the given PIF_metrics.
func (pIFMetrics) GetDeviceID(session *Session, self PIFMetricsRef) (retval string, err error) {
	method := "PIF_metrics.get_device_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDeviceName: Get the device_name field of the given PIF_metrics.
func (pIFMetrics) GetDeviceName(session *Session, self PIFMetricsRef) (retval string, err error) {
	method := "PIF_metrics.get_device_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSpeed: Get the speed field of the given PIF_metrics.
func (pIFMetrics) GetSpeed(session *Session, self PIFMetricsRef) (retval int, err error) {
	method := "PIF_metrics.get_speed"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDuplex: Get the duplex field of the given PIF_metrics.
func (pIFMetrics) GetDuplex(session *Session, self PIFMetricsRef) (retval bool, err error) {
	method := "PIF_metrics.get_duplex"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPciBusPath: Get the pci_bus_path field of the given PIF_metrics.
func (pIFMetrics) GetPciBusPath(session *Session, self PIFMetricsRef) (retval string, err error) {
	method := "PIF_metrics.get_pci_bus_path"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastUpdated: Get the last_updated field of the given PIF_metrics.
func (pIFMetrics) GetLastUpdated(session *Session, self PIFMetricsRef) (retval time.Time, err error) {
	method := "PIF_metrics.get_last_updated"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given PIF_metrics.
func (pIFMetrics) GetOtherConfig(session *Session, self PIFMetricsRef) (retval map[string]string, err error) {
	method := "PIF_metrics.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given PIF_metrics.
func (pIFMetrics) SetOtherConfig(session *Session, self PIFMetricsRef, value map[string]string) (err error) {
	method := "PIF_metrics.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given PIF_metrics.
func (pIFMetrics) AddToOtherConfig(session *Session, self PIFMetricsRef, key string, value string) (err error) {
	method := "PIF_metrics.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PIF_metrics.  If the key is not in that Map, then do nothing.
func (pIFMetrics) RemoveFromOtherConfig(session *Session, self PIFMetricsRef, key string) (err error) {
	method := "PIF_metrics.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePIFMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the PIF_metrics instances known to the system.
func (pIFMetrics) GetAll(session *Session) (retval []PIFMetricsRef, err error) {
	method := "PIF_metrics.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePIFMetricsRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of PIF_metrics references to PIF_metrics records for all PIF_metrics instances known to the system.
func (pIFMetrics) GetAllRecords(session *Session) (retval map[PIFMetricsRef]PIFMetricsRecord, err error) {
	method := "PIF_metrics.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePIFMetricsRefToPIFMetricsRecordMap(method+" -> ", result)
	return
}
