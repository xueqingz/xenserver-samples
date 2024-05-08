package xenapi

import (
	"fmt"
)

type PGPURecord struct {
	// Unique identifier/object reference
	UUID string
	// Link to underlying PCI device
	PCI PCIRef
	// GPU group the pGPU is contained in
	GPUGroup GPUGroupRef
	// Host that owns the GPU
	Host HostRef
	// Additional configuration
	OtherConfig map[string]string
	// List of VGPU types supported by the underlying hardware
	SupportedVGPUTypes []VGPUTypeRef
	// List of VGPU types which have been enabled for this PGPU
	EnabledVGPUTypes []VGPUTypeRef
	// List of VGPUs running on this PGPU
	ResidentVGPUs []VGPURef
	// A map relating each VGPU type supported on this GPU to the maximum number of VGPUs of that type which can run simultaneously on this GPU
	SupportedVGPUMaxCapacities map[VGPUTypeRef]int
	// The accessibility of this device from dom0
	Dom0Access PgpuDom0Access
	// Is this device the system display device
	IsSystemDisplayDevice bool
	// PGPU metadata to determine whether a VGPU can migrate between two PGPUs
	CompatibilityMetadata map[string]string
}

type PGPURef string

// A physical GPU (pGPU)
type pGPU struct{}

var PGPU pGPU

// GetRecord: Get a record containing the current state of the given PGPU.
func (pGPU) GetRecord(session *Session, self PGPURef) (retval PGPURecord, err error) {
	method := "PGPU.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePGPURecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the PGPU instance with the specified UUID.
func (pGPU) GetByUUID(session *Session, uUID string) (retval PGPURef, err error) {
	method := "PGPU.get_by_uuid"
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
	retval, err = deserializePGPURef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given PGPU.
func (pGPU) GetUUID(session *Session, self PGPURef) (retval string, err error) {
	method := "PGPU.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPCI: Get the PCI field of the given PGPU.
func (pGPU) GetPCI(session *Session, self PGPURef) (retval PCIRef, err error) {
	method := "PGPU.get_PCI"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetGPUGroup: Get the GPU_group field of the given PGPU.
func (pGPU) GetGPUGroup(session *Session, self PGPURef) (retval GPUGroupRef, err error) {
	method := "PGPU.get_GPU_group"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost: Get the host field of the given PGPU.
func (pGPU) GetHost(session *Session, self PGPURef) (retval HostRef, err error) {
	method := "PGPU.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given PGPU.
func (pGPU) GetOtherConfig(session *Session, self PGPURef) (retval map[string]string, err error) {
	method := "PGPU.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSupportedVGPUTypes: Get the supported_VGPU_types field of the given PGPU.
func (pGPU) GetSupportedVGPUTypes(session *Session, self PGPURef) (retval []VGPUTypeRef, err error) {
	method := "PGPU.get_supported_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRefSet(method+" -> ", result)
	return
}

// GetEnabledVGPUTypes: Get the enabled_VGPU_types field of the given PGPU.
func (pGPU) GetEnabledVGPUTypes(session *Session, self PGPURef) (retval []VGPUTypeRef, err error) {
	method := "PGPU.get_enabled_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRefSet(method+" -> ", result)
	return
}

// GetResidentVGPUs: Get the resident_VGPUs field of the given PGPU.
func (pGPU) GetResidentVGPUs(session *Session, self PGPURef) (retval []VGPURef, err error) {
	method := "PGPU.get_resident_VGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPURefSet(method+" -> ", result)
	return
}

// GetSupportedVGPUMaxCapacities: Get the supported_VGPU_max_capacities field of the given PGPU.
func (pGPU) GetSupportedVGPUMaxCapacities(session *Session, self PGPURef) (retval map[VGPUTypeRef]int, err error) {
	method := "PGPU.get_supported_VGPU_max_capacities"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRefToIntMap(method+" -> ", result)
	return
}

// GetDom0Access: Get the dom0_access field of the given PGPU.
func (pGPU) GetDom0Access(session *Session, self PGPURef) (retval PgpuDom0Access, err error) {
	method := "PGPU.get_dom0_access"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumPgpuDom0Access(method+" -> ", result)
	return
}

// GetIsSystemDisplayDevice: Get the is_system_display_device field of the given PGPU.
func (pGPU) GetIsSystemDisplayDevice(session *Session, self PGPURef) (retval bool, err error) {
	method := "PGPU.get_is_system_display_device"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCompatibilityMetadata: Get the compatibility_metadata field of the given PGPU.
func (pGPU) GetCompatibilityMetadata(session *Session, self PGPURef) (retval map[string]string, err error) {
	method := "PGPU.get_compatibility_metadata"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given PGPU.
func (pGPU) SetOtherConfig(session *Session, self PGPURef, value map[string]string) (err error) {
	method := "PGPU.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given PGPU.
func (pGPU) AddToOtherConfig(session *Session, self PGPURef, key string, value string) (err error) {
	method := "PGPU.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PGPU.  If the key is not in that Map, then do nothing.
func (pGPU) RemoveFromOtherConfig(session *Session, self PGPURef, key string) (err error) {
	method := "PGPU.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddEnabledVGPUTypes:
func (pGPU) AddEnabledVGPUTypes(session *Session, self PGPURef, value VGPUTypeRef) (err error) {
	method := "PGPU.add_enabled_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncAddEnabledVGPUTypes:
func (pGPU) AsyncAddEnabledVGPUTypes(session *Session, self PGPURef, value VGPUTypeRef) (retval TaskRef, err error) {
	method := "Async.PGPU.add_enabled_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// RemoveEnabledVGPUTypes:
func (pGPU) RemoveEnabledVGPUTypes(session *Session, self PGPURef, value VGPUTypeRef) (err error) {
	method := "PGPU.remove_enabled_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncRemoveEnabledVGPUTypes:
func (pGPU) AsyncRemoveEnabledVGPUTypes(session *Session, self PGPURef, value VGPUTypeRef) (retval TaskRef, err error) {
	method := "Async.PGPU.remove_enabled_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetEnabledVGPUTypes:
func (pGPU) SetEnabledVGPUTypes(session *Session, self PGPURef, value []VGPUTypeRef) (err error) {
	method := "PGPU.set_enabled_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVGPUTypeRefSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetEnabledVGPUTypes:
func (pGPU) AsyncSetEnabledVGPUTypes(session *Session, self PGPURef, value []VGPUTypeRef) (retval TaskRef, err error) {
	method := "Async.PGPU.set_enabled_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVGPUTypeRefSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetGPUGroup:
func (pGPU) SetGPUGroup(session *Session, self PGPURef, value GPUGroupRef) (err error) {
	method := "PGPU.set_GPU_group"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetGPUGroup:
func (pGPU) AsyncSetGPUGroup(session *Session, self PGPURef, value GPUGroupRef) (retval TaskRef, err error) {
	method := "Async.PGPU.set_GPU_group"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetRemainingCapacity:
func (pGPU) GetRemainingCapacity(session *Session, self PGPURef, vgpuType VGPUTypeRef) (retval int, err error) {
	method := "PGPU.get_remaining_capacity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	vgpuTypeArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "vgpu_type"), vgpuType)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, vgpuTypeArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// AsyncGetRemainingCapacity:
func (pGPU) AsyncGetRemainingCapacity(session *Session, self PGPURef, vgpuType VGPUTypeRef) (retval TaskRef, err error) {
	method := "Async.PGPU.get_remaining_capacity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	vgpuTypeArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "vgpu_type"), vgpuType)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, vgpuTypeArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EnableDom0Access:
func (pGPU) EnableDom0Access(session *Session, self PGPURef) (retval PgpuDom0Access, err error) {
	method := "PGPU.enable_dom0_access"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumPgpuDom0Access(method+" -> ", result)
	return
}

// AsyncEnableDom0Access:
func (pGPU) AsyncEnableDom0Access(session *Session, self PGPURef) (retval TaskRef, err error) {
	method := "Async.PGPU.enable_dom0_access"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// DisableDom0Access:
func (pGPU) DisableDom0Access(session *Session, self PGPURef) (retval PgpuDom0Access, err error) {
	method := "PGPU.disable_dom0_access"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumPgpuDom0Access(method+" -> ", result)
	return
}

// AsyncDisableDom0Access:
func (pGPU) AsyncDisableDom0Access(session *Session, self PGPURef) (retval TaskRef, err error) {
	method := "Async.PGPU.disable_dom0_access"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePGPURef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the PGPUs known to the system.
func (pGPU) GetAll(session *Session) (retval []PGPURef, err error) {
	method := "PGPU.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePGPURefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of PGPU references to PGPU records for all PGPUs known to the system.
func (pGPU) GetAllRecords(session *Session) (retval map[PGPURef]PGPURecord, err error) {
	method := "PGPU.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePGPURefToPGPURecordMap(method+" -> ", result)
	return
}

