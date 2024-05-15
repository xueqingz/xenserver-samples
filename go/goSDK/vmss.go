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

// GetAllRecords: Return a map of VMSS references to VMSS records for all VMSSs known to the system.
// Version: falcon
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

// GetAllRecords1: Return a map of VMSS references to VMSS records for all VMSSs known to the system.
// Version: falcon
func (vMSS) GetAllRecords1(session *Session) (retval map[VMSSRef]VMSSRecord, err error) {
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

// GetAll: Return a list of all the VMSSs known to the system.
// Version: falcon
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

// GetAll1: Return a list of all the VMSSs known to the system.
// Version: falcon
func (vMSS) GetAll1(session *Session) (retval []VMSSRef, err error) {
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

// SetType:
// Version: falcon
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

// SetType3:
// Version: falcon
func (vMSS) SetType3(session *Session, self VMSSRef, value VmssType) (err error) {
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

// SetLastRunTime:
// Version: falcon
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

// SetLastRunTime3:
// Version: falcon
func (vMSS) SetLastRunTime3(session *Session, self VMSSRef, value time.Time) (err error) {
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

// RemoveFromSchedule:
// Version: falcon
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

// RemoveFromSchedule3:
// Version: falcon
func (vMSS) RemoveFromSchedule3(session *Session, self VMSSRef, key string) (err error) {
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

// AddToSchedule:
// Version: falcon
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

// AddToSchedule4:
// Version: falcon
func (vMSS) AddToSchedule4(session *Session, self VMSSRef, key string, value string) (err error) {
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

// SetSchedule:
// Version: falcon
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

// SetSchedule3:
// Version: falcon
func (vMSS) SetSchedule3(session *Session, self VMSSRef, value map[string]string) (err error) {
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

// SetFrequency: Set the value of the frequency field
// Version: falcon
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

// SetFrequency3: Set the value of the frequency field
// Version: falcon
func (vMSS) SetFrequency3(session *Session, self VMSSRef, value VmssFrequency) (err error) {
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

// SetRetainedSnapshots:
// Version: falcon
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

// SetRetainedSnapshots3:
// Version: falcon
func (vMSS) SetRetainedSnapshots3(session *Session, self VMSSRef, value int) (err error) {
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

// SnapshotNow: This call executes the snapshot schedule immediately
// Version: falcon
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

// SnapshotNow2: This call executes the snapshot schedule immediately
// Version: falcon
func (vMSS) SnapshotNow2(session *Session, vmss VMSSRef) (retval string, err error) {
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

// SetEnabled: Set the enabled field of the given VMSS.
// Version: falcon
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

// SetEnabled3: Set the enabled field of the given VMSS.
// Version: falcon
func (vMSS) SetEnabled3(session *Session, self VMSSRef, value bool) (err error) {
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

// SetEnabled2: Set the enabled field of the given VMSS.
// Version: rio
func (vMSS) SetEnabled2(session *Session, value bool) (err error) {
	method := "VMSS.set_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, valueArg)
	return
}

// SetNameDescription: Set the name/description field of the given VMSS.
// Version: falcon
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

// SetNameDescription3: Set the name/description field of the given VMSS.
// Version: falcon
func (vMSS) SetNameDescription3(session *Session, self VMSSRef, value string) (err error) {
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

// SetNameDescription2: Set the name/description field of the given VMSS.
// Version: rio
func (vMSS) SetNameDescription2(session *Session, value string) (err error) {
	method := "VMSS.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, valueArg)
	return
}

// SetNameLabel: Set the name/label field of the given VMSS.
// Version: falcon
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

// SetNameLabel3: Set the name/label field of the given VMSS.
// Version: falcon
func (vMSS) SetNameLabel3(session *Session, self VMSSRef, value string) (err error) {
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

// SetNameLabel2: Set the name/label field of the given VMSS.
// Version: rio
func (vMSS) SetNameLabel2(session *Session, value string) (err error) {
	method := "VMSS.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, valueArg)
	return
}

// GetVMs: Get the VMs field of the given VMSS.
// Version: falcon
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

// GetVMs2: Get the VMs field of the given VMSS.
// Version: falcon
func (vMSS) GetVMs2(session *Session, self VMSSRef) (retval []VMRef, err error) {
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

// GetLastRunTime: Get the last_run_time field of the given VMSS.
// Version: falcon
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

// GetLastRunTime2: Get the last_run_time field of the given VMSS.
// Version: falcon
func (vMSS) GetLastRunTime2(session *Session, self VMSSRef) (retval time.Time, err error) {
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

// GetSchedule: Get the schedule field of the given VMSS.
// Version: falcon
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

// GetSchedule2: Get the schedule field of the given VMSS.
// Version: falcon
func (vMSS) GetSchedule2(session *Session, self VMSSRef) (retval map[string]string, err error) {
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

// GetFrequency: Get the frequency field of the given VMSS.
// Version: falcon
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

// GetFrequency2: Get the frequency field of the given VMSS.
// Version: falcon
func (vMSS) GetFrequency2(session *Session, self VMSSRef) (retval VmssFrequency, err error) {
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

// GetRetainedSnapshots: Get the retained_snapshots field of the given VMSS.
// Version: falcon
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

// GetRetainedSnapshots2: Get the retained_snapshots field of the given VMSS.
// Version: falcon
func (vMSS) GetRetainedSnapshots2(session *Session, self VMSSRef) (retval int, err error) {
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

// GetType: Get the type field of the given VMSS.
// Version: falcon
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

// GetType2: Get the type field of the given VMSS.
// Version: falcon
func (vMSS) GetType2(session *Session, self VMSSRef) (retval VmssType, err error) {
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

// GetEnabled: Get the enabled field of the given VMSS.
// Version: falcon
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

// GetEnabled2: Get the enabled field of the given VMSS.
// Version: falcon
func (vMSS) GetEnabled2(session *Session, self VMSSRef) (retval bool, err error) {
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

// GetNameDescription: Get the name/description field of the given VMSS.
// Version: falcon
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

// GetNameDescription2: Get the name/description field of the given VMSS.
// Version: falcon
func (vMSS) GetNameDescription2(session *Session, self VMSSRef) (retval string, err error) {
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

// GetNameLabel: Get the name/label field of the given VMSS.
// Version: falcon
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

// GetNameLabel2: Get the name/label field of the given VMSS.
// Version: falcon
func (vMSS) GetNameLabel2(session *Session, self VMSSRef) (retval string, err error) {
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

// GetUUID: Get the uuid field of the given VMSS.
// Version: falcon
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

// GetUUID2: Get the uuid field of the given VMSS.
// Version: falcon
func (vMSS) GetUUID2(session *Session, self VMSSRef) (retval string, err error) {
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

// GetByNameLabel: Get all the VMSS instances with the given label.
// Version: falcon
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

// GetByNameLabel2: Get all the VMSS instances with the given label.
// Version: falcon
func (vMSS) GetByNameLabel2(session *Session, label string) (retval []VMSSRef, err error) {
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

// Destroy: Destroy the specified VMSS instance.
// Version: falcon
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
// Version: falcon
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

// Destroy2: Destroy the specified VMSS instance.
// Version: falcon
func (vMSS) Destroy2(session *Session, self VMSSRef) (err error) {
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

// AsyncDestroy2: Destroy the specified VMSS instance.
// Version: falcon
func (vMSS) AsyncDestroy2(session *Session, self VMSSRef) (retval TaskRef, err error) {
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

// Create: Create a new VMSS instance, and return its handle. The constructor args are: name_label, name_description, enabled, type*, retained_snapshots, frequency*, schedule (* = non-optional).
// Version: falcon
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
// Version: falcon
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

// Create2: Create a new VMSS instance, and return its handle. The constructor args are: name_label, name_description, enabled, type*, retained_snapshots, frequency*, schedule (* = non-optional).
// Version: falcon
func (vMSS) Create2(session *Session, args VMSSRecord) (retval VMSSRef, err error) {
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

// AsyncCreate2: Create a new VMSS instance, and return its handle. The constructor args are: name_label, name_description, enabled, type*, retained_snapshots, frequency*, schedule (* = non-optional).
// Version: falcon
func (vMSS) AsyncCreate2(session *Session, args VMSSRecord) (retval TaskRef, err error) {
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

// GetByUUID: Get a reference to the VMSS instance with the specified UUID.
// Version: falcon
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

// GetByUUID2: Get a reference to the VMSS instance with the specified UUID.
// Version: falcon
func (vMSS) GetByUUID2(session *Session, uUID string) (retval VMSSRef, err error) {
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

// GetRecord: Get a record containing the current state of the given VMSS.
// Version: falcon
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

// GetRecord2: Get a record containing the current state of the given VMSS.
// Version: falcon
func (vMSS) GetRecord2(session *Session, self VMSSRef) (retval VMSSRecord, err error) {
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
