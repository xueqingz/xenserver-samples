package xenapi

import (
	"fmt"
	"time"
)

type VDIRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VdiOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VdiOperations
	// storage repository in which the VDI resides
	SR SRRef
	// list of vbds that refer to this disk
	VBDs []VBDRef
	// list of crash dumps that refer to this disk
	CrashDumps []CrashdumpRef
	// size of disk as presented to the guest (in bytes). Note that, depending on storage backend type, requested size may not be respected exactly
	VirtualSize int
	// amount of physical space that the disk image is currently taking up on the storage repository (in bytes)
	PhysicalUtilisation int
	// type of the VDI
	Type VdiType
	// true if this disk may be shared
	Sharable bool
	// true if this disk may ONLY be mounted read-only
	ReadOnly bool
	// additional configuration
	OtherConfig map[string]string
	// true if this disk is locked at the storage level
	StorageLock bool
	// location information
	Location string
	//
	Managed bool
	// true if SR scan operation reported this VDI as not present on disk
	Missing bool
	// This field is always null. Deprecated
	Parent VDIRef
	// data to be inserted into the xenstore tree (/local/domain/0/backend/vbd/&lt;domid&gt;/&lt;device-id&gt;/sm-data) after the VDI is attached. This is generally set by the SM backends on vdi_attach.
	XenstoreData map[string]string
	// SM dependent data
	SmConfig map[string]string
	// true if this is a snapshot.
	IsASnapshot bool
	// Ref pointing to the VDI this snapshot is of.
	SnapshotOf VDIRef
	// List pointing to all the VDIs snapshots.
	Snapshots []VDIRef
	// Date/time when this snapshot was created.
	SnapshotTime time.Time
	// user-specified tags for categorization purposes
	Tags []string
	// true if this VDI is to be cached in the local cache SR
	AllowCaching bool
	// The behaviour of this VDI on a VM boot
	OnBoot OnBoot
	// The pool whose metadata is contained in this VDI
	MetadataOfPool PoolRef
	// Whether this VDI contains the latest known accessible metadata for the pool
	MetadataLatest bool
	// Whether this VDI is a Tools ISO
	IsToolsIso bool
	// True if changed blocks are tracked for this VDI
	CbtEnabled bool
}

type VDIRef string

// A virtual disk image
type vDI struct{}

var VDI vDI

// GetRecord: Get a record containing the current state of the given VDI.
func (vDI) GetRecord(session *Session, self VDIRef) (retval VDIRecord, err error) {
	method := "VDI.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VDI instance with the specified UUID.
func (vDI) GetByUUID(session *Session, uUID string) (retval VDIRef, err error) {
	method := "VDI.get_by_uuid"
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
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// Create: Create a new VDI instance, and return its handle. The constructor args are: name_label, name_description, SR*, virtual_size*, type*, sharable*, read_only*, other_config*, xenstore_data, sm_config, tags (* = non-optional).
func (vDI) Create(session *Session, args VDIRecord) (retval VDIRef, err error) {
	method := "VDI.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVDIRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new VDI instance, and return its handle. The constructor args are: name_label, name_description, SR*, virtual_size*, type*, sharable*, read_only*, other_config*, xenstore_data, sm_config, tags (* = non-optional).
func (vDI) AsyncCreate(session *Session, args VDIRecord) (retval TaskRef, err error) {
	method := "Async.VDI.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVDIRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Destroy the specified VDI instance.
func (vDI) Destroy(session *Session, self VDIRef) (err error) {
	method := "VDI.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified VDI instance.
func (vDI) AsyncDestroy(session *Session, self VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByNameLabel: Get all the VDI instances with the given label.
func (vDI) GetByNameLabel(session *Session, label string) (retval []VDIRef, err error) {
	method := "VDI.get_by_name_label"
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
	retval, err = deserializeVDIRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VDI.
func (vDI) GetUUID(session *Session, self VDIRef) (retval string, err error) {
	method := "VDI.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given VDI.
func (vDI) GetNameLabel(session *Session, self VDIRef) (retval string, err error) {
	method := "VDI.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given VDI.
func (vDI) GetNameDescription(session *Session, self VDIRef) (retval string, err error) {
	method := "VDI.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given VDI.
func (vDI) GetAllowedOperations(session *Session, self VDIRef) (retval []VdiOperations, err error) {
	method := "VDI.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVdiOperationsSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given VDI.
func (vDI) GetCurrentOperations(session *Session, self VDIRef) (retval map[string]VdiOperations, err error) {
	method := "VDI.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumVdiOperationsMap(method+" -> ", result)
	return
}

// GetSR: Get the SR field of the given VDI.
func (vDI) GetSR(session *Session, self VDIRef) (retval SRRef, err error) {
	method := "VDI.get_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRef(method+" -> ", result)
	return
}

// GetVBDs: Get the VBDs field of the given VDI.
func (vDI) GetVBDs(session *Session, self VDIRef) (retval []VBDRef, err error) {
	method := "VDI.get_VBDs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRefSet(method+" -> ", result)
	return
}

// GetCrashDumps: Get the crash_dumps field of the given VDI.
func (vDI) GetCrashDumps(session *Session, self VDIRef) (retval []CrashdumpRef, err error) {
	method := "VDI.get_crash_dumps"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeCrashdumpRefSet(method+" -> ", result)
	return
}

// GetVirtualSize: Get the virtual_size field of the given VDI.
func (vDI) GetVirtualSize(session *Session, self VDIRef) (retval int, err error) {
	method := "VDI.get_virtual_size"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPhysicalUtilisation: Get the physical_utilisation field of the given VDI.
func (vDI) GetPhysicalUtilisation(session *Session, self VDIRef) (retval int, err error) {
	method := "VDI.get_physical_utilisation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetType: Get the type field of the given VDI.
func (vDI) GetType(session *Session, self VDIRef) (retval VdiType, err error) {
	method := "VDI.get_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVdiType(method+" -> ", result)
	return
}

// GetSharable: Get the sharable field of the given VDI.
func (vDI) GetSharable(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_sharable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetReadOnly: Get the read_only field of the given VDI.
func (vDI) GetReadOnly(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_read_only"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given VDI.
func (vDI) GetOtherConfig(session *Session, self VDIRef) (retval map[string]string, err error) {
	method := "VDI.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStorageLock: Get the storage_lock field of the given VDI.
func (vDI) GetStorageLock(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_storage_lock"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLocation: Get the location field of the given VDI.
func (vDI) GetLocation(session *Session, self VDIRef) (retval string, err error) {
	method := "VDI.get_location"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetManaged: Get the managed field of the given VDI.
func (vDI) GetManaged(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_managed"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMissing: Get the missing field of the given VDI.
func (vDI) GetMissing(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_missing"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetParent: Get the parent field of the given VDI.
func (vDI) GetParent(session *Session, self VDIRef) (retval VDIRef, err error) {
	method := "VDI.get_parent"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// GetXenstoreData: Get the xenstore_data field of the given VDI.
func (vDI) GetXenstoreData(session *Session, self VDIRef) (retval map[string]string, err error) {
	method := "VDI.get_xenstore_data"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSmConfig: Get the sm_config field of the given VDI.
func (vDI) GetSmConfig(session *Session, self VDIRef) (retval map[string]string, err error) {
	method := "VDI.get_sm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsASnapshot: Get the is_a_snapshot field of the given VDI.
func (vDI) GetIsASnapshot(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_is_a_snapshot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSnapshotOf: Get the snapshot_of field of the given VDI.
func (vDI) GetSnapshotOf(session *Session, self VDIRef) (retval VDIRef, err error) {
	method := "VDI.get_snapshot_of"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// GetSnapshots: Get the snapshots field of the given VDI.
func (vDI) GetSnapshots(session *Session, self VDIRef) (retval []VDIRef, err error) {
	method := "VDI.get_snapshots"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSnapshotTime: Get the snapshot_time field of the given VDI.
func (vDI) GetSnapshotTime(session *Session, self VDIRef) (retval time.Time, err error) {
	method := "VDI.get_snapshot_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTags: Get the tags field of the given VDI.
func (vDI) GetTags(session *Session, self VDIRef) (retval []string, err error) {
	method := "VDI.get_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowCaching: Get the allow_caching field of the given VDI.
func (vDI) GetAllowCaching(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_allow_caching"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOnBoot: Get the on_boot field of the given VDI.
func (vDI) GetOnBoot(session *Session, self VDIRef) (retval OnBoot, err error) {
	method := "VDI.get_on_boot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumOnBoot(method+" -> ", result)
	return
}

// GetMetadataOfPool: Get the metadata_of_pool field of the given VDI.
func (vDI) GetMetadataOfPool(session *Session, self VDIRef) (retval PoolRef, err error) {
	method := "VDI.get_metadata_of_pool"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolRef(method+" -> ", result)
	return
}

// GetMetadataLatest: Get the metadata_latest field of the given VDI.
func (vDI) GetMetadataLatest(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_metadata_latest"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsToolsIso: Get the is_tools_iso field of the given VDI.
func (vDI) GetIsToolsIso(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_is_tools_iso"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCbtEnabled: Get the cbt_enabled field of the given VDI.
func (vDI) GetCbtEnabled(session *Session, self VDIRef) (retval bool, err error) {
	method := "VDI.get_cbt_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given VDI.
func (vDI) SetOtherConfig(session *Session, self VDIRef, value map[string]string) (err error) {
	method := "VDI.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VDI.
func (vDI) AddToOtherConfig(session *Session, self VDIRef, key string, value string) (err error) {
	method := "VDI.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VDI.  If the key is not in that Map, then do nothing.
func (vDI) RemoveFromOtherConfig(session *Session, self VDIRef, key string) (err error) {
	method := "VDI.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetXenstoreData: Set the xenstore_data field of the given VDI.
func (vDI) SetXenstoreData(session *Session, self VDIRef, value map[string]string) (err error) {
	method := "VDI.set_xenstore_data"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToXenstoreData: Add the given key-value pair to the xenstore_data field of the given VDI.
func (vDI) AddToXenstoreData(session *Session, self VDIRef, key string, value string) (err error) {
	method := "VDI.add_to_xenstore_data"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromXenstoreData: Remove the given key and its corresponding value from the xenstore_data field of the given VDI.  If the key is not in that Map, then do nothing.
func (vDI) RemoveFromXenstoreData(session *Session, self VDIRef, key string) (err error) {
	method := "VDI.remove_from_xenstore_data"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetSmConfig: Set the sm_config field of the given VDI.
func (vDI) SetSmConfig(session *Session, self VDIRef, value map[string]string) (err error) {
	method := "VDI.set_sm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToSmConfig: Add the given key-value pair to the sm_config field of the given VDI.
func (vDI) AddToSmConfig(session *Session, self VDIRef, key string, value string) (err error) {
	method := "VDI.add_to_sm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromSmConfig: Remove the given key and its corresponding value from the sm_config field of the given VDI.  If the key is not in that Map, then do nothing.
func (vDI) RemoveFromSmConfig(session *Session, self VDIRef, key string) (err error) {
	method := "VDI.remove_from_sm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetTags: Set the tags field of the given VDI.
func (vDI) SetTags(session *Session, self VDIRef, value []string) (err error) {
	method := "VDI.set_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddTags: Add the given value to the tags field of the given VDI.  If the value is already in that Set, then do nothing.
func (vDI) AddTags(session *Session, self VDIRef, value string) (err error) {
	method := "VDI.add_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveTags: Remove the given value from the tags field of the given VDI.  If the value is not in that Set, then do nothing.
func (vDI) RemoveTags(session *Session, self VDIRef, value string) (err error) {
	method := "VDI.remove_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Snapshot: Take a read-only snapshot of the VDI, returning a reference to the snapshot. If any driver_params are specified then these are passed through to the storage-specific substrate driver that takes the snapshot. NB the snapshot lives in the same Storage Repository as its parent.
func (vDI) Snapshot(session *Session, vdi VDIRef, driverParams map[string]string) (retval VDIRef, err error) {
	method := "VDI.snapshot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	driverParamsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "driver_params"), driverParams)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, driverParamsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// AsyncSnapshot: Take a read-only snapshot of the VDI, returning a reference to the snapshot. If any driver_params are specified then these are passed through to the storage-specific substrate driver that takes the snapshot. NB the snapshot lives in the same Storage Repository as its parent.
func (vDI) AsyncSnapshot(session *Session, vdi VDIRef, driverParams map[string]string) (retval TaskRef, err error) {
	method := "Async.VDI.snapshot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	driverParamsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "driver_params"), driverParams)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, driverParamsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Clone: Take an exact copy of the VDI and return a reference to the new disk. If any driver_params are specified then these are passed through to the storage-specific substrate driver that implements the clone operation. NB the clone lives in the same Storage Repository as its parent.
func (vDI) Clone(session *Session, vdi VDIRef, driverParams map[string]string) (retval VDIRef, err error) {
	method := "VDI.clone"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	driverParamsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "driver_params"), driverParams)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, driverParamsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// AsyncClone: Take an exact copy of the VDI and return a reference to the new disk. If any driver_params are specified then these are passed through to the storage-specific substrate driver that implements the clone operation. NB the clone lives in the same Storage Repository as its parent.
func (vDI) AsyncClone(session *Session, vdi VDIRef, driverParams map[string]string) (retval TaskRef, err error) {
	method := "Async.VDI.clone"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	driverParamsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "driver_params"), driverParams)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, driverParamsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Resize: Resize the VDI.
func (vDI) Resize(session *Session, vdi VDIRef, size int) (err error) {
	method := "VDI.resize"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	sizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "size"), size)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vdiArg, sizeArg)
	return
}

// AsyncResize: Resize the VDI.
func (vDI) AsyncResize(session *Session, vdi VDIRef, size int) (retval TaskRef, err error) {
	method := "Async.VDI.resize"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	sizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "size"), size)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, sizeArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// ResizeOnline: Resize the VDI which may or may not be attached to running guests.
func (vDI) ResizeOnline(session *Session, vdi VDIRef, size int) (err error) {
	method := "VDI.resize_online"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	sizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "size"), size)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vdiArg, sizeArg)
	return
}

// AsyncResizeOnline: Resize the VDI which may or may not be attached to running guests.
func (vDI) AsyncResizeOnline(session *Session, vdi VDIRef, size int) (retval TaskRef, err error) {
	method := "Async.VDI.resize_online"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	sizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "size"), size)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, sizeArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Introduce: Create a new VDI record in the database only
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
func (vDI) Introduce(session *Session, uUID string, nameLabel string, nameDescription string, sR SRRef, typeKey VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef) (retval VDIRef, err error) {
	method := "VDI.introduce"
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
	sRArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "SR"), sR)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeEnumVdiType(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	sharableArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "sharable"), sharable)
	if err != nil {
		return
	}
	readOnlyArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "read_only"), readOnly)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	locationArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "location"), location)
	if err != nil {
		return
	}
	xenstoreDataArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "xenstore_data"), xenstoreData)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	managedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "managed"), managed)
	if err != nil {
		return
	}
	virtualSizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "virtual_size"), virtualSize)
	if err != nil {
		return
	}
	physicalUtilisationArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "physical_utilisation"), physicalUtilisation)
	if err != nil {
		return
	}
	metadataOfPoolArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "metadata_of_pool"), metadataOfPool)
	if err != nil {
		return
	}
	isASnapshotArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "is_a_snapshot"), isASnapshot)
	if err != nil {
		return
	}
	snapshotTimeArg, err := serializeTime(fmt.Sprintf("%s(%s)", method, "snapshot_time"), snapshotTime)
	if err != nil {
		return
	}
	snapshotOfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "snapshot_of"), snapshotOf)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uUIDArg, nameLabelArg, nameDescriptionArg, sRArg, typeKeyArg, sharableArg, readOnlyArg, otherConfigArg, locationArg, xenstoreDataArg, smConfigArg, managedArg, virtualSizeArg, physicalUtilisationArg, metadataOfPoolArg, isASnapshotArg, snapshotTimeArg, snapshotOfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// AsyncIntroduce: Create a new VDI record in the database only
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
func (vDI) AsyncIntroduce(session *Session, uUID string, nameLabel string, nameDescription string, sR SRRef, typeKey VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.introduce"
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
	sRArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "SR"), sR)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeEnumVdiType(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	sharableArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "sharable"), sharable)
	if err != nil {
		return
	}
	readOnlyArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "read_only"), readOnly)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	locationArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "location"), location)
	if err != nil {
		return
	}
	xenstoreDataArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "xenstore_data"), xenstoreData)
	if err != nil {
		return
	}
	smConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "sm_config"), smConfig)
	if err != nil {
		return
	}
	managedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "managed"), managed)
	if err != nil {
		return
	}
	virtualSizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "virtual_size"), virtualSize)
	if err != nil {
		return
	}
	physicalUtilisationArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "physical_utilisation"), physicalUtilisation)
	if err != nil {
		return
	}
	metadataOfPoolArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "metadata_of_pool"), metadataOfPool)
	if err != nil {
		return
	}
	isASnapshotArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "is_a_snapshot"), isASnapshot)
	if err != nil {
		return
	}
	snapshotTimeArg, err := serializeTime(fmt.Sprintf("%s(%s)", method, "snapshot_time"), snapshotTime)
	if err != nil {
		return
	}
	snapshotOfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "snapshot_of"), snapshotOf)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uUIDArg, nameLabelArg, nameDescriptionArg, sRArg, typeKeyArg, sharableArg, readOnlyArg, otherConfigArg, locationArg, xenstoreDataArg, smConfigArg, managedArg, virtualSizeArg, physicalUtilisationArg, metadataOfPoolArg, isASnapshotArg, snapshotTimeArg, snapshotOfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Update: Ask the storage backend to refresh the fields in the VDI object
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
func (vDI) Update(session *Session, vdi VDIRef) (err error) {
	method := "VDI.update"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vdiArg)
	return
}

// AsyncUpdate: Ask the storage backend to refresh the fields in the VDI object
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
func (vDI) AsyncUpdate(session *Session, vdi VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.update"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Copy: Copy either a full VDI or the block differences between two VDIs into either a fresh VDI or an existing VDI.
//
// Errors:
// VDI_READONLY - The operation required write access but this VDI is read-only
// VDI_TOO_SMALL - The VDI is too small. Please resize it to at least the minimum size.
// VDI_NOT_SPARSE - The VDI is not stored using a sparse format. It is not possible to query and manipulate only the changed blocks (or &apos;block differences&apos; or &apos;disk deltas&apos;) between two VDIs. Please select a VDI which uses a sparse-aware technology such as VHD.
func (vDI) Copy(session *Session, vdi VDIRef, sr SRRef, baseVdi VDIRef, intoVdi VDIRef) (retval VDIRef, err error) {
	method := "VDI.copy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	baseVdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "base_vdi"), baseVdi)
	if err != nil {
		return
	}
	intoVdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "into_vdi"), intoVdi)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, srArg, baseVdiArg, intoVdiArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// AsyncCopy: Copy either a full VDI or the block differences between two VDIs into either a fresh VDI or an existing VDI.
//
// Errors:
// VDI_READONLY - The operation required write access but this VDI is read-only
// VDI_TOO_SMALL - The VDI is too small. Please resize it to at least the minimum size.
// VDI_NOT_SPARSE - The VDI is not stored using a sparse format. It is not possible to query and manipulate only the changed blocks (or &apos;block differences&apos; or &apos;disk deltas&apos;) between two VDIs. Please select a VDI which uses a sparse-aware technology such as VHD.
func (vDI) AsyncCopy(session *Session, vdi VDIRef, sr SRRef, baseVdi VDIRef, intoVdi VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.copy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	baseVdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "base_vdi"), baseVdi)
	if err != nil {
		return
	}
	intoVdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "into_vdi"), intoVdi)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, srArg, baseVdiArg, intoVdiArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Forget: Removes a VDI record from the database
func (vDI) Forget(session *Session, vdi VDIRef) (err error) {
	method := "VDI.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vdiArg)
	return
}

// AsyncForget: Removes a VDI record from the database
func (vDI) AsyncForget(session *Session, vdi VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetSharable: Sets the VDI&apos;s sharable field
func (vDI) SetSharable(session *Session, self VDIRef, value bool) (err error) {
	method := "VDI.set_sharable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetReadOnly: Sets the VDI&apos;s read_only field
func (vDI) SetReadOnly(session *Session, self VDIRef, value bool) (err error) {
	method := "VDI.set_read_only"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetNameLabel: Set the name label of the VDI. This can only happen when then its SR is currently attached.
func (vDI) SetNameLabel(session *Session, self VDIRef, value string) (err error) {
	method := "VDI.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetNameLabel: Set the name label of the VDI. This can only happen when then its SR is currently attached.
func (vDI) AsyncSetNameLabel(session *Session, self VDIRef, value string) (retval TaskRef, err error) {
	method := "Async.VDI.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetNameDescription: Set the name description of the VDI. This can only happen when its SR is currently attached.
func (vDI) SetNameDescription(session *Session, self VDIRef, value string) (err error) {
	method := "VDI.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetNameDescription: Set the name description of the VDI. This can only happen when its SR is currently attached.
func (vDI) AsyncSetNameDescription(session *Session, self VDIRef, value string) (retval TaskRef, err error) {
	method := "Async.VDI.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetOnBoot: Set the value of the on_boot parameter. This value can only be changed when the VDI is not attached to a running VM.
func (vDI) SetOnBoot(session *Session, self VDIRef, value OnBoot) (err error) {
	method := "VDI.set_on_boot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumOnBoot(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetOnBoot: Set the value of the on_boot parameter. This value can only be changed when the VDI is not attached to a running VM.
func (vDI) AsyncSetOnBoot(session *Session, self VDIRef, value OnBoot) (retval TaskRef, err error) {
	method := "Async.VDI.set_on_boot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumOnBoot(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetAllowCaching: Set the value of the allow_caching parameter. This value can only be changed when the VDI is not attached to a running VM. The caching behaviour is only affected by this flag for VHD-based VDIs that have one parent and no child VHDs. Moreover, caching only takes place when the host running the VM containing this VDI has a nominated SR for local caching.
func (vDI) SetAllowCaching(session *Session, self VDIRef, value bool) (err error) {
	method := "VDI.set_allow_caching"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetAllowCaching: Set the value of the allow_caching parameter. This value can only be changed when the VDI is not attached to a running VM. The caching behaviour is only affected by this flag for VHD-based VDIs that have one parent and no child VHDs. Moreover, caching only takes place when the host running the VM containing this VDI has a nominated SR for local caching.
func (vDI) AsyncSetAllowCaching(session *Session, self VDIRef, value bool) (retval TaskRef, err error) {
	method := "Async.VDI.set_allow_caching"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// OpenDatabase: Load the metadata found on the supplied VDI and return a session reference which can be used in API calls to query its contents.
func (vDI) OpenDatabase(session *Session, self VDIRef) (retval SessionRef, err error) {
	method := "VDI.open_database"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSessionRef(method+" -> ", result)
	return
}

// AsyncOpenDatabase: Load the metadata found on the supplied VDI and return a session reference which can be used in API calls to query its contents.
func (vDI) AsyncOpenDatabase(session *Session, self VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.open_database"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// ReadDatabasePoolUUID: Check the VDI cache for the pool UUID of the database on this VDI.
func (vDI) ReadDatabasePoolUUID(session *Session, self VDIRef) (retval string, err error) {
	method := "VDI.read_database_pool_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncReadDatabasePoolUUID: Check the VDI cache for the pool UUID of the database on this VDI.
func (vDI) AsyncReadDatabasePoolUUID(session *Session, self VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.read_database_pool_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// PoolMigrate: Migrate a VDI, which may be attached to a running guest, to a different SR. The destination SR must be visible to the guest.
func (vDI) PoolMigrate(session *Session, vdi VDIRef, sr SRRef, options map[string]string) (retval VDIRef, err error) {
	method := "VDI.pool_migrate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, srArg, optionsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// AsyncPoolMigrate: Migrate a VDI, which may be attached to a running guest, to a different SR. The destination SR must be visible to the guest.
func (vDI) AsyncPoolMigrate(session *Session, vdi VDIRef, sr SRRef, options map[string]string) (retval TaskRef, err error) {
	method := "Async.VDI.pool_migrate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiArg, srArg, optionsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EnableCbt: Enable changed block tracking for the VDI. This call is idempotent - enabling CBT for a VDI for which CBT is already enabled results in a no-op, and no error will be thrown.
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
// VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
// SR_NOT_ATTACHED - The SR is not attached.
// SR_HAS_NO_PBDS - The SR has no attached PBDs
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
// VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION - This operation is not permitted on VDIs in the &apos;on-boot=reset&apos; mode, or on VMs having such VDIs.
func (vDI) EnableCbt(session *Session, self VDIRef) (err error) {
	method := "VDI.enable_cbt"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncEnableCbt: Enable changed block tracking for the VDI. This call is idempotent - enabling CBT for a VDI for which CBT is already enabled results in a no-op, and no error will be thrown.
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
// VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
// SR_NOT_ATTACHED - The SR is not attached.
// SR_HAS_NO_PBDS - The SR has no attached PBDs
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
// VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION - This operation is not permitted on VDIs in the &apos;on-boot=reset&apos; mode, or on VMs having such VDIs.
func (vDI) AsyncEnableCbt(session *Session, self VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.enable_cbt"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// DisableCbt: Disable changed block tracking for the VDI. This call is only allowed on VDIs that support enabling CBT. It is an idempotent operation - disabling CBT for a VDI for which CBT is not enabled results in a no-op, and no error will be thrown.
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
// VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
// SR_NOT_ATTACHED - The SR is not attached.
// SR_HAS_NO_PBDS - The SR has no attached PBDs
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
// VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION - This operation is not permitted on VDIs in the &apos;on-boot=reset&apos; mode, or on VMs having such VDIs.
func (vDI) DisableCbt(session *Session, self VDIRef) (err error) {
	method := "VDI.disable_cbt"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDisableCbt: Disable changed block tracking for the VDI. This call is only allowed on VDIs that support enabling CBT. It is an idempotent operation - disabling CBT for a VDI for which CBT is not enabled results in a no-op, and no error will be thrown.
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
// VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
// SR_NOT_ATTACHED - The SR is not attached.
// SR_HAS_NO_PBDS - The SR has no attached PBDs
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
// VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION - This operation is not permitted on VDIs in the &apos;on-boot=reset&apos; mode, or on VMs having such VDIs.
func (vDI) AsyncDisableCbt(session *Session, self VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.disable_cbt"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// DataDestroy: Delete the data of the snapshot VDI, but keep its changed block tracking metadata. When successful, this call changes the type of the VDI to cbt_metadata. This operation is idempotent: calling it on a VDI of type cbt_metadata results in a no-op, and no error will be thrown.
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
// VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
// SR_NOT_ATTACHED - The SR is not attached.
// SR_HAS_NO_PBDS - The SR has no attached PBDs
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
// VDI_NO_CBT_METADATA - The requested operation is not allowed because the specified VDI does not have changed block tracking metadata.
// VDI_IN_USE - This operation cannot be performed because this VDI is in use by some other operation
// VDI_IS_A_PHYSICAL_DEVICE - The operation cannot be performed on physical device
func (vDI) DataDestroy(session *Session, self VDIRef) (err error) {
	method := "VDI.data_destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDataDestroy: Delete the data of the snapshot VDI, but keep its changed block tracking metadata. When successful, this call changes the type of the VDI to cbt_metadata. This operation is idempotent: calling it on a VDI of type cbt_metadata results in a no-op, and no error will be thrown.
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
// VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
// SR_NOT_ATTACHED - The SR is not attached.
// SR_HAS_NO_PBDS - The SR has no attached PBDs
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
// VDI_NO_CBT_METADATA - The requested operation is not allowed because the specified VDI does not have changed block tracking metadata.
// VDI_IN_USE - This operation cannot be performed because this VDI is in use by some other operation
// VDI_IS_A_PHYSICAL_DEVICE - The operation cannot be performed on physical device
func (vDI) AsyncDataDestroy(session *Session, self VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.data_destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// ListChangedBlocks: Compare two VDIs in 64k block increments and report which blocks differ. This operation is not allowed when vdi_to is attached to a VM.
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
// VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
// SR_NOT_ATTACHED - The SR is not attached.
// SR_HAS_NO_PBDS - The SR has no attached PBDs
// VDI_IN_USE - This operation cannot be performed because this VDI is in use by some other operation
func (vDI) ListChangedBlocks(session *Session, vdiFrom VDIRef, vdiTo VDIRef) (retval string, err error) {
	method := "VDI.list_changed_blocks"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiFromArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi_from"), vdiFrom)
	if err != nil {
		return
	}
	vdiToArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi_to"), vdiTo)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiFromArg, vdiToArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncListChangedBlocks: Compare two VDIs in 64k block increments and report which blocks differ. This operation is not allowed when vdi_to is attached to a VM.
//
// Errors:
// SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR&apos;s allowed operations)
// VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
// SR_NOT_ATTACHED - The SR is not attached.
// SR_HAS_NO_PBDS - The SR has no attached PBDs
// VDI_IN_USE - This operation cannot be performed because this VDI is in use by some other operation
func (vDI) AsyncListChangedBlocks(session *Session, vdiFrom VDIRef, vdiTo VDIRef) (retval TaskRef, err error) {
	method := "Async.VDI.list_changed_blocks"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vdiFromArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi_from"), vdiFrom)
	if err != nil {
		return
	}
	vdiToArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi_to"), vdiTo)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vdiFromArg, vdiToArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetNbdInfo: Get details specifying how to access this VDI via a Network Block Device server. For each of a set of NBD server addresses on which the VDI is available, the return value set contains a vdi_nbd_server_info object that contains an exportname to request once the NBD connection is established, and connection details for the address. An empty list is returned if there is no network that has a PIF on a host with access to the relevant SR, or if no such network has been assigned an NBD-related purpose in its purpose field. To access the given VDI, any of the vdi_nbd_server_info objects can be used to make a connection to a server, and then the VDI will be available by requesting the exportname.
//
// Errors:
// VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
func (vDI) GetNbdInfo(session *Session, self VDIRef) (retval []VdiNbdServerInfoRecord, err error) {
	method := "VDI.get_nbd_info"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVdiNbdServerInfoRecordSet(method+" -> ", result)
	return
}

// GetAll: Return a list of all the VDIs known to the system.
func (vDI) GetAll(session *Session) (retval []VDIRef, err error) {
	method := "VDI.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VDI references to VDI records for all VDIs known to the system.
func (vDI) GetAllRecords(session *Session) (retval map[VDIRef]VDIRecord, err error) {
	method := "VDI.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRefToVDIRecordMap(method+" -> ", result)
	return
}
