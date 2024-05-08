package xenapi

import (
	"fmt"
	"time"
)

type VMGuestMetricsRecord struct {
	// Unique identifier/object reference
	UUID string
	// version of the OS
	OsVersion map[string]string
	// version of the PV drivers
	PVDriversVersion map[string]string
	// Logically equivalent to PV_drivers_detected
	PVDriversUpToDate bool
	// This field exists but has no data. Use the memory and memory_internal_free RRD data-sources instead.
	Memory map[string]string
	// This field exists but has no data.
	Disks map[string]string
	// network configuration
	Networks map[string]string
	// anything else
	Other map[string]string
	// Time at which this information was last updated
	LastUpdated time.Time
	// additional configuration
	OtherConfig map[string]string
	// True if the guest is sending heartbeat messages via the guest agent
	Live bool
	// The guest&apos;s statement of whether it supports VBD hotplug, i.e. whether it is capable of responding immediately to instantiation of a new VBD by bringing online a new PV block device. If the guest states that it is not capable, then the VBD plug and unplug operations will not be allowed while the guest is running.
	CanUseHotplugVbd TristateType
	// The guest&apos;s statement of whether it supports VIF hotplug, i.e. whether it is capable of responding immediately to instantiation of a new VIF by bringing online a new PV network device. If the guest states that it is not capable, then the VIF plug and unplug operations will not be allowed while the guest is running.
	CanUseHotplugVif TristateType
	// At least one of the guest&apos;s devices has successfully connected to the backend.
	PVDriversDetected bool
}

type VMGuestMetricsRef string

// The metrics reported by the guest (as opposed to inferred from outside)
type vMGuestMetrics struct{}

var VMGuestMetrics vMGuestMetrics

// GetRecord: Get a record containing the current state of the given VM_guest_metrics.
func (vMGuestMetrics) GetRecord(session *Session, self VMGuestMetricsRef) (retval VMGuestMetricsRecord, err error) {
	method := "VM_guest_metrics.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMGuestMetricsRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VM_guest_metrics instance with the specified UUID.
func (vMGuestMetrics) GetByUUID(session *Session, uUID string) (retval VMGuestMetricsRef, err error) {
	method := "VM_guest_metrics.get_by_uuid"
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
	retval, err = deserializeVMGuestMetricsRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VM_guest_metrics.
func (vMGuestMetrics) GetUUID(session *Session, self VMGuestMetricsRef) (retval string, err error) {
	method := "VM_guest_metrics.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOsVersion: Get the os_version field of the given VM_guest_metrics.
func (vMGuestMetrics) GetOsVersion(session *Session, self VMGuestMetricsRef) (retval map[string]string, err error) {
	method := "VM_guest_metrics.get_os_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPVDriversVersion: Get the PV_drivers_version field of the given VM_guest_metrics.
func (vMGuestMetrics) GetPVDriversVersion(session *Session, self VMGuestMetricsRef) (retval map[string]string, err error) {
	method := "VM_guest_metrics.get_PV_drivers_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPVDriversUpToDate: Get the PV_drivers_up_to_date field of the given VM_guest_metrics.
func (vMGuestMetrics) GetPVDriversUpToDate(session *Session, self VMGuestMetricsRef) (retval bool, err error) {
	method := "VM_guest_metrics.get_PV_drivers_up_to_date"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemory: Get the memory field of the given VM_guest_metrics.
func (vMGuestMetrics) GetMemory(session *Session, self VMGuestMetricsRef) (retval map[string]string, err error) {
	method := "VM_guest_metrics.get_memory"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDisks: Get the disks field of the given VM_guest_metrics.
func (vMGuestMetrics) GetDisks(session *Session, self VMGuestMetricsRef) (retval map[string]string, err error) {
	method := "VM_guest_metrics.get_disks"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNetworks: Get the networks field of the given VM_guest_metrics.
func (vMGuestMetrics) GetNetworks(session *Session, self VMGuestMetricsRef) (retval map[string]string, err error) {
	method := "VM_guest_metrics.get_networks"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOther: Get the other field of the given VM_guest_metrics.
func (vMGuestMetrics) GetOther(session *Session, self VMGuestMetricsRef) (retval map[string]string, err error) {
	method := "VM_guest_metrics.get_other"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastUpdated: Get the last_updated field of the given VM_guest_metrics.
func (vMGuestMetrics) GetLastUpdated(session *Session, self VMGuestMetricsRef) (retval time.Time, err error) {
	method := "VM_guest_metrics.get_last_updated"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given VM_guest_metrics.
func (vMGuestMetrics) GetOtherConfig(session *Session, self VMGuestMetricsRef) (retval map[string]string, err error) {
	method := "VM_guest_metrics.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLive: Get the live field of the given VM_guest_metrics.
func (vMGuestMetrics) GetLive(session *Session, self VMGuestMetricsRef) (retval bool, err error) {
	method := "VM_guest_metrics.get_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCanUseHotplugVbd: Get the can_use_hotplug_vbd field of the given VM_guest_metrics.
func (vMGuestMetrics) GetCanUseHotplugVbd(session *Session, self VMGuestMetricsRef) (retval TristateType, err error) {
	method := "VM_guest_metrics.get_can_use_hotplug_vbd"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumTristateType(method+" -> ", result)
	return
}

// GetCanUseHotplugVif: Get the can_use_hotplug_vif field of the given VM_guest_metrics.
func (vMGuestMetrics) GetCanUseHotplugVif(session *Session, self VMGuestMetricsRef) (retval TristateType, err error) {
	method := "VM_guest_metrics.get_can_use_hotplug_vif"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumTristateType(method+" -> ", result)
	return
}

// GetPVDriversDetected: Get the PV_drivers_detected field of the given VM_guest_metrics.
func (vMGuestMetrics) GetPVDriversDetected(session *Session, self VMGuestMetricsRef) (retval bool, err error) {
	method := "VM_guest_metrics.get_PV_drivers_detected"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given VM_guest_metrics.
func (vMGuestMetrics) SetOtherConfig(session *Session, self VMGuestMetricsRef, value map[string]string) (err error) {
	method := "VM_guest_metrics.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VM_guest_metrics.
func (vMGuestMetrics) AddToOtherConfig(session *Session, self VMGuestMetricsRef, key string, value string) (err error) {
	method := "VM_guest_metrics.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VM_guest_metrics.  If the key is not in that Map, then do nothing.
func (vMGuestMetrics) RemoveFromOtherConfig(session *Session, self VMGuestMetricsRef, key string) (err error) {
	method := "VM_guest_metrics.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMGuestMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the VM_guest_metrics instances known to the system.
func (vMGuestMetrics) GetAll(session *Session) (retval []VMGuestMetricsRef, err error) {
	method := "VM_guest_metrics.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMGuestMetricsRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VM_guest_metrics references to VM_guest_metrics records for all VM_guest_metrics instances known to the system.
func (vMGuestMetrics) GetAllRecords(session *Session) (retval map[VMGuestMetricsRef]VMGuestMetricsRecord, err error) {
	method := "VM_guest_metrics.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMGuestMetricsRefToVMGuestMetricsRecordMap(method+" -> ", result)
	return
}

