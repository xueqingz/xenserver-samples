package xenapi

import (
	"fmt"
)

type HostCPURecord struct {
	// Unique identifier/object reference
	UUID string
	// the host the CPU is in
	Host HostRef
	// the number of the physical CPU within the host
	Number int
	// the vendor of the physical CPU
	Vendor string
	// the speed of the physical CPU
	Speed int
	// the model name of the physical CPU
	Modelname string
	// the family (number) of the physical CPU
	Family int
	// the model number of the physical CPU
	Model int
	// the stepping of the physical CPU
	Stepping string
	// the flags of the physical CPU (a decoded version of the features field)
	Flags string
	// the physical CPU feature bitmap
	Features string
	// the current CPU utilisation
	Utilisation float64
	// additional configuration
	OtherConfig map[string]string
}

type HostCPURef string

// A physical CPU
type hostCPU struct{}

var HostCPU hostCPU

// GetRecord: Get a record containing the current state of the given host_cpu.
func (hostCPU) GetRecord(session *Session, self HostCPURef) (retval HostCPURecord, err error) {
	method := "host_cpu.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostCPURecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the host_cpu instance with the specified UUID.
func (hostCPU) GetByUUID(session *Session, uUID string) (retval HostCPURef, err error) {
	method := "host_cpu.get_by_uuid"
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
	retval, err = deserializeHostCPURef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given host_cpu.
func (hostCPU) GetUUID(session *Session, self HostCPURef) (retval string, err error) {
	method := "host_cpu.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost: Get the host field of the given host_cpu.
func (hostCPU) GetHost(session *Session, self HostCPURef) (retval HostRef, err error) {
	method := "host_cpu.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNumber: Get the number field of the given host_cpu.
func (hostCPU) GetNumber(session *Session, self HostCPURef) (retval int, err error) {
	method := "host_cpu.get_number"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVendor: Get the vendor field of the given host_cpu.
func (hostCPU) GetVendor(session *Session, self HostCPURef) (retval string, err error) {
	method := "host_cpu.get_vendor"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSpeed: Get the speed field of the given host_cpu.
func (hostCPU) GetSpeed(session *Session, self HostCPURef) (retval int, err error) {
	method := "host_cpu.get_speed"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetModelname: Get the modelname field of the given host_cpu.
func (hostCPU) GetModelname(session *Session, self HostCPURef) (retval string, err error) {
	method := "host_cpu.get_modelname"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetFamily: Get the family field of the given host_cpu.
func (hostCPU) GetFamily(session *Session, self HostCPURef) (retval int, err error) {
	method := "host_cpu.get_family"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetModel: Get the model field of the given host_cpu.
func (hostCPU) GetModel(session *Session, self HostCPURef) (retval int, err error) {
	method := "host_cpu.get_model"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStepping: Get the stepping field of the given host_cpu.
func (hostCPU) GetStepping(session *Session, self HostCPURef) (retval string, err error) {
	method := "host_cpu.get_stepping"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetFlags: Get the flags field of the given host_cpu.
func (hostCPU) GetFlags(session *Session, self HostCPURef) (retval string, err error) {
	method := "host_cpu.get_flags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetFeatures: Get the features field of the given host_cpu.
func (hostCPU) GetFeatures(session *Session, self HostCPURef) (retval string, err error) {
	method := "host_cpu.get_features"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUtilisation: Get the utilisation field of the given host_cpu.
func (hostCPU) GetUtilisation(session *Session, self HostCPURef) (retval float64, err error) {
	method := "host_cpu.get_utilisation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given host_cpu.
func (hostCPU) GetOtherConfig(session *Session, self HostCPURef) (retval map[string]string, err error) {
	method := "host_cpu.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given host_cpu.
func (hostCPU) SetOtherConfig(session *Session, self HostCPURef, value map[string]string) (err error) {
	method := "host_cpu.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given host_cpu.
func (hostCPU) AddToOtherConfig(session *Session, self HostCPURef, key string, value string) (err error) {
	method := "host_cpu.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given host_cpu.  If the key is not in that Map, then do nothing.
func (hostCPU) RemoveFromOtherConfig(session *Session, self HostCPURef, key string) (err error) {
	method := "host_cpu.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostCPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the host_cpus known to the system.
func (hostCPU) GetAll(session *Session) (retval []HostCPURef, err error) {
	method := "host_cpu.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostCPURefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of host_cpu references to host_cpu records for all host_cpus known to the system.
func (hostCPU) GetAllRecords(session *Session) (retval map[HostCPURef]HostCPURecord, err error) {
	method := "host_cpu.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostCPURefToHostCPURecordMap(method+" -> ", result)
	return
}

