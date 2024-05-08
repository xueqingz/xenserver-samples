package xenapi

import (
	"fmt"
	"time"
)

type VMSSRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// enable or disable this snapshot schedule
	Enabled bool
	// type of the snapshot schedule
	Type VmssType
	// maximum number of snapshots that should be stored at any time
	RetainedSnapshots int
	// frequency of taking snapshot from snapshot schedule
	Frequency VmssFrequency
	// schedule of the snapshot containing &apos;hour&apos;, &apos;min&apos;, &apos;days&apos;. Date/time-related information is in Local Timezone
	Schedule map[string]string
	// time of the last snapshot
	LastRunTime time.Time
	// all VMs attached to this snapshot schedule
	VMs []VMRef
}

type VMSSRef string

// VM Snapshot Schedule
type vMSS struct{}

var VMSS vMSS

// GetRecord: Get a record containing the current state of the given VMSS.
func (vMSS) GetRecord(session *Session, self VMSSRef) (retval VMSSRecord, err error) {
	method := "VMSS.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMSSRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VMSS instance with the specified UUID.
func (vMSS) GetByUUID(session *Session, uUID string) (retval VMSSRef, err error) {
	method := "VMSS.get_by_uuid"
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
	retval, err = deserializeVMSSRef(method+" -> ", result)
	return
}

// Create: Create a new VMSS instance, and return its handle. The constructor args are: name_label, name_description, enabled, type*, retained_snapshots, frequency*, schedule (* = non-optional).
func (vMSS) Create(session *Session, args VMSSRecord) (retval VMSSRef, err error) {
	method := "VMSS.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVMSSRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMSSRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new VMSS instance, and return its handle. The constructor args are: name_label, name_description, enabled, type*, retained_snapshots, frequency*, schedule (* = non-optional).
func (vMSS) AsyncCreate(session *Session, args VMSSRecord) (retval TaskRef, err error) {
	method := "Async.VMSS.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVMSSRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// Destroy: Destroy the specified VMSS instance.
func (vMSS) Destroy(session *Session, self VMSSRef) (err error) {
	method := "VMSS.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified VMSS instance.
func (vMSS) AsyncDestroy(session *Session, self VMSSRef) (retval TaskRef, err error) {
	method := "Async.VMSS.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByNameLabel: Get all the VMSS instances with the given label.
func (vMSS) GetByNameLabel(session *Session, label string) (retval []VMSSRef, err error) {
	method := "VMSS.get_by_name_label"
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
	retval, err = deserializeVMSSRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VMSS.
func (vMSS) GetUUID(session *Session, self VMSSRef) (retval string, err error) {
	method := "VMSS.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given VMSS.
func (vMSS) GetNameLabel(session *Session, self VMSSRef) (retval string, err error) {
	method := "VMSS.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given VMSS.
func (vMSS) GetNameDescription(session *Session, self VMSSRef) (retval string, err error) {
	method := "VMSS.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnabled: Get the enabled field of the given VMSS.
func (vMSS) GetEnabled(session *Session, self VMSSRef) (retval bool, err error) {
	method := "VMSS.get_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetType: Get the type field of the given VMSS.
func (vMSS) GetType(session *Session, self VMSSRef) (retval VmssType, err error) {
	method := "VMSS.get_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVmssType(method+" -> ", result)
	return
}

// GetRetainedSnapshots: Get the retained_snapshots field of the given VMSS.
func (vMSS) GetRetainedSnapshots(session *Session, self VMSSRef) (retval int, err error) {
	method := "VMSS.get_retained_snapshots"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetFrequency: Get the frequency field of the given VMSS.
func (vMSS) GetFrequency(session *Session, self VMSSRef) (retval VmssFrequency, err error) {
	method := "VMSS.get_frequency"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVmssFrequency(method+" -> ", result)
	return
}

// GetSchedule: Get the schedule field of the given VMSS.
func (vMSS) GetSchedule(session *Session, self VMSSRef) (retval map[string]string, err error) {
	method := "VMSS.get_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastRunTime: Get the last_run_time field of the given VMSS.
func (vMSS) GetLastRunTime(session *Session, self VMSSRef) (retval time.Time, err error) {
	method := "VMSS.get_last_run_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVMs: Get the VMs field of the given VMSS.
func (vMSS) GetVMs(session *Session, self VMSSRef) (retval []VMRef, err error) {
	method := "VMSS.get_VMs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefSet(method+" -> ", result)
	return
}

// SetNameLabel: Set the name/label field of the given VMSS.
func (vMSS) SetNameLabel(session *Session, self VMSSRef, value string) (err error) {
	method := "VMSS.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription: Set the name/description field of the given VMSS.
func (vMSS) SetNameDescription(session *Session, self VMSSRef, value string) (err error) {
	method := "VMSS.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetEnabled: Set the enabled field of the given VMSS.
func (vMSS) SetEnabled(session *Session, self VMSSRef, value bool) (err error) {
	method := "VMSS.set_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SnapshotNow: This call executes the snapshot schedule immediately
func (vMSS) SnapshotNow(session *Session, vmss VMSSRef) (retval string, err error) {
	method := "VMSS.snapshot_now"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vmssArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "vmss"), vmss)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vmssArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// SetRetainedSnapshots:
func (vMSS) SetRetainedSnapshots(session *Session, self VMSSRef, value int) (err error) {
	method := "VMSS.set_retained_snapshots"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetFrequency: Set the value of the frequency field
func (vMSS) SetFrequency(session *Session, self VMSSRef, value VmssFrequency) (err error) {
	method := "VMSS.set_frequency"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVmssFrequency(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetSchedule:
func (vMSS) SetSchedule(session *Session, self VMSSRef, value map[string]string) (err error) {
	method := "VMSS.set_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToSchedule:
func (vMSS) AddToSchedule(session *Session, self VMSSRef, key string, value string) (err error) {
	method := "VMSS.add_to_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromSchedule:
func (vMSS) RemoveFromSchedule(session *Session, self VMSSRef, key string) (err error) {
	method := "VMSS.remove_from_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetLastRunTime:
func (vMSS) SetLastRunTime(session *Session, self VMSSRef, value time.Time) (err error) {
	method := "VMSS.set_last_run_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeTime(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetType:
func (vMSS) SetType(session *Session, self VMSSRef, value VmssType) (err error) {
	method := "VMSS.set_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVmssType(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// GetAll: Return a list of all the VMSSs known to the system.
func (vMSS) GetAll(session *Session) (retval []VMSSRef, err error) {
	method := "VMSS.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMSSRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VMSS references to VMSS records for all VMSSs known to the system.
func (vMSS) GetAllRecords(session *Session) (retval map[VMSSRef]VMSSRecord, err error) {
	method := "VMSS.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMSSRefToVMSSRecordMap(method+" -> ", result)
	return
}

