package xenapi

import (
	"fmt"
)

type PoolUpdateRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// Update version number
	Version string
	// Size of the update in bytes
	InstallationSize int
	// GPG key of the update
	Key string
	// What the client should do after this update has been applied.
	AfterApplyGuidance []UpdateAfterApplyGuidance
	// VDI the update was uploaded to
	Vdi VDIRef
	// The hosts that have applied this update.
	Hosts []HostRef
	// additional configuration
	OtherConfig map[string]string
	// Flag - if true, all hosts in a pool must apply this update
	EnforceHomogeneity bool
}

type PoolUpdateRef string

// Pool-wide updates to the host software
type poolUpdate struct{}

var PoolUpdate poolUpdate

// GetRecord: Get a record containing the current state of the given pool_update.
func (poolUpdate) GetRecord(session *Session, self PoolUpdateRef) (retval PoolUpdateRecord, err error) {
	method := "pool_update.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolUpdateRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the pool_update instance with the specified UUID.
func (poolUpdate) GetByUUID(session *Session, uUID string) (retval PoolUpdateRef, err error) {
	method := "pool_update.get_by_uuid"
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
	retval, err = deserializePoolUpdateRef(method+" -> ", result)
	return
}

// GetByNameLabel: Get all the pool_update instances with the given label.
func (poolUpdate) GetByNameLabel(session *Session, label string) (retval []PoolUpdateRef, err error) {
	method := "pool_update.get_by_name_label"
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
	retval, err = deserializePoolUpdateRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given pool_update.
func (poolUpdate) GetUUID(session *Session, self PoolUpdateRef) (retval string, err error) {
	method := "pool_update.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given pool_update.
func (poolUpdate) GetNameLabel(session *Session, self PoolUpdateRef) (retval string, err error) {
	method := "pool_update.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given pool_update.
func (poolUpdate) GetNameDescription(session *Session, self PoolUpdateRef) (retval string, err error) {
	method := "pool_update.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVersion: Get the version field of the given pool_update.
func (poolUpdate) GetVersion(session *Session, self PoolUpdateRef) (retval string, err error) {
	method := "pool_update.get_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetInstallationSize: Get the installation_size field of the given pool_update.
func (poolUpdate) GetInstallationSize(session *Session, self PoolUpdateRef) (retval int, err error) {
	method := "pool_update.get_installation_size"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetKey: Get the key field of the given pool_update.
func (poolUpdate) GetKey(session *Session, self PoolUpdateRef) (retval string, err error) {
	method := "pool_update.get_key"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAfterApplyGuidance: Get the after_apply_guidance field of the given pool_update.
func (poolUpdate) GetAfterApplyGuidance(session *Session, self PoolUpdateRef) (retval []UpdateAfterApplyGuidance, err error) {
	method := "pool_update.get_after_apply_guidance"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumUpdateAfterApplyGuidanceSet(method+" -> ", result)
	return
}

// GetVdi: Get the vdi field of the given pool_update.
func (poolUpdate) GetVdi(session *Session, self PoolUpdateRef) (retval VDIRef, err error) {
	method := "pool_update.get_vdi"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHosts: Get the hosts field of the given pool_update.
func (poolUpdate) GetHosts(session *Session, self PoolUpdateRef) (retval []HostRef, err error) {
	method := "pool_update.get_hosts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRefSet(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given pool_update.
func (poolUpdate) GetOtherConfig(session *Session, self PoolUpdateRef) (retval map[string]string, err error) {
	method := "pool_update.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnforceHomogeneity: Get the enforce_homogeneity field of the given pool_update.
func (poolUpdate) GetEnforceHomogeneity(session *Session, self PoolUpdateRef) (retval bool, err error) {
	method := "pool_update.get_enforce_homogeneity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given pool_update.
func (poolUpdate) SetOtherConfig(session *Session, self PoolUpdateRef, value map[string]string) (err error) {
	method := "pool_update.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given pool_update.
func (poolUpdate) AddToOtherConfig(session *Session, self PoolUpdateRef, key string, value string) (err error) {
	method := "pool_update.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given pool_update.  If the key is not in that Map, then do nothing.
func (poolUpdate) RemoveFromOtherConfig(session *Session, self PoolUpdateRef, key string) (err error) {
	method := "pool_update.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Introduce: Introduce update VDI
func (poolUpdate) Introduce(session *Session, vdi VDIRef) (retval PoolUpdateRef, err error) {
	method := "pool_update.introduce"
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
	retval, err = deserializePoolUpdateRef(method+" -> ", result)
	return
}

// AsyncIntroduce: Introduce update VDI
func (poolUpdate) AsyncIntroduce(session *Session, vdi VDIRef) (retval TaskRef, err error) {
	method := "Async.pool_update.introduce"
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

// Precheck: Execute the precheck stage of the selected update on a host
func (poolUpdate) Precheck(session *Session, self PoolUpdateRef, host HostRef) (retval LivepatchStatus, err error) {
	method := "pool_update.precheck"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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
	retval, err = deserializeEnumLivepatchStatus(method+" -> ", result)
	return
}

// AsyncPrecheck: Execute the precheck stage of the selected update on a host
func (poolUpdate) AsyncPrecheck(session *Session, self PoolUpdateRef, host HostRef) (retval TaskRef, err error) {
	method := "Async.pool_update.precheck"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Apply: Apply the selected update to a host
func (poolUpdate) Apply(session *Session, self PoolUpdateRef, host HostRef) (err error) {
	method := "pool_update.apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncApply: Apply the selected update to a host
func (poolUpdate) AsyncApply(session *Session, self PoolUpdateRef, host HostRef) (retval TaskRef, err error) {
	method := "Async.pool_update.apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// PoolApply: Apply the selected update to all hosts in the pool
func (poolUpdate) PoolApply(session *Session, self PoolUpdateRef) (err error) {
	method := "pool_update.pool_apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPoolApply: Apply the selected update to all hosts in the pool
func (poolUpdate) AsyncPoolApply(session *Session, self PoolUpdateRef) (retval TaskRef, err error) {
	method := "Async.pool_update.pool_apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// PoolClean: Removes the update&apos;s files from all hosts in the pool, but does not revert the update
func (poolUpdate) PoolClean(session *Session, self PoolUpdateRef) (err error) {
	method := "pool_update.pool_clean"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPoolClean: Removes the update&apos;s files from all hosts in the pool, but does not revert the update
func (poolUpdate) AsyncPoolClean(session *Session, self PoolUpdateRef) (retval TaskRef, err error) {
	method := "Async.pool_update.pool_clean"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy: Removes the database entry. Only works on unapplied update.
func (poolUpdate) Destroy(session *Session, self PoolUpdateRef) (err error) {
	method := "pool_update.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Removes the database entry. Only works on unapplied update.
func (poolUpdate) AsyncDestroy(session *Session, self PoolUpdateRef) (retval TaskRef, err error) {
	method := "Async.pool_update.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolUpdateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the pool_updates known to the system.
func (poolUpdate) GetAll(session *Session) (retval []PoolUpdateRef, err error) {
	method := "pool_update.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolUpdateRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of pool_update references to pool_update records for all pool_updates known to the system.
func (poolUpdate) GetAllRecords(session *Session) (retval map[PoolUpdateRef]PoolUpdateRecord, err error) {
	method := "pool_update.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolUpdateRefToPoolUpdateRecordMap(method+" -> ", result)
	return
}
