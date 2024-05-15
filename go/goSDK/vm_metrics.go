package xenapi

import (
	"fmt"
	"time"
)

type VMMetricsRecord struct {
	// Unique identifier/object reference
	UUID string
	// Guest&apos;s actual memory (bytes)
	MemoryActual int
	// Current number of VCPUs
	VCPUsNumber int
	// Utilisation for all of guest&apos;s current VCPUs
	VCPUsUtilisation map[int]float64
	// VCPU to PCPU map
	VCPUsCPU map[int]int
	// The live equivalent to VM.VCPUs_params
	VCPUsParams map[string]string
	// CPU flags (blocked,online,running)
	VCPUsFlags map[int][]string
	// The state of the guest, eg blocked, dying etc
	State []string
	// Time at which this VM was last booted
	StartTime time.Time
	// Time at which the VM was installed
	InstallTime time.Time
	// Time at which this information was last updated
	LastUpdated time.Time
	// additional configuration
	OtherConfig map[string]string
	// hardware virtual machine
	Hvm bool
	// VM supports nested virtualisation
	NestedVirt bool
	// VM is immobile and can&apos;t migrate between hosts
	Nomigrate bool
	// The current domain type of the VM (for running,suspended, or paused VMs). The last-known domain type for halted VMs.
	CurrentDomainType DomainType
}

type VMMetricsRef string

// The metrics associated with a VM
type vMMetrics struct{}

var VMMetrics vMMetrics

// GetAllRecords: Return a map of VM_metrics references to VM_metrics records for all VM_metrics instances known to the system.
// Version: rio
func (vMMetrics) GetAllRecords(session *Session) (retval map[VMMetricsRef]VMMetricsRecord, err error) {
	method := "VM_metrics.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMMetricsRefToVMMetricsRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of VM_metrics references to VM_metrics records for all VM_metrics instances known to the system.
// Version: rio
func (vMMetrics) GetAllRecords1(session *Session) (retval map[VMMetricsRef]VMMetricsRecord, err error) {
	method := "VM_metrics.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMMetricsRefToVMMetricsRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the VM_metrics instances known to the system.
// Version: rio
func (vMMetrics) GetAll(session *Session) (retval []VMMetricsRef, err error) {
	method := "VM_metrics.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMMetricsRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the VM_metrics instances known to the system.
// Version: rio
func (vMMetrics) GetAll1(session *Session) (retval []VMMetricsRef, err error) {
	method := "VM_metrics.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMMetricsRefSet(method+" -> ", result)
	return
}

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VM_metrics.  If the key is not in that Map, then do nothing.
// Version: orlando
func (vMMetrics) RemoveFromOtherConfig(session *Session, self VMMetricsRef, key string) (err error) {
	method := "VM_metrics.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig3: Remove the given key and its corresponding value from the other_config field of the given VM_metrics.  If the key is not in that Map, then do nothing.
// Version: orlando
func (vMMetrics) RemoveFromOtherConfig3(session *Session, self VMMetricsRef, key string) (err error) {
	method := "VM_metrics.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig2: Remove the given key and its corresponding value from the other_config field of the given VM_metrics.  If the key is not in that Map, then do nothing.
// Version: rio
func (vMMetrics) RemoveFromOtherConfig2(session *Session, self VMMetricsRef) (err error) {
	method := "VM_metrics.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VM_metrics.
// Version: orlando
func (vMMetrics) AddToOtherConfig(session *Session, self VMMetricsRef, key string, value string) (err error) {
	method := "VM_metrics.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig4: Add the given key-value pair to the other_config field of the given VM_metrics.
// Version: orlando
func (vMMetrics) AddToOtherConfig4(session *Session, self VMMetricsRef, key string, value string) (err error) {
	method := "VM_metrics.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig2: Add the given key-value pair to the other_config field of the given VM_metrics.
// Version: rio
func (vMMetrics) AddToOtherConfig2(session *Session, self VMMetricsRef) (err error) {
	method := "VM_metrics.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// SetOtherConfig: Set the other_config field of the given VM_metrics.
// Version: orlando
func (vMMetrics) SetOtherConfig(session *Session, self VMMetricsRef, value map[string]string) (err error) {
	method := "VM_metrics.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig3: Set the other_config field of the given VM_metrics.
// Version: orlando
func (vMMetrics) SetOtherConfig3(session *Session, self VMMetricsRef, value map[string]string) (err error) {
	method := "VM_metrics.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig2: Set the other_config field of the given VM_metrics.
// Version: rio
func (vMMetrics) SetOtherConfig2(session *Session, self VMMetricsRef) (err error) {
	method := "VM_metrics.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// GetCurrentDomainType: Get the current_domain_type field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetCurrentDomainType(session *Session, self VMMetricsRef) (retval DomainType, err error) {
	method := "VM_metrics.get_current_domain_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumDomainType(method+" -> ", result)
	return
}

// GetCurrentDomainType2: Get the current_domain_type field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetCurrentDomainType2(session *Session, self VMMetricsRef) (retval DomainType, err error) {
	method := "VM_metrics.get_current_domain_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumDomainType(method+" -> ", result)
	return
}

// GetNomigrate: Get the nomigrate field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetNomigrate(session *Session, self VMMetricsRef) (retval bool, err error) {
	method := "VM_metrics.get_nomigrate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNomigrate2: Get the nomigrate field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetNomigrate2(session *Session, self VMMetricsRef) (retval bool, err error) {
	method := "VM_metrics.get_nomigrate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNestedVirt: Get the nested_virt field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetNestedVirt(session *Session, self VMMetricsRef) (retval bool, err error) {
	method := "VM_metrics.get_nested_virt"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNestedVirt2: Get the nested_virt field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetNestedVirt2(session *Session, self VMMetricsRef) (retval bool, err error) {
	method := "VM_metrics.get_nested_virt"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHvm: Get the hvm field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetHvm(session *Session, self VMMetricsRef) (retval bool, err error) {
	method := "VM_metrics.get_hvm"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHvm2: Get the hvm field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetHvm2(session *Session, self VMMetricsRef) (retval bool, err error) {
	method := "VM_metrics.get_hvm"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetOtherConfig(session *Session, self VMMetricsRef) (retval map[string]string, err error) {
	method := "VM_metrics.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig2: Get the other_config field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetOtherConfig2(session *Session, self VMMetricsRef) (retval map[string]string, err error) {
	method := "VM_metrics.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastUpdated: Get the last_updated field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetLastUpdated(session *Session, self VMMetricsRef) (retval time.Time, err error) {
	method := "VM_metrics.get_last_updated"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastUpdated2: Get the last_updated field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetLastUpdated2(session *Session, self VMMetricsRef) (retval time.Time, err error) {
	method := "VM_metrics.get_last_updated"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetInstallTime: Get the install_time field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetInstallTime(session *Session, self VMMetricsRef) (retval time.Time, err error) {
	method := "VM_metrics.get_install_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetInstallTime2: Get the install_time field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetInstallTime2(session *Session, self VMMetricsRef) (retval time.Time, err error) {
	method := "VM_metrics.get_install_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStartTime: Get the start_time field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetStartTime(session *Session, self VMMetricsRef) (retval time.Time, err error) {
	method := "VM_metrics.get_start_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStartTime2: Get the start_time field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetStartTime2(session *Session, self VMMetricsRef) (retval time.Time, err error) {
	method := "VM_metrics.get_start_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetState: Get the state field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetState(session *Session, self VMMetricsRef) (retval []string, err error) {
	method := "VM_metrics.get_state"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// GetState2: Get the state field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetState2(session *Session, self VMMetricsRef) (retval []string, err error) {
	method := "VM_metrics.get_state"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// GetVCPUsFlags: Get the VCPUs/flags field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsFlags(session *Session, self VMMetricsRef) (retval map[int][]string, err error) {
	method := "VM_metrics.get_VCPUs_flags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeIntToStringSetMap(method+" -> ", result)
	return
}

// GetVCPUsFlags2: Get the VCPUs/flags field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsFlags2(session *Session, self VMMetricsRef) (retval map[int][]string, err error) {
	method := "VM_metrics.get_VCPUs_flags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeIntToStringSetMap(method+" -> ", result)
	return
}

// GetVCPUsParams: Get the VCPUs/params field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsParams(session *Session, self VMMetricsRef) (retval map[string]string, err error) {
	method := "VM_metrics.get_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVCPUsParams2: Get the VCPUs/params field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsParams2(session *Session, self VMMetricsRef) (retval map[string]string, err error) {
	method := "VM_metrics.get_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVCPUsCPU: Get the VCPUs/CPU field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsCPU(session *Session, self VMMetricsRef) (retval map[int]int, err error) {
	method := "VM_metrics.get_VCPUs_CPU"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeIntToIntMap(method+" -> ", result)
	return
}

// GetVCPUsCPU2: Get the VCPUs/CPU field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsCPU2(session *Session, self VMMetricsRef) (retval map[int]int, err error) {
	method := "VM_metrics.get_VCPUs_CPU"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeIntToIntMap(method+" -> ", result)
	return
}

// GetVCPUsUtilisation: Get the VCPUs/utilisation field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsUtilisation(session *Session, self VMMetricsRef) (retval map[int]float64, err error) {
	method := "VM_metrics.get_VCPUs_utilisation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeIntToFloatMap(method+" -> ", result)
	return
}

// GetVCPUsUtilisation2: Get the VCPUs/utilisation field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsUtilisation2(session *Session, self VMMetricsRef) (retval map[int]float64, err error) {
	method := "VM_metrics.get_VCPUs_utilisation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeIntToFloatMap(method+" -> ", result)
	return
}

// GetVCPUsNumber: Get the VCPUs/number field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsNumber(session *Session, self VMMetricsRef) (retval int, err error) {
	method := "VM_metrics.get_VCPUs_number"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVCPUsNumber2: Get the VCPUs/number field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetVCPUsNumber2(session *Session, self VMMetricsRef) (retval int, err error) {
	method := "VM_metrics.get_VCPUs_number"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryActual: Get the memory/actual field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetMemoryActual(session *Session, self VMMetricsRef) (retval int, err error) {
	method := "VM_metrics.get_memory_actual"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryActual2: Get the memory/actual field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetMemoryActual2(session *Session, self VMMetricsRef) (retval int, err error) {
	method := "VM_metrics.get_memory_actual"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetUUID(session *Session, self VMMetricsRef) (retval string, err error) {
	method := "VM_metrics.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given VM_metrics.
// Version: rio
func (vMMetrics) GetUUID2(session *Session, self VMMetricsRef) (retval string, err error) {
	method := "VM_metrics.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByUUID: Get a reference to the VM_metrics instance with the specified UUID.
// Version: rio
func (vMMetrics) GetByUUID(session *Session, uUID string) (retval VMMetricsRef, err error) {
	method := "VM_metrics.get_by_uuid"
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
	retval, err = deserializeVMMetricsRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the VM_metrics instance with the specified UUID.
// Version: rio
func (vMMetrics) GetByUUID2(session *Session, uUID string) (retval VMMetricsRef, err error) {
	method := "VM_metrics.get_by_uuid"
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
	retval, err = deserializeVMMetricsRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given VM_metrics.
// Version: rio
func (vMMetrics) GetRecord(session *Session, self VMMetricsRef) (retval VMMetricsRecord, err error) {
	method := "VM_metrics.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMMetricsRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given VM_metrics.
// Version: rio
func (vMMetrics) GetRecord2(session *Session, self VMMetricsRef) (retval VMMetricsRecord, err error) {
	method := "VM_metrics.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMMetricsRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMMetricsRecord(method+" -> ", result)
	return
}
