package xenapi

import (
	"fmt"
)

type VGPURecord struct {
	// Unique identifier/object reference
	UUID string
	// VM that owns the vGPU
	VM VMRef
	// GPU group used by the vGPU
	GPUGroup GPUGroupRef
	// Order in which the devices are plugged into the VM
	Device string
	// Reflects whether the virtual device is currently connected to a physical device
	CurrentlyAttached bool
	// Additional configuration
	OtherConfig map[string]string
	// Preset type for this VGPU
	Type VGPUTypeRef
	// The PGPU on which this VGPU is running
	ResidentOn PGPURef
	// The PGPU on which this VGPU is scheduled to run
	ScheduledToBeResidentOn PGPURef
	// VGPU metadata to determine whether a VGPU can migrate between two PGPUs
	CompatibilityMetadata map[string]string
	// Extra arguments for vGPU and passed to demu
	ExtraArgs string
	// Device passed trough to VM, either as full device or SR-IOV virtual function
	PCI PCIRef
}

type VGPURef string

// A virtual GPU (vGPU)
type vGPU struct{}

var VGPU vGPU

// GetRecord: Get a record containing the current state of the given VGPU.
func (vGPU) GetRecord(session *Session, self VGPURef) (retval VGPURecord, err error) {
	method := "VGPU.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPURecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VGPU instance with the specified UUID.
func (vGPU) GetByUUID(session *Session, uUID string) (retval VGPURef, err error) {
	method := "VGPU.get_by_uuid"
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
	retval, err = deserializeVGPURef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VGPU.
func (vGPU) GetUUID(session *Session, self VGPURef) (retval string, err error) {
	method := "VGPU.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVM: Get the VM field of the given VGPU.
func (vGPU) GetVM(session *Session, self VGPURef) (retval VMRef, err error) {
	method := "VGPU.get_VM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetGPUGroup: Get the GPU_group field of the given VGPU.
func (vGPU) GetGPUGroup(session *Session, self VGPURef) (retval GPUGroupRef, err error) {
	method := "VGPU.get_GPU_group"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRef(method+" -> ", result)
	return
}

// GetDevice: Get the device field of the given VGPU.
func (vGPU) GetDevice(session *Session, self VGPURef) (retval string, err error) {
	method := "VGPU.get_device"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCurrentlyAttached: Get the currently_attached field of the given VGPU.
func (vGPU) GetCurrentlyAttached(session *Session, self VGPURef) (retval bool, err error) {
	method := "VGPU.get_currently_attached"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given VGPU.
func (vGPU) GetOtherConfig(session *Session, self VGPURef) (retval map[string]string, err error) {
	method := "VGPU.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetType: Get the type field of the given VGPU.
func (vGPU) GetType(session *Session, self VGPURef) (retval VGPUTypeRef, err error) {
	method := "VGPU.get_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRef(method+" -> ", result)
	return
}

// GetResidentOn: Get the resident_on field of the given VGPU.
func (vGPU) GetResidentOn(session *Session, self VGPURef) (retval PGPURef, err error) {
	method := "VGPU.get_resident_on"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePGPURef(method+" -> ", result)
	return
}

// GetScheduledToBeResidentOn: Get the scheduled_to_be_resident_on field of the given VGPU.
func (vGPU) GetScheduledToBeResidentOn(session *Session, self VGPURef) (retval PGPURef, err error) {
	method := "VGPU.get_scheduled_to_be_resident_on"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePGPURef(method+" -> ", result)
	return
}

// GetCompatibilityMetadata: Get the compatibility_metadata field of the given VGPU.
func (vGPU) GetCompatibilityMetadata(session *Session, self VGPURef) (retval map[string]string, err error) {
	method := "VGPU.get_compatibility_metadata"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetExtraArgs: Get the extra_args field of the given VGPU.
func (vGPU) GetExtraArgs(session *Session, self VGPURef) (retval string, err error) {
	method := "VGPU.get_extra_args"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPCI: Get the PCI field of the given VGPU.
func (vGPU) GetPCI(session *Session, self VGPURef) (retval PCIRef, err error) {
	method := "VGPU.get_PCI"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePCIRef(method+" -> ", result)
	return
}

// SetOtherConfig: Set the other_config field of the given VGPU.
func (vGPU) SetOtherConfig(session *Session, self VGPURef, value map[string]string) (err error) {
	method := "VGPU.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VGPU.
func (vGPU) AddToOtherConfig(session *Session, self VGPURef, key string, value string) (err error) {
	method := "VGPU.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VGPU.  If the key is not in that Map, then do nothing.
func (vGPU) RemoveFromOtherConfig(session *Session, self VGPURef, key string) (err error) {
	method := "VGPU.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetExtraArgs: Set the extra_args field of the given VGPU.
func (vGPU) SetExtraArgs(session *Session, self VGPURef, value string) (err error) {
	method := "VGPU.set_extra_args"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// Create:
func (vGPU) Create(session *Session, vM VMRef, gPUGroup GPUGroupRef, device string, otherConfig map[string]string, typeKey VGPUTypeRef) (retval VGPURef, err error) {
	method := "VGPU.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "VM"), vM)
	if err != nil {
		return
	}
	gPUGroupArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "GPU_group"), gPUGroup)
	if err != nil {
		return
	}
	deviceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "device"), device)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, gPUGroupArg, deviceArg, otherConfigArg, typeKeyArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPURef(method+" -> ", result)
	return
}

// AsyncCreate:
func (vGPU) AsyncCreate(session *Session, vM VMRef, gPUGroup GPUGroupRef, device string, otherConfig map[string]string, typeKey VGPUTypeRef) (retval TaskRef, err error) {
	method := "Async.VGPU.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "VM"), vM)
	if err != nil {
		return
	}
	gPUGroupArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "GPU_group"), gPUGroup)
	if err != nil {
		return
	}
	deviceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "device"), device)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, gPUGroupArg, deviceArg, otherConfigArg, typeKeyArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy:
func (vGPU) Destroy(session *Session, self VGPURef) (err error) {
	method := "VGPU.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy:
func (vGPU) AsyncDestroy(session *Session, self VGPURef) (retval TaskRef, err error) {
	method := "Async.VGPU.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the VGPUs known to the system.
func (vGPU) GetAll(session *Session) (retval []VGPURef, err error) {
	method := "VGPU.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPURefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VGPU references to VGPU records for all VGPUs known to the system.
func (vGPU) GetAllRecords(session *Session) (retval map[VGPURef]VGPURecord, err error) {
	method := "VGPU.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPURefToVGPURecordMap(method+" -> ", result)
	return
}

