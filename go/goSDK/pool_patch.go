package xenapi

import (
	"fmt"
)

type PoolPatchRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// Patch version number
	Version string
	// Size of the patch
	Size int
	// This patch should be applied across the entire pool
	PoolApplied bool
	// This hosts this patch is applied to.
	HostPatches []HostPatchRef
	// What the client should do after this patch has been applied.
	AfterApplyGuidance []AfterApplyGuidance
	// A reference to the associated pool_update object
	PoolUpdate PoolUpdateRef
	// additional configuration
	OtherConfig map[string]string
}

type PoolPatchRef string

// Pool-wide patches
type poolPatch struct{}

var PoolPatch poolPatch

// GetRecord: Get a record containing the current state of the given pool_patch.
func (poolPatch) GetRecord(session *Session, self PoolPatchRef) (retval PoolPatchRecord, err error) {
	method := "pool_patch.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolPatchRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the pool_patch instance with the specified UUID.
func (poolPatch) GetByUUID(session *Session, uUID string) (retval PoolPatchRef, err error) {
	method := "pool_patch.get_by_uuid"
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
	retval, err = deserializePoolPatchRef(method+" -> ", result)
	return
}

// GetByNameLabel: Get all the pool_patch instances with the given label.
func (poolPatch) GetByNameLabel(session *Session, label string) (retval []PoolPatchRef, err error) {
	method := "pool_patch.get_by_name_label"
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
	retval, err = deserializePoolPatchRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given pool_patch.
func (poolPatch) GetUUID(session *Session, self PoolPatchRef) (retval string, err error) {
	method := "pool_patch.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given pool_patch.
func (poolPatch) GetNameLabel(session *Session, self PoolPatchRef) (retval string, err error) {
	method := "pool_patch.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given pool_patch.
func (poolPatch) GetNameDescription(session *Session, self PoolPatchRef) (retval string, err error) {
	method := "pool_patch.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVersion: Get the version field of the given pool_patch.
func (poolPatch) GetVersion(session *Session, self PoolPatchRef) (retval string, err error) {
	method := "pool_patch.get_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSize: Get the size field of the given pool_patch.
func (poolPatch) GetSize(session *Session, self PoolPatchRef) (retval int, err error) {
	method := "pool_patch.get_size"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPoolApplied: Get the pool_applied field of the given pool_patch.
func (poolPatch) GetPoolApplied(session *Session, self PoolPatchRef) (retval bool, err error) {
	method := "pool_patch.get_pool_applied"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHostPatches: Get the host_patches field of the given pool_patch.
func (poolPatch) GetHostPatches(session *Session, self PoolPatchRef) (retval []HostPatchRef, err error) {
	method := "pool_patch.get_host_patches"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostPatchRefSet(method+" -> ", result)
	return
}

// GetAfterApplyGuidance: Get the after_apply_guidance field of the given pool_patch.
func (poolPatch) GetAfterApplyGuidance(session *Session, self PoolPatchRef) (retval []AfterApplyGuidance, err error) {
	method := "pool_patch.get_after_apply_guidance"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumAfterApplyGuidanceSet(method+" -> ", result)
	return
}

// GetPoolUpdate: Get the pool_update field of the given pool_patch.
func (poolPatch) GetPoolUpdate(session *Session, self PoolPatchRef) (retval PoolUpdateRef, err error) {
	method := "pool_patch.get_pool_update"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolUpdateRef(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given pool_patch.
func (poolPatch) GetOtherConfig(session *Session, self PoolPatchRef) (retval map[string]string, err error) {
	method := "pool_patch.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given pool_patch.
func (poolPatch) SetOtherConfig(session *Session, self PoolPatchRef, value map[string]string) (err error) {
	method := "pool_patch.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given pool_patch.
func (poolPatch) AddToOtherConfig(session *Session, self PoolPatchRef, key string, value string) (err error) {
	method := "pool_patch.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given pool_patch.  If the key is not in that Map, then do nothing.
func (poolPatch) RemoveFromOtherConfig(session *Session, self PoolPatchRef, key string) (err error) {
	method := "pool_patch.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Apply: Apply the selected patch to a host and return its output
func (poolPatch) Apply(session *Session, self PoolPatchRef, host HostRef) (retval string, err error) {
	method := "pool_patch.apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncApply: Apply the selected patch to a host and return its output
func (poolPatch) AsyncApply(session *Session, self PoolPatchRef, host HostRef) (retval TaskRef, err error) {
	method := "Async.pool_patch.apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// PoolApply: Apply the selected patch to all hosts in the pool and return a map of host_ref -&gt; patch output
func (poolPatch) PoolApply(session *Session, self PoolPatchRef) (err error) {
	method := "pool_patch.pool_apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPoolApply: Apply the selected patch to all hosts in the pool and return a map of host_ref -&gt; patch output
func (poolPatch) AsyncPoolApply(session *Session, self PoolPatchRef) (retval TaskRef, err error) {
	method := "Async.pool_patch.pool_apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Precheck: Execute the precheck stage of the selected patch on a host and return its output
func (poolPatch) Precheck(session *Session, self PoolPatchRef, host HostRef) (retval string, err error) {
	method := "pool_patch.precheck"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncPrecheck: Execute the precheck stage of the selected patch on a host and return its output
func (poolPatch) AsyncPrecheck(session *Session, self PoolPatchRef, host HostRef) (retval TaskRef, err error) {
	method := "Async.pool_patch.precheck"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Clean: Removes the patch&apos;s files from the server
func (poolPatch) Clean(session *Session, self PoolPatchRef) (err error) {
	method := "pool_patch.clean"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncClean: Removes the patch&apos;s files from the server
func (poolPatch) AsyncClean(session *Session, self PoolPatchRef) (retval TaskRef, err error) {
	method := "Async.pool_patch.clean"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// PoolClean: Removes the patch&apos;s files from all hosts in the pool, but does not remove the database entries
func (poolPatch) PoolClean(session *Session, self PoolPatchRef) (err error) {
	method := "pool_patch.pool_clean"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPoolClean: Removes the patch&apos;s files from all hosts in the pool, but does not remove the database entries
func (poolPatch) AsyncPoolClean(session *Session, self PoolPatchRef) (retval TaskRef, err error) {
	method := "Async.pool_patch.pool_clean"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy: Removes the patch&apos;s files from all hosts in the pool, and removes the database entries.  Only works on unapplied patches.
func (poolPatch) Destroy(session *Session, self PoolPatchRef) (err error) {
	method := "pool_patch.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Removes the patch&apos;s files from all hosts in the pool, and removes the database entries.  Only works on unapplied patches.
func (poolPatch) AsyncDestroy(session *Session, self PoolPatchRef) (retval TaskRef, err error) {
	method := "Async.pool_patch.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// CleanOnHost: Removes the patch&apos;s files from the specified host
func (poolPatch) CleanOnHost(session *Session, self PoolPatchRef, host HostRef) (err error) {
	method := "pool_patch.clean_on_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, hostArg)
	return
}

// AsyncCleanOnHost: Removes the patch&apos;s files from the specified host
func (poolPatch) AsyncCleanOnHost(session *Session, self PoolPatchRef, host HostRef) (retval TaskRef, err error) {
	method := "Async.pool_patch.clean_on_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolPatchRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetAll: Return a list of all the pool_patchs known to the system.
func (poolPatch) GetAll(session *Session) (retval []PoolPatchRef, err error) {
	method := "pool_patch.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolPatchRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of pool_patch references to pool_patch records for all pool_patchs known to the system.
func (poolPatch) GetAllRecords(session *Session) (retval map[PoolPatchRef]PoolPatchRecord, err error) {
	method := "pool_patch.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolPatchRefToPoolPatchRecordMap(method+" -> ", result)
	return
}

