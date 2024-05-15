package xenapi

import (
	"fmt"
)

type GPUGroupRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// List of pGPUs in the group
	PGPUs []PGPURef
	// List of vGPUs using the group
	VGPUs []VGPURef
	// List of GPU types (vendor+device ID) that can be in this group
	GPUTypes []string
	// Additional configuration
	OtherConfig map[string]string
	// Current allocation of vGPUs to pGPUs for this group
	AllocationAlgorithm AllocationAlgorithm
	// vGPU types supported on at least one of the pGPUs in this group
	SupportedVGPUTypes []VGPUTypeRef
	// vGPU types supported on at least one of the pGPUs in this group
	EnabledVGPUTypes []VGPUTypeRef
}

type GPUGroupRef string

// A group of compatible GPUs across the resource pool
type gPUGroup struct{}

var GPUGroup gPUGroup

// GetAllRecords: Return a map of GPU_group references to GPU_group records for all GPU_groups known to the system.
// Version: boston
func (gPUGroup) GetAllRecords(session *Session) (retval map[GPUGroupRef]GPUGroupRecord, err error) {
	method := "GPU_group.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefToGPUGroupRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of GPU_group references to GPU_group records for all GPU_groups known to the system.
// Version: boston
func (gPUGroup) GetAllRecords1(session *Session) (retval map[GPUGroupRef]GPUGroupRecord, err error) {
	method := "GPU_group.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefToGPUGroupRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the GPU_groups known to the system.
// Version: boston
func (gPUGroup) GetAll(session *Session) (retval []GPUGroupRef, err error) {
	method := "GPU_group.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the GPU_groups known to the system.
// Version: boston
func (gPUGroup) GetAll1(session *Session) (retval []GPUGroupRef, err error) {
	method := "GPU_group.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefSet(method+" -> ", result)
	return
}

// GetRemainingCapacity:
// Version: vgpu-tech-preview
func (gPUGroup) GetRemainingCapacity(session *Session, self GPUGroupRef, vgpuType VGPUTypeRef) (retval int, err error) {
	method := "GPU_group.get_remaining_capacity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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
// Version: vgpu-tech-preview
func (gPUGroup) AsyncGetRemainingCapacity(session *Session, self GPUGroupRef, vgpuType VGPUTypeRef) (retval TaskRef, err error) {
	method := "Async.GPU_group.get_remaining_capacity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRemainingCapacity3:
// Version: vgpu-tech-preview
func (gPUGroup) GetRemainingCapacity3(session *Session, self GPUGroupRef, vgpuType VGPUTypeRef) (retval int, err error) {
	method := "GPU_group.get_remaining_capacity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncGetRemainingCapacity3:
// Version: vgpu-tech-preview
func (gPUGroup) AsyncGetRemainingCapacity3(session *Session, self GPUGroupRef, vgpuType VGPUTypeRef) (retval TaskRef, err error) {
	method := "Async.GPU_group.get_remaining_capacity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy:
// Version: boston
func (gPUGroup) Destroy(session *Session, self GPUGroupRef) (err error) {
	method := "GPU_group.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy:
// Version: boston
func (gPUGroup) AsyncDestroy(session *Session, self GPUGroupRef) (retval TaskRef, err error) {
	method := "Async.GPU_group.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy2:
// Version: boston
func (gPUGroup) Destroy2(session *Session, self GPUGroupRef) (err error) {
	method := "GPU_group.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy2:
// Version: boston
func (gPUGroup) AsyncDestroy2(session *Session, self GPUGroupRef) (retval TaskRef, err error) {
	method := "Async.GPU_group.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create:
// Version: boston
func (gPUGroup) Create(session *Session, nameLabel string, nameDescription string, otherConfig map[string]string) (retval GPUGroupRef, err error) {
	method := "GPU_group.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, otherConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRef(method+" -> ", result)
	return
}

// AsyncCreate:
// Version: boston
func (gPUGroup) AsyncCreate(session *Session, nameLabel string, nameDescription string, otherConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.GPU_group.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, otherConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Create4:
// Version: boston
func (gPUGroup) Create4(session *Session, nameLabel string, nameDescription string, otherConfig map[string]string) (retval GPUGroupRef, err error) {
	method := "GPU_group.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, otherConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRef(method+" -> ", result)
	return
}

// AsyncCreate4:
// Version: boston
func (gPUGroup) AsyncCreate4(session *Session, nameLabel string, nameDescription string, otherConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.GPU_group.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, otherConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetAllocationAlgorithm: Set the allocation_algorithm field of the given GPU_group.
// Version: vgpu-tech-preview
func (gPUGroup) SetAllocationAlgorithm(session *Session, self GPUGroupRef, value AllocationAlgorithm) (err error) {
	method := "GPU_group.set_allocation_algorithm"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumAllocationAlgorithm(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetAllocationAlgorithm3: Set the allocation_algorithm field of the given GPU_group.
// Version: vgpu-tech-preview
func (gPUGroup) SetAllocationAlgorithm3(session *Session, self GPUGroupRef, value AllocationAlgorithm) (err error) {
	method := "GPU_group.set_allocation_algorithm"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumAllocationAlgorithm(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetAllocationAlgorithm2: Set the allocation_algorithm field of the given GPU_group.
// Version: boston
func (gPUGroup) SetAllocationAlgorithm2(session *Session, self GPUGroupRef) (err error) {
	method := "GPU_group.set_allocation_algorithm"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given GPU_group.  If the key is not in that Map, then do nothing.
// Version: boston
func (gPUGroup) RemoveFromOtherConfig(session *Session, self GPUGroupRef, key string) (err error) {
	method := "GPU_group.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig3: Remove the given key and its corresponding value from the other_config field of the given GPU_group.  If the key is not in that Map, then do nothing.
// Version: boston
func (gPUGroup) RemoveFromOtherConfig3(session *Session, self GPUGroupRef, key string) (err error) {
	method := "GPU_group.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given GPU_group.
// Version: boston
func (gPUGroup) AddToOtherConfig(session *Session, self GPUGroupRef, key string, value string) (err error) {
	method := "GPU_group.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig4: Add the given key-value pair to the other_config field of the given GPU_group.
// Version: boston
func (gPUGroup) AddToOtherConfig4(session *Session, self GPUGroupRef, key string, value string) (err error) {
	method := "GPU_group.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given GPU_group.
// Version: boston
func (gPUGroup) SetOtherConfig(session *Session, self GPUGroupRef, value map[string]string) (err error) {
	method := "GPU_group.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig3: Set the other_config field of the given GPU_group.
// Version: boston
func (gPUGroup) SetOtherConfig3(session *Session, self GPUGroupRef, value map[string]string) (err error) {
	method := "GPU_group.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription: Set the name/description field of the given GPU_group.
// Version: boston
func (gPUGroup) SetNameDescription(session *Session, self GPUGroupRef, value string) (err error) {
	method := "GPU_group.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription3: Set the name/description field of the given GPU_group.
// Version: boston
func (gPUGroup) SetNameDescription3(session *Session, self GPUGroupRef, value string) (err error) {
	method := "GPU_group.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel: Set the name/label field of the given GPU_group.
// Version: boston
func (gPUGroup) SetNameLabel(session *Session, self GPUGroupRef, value string) (err error) {
	method := "GPU_group.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel3: Set the name/label field of the given GPU_group.
// Version: boston
func (gPUGroup) SetNameLabel3(session *Session, self GPUGroupRef, value string) (err error) {
	method := "GPU_group.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnabledVGPUTypes: Get the enabled_VGPU_types field of the given GPU_group.
// Version: boston
func (gPUGroup) GetEnabledVGPUTypes(session *Session, self GPUGroupRef) (retval []VGPUTypeRef, err error) {
	method := "GPU_group.get_enabled_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnabledVGPUTypes2: Get the enabled_VGPU_types field of the given GPU_group.
// Version: boston
func (gPUGroup) GetEnabledVGPUTypes2(session *Session, self GPUGroupRef) (retval []VGPUTypeRef, err error) {
	method := "GPU_group.get_enabled_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSupportedVGPUTypes: Get the supported_VGPU_types field of the given GPU_group.
// Version: boston
func (gPUGroup) GetSupportedVGPUTypes(session *Session, self GPUGroupRef) (retval []VGPUTypeRef, err error) {
	method := "GPU_group.get_supported_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSupportedVGPUTypes2: Get the supported_VGPU_types field of the given GPU_group.
// Version: boston
func (gPUGroup) GetSupportedVGPUTypes2(session *Session, self GPUGroupRef) (retval []VGPUTypeRef, err error) {
	method := "GPU_group.get_supported_VGPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllocationAlgorithm: Get the allocation_algorithm field of the given GPU_group.
// Version: boston
func (gPUGroup) GetAllocationAlgorithm(session *Session, self GPUGroupRef) (retval AllocationAlgorithm, err error) {
	method := "GPU_group.get_allocation_algorithm"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumAllocationAlgorithm(method+" -> ", result)
	return
}

// GetAllocationAlgorithm2: Get the allocation_algorithm field of the given GPU_group.
// Version: boston
func (gPUGroup) GetAllocationAlgorithm2(session *Session, self GPUGroupRef) (retval AllocationAlgorithm, err error) {
	method := "GPU_group.get_allocation_algorithm"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumAllocationAlgorithm(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given GPU_group.
// Version: boston
func (gPUGroup) GetOtherConfig(session *Session, self GPUGroupRef) (retval map[string]string, err error) {
	method := "GPU_group.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig2: Get the other_config field of the given GPU_group.
// Version: boston
func (gPUGroup) GetOtherConfig2(session *Session, self GPUGroupRef) (retval map[string]string, err error) {
	method := "GPU_group.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetGPUTypes: Get the GPU_types field of the given GPU_group.
// Version: boston
func (gPUGroup) GetGPUTypes(session *Session, self GPUGroupRef) (retval []string, err error) {
	method := "GPU_group.get_GPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetGPUTypes2: Get the GPU_types field of the given GPU_group.
// Version: boston
func (gPUGroup) GetGPUTypes2(session *Session, self GPUGroupRef) (retval []string, err error) {
	method := "GPU_group.get_GPU_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVGPUs: Get the VGPUs field of the given GPU_group.
// Version: boston
func (gPUGroup) GetVGPUs(session *Session, self GPUGroupRef) (retval []VGPURef, err error) {
	method := "GPU_group.get_VGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVGPUs2: Get the VGPUs field of the given GPU_group.
// Version: boston
func (gPUGroup) GetVGPUs2(session *Session, self GPUGroupRef) (retval []VGPURef, err error) {
	method := "GPU_group.get_VGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPGPUs: Get the PGPUs field of the given GPU_group.
// Version: boston
func (gPUGroup) GetPGPUs(session *Session, self GPUGroupRef) (retval []PGPURef, err error) {
	method := "GPU_group.get_PGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePGPURefSet(method+" -> ", result)
	return
}

// GetPGPUs2: Get the PGPUs field of the given GPU_group.
// Version: boston
func (gPUGroup) GetPGPUs2(session *Session, self GPUGroupRef) (retval []PGPURef, err error) {
	method := "GPU_group.get_PGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePGPURefSet(method+" -> ", result)
	return
}

// GetNameDescription: Get the name/description field of the given GPU_group.
// Version: boston
func (gPUGroup) GetNameDescription(session *Session, self GPUGroupRef) (retval string, err error) {
	method := "GPU_group.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription2: Get the name/description field of the given GPU_group.
// Version: boston
func (gPUGroup) GetNameDescription2(session *Session, self GPUGroupRef) (retval string, err error) {
	method := "GPU_group.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given GPU_group.
// Version: boston
func (gPUGroup) GetNameLabel(session *Session, self GPUGroupRef) (retval string, err error) {
	method := "GPU_group.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel2: Get the name/label field of the given GPU_group.
// Version: boston
func (gPUGroup) GetNameLabel2(session *Session, self GPUGroupRef) (retval string, err error) {
	method := "GPU_group.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given GPU_group.
// Version: boston
func (gPUGroup) GetUUID(session *Session, self GPUGroupRef) (retval string, err error) {
	method := "GPU_group.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given GPU_group.
// Version: boston
func (gPUGroup) GetUUID2(session *Session, self GPUGroupRef) (retval string, err error) {
	method := "GPU_group.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByNameLabel: Get all the GPU_group instances with the given label.
// Version: boston
func (gPUGroup) GetByNameLabel(session *Session, label string) (retval []GPUGroupRef, err error) {
	method := "GPU_group.get_by_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	labelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "label"), label)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, labelArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefSet(method+" -> ", result)
	return
}

// GetByNameLabel2: Get all the GPU_group instances with the given label.
// Version: boston
func (gPUGroup) GetByNameLabel2(session *Session, label string) (retval []GPUGroupRef, err error) {
	method := "GPU_group.get_by_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	labelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "label"), label)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, labelArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefSet(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the GPU_group instance with the specified UUID.
// Version: boston
func (gPUGroup) GetByUUID(session *Session, uUID string) (retval GPUGroupRef, err error) {
	method := "GPU_group.get_by_uuid"
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
	retval, err = deserializeGPUGroupRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the GPU_group instance with the specified UUID.
// Version: boston
func (gPUGroup) GetByUUID2(session *Session, uUID string) (retval GPUGroupRef, err error) {
	method := "GPU_group.get_by_uuid"
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
	retval, err = deserializeGPUGroupRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given GPU_group.
// Version: boston
func (gPUGroup) GetRecord(session *Session, self GPUGroupRef) (retval GPUGroupRecord, err error) {
	method := "GPU_group.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given GPU_group.
// Version: boston
func (gPUGroup) GetRecord2(session *Session, self GPUGroupRef) (retval GPUGroupRecord, err error) {
	method := "GPU_group.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeGPUGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRecord(method+" -> ", result)
	return
}
