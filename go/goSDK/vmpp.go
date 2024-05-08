package xenapi

import (
	"fmt"
	"time"
)

type VMPPRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// enable or disable this policy
	IsPolicyEnabled bool
	// type of the backup sub-policy
	BackupType VmppBackupType
	// maximum number of backups that should be stored at any time
	BackupRetentionValue int
	// frequency of the backup schedule
	BackupFrequency VmppBackupFrequency
	// schedule of the backup containing &apos;hour&apos;, &apos;min&apos;, &apos;days&apos;. Date/time-related information is in Local Timezone
	BackupSchedule map[string]string
	// true if this protection policy&apos;s backup is running
	IsBackupRunning bool
	// time of the last backup
	BackupLastRunTime time.Time
	// type of the archive target config
	ArchiveTargetType VmppArchiveTargetType
	// configuration for the archive, including its &apos;location&apos;, &apos;username&apos;, &apos;password&apos;
	ArchiveTargetConfig map[string]string
	// frequency of the archive schedule
	ArchiveFrequency VmppArchiveFrequency
	// schedule of the archive containing &apos;hour&apos;, &apos;min&apos;, &apos;days&apos;. Date/time-related information is in Local Timezone
	ArchiveSchedule map[string]string
	// true if this protection policy&apos;s archive is running
	IsArchiveRunning bool
	// time of the last archive
	ArchiveLastRunTime time.Time
	// all VMs attached to this protection policy
	VMs []VMRef
	// true if alarm is enabled for this policy
	IsAlarmEnabled bool
	// configuration for the alarm
	AlarmConfig map[string]string
	// recent alerts
	RecentAlerts []string
}

type VMPPRef string

// VM Protection Policy
type vMPP struct{}

var VMPP vMPP

// GetRecord: Get a record containing the current state of the given VMPP.
func (vMPP) GetRecord(session *Session, self VMPPRef) (retval VMPPRecord, err error) {
	method := "VMPP.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMPPRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VMPP instance with the specified UUID.
func (vMPP) GetByUUID(session *Session, uUID string) (retval VMPPRef, err error) {
	method := "VMPP.get_by_uuid"
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
	retval, err = deserializeVMPPRef(method+" -> ", result)
	return
}

// Create: Create a new VMPP instance, and return its handle. The constructor args are: name_label, name_description, is_policy_enabled, backup_type, backup_retention_value, backup_frequency, backup_schedule, archive_target_type, archive_target_config, archive_frequency, archive_schedule, is_alarm_enabled, alarm_config (* = non-optional).
func (vMPP) Create(session *Session, args VMPPRecord) (retval VMPPRef, err error) {
	method := "VMPP.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVMPPRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMPPRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new VMPP instance, and return its handle. The constructor args are: name_label, name_description, is_policy_enabled, backup_type, backup_retention_value, backup_frequency, backup_schedule, archive_target_type, archive_target_config, archive_frequency, archive_schedule, is_alarm_enabled, alarm_config (* = non-optional).
func (vMPP) AsyncCreate(session *Session, args VMPPRecord) (retval TaskRef, err error) {
	method := "Async.VMPP.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVMPPRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// Destroy: Destroy the specified VMPP instance.
func (vMPP) Destroy(session *Session, self VMPPRef) (err error) {
	method := "VMPP.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified VMPP instance.
func (vMPP) AsyncDestroy(session *Session, self VMPPRef) (retval TaskRef, err error) {
	method := "Async.VMPP.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByNameLabel: Get all the VMPP instances with the given label.
func (vMPP) GetByNameLabel(session *Session, label string) (retval []VMPPRef, err error) {
	method := "VMPP.get_by_name_label"
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
	retval, err = deserializeVMPPRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VMPP.
func (vMPP) GetUUID(session *Session, self VMPPRef) (retval string, err error) {
	method := "VMPP.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given VMPP.
func (vMPP) GetNameLabel(session *Session, self VMPPRef) (retval string, err error) {
	method := "VMPP.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given VMPP.
func (vMPP) GetNameDescription(session *Session, self VMPPRef) (retval string, err error) {
	method := "VMPP.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsPolicyEnabled: Get the is_policy_enabled field of the given VMPP.
func (vMPP) GetIsPolicyEnabled(session *Session, self VMPPRef) (retval bool, err error) {
	method := "VMPP.get_is_policy_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBackupType: Get the backup_type field of the given VMPP.
func (vMPP) GetBackupType(session *Session, self VMPPRef) (retval VmppBackupType, err error) {
	method := "VMPP.get_backup_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVmppBackupType(method+" -> ", result)
	return
}

// GetBackupRetentionValue: Get the backup_retention_value field of the given VMPP.
func (vMPP) GetBackupRetentionValue(session *Session, self VMPPRef) (retval int, err error) {
	method := "VMPP.get_backup_retention_value"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBackupFrequency: Get the backup_frequency field of the given VMPP.
func (vMPP) GetBackupFrequency(session *Session, self VMPPRef) (retval VmppBackupFrequency, err error) {
	method := "VMPP.get_backup_frequency"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVmppBackupFrequency(method+" -> ", result)
	return
}

// GetBackupSchedule: Get the backup_schedule field of the given VMPP.
func (vMPP) GetBackupSchedule(session *Session, self VMPPRef) (retval map[string]string, err error) {
	method := "VMPP.get_backup_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsBackupRunning: Get the is_backup_running field of the given VMPP.
func (vMPP) GetIsBackupRunning(session *Session, self VMPPRef) (retval bool, err error) {
	method := "VMPP.get_is_backup_running"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBackupLastRunTime: Get the backup_last_run_time field of the given VMPP.
func (vMPP) GetBackupLastRunTime(session *Session, self VMPPRef) (retval time.Time, err error) {
	method := "VMPP.get_backup_last_run_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetArchiveTargetType: Get the archive_target_type field of the given VMPP.
func (vMPP) GetArchiveTargetType(session *Session, self VMPPRef) (retval VmppArchiveTargetType, err error) {
	method := "VMPP.get_archive_target_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVmppArchiveTargetType(method+" -> ", result)
	return
}

// GetArchiveTargetConfig: Get the archive_target_config field of the given VMPP.
func (vMPP) GetArchiveTargetConfig(session *Session, self VMPPRef) (retval map[string]string, err error) {
	method := "VMPP.get_archive_target_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetArchiveFrequency: Get the archive_frequency field of the given VMPP.
func (vMPP) GetArchiveFrequency(session *Session, self VMPPRef) (retval VmppArchiveFrequency, err error) {
	method := "VMPP.get_archive_frequency"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVmppArchiveFrequency(method+" -> ", result)
	return
}

// GetArchiveSchedule: Get the archive_schedule field of the given VMPP.
func (vMPP) GetArchiveSchedule(session *Session, self VMPPRef) (retval map[string]string, err error) {
	method := "VMPP.get_archive_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsArchiveRunning: Get the is_archive_running field of the given VMPP.
func (vMPP) GetIsArchiveRunning(session *Session, self VMPPRef) (retval bool, err error) {
	method := "VMPP.get_is_archive_running"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetArchiveLastRunTime: Get the archive_last_run_time field of the given VMPP.
func (vMPP) GetArchiveLastRunTime(session *Session, self VMPPRef) (retval time.Time, err error) {
	method := "VMPP.get_archive_last_run_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVMs: Get the VMs field of the given VMPP.
func (vMPP) GetVMs(session *Session, self VMPPRef) (retval []VMRef, err error) {
	method := "VMPP.get_VMs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsAlarmEnabled: Get the is_alarm_enabled field of the given VMPP.
func (vMPP) GetIsAlarmEnabled(session *Session, self VMPPRef) (retval bool, err error) {
	method := "VMPP.get_is_alarm_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAlarmConfig: Get the alarm_config field of the given VMPP.
func (vMPP) GetAlarmConfig(session *Session, self VMPPRef) (retval map[string]string, err error) {
	method := "VMPP.get_alarm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRecentAlerts: Get the recent_alerts field of the given VMPP.
func (vMPP) GetRecentAlerts(session *Session, self VMPPRef) (retval []string, err error) {
	method := "VMPP.get_recent_alerts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel: Set the name/label field of the given VMPP.
func (vMPP) SetNameLabel(session *Session, self VMPPRef, value string) (err error) {
	method := "VMPP.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription: Set the name/description field of the given VMPP.
func (vMPP) SetNameDescription(session *Session, self VMPPRef, value string) (err error) {
	method := "VMPP.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetIsPolicyEnabled: Set the is_policy_enabled field of the given VMPP.
func (vMPP) SetIsPolicyEnabled(session *Session, self VMPPRef, value bool) (err error) {
	method := "VMPP.set_is_policy_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetBackupType: Set the backup_type field of the given VMPP.
func (vMPP) SetBackupType(session *Session, self VMPPRef, value VmppBackupType) (err error) {
	method := "VMPP.set_backup_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVmppBackupType(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// ProtectNow: This call executes the protection policy immediately
func (vMPP) ProtectNow(session *Session, vmpp VMPPRef) (retval string, err error) {
	method := "VMPP.protect_now"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vmppArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "vmpp"), vmpp)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vmppArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// ArchiveNow: This call archives the snapshot provided as a parameter
func (vMPP) ArchiveNow(session *Session, snapshot VMRef) (retval string, err error) {
	method := "VMPP.archive_now"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	snapshotArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "snapshot"), snapshot)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, snapshotArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// GetAlerts: This call fetches a history of alerts for a given protection policy
func (vMPP) GetAlerts(session *Session, vmpp VMPPRef, hoursFromNow int) (retval []string, err error) {
	method := "VMPP.get_alerts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vmppArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "vmpp"), vmpp)
	if err != nil {
		return
	}
	hoursFromNowArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "hours_from_now"), hoursFromNow)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vmppArg, hoursFromNowArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// SetBackupRetentionValue:
func (vMPP) SetBackupRetentionValue(session *Session, self VMPPRef, value int) (err error) {
	method := "VMPP.set_backup_retention_value"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetBackupFrequency: Set the value of the backup_frequency field
func (vMPP) SetBackupFrequency(session *Session, self VMPPRef, value VmppBackupFrequency) (err error) {
	method := "VMPP.set_backup_frequency"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVmppBackupFrequency(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetBackupSchedule:
func (vMPP) SetBackupSchedule(session *Session, self VMPPRef, value map[string]string) (err error) {
	method := "VMPP.set_backup_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetArchiveFrequency: Set the value of the archive_frequency field
func (vMPP) SetArchiveFrequency(session *Session, self VMPPRef, value VmppArchiveFrequency) (err error) {
	method := "VMPP.set_archive_frequency"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVmppArchiveFrequency(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetArchiveSchedule:
func (vMPP) SetArchiveSchedule(session *Session, self VMPPRef, value map[string]string) (err error) {
	method := "VMPP.set_archive_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetArchiveTargetType: Set the value of the archive_target_config_type field
func (vMPP) SetArchiveTargetType(session *Session, self VMPPRef, value VmppArchiveTargetType) (err error) {
	method := "VMPP.set_archive_target_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVmppArchiveTargetType(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetArchiveTargetConfig:
func (vMPP) SetArchiveTargetConfig(session *Session, self VMPPRef, value map[string]string) (err error) {
	method := "VMPP.set_archive_target_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetIsAlarmEnabled: Set the value of the is_alarm_enabled field
func (vMPP) SetIsAlarmEnabled(session *Session, self VMPPRef, value bool) (err error) {
	method := "VMPP.set_is_alarm_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetAlarmConfig:
func (vMPP) SetAlarmConfig(session *Session, self VMPPRef, value map[string]string) (err error) {
	method := "VMPP.set_alarm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToBackupSchedule:
func (vMPP) AddToBackupSchedule(session *Session, self VMPPRef, key string, value string) (err error) {
	method := "VMPP.add_to_backup_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToArchiveTargetConfig:
func (vMPP) AddToArchiveTargetConfig(session *Session, self VMPPRef, key string, value string) (err error) {
	method := "VMPP.add_to_archive_target_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToArchiveSchedule:
func (vMPP) AddToArchiveSchedule(session *Session, self VMPPRef, key string, value string) (err error) {
	method := "VMPP.add_to_archive_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToAlarmConfig:
func (vMPP) AddToAlarmConfig(session *Session, self VMPPRef, key string, value string) (err error) {
	method := "VMPP.add_to_alarm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromBackupSchedule:
func (vMPP) RemoveFromBackupSchedule(session *Session, self VMPPRef, key string) (err error) {
	method := "VMPP.remove_from_backup_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromArchiveTargetConfig:
func (vMPP) RemoveFromArchiveTargetConfig(session *Session, self VMPPRef, key string) (err error) {
	method := "VMPP.remove_from_archive_target_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromArchiveSchedule:
func (vMPP) RemoveFromArchiveSchedule(session *Session, self VMPPRef, key string) (err error) {
	method := "VMPP.remove_from_archive_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromAlarmConfig:
func (vMPP) RemoveFromAlarmConfig(session *Session, self VMPPRef, key string) (err error) {
	method := "VMPP.remove_from_alarm_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetBackupLastRunTime:
func (vMPP) SetBackupLastRunTime(session *Session, self VMPPRef, value time.Time) (err error) {
	method := "VMPP.set_backup_last_run_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetArchiveLastRunTime:
func (vMPP) SetArchiveLastRunTime(session *Session, self VMPPRef, value time.Time) (err error) {
	method := "VMPP.set_archive_last_run_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the VMPPs known to the system.
func (vMPP) GetAll(session *Session) (retval []VMPPRef, err error) {
	method := "VMPP.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMPPRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VMPP references to VMPP records for all VMPPs known to the system.
func (vMPP) GetAllRecords(session *Session) (retval map[VMPPRef]VMPPRecord, err error) {
	method := "VMPP.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMPPRefToVMPPRecordMap(method+" -> ", result)
	return
}
