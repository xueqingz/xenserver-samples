package xenapi

import (
	"fmt"
)

type SRRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []StorageOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]StorageOperations
	// all virtual disks known to this storage repository
	VDIs []VDIRef
	// describes how particular hosts can see this storage repository
	PBDs []PBDRef
	// sum of virtual_sizes of all VDIs in this storage repository (in bytes)
	VirtualAllocation int
	// physical space currently utilised on this storage repository (in bytes). Note that for sparse disk formats, physical_utilisation may be less than virtual_allocation
	PhysicalUtilisation int
	// total physical size of the repository (in bytes)
	PhysicalSize int
	// type of the storage repository
	Type string
	// the type of the SR&apos;s content, if required (e.g. ISOs)
	ContentType string
	// true if this SR is (capable of being) shared between multiple hosts
	Shared bool
	// additional configuration
	OtherConfig map[string]string
	// user-specified tags for categorization purposes
	Tags []string
	// SM dependent data
	SmConfig map[string]string
	// Binary blobs associated with this SR
	Blobs map[string]BlobRef
	// True if this SR is assigned to be the local cache for its host
	LocalCacheEnabled bool
	// The disaster recovery task which introduced this SR
	IntroducedBy DRTaskRef
	// True if the SR is using aggregated local storage
	Clustered bool
	// True if this is the SR that contains the Tools ISO VDIs
	IsToolsSr bool
}

type SRRef string

// A storage repository
type sR struct{}

var SR sR

// GetRecord: Get a record containing the current state of the given SR.
func (sR) GetRecord(session *Session, self SRRef) (retval SRRecord, err error) {
	method := "SR.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the SR instance with the specified UUID.
func (sR) GetByUUID(session *Session, uUID string) (retval SRRef, err error) {
	method := "SR.get_by_uuid"
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
	retval, err = deserializeSRRef(method+" -> ", result)
	return
}

// GetByNameLabel: Get all the SR instances with the given label.
func (sR) GetByNameLabel(session *Session, label string) (retval []SRRef, err error) {
	method := "SR.get_by_name_label"
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
	retval, err = deserializeSRRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given SR.
func (sR) GetUUID(session *Session, self SRRef) (retval string, err error) {
	method := "SR.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given SR.
func (sR) GetNameLabel(session *Session, self SRRef) (retval string, err error) {
	method := "SR.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given SR.
func (sR) GetNameDescription(session *Session, self SRRef) (retval string, err error) {
	method := "SR.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given SR.
func (sR) GetAllowedOperations(session *Session, self SRRef) (retval []StorageOperations, err error) {
	method := "SR.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumStorageOperationsSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given SR.
func (sR) GetCurrentOperations(session *Session, self SRRef) (retval map[string]StorageOperations, err error) {
	method := "SR.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumStorageOperationsMap(method+" -> ", result)
	return
}

// GetVDIs: Get the VDIs field of the given SR.
func (sR) GetVDIs(session *Session, self SRRef) (retval []VDIRef, err error) {
	method := "SR.get_VDIs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRefSet(method+" -> ", result)
	return
}

// GetPBDs: Get the PBDs field of the given SR.
func (sR) GetPBDs(session *Session, self SRRef) (retval []PBDRef, err error) {
	method := "SR.get_PBDs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePBDRefSet(method+" -> ", result)
	return
}

// GetVirtualAllocation: Get the virtual_allocation field of the given SR.
func (sR) GetVirtualAllocation(session *Session, self SRRef) (retval int, err error) {
	method := "SR.get_virtual_allocation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPhysicalUtilisation: Get the physical_utilisation field of the given SR.
func (sR) GetPhysicalUtilisation(session *Session, self SRRef) (retval int, err error) {
	method := "SR.get_physical_utilisation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPhysicalSize: Get the physical_size field of the given SR.
func (sR) GetPhysicalSize(session *Session, self SRRef) (retval int, err error) {
	method := "SR.get_physical_size"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetType: Get the type field of the given SR.
func (sR) GetType(session *Session, self SRRef) (retval string, err error) {
	method := "SR.get_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetContentType: Get the content_type field of the given SR.
func (sR) GetContentType(session *Session, self SRRef) (retval string, err error) {
	method := "SR.get_content_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetShared: Get the shared field of the given SR.
func (sR) GetShared(session *Session, self SRRef) (retval bool, err error) {
	method := "SR.get_shared"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given SR.
func (sR) GetOtherConfig(session *Session, self SRRef) (retval map[string]string, err error) {
	method := "SR.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTags: Get the tags field of the given SR.
func (sR) GetTags(session *Session, self SRRef) (retval []string, err error) {
	method := "SR.get_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSmConfig: Get the sm_config field of the given SR.
func (sR) GetSmConfig(session *Session, self SRRef) (retval map[string]string, err error) {
	method := "SR.get_sm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBlobs: Get the blobs field of the given SR.
func (sR) GetBlobs(session *Session, self SRRef) (retval map[string]BlobRef, err error) {
	method := "SR.get_blobs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToBlobRefMap(method+" -> ", result)
	return
}

// GetLocalCacheEnabled: Get the local_cache_enabled field of the given SR.
func (sR) GetLocalCacheEnabled(session *Session, self SRRef) (retval bool, err error) {
	method := "SR.get_local_cache_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIntroducedBy: Get the introduced_by field of the given SR.
func (sR) GetIntroducedBy(session *Session, self SRRef) (retval DRTaskRef, err error) {
	method := "SR.get_introduced_by"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeDRTaskRef(method+" -> ", result)
	return
}

// GetClustered: Get the clustered field of the given SR.
func (sR) GetClustered(session *Session, self SRRef) (retval bool, err error) {
	method := "SR.get_clustered"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsToolsSr: Get the is_tools_sr field of the given SR.
func (sR) GetIsToolsSr(session *Session, self SRRef) (retval bool, err error) {
	method := "SR.get_is_tools_sr"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given SR.
func (sR) SetOtherConfig(session *Session, self SRRef, value map[string]string) (err error) {
	method := "SR.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given SR.
func (sR) AddToOtherConfig(session *Session, self SRRef, key string, value string) (err error) {
	method := "SR.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given SR.  If the key is not in that Map, then do nothing.
func (sR) RemoveFromOtherConfig(session *Session, self SRRef, key string) (err error) {
	method := "SR.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetTags: Set the tags field of the given SR.
func (sR) SetTags(session *Session, self SRRef, value []string) (err error) {
	method := "SR.set_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AddTags: Add the given value to the tags field of the given SR.  If the value is already in that Set, then do nothing.
func (sR) AddTags(session *Session, self SRRef, value string) (err error) {
	method := "SR.add_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveTags: Remove the given value from the tags field of the given SR.  If the value is not in that Set, then do nothing.
func (sR) RemoveTags(session *Session, self SRRef, value string) (err error) {
	method := "SR.remove_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetSmConfig: Set the sm_config field of the given SR.
func (sR) SetSmConfig(session *Session, self SRRef, value map[string]string) (err error) {
	method := "SR.set_sm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToSmConfig: Add the given key-value pair to the sm_config field of the given SR.
func (sR) AddToSmConfig(session *Session, self SRRef, key string, value string) (err error) {
	method := "SR.add_to_sm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromSmConfig: Remove the given key and its corresponding value from the sm_config field of the given SR.  If the key is not in that Map, then do nothing.
func (sR) RemoveFromSmConfig(session *Session, self SRRef, key string) (err error) {
	method := "SR.remove_from_sm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Create a new Storage Repository and introduce it into the managed system, creating both SR record and PBD record to attach it to current host (with specified device_config parameters)
//
// Errors:
// SR_UNKNOWN_DRIVER - The SR could not be connected because the driver was not recognised.
func (sR) Create(session *Session, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, typeKey string, contentType string, shared bool, smConfig map[string]string) (retval SRRef, err error) {
	method := "SR.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	physicalSizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "physical_size"), physicalSize)
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
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	contentTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "content_type"), contentType)
	if err != nil {
		return
	}
	sharedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "shared"), shared)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, deviceConfigArg, physicalSizeArg, nameLabelArg, nameDescriptionArg, typeKeyArg, contentTypeArg, sharedArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new Storage Repository and introduce it into the managed system, creating both SR record and PBD record to attach it to current host (with specified device_config parameters)
//
// Errors:
// SR_UNKNOWN_DRIVER - The SR could not be connected because the driver was not recognised.
func (sR) AsyncCreate(session *Session, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, typeKey string, contentType string, shared bool, smConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.SR.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	physicalSizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "physical_size"), physicalSize)
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
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	contentTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "content_type"), contentType)
	if err != nil {
		return
	}
	sharedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "shared"), shared)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, deviceConfigArg, physicalSizeArg, nameLabelArg, nameDescriptionArg, typeKeyArg, contentTypeArg, sharedArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Introduce: Introduce a new Storage Repository into the managed system
func (sR) Introduce(session *Session, uUID string, nameLabel string, nameDescription string, typeKey string, contentType string, shared bool, smConfig map[string]string) (retval SRRef, err error) {
	method := "SR.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	uUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uuid"), uUID)
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
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	contentTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "content_type"), contentType)
	if err != nil {
		return
	}
	sharedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "shared"), shared)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uUIDArg, nameLabelArg, nameDescriptionArg, typeKeyArg, contentTypeArg, sharedArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRef(method+" -> ", result)
	return
}

// AsyncIntroduce: Introduce a new Storage Repository into the managed system
func (sR) AsyncIntroduce(session *Session, uUID string, nameLabel string, nameDescription string, typeKey string, contentType string, shared bool, smConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.SR.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	uUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uuid"), uUID)
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
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	contentTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "content_type"), contentType)
	if err != nil {
		return
	}
	sharedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "shared"), shared)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uUIDArg, nameLabelArg, nameDescriptionArg, typeKeyArg, contentTypeArg, sharedArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Make: Create a new Storage Repository on disk. This call is deprecated: use SR.create instead.
func (sR) Make(session *Session, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, typeKey string, contentType string, smConfig map[string]string) (retval string, err error) {
	method := "SR.make"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	physicalSizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "physical_size"), physicalSize)
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
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	contentTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "content_type"), contentType)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, deviceConfigArg, physicalSizeArg, nameLabelArg, nameDescriptionArg, typeKeyArg, contentTypeArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncMake: Create a new Storage Repository on disk. This call is deprecated: use SR.create instead.
func (sR) AsyncMake(session *Session, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, typeKey string, contentType string, smConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.SR.make"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	physicalSizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "physical_size"), physicalSize)
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
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	contentTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "content_type"), contentType)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, deviceConfigArg, physicalSizeArg, nameLabelArg, nameDescriptionArg, typeKeyArg, contentTypeArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Destroy specified SR, removing SR-record from database and remove SR from disk. (In order to affect this operation the appropriate device_config is read from the specified SR&apos;s PBD on current host)
//
// Errors:
// SR_HAS_PBD - The SR is still connected to a host via a PBD. It cannot be destroyed or forgotten.
func (sR) Destroy(session *Session, sr SRRef) (err error) {
	method := "SR.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg)
	return
}

// AsyncDestroy: Destroy specified SR, removing SR-record from database and remove SR from disk. (In order to affect this operation the appropriate device_config is read from the specified SR&apos;s PBD on current host)
//
// Errors:
// SR_HAS_PBD - The SR is still connected to a host via a PBD. It cannot be destroyed or forgotten.
func (sR) AsyncDestroy(session *Session, sr SRRef) (retval TaskRef, err error) {
	method := "Async.SR.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Forget: Removing specified SR-record from database, without attempting to remove SR from disk
//
// Errors:
// SR_HAS_PBD - The SR is still connected to a host via a PBD. It cannot be destroyed or forgotten.
func (sR) Forget(session *Session, sr SRRef) (err error) {
	method := "SR.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg)
	return
}

// AsyncForget: Removing specified SR-record from database, without attempting to remove SR from disk
//
// Errors:
// SR_HAS_PBD - The SR is still connected to a host via a PBD. It cannot be destroyed or forgotten.
func (sR) AsyncForget(session *Session, sr SRRef) (retval TaskRef, err error) {
	method := "Async.SR.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Update: Refresh the fields on the SR object
func (sR) Update(session *Session, sr SRRef) (err error) {
	method := "SR.update"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg)
	return
}

// AsyncUpdate: Refresh the fields on the SR object
func (sR) AsyncUpdate(session *Session, sr SRRef) (retval TaskRef, err error) {
	method := "Async.SR.update"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetSupportedTypes: Return a set of all the SR types supported by the system
func (sR) GetSupportedTypes(session *Session) (retval []string, err error) {
	method := "SR.get_supported_types"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// Scan: Refreshes the list of VDIs associated with an SR
func (sR) Scan(session *Session, sr SRRef) (err error) {
	method := "SR.scan"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg)
	return
}

// AsyncScan: Refreshes the list of VDIs associated with an SR
func (sR) AsyncScan(session *Session, sr SRRef) (retval TaskRef, err error) {
	method := "Async.SR.scan"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Probe: Perform a backend-specific scan, using the given device_config.  If the device_config is complete, then this will return a list of the SRs present of this type on the device, if any.  If the device_config is partial, then a backend-specific scan will be performed, returning results that will guide the user in improving the device_config.
func (sR) Probe(session *Session, host HostRef, deviceConfig map[string]string, typeKey string, smConfig map[string]string) (retval string, err error) {
	method := "SR.probe"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, deviceConfigArg, typeKeyArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncProbe: Perform a backend-specific scan, using the given device_config.  If the device_config is complete, then this will return a list of the SRs present of this type on the device, if any.  If the device_config is partial, then a backend-specific scan will be performed, returning results that will guide the user in improving the device_config.
func (sR) AsyncProbe(session *Session, host HostRef, deviceConfig map[string]string, typeKey string, smConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.SR.probe"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, deviceConfigArg, typeKeyArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// ProbeExt: Perform a backend-specific scan, using the given device_config.  If the device_config is complete, then this will return a list of the SRs present of this type on the device, if any.  If the device_config is partial, then a backend-specific scan will be performed, returning results that will guide the user in improving the device_config.
func (sR) ProbeExt(session *Session, host HostRef, deviceConfig map[string]string, typeKey string, smConfig map[string]string) (retval []ProbeResultRecord, err error) {
	method := "SR.probe_ext"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, deviceConfigArg, typeKeyArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeProbeResultRecordSet(method+" -> ", result)
	return
}

// AsyncProbeExt: Perform a backend-specific scan, using the given device_config.  If the device_config is complete, then this will return a list of the SRs present of this type on the device, if any.  If the device_config is partial, then a backend-specific scan will be performed, returning results that will guide the user in improving the device_config.
func (sR) AsyncProbeExt(session *Session, host HostRef, deviceConfig map[string]string, typeKey string, smConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.SR.probe_ext"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, deviceConfigArg, typeKeyArg, smConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetShared: Sets the shared flag on the SR
func (sR) SetShared(session *Session, sr SRRef, value bool) (err error) {
	method := "SR.set_shared"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg, valueArg)
	return
}

// AsyncSetShared: Sets the shared flag on the SR
func (sR) AsyncSetShared(session *Session, sr SRRef, value bool) (retval TaskRef, err error) {
	method := "Async.SR.set_shared"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetNameLabel: Set the name label of the SR
func (sR) SetNameLabel(session *Session, sr SRRef, value string) (err error) {
	method := "SR.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg, valueArg)
	return
}

// AsyncSetNameLabel: Set the name label of the SR
func (sR) AsyncSetNameLabel(session *Session, sr SRRef, value string) (retval TaskRef, err error) {
	method := "Async.SR.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetNameDescription: Set the name description of the SR
func (sR) SetNameDescription(session *Session, sr SRRef, value string) (err error) {
	method := "SR.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg, valueArg)
	return
}

// AsyncSetNameDescription: Set the name description of the SR
func (sR) AsyncSetNameDescription(session *Session, sr SRRef, value string) (retval TaskRef, err error) {
	method := "Async.SR.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this SR
func (sR) CreateNewBlob(session *Session, sr SRRef, name string, mimeType string, public bool) (retval BlobRef, err error) {
	method := "SR.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	mimeTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "mime_type"), mimeType)
	if err != nil {
		return
	}
	publicArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "public"), public)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeBlobRef(method+" -> ", result)
	return
}

// AsyncCreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this SR
func (sR) AsyncCreateNewBlob(session *Session, sr SRRef, name string, mimeType string, public bool) (retval TaskRef, err error) {
	method := "Async.SR.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	mimeTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "mime_type"), mimeType)
	if err != nil {
		return
	}
	publicArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "public"), public)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetPhysicalSize: Sets the SR&apos;s physical_size field
func (sR) SetPhysicalSize(session *Session, self SRRef, value int) (err error) {
	method := "SR.set_physical_size"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AssertCanHostHaStatefile: Returns successfully if the given SR can host an HA statefile. Otherwise returns an error to explain why not
func (sR) AssertCanHostHaStatefile(session *Session, sr SRRef) (err error) {
	method := "SR.assert_can_host_ha_statefile"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg)
	return
}

// AsyncAssertCanHostHaStatefile: Returns successfully if the given SR can host an HA statefile. Otherwise returns an error to explain why not
func (sR) AsyncAssertCanHostHaStatefile(session *Session, sr SRRef) (retval TaskRef, err error) {
	method := "Async.SR.assert_can_host_ha_statefile"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// AssertSupportsDatabaseReplication: Returns successfully if the given SR supports database replication. Otherwise returns an error to explain why not.
func (sR) AssertSupportsDatabaseReplication(session *Session, sr SRRef) (err error) {
	method := "SR.assert_supports_database_replication"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg)
	return
}

// AsyncAssertSupportsDatabaseReplication: Returns successfully if the given SR supports database replication. Otherwise returns an error to explain why not.
func (sR) AsyncAssertSupportsDatabaseReplication(session *Session, sr SRRef) (retval TaskRef, err error) {
	method := "Async.SR.assert_supports_database_replication"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EnableDatabaseReplication:
func (sR) EnableDatabaseReplication(session *Session, sr SRRef) (err error) {
	method := "SR.enable_database_replication"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg)
	return
}

// AsyncEnableDatabaseReplication:
func (sR) AsyncEnableDatabaseReplication(session *Session, sr SRRef) (retval TaskRef, err error) {
	method := "Async.SR.enable_database_replication"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// DisableDatabaseReplication:
func (sR) DisableDatabaseReplication(session *Session, sr SRRef) (err error) {
	method := "SR.disable_database_replication"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg)
	return
}

// AsyncDisableDatabaseReplication:
func (sR) AsyncDisableDatabaseReplication(session *Session, sr SRRef) (retval TaskRef, err error) {
	method := "Async.SR.disable_database_replication"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetDataSources:
func (sR) GetDataSources(session *Session, sr SRRef) (retval []DataSourceRecord, err error) {
	method := "SR.get_data_sources"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeDataSourceRecordSet(method+" -> ", result)
	return
}

// RecordDataSource: Start recording the specified data source
func (sR) RecordDataSource(session *Session, sr SRRef, dataSource string) (err error) {
	method := "SR.record_data_source"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	dataSourceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "data_source"), dataSource)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg, dataSourceArg)
	return
}

// QueryDataSource: Query the latest value of the specified data source
func (sR) QueryDataSource(session *Session, sr SRRef, dataSource string) (retval float64, err error) {
	method := "SR.query_data_source"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	dataSourceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "data_source"), dataSource)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg, dataSourceArg)
	if err != nil {
		return
	}
	retval, err = deserializeFloat(method+" -> ", result)
	return
}

// ForgetDataSourceArchives: Forget the recorded statistics related to the specified data source
func (sR) ForgetDataSourceArchives(session *Session, sr SRRef, dataSource string) (err error) {
	method := "SR.forget_data_source_archives"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	dataSourceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "data_source"), dataSource)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg, dataSourceArg)
	return
}

// GetAll: Return a list of all the SRs known to the system.
func (sR) GetAll(session *Session) (retval []SRRef, err error) {
	method := "SR.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of SR references to SR records for all SRs known to the system.
func (sR) GetAllRecords(session *Session) (retval map[SRRef]SRRecord, err error) {
	method := "SR.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRefToSRRecordMap(method+" -> ", result)
	return
}

