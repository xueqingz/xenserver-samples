package xenapi

import (
	"fmt"
)

type VGPUTypeRecord struct {
	// Unique identifier/object reference
	UUID string
	// Name of VGPU vendor
	VendorName string
	// Model name associated with the VGPU type
	ModelName string
	// Framebuffer size of the VGPU type, in bytes
	FramebufferSize int
	// Maximum number of displays supported by the VGPU type
	MaxHeads int
	// Maximum resolution (width) supported by the VGPU type
	MaxResolutionX int
	// Maximum resolution (height) supported by the VGPU type
	MaxResolutionY int
	// List of PGPUs that support this VGPU type
	SupportedOnPGPUs []PGPURef
	// List of PGPUs that have this VGPU type enabled
	EnabledOnPGPUs []PGPURef
	// List of VGPUs of this type
	VGPUs []VGPURef
	// List of GPU groups in which at least one PGPU supports this VGPU type
	SupportedOnGPUGroups []GPUGroupRef
	// List of GPU groups in which at least one have this VGPU type enabled
	EnabledOnGPUGroups []GPUGroupRef
	// The internal implementation of this VGPU type
	Implementation VgpuTypeImplementation
	// Key used to identify VGPU types and avoid creating duplicates - this field is used internally and not intended for interpretation by API clients
	Identifier string
	// Indicates whether VGPUs of this type should be considered experimental
	Experimental bool
	// List of VGPU types which are compatible in one VM
	CompatibleTypesInVM []VGPUTypeRef
}

type VGPUTypeRef string

// A type of virtual GPU
type vgpuType struct{}

var VGPUType vgpuType

// GetAllRecords: Return a map of VGPU_type references to VGPU_type records for all VGPU_types known to the system.
// Version: vgpu-tech-preview
func (vgpuType) GetAllRecords(session *Session) (retval map[VGPUTypeRef]VGPUTypeRecord, err error) {
	method := "VGPU_type.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRefToVGPUTypeRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of VGPU_type references to VGPU_type records for all VGPU_types known to the system.
// Version: vgpu-tech-preview
func (vgpuType) GetAllRecords1(session *Session) (retval map[VGPUTypeRef]VGPUTypeRecord, err error) {
	method := "VGPU_type.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRefToVGPUTypeRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the VGPU_types known to the system.
// Version: vgpu-tech-preview
func (vgpuType) GetAll(session *Session) (retval []VGPUTypeRef, err error) {
	method := "VGPU_type.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the VGPU_types known to the system.
// Version: vgpu-tech-preview
func (vgpuType) GetAll1(session *Session) (retval []VGPUTypeRef, err error) {
	method := "VGPU_type.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRefSet(method+" -> ", result)
	return
}

// GetCompatibleTypesInVM: Get the compatible_types_in_vm field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetCompatibleTypesInVM(session *Session, self VGPUTypeRef) (retval []VGPUTypeRef, err error) {
	method := "VGPU_type.get_compatible_types_in_vm"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCompatibleTypesInVM2: Get the compatible_types_in_vm field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetCompatibleTypesInVM2(session *Session, self VGPUTypeRef) (retval []VGPUTypeRef, err error) {
	method := "VGPU_type.get_compatible_types_in_vm"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetExperimental: Get the experimental field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetExperimental(session *Session, self VGPUTypeRef) (retval bool, err error) {
	method := "VGPU_type.get_experimental"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetExperimental2: Get the experimental field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetExperimental2(session *Session, self VGPUTypeRef) (retval bool, err error) {
	method := "VGPU_type.get_experimental"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIdentifier: Get the identifier field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetIdentifier(session *Session, self VGPUTypeRef) (retval string, err error) {
	method := "VGPU_type.get_identifier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIdentifier2: Get the identifier field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetIdentifier2(session *Session, self VGPUTypeRef) (retval string, err error) {
	method := "VGPU_type.get_identifier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetImplementation: Get the implementation field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetImplementation(session *Session, self VGPUTypeRef) (retval VgpuTypeImplementation, err error) {
	method := "VGPU_type.get_implementation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVgpuTypeImplementation(method+" -> ", result)
	return
}

// GetImplementation2: Get the implementation field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetImplementation2(session *Session, self VGPUTypeRef) (retval VgpuTypeImplementation, err error) {
	method := "VGPU_type.get_implementation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVgpuTypeImplementation(method+" -> ", result)
	return
}

// GetEnabledOnGPUGroups: Get the enabled_on_GPU_groups field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetEnabledOnGPUGroups(session *Session, self VGPUTypeRef) (retval []GPUGroupRef, err error) {
	method := "VGPU_type.get_enabled_on_GPU_groups"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefSet(method+" -> ", result)
	return
}

// GetEnabledOnGPUGroups2: Get the enabled_on_GPU_groups field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetEnabledOnGPUGroups2(session *Session, self VGPUTypeRef) (retval []GPUGroupRef, err error) {
	method := "VGPU_type.get_enabled_on_GPU_groups"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefSet(method+" -> ", result)
	return
}

// GetSupportedOnGPUGroups: Get the supported_on_GPU_groups field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetSupportedOnGPUGroups(session *Session, self VGPUTypeRef) (retval []GPUGroupRef, err error) {
	method := "VGPU_type.get_supported_on_GPU_groups"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefSet(method+" -> ", result)
	return
}

// GetSupportedOnGPUGroups2: Get the supported_on_GPU_groups field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetSupportedOnGPUGroups2(session *Session, self VGPUTypeRef) (retval []GPUGroupRef, err error) {
	method := "VGPU_type.get_supported_on_GPU_groups"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeGPUGroupRefSet(method+" -> ", result)
	return
}

// GetVGPUs: Get the VGPUs field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetVGPUs(session *Session, self VGPUTypeRef) (retval []VGPURef, err error) {
	method := "VGPU_type.get_VGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVGPUs2: Get the VGPUs field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetVGPUs2(session *Session, self VGPUTypeRef) (retval []VGPURef, err error) {
	method := "VGPU_type.get_VGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnabledOnPGPUs: Get the enabled_on_PGPUs field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetEnabledOnPGPUs(session *Session, self VGPUTypeRef) (retval []PGPURef, err error) {
	method := "VGPU_type.get_enabled_on_PGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnabledOnPGPUs2: Get the enabled_on_PGPUs field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetEnabledOnPGPUs2(session *Session, self VGPUTypeRef) (retval []PGPURef, err error) {
	method := "VGPU_type.get_enabled_on_PGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSupportedOnPGPUs: Get the supported_on_PGPUs field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetSupportedOnPGPUs(session *Session, self VGPUTypeRef) (retval []PGPURef, err error) {
	method := "VGPU_type.get_supported_on_PGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSupportedOnPGPUs2: Get the supported_on_PGPUs field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetSupportedOnPGPUs2(session *Session, self VGPUTypeRef) (retval []PGPURef, err error) {
	method := "VGPU_type.get_supported_on_PGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMaxResolutionY: Get the max_resolution_y field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetMaxResolutionY(session *Session, self VGPUTypeRef) (retval int, err error) {
	method := "VGPU_type.get_max_resolution_y"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMaxResolutionY2: Get the max_resolution_y field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetMaxResolutionY2(session *Session, self VGPUTypeRef) (retval int, err error) {
	method := "VGPU_type.get_max_resolution_y"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMaxResolutionX: Get the max_resolution_x field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetMaxResolutionX(session *Session, self VGPUTypeRef) (retval int, err error) {
	method := "VGPU_type.get_max_resolution_x"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMaxResolutionX2: Get the max_resolution_x field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetMaxResolutionX2(session *Session, self VGPUTypeRef) (retval int, err error) {
	method := "VGPU_type.get_max_resolution_x"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMaxHeads: Get the max_heads field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetMaxHeads(session *Session, self VGPUTypeRef) (retval int, err error) {
	method := "VGPU_type.get_max_heads"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMaxHeads2: Get the max_heads field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetMaxHeads2(session *Session, self VGPUTypeRef) (retval int, err error) {
	method := "VGPU_type.get_max_heads"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetFramebufferSize: Get the framebuffer_size field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetFramebufferSize(session *Session, self VGPUTypeRef) (retval int, err error) {
	method := "VGPU_type.get_framebuffer_size"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetFramebufferSize2: Get the framebuffer_size field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetFramebufferSize2(session *Session, self VGPUTypeRef) (retval int, err error) {
	method := "VGPU_type.get_framebuffer_size"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetModelName: Get the model_name field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetModelName(session *Session, self VGPUTypeRef) (retval string, err error) {
	method := "VGPU_type.get_model_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetModelName2: Get the model_name field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetModelName2(session *Session, self VGPUTypeRef) (retval string, err error) {
	method := "VGPU_type.get_model_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVendorName: Get the vendor_name field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetVendorName(session *Session, self VGPUTypeRef) (retval string, err error) {
	method := "VGPU_type.get_vendor_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVendorName2: Get the vendor_name field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetVendorName2(session *Session, self VGPUTypeRef) (retval string, err error) {
	method := "VGPU_type.get_vendor_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetUUID(session *Session, self VGPUTypeRef) (retval string, err error) {
	method := "VGPU_type.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetUUID2(session *Session, self VGPUTypeRef) (retval string, err error) {
	method := "VGPU_type.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByUUID: Get a reference to the VGPU_type instance with the specified UUID.
// Version: vgpu-tech-preview
func (vgpuType) GetByUUID(session *Session, uuid string) (retval VGPUTypeRef, err error) {
	method := "VGPU_type.get_by_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	uuidArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uuid"), uuid)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uuidArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the VGPU_type instance with the specified UUID.
// Version: vgpu-tech-preview
func (vgpuType) GetByUUID2(session *Session, uuid string) (retval VGPUTypeRef, err error) {
	method := "VGPU_type.get_by_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	uuidArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uuid"), uuid)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uuidArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetRecord(session *Session, self VGPUTypeRef) (retval VGPUTypeRecord, err error) {
	method := "VGPU_type.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given VGPU_type.
// Version: vgpu-tech-preview
func (vgpuType) GetRecord2(session *Session, self VGPUTypeRef) (retval VGPUTypeRecord, err error) {
	method := "VGPU_type.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVGPUTypeRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPUTypeRecord(method+" -> ", result)
	return
}
