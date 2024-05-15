package xenapi

import (
	"fmt"
)

type DRTaskRecord struct {
	// Unique identifier/object reference
	UUID string
	// All SRs introduced by this appliance
	IntroducedSRs []SRRef
}

type DRTaskRef string

// DR task
type dRTask struct{}

var DRTask dRTask

// GetAllRecords: Return a map of DR_task references to DR_task records for all DR_tasks known to the system.
// Version: boston
func (dRTask) GetAllRecords(session *Session) (retval map[DRTaskRef]DRTaskRecord, err error) {
	method := "DR_task.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeDRTaskRefToDRTaskRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of DR_task references to DR_task records for all DR_tasks known to the system.
// Version: boston
func (dRTask) GetAllRecords1(session *Session) (retval map[DRTaskRef]DRTaskRecord, err error) {
	method := "DR_task.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeDRTaskRefToDRTaskRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the DR_tasks known to the system.
// Version: boston
func (dRTask) GetAll(session *Session) (retval []DRTaskRef, err error) {
	method := "DR_task.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeDRTaskRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the DR_tasks known to the system.
// Version: boston
func (dRTask) GetAll1(session *Session) (retval []DRTaskRef, err error) {
	method := "DR_task.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeDRTaskRefSet(method+" -> ", result)
	return
}

// Destroy: Destroy the disaster recovery task, detaching and forgetting any SRs introduced which are no longer required
// Version: boston
func (dRTask) Destroy(session *Session, self DRTaskRef) (err error) {
	method := "DR_task.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the disaster recovery task, detaching and forgetting any SRs introduced which are no longer required
// Version: boston
func (dRTask) AsyncDestroy(session *Session, self DRTaskRef) (retval TaskRef, err error) {
	method := "Async.DR_task.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy2: Destroy the disaster recovery task, detaching and forgetting any SRs introduced which are no longer required
// Version: boston
func (dRTask) Destroy2(session *Session, self DRTaskRef) (err error) {
	method := "DR_task.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy2: Destroy the disaster recovery task, detaching and forgetting any SRs introduced which are no longer required
// Version: boston
func (dRTask) AsyncDestroy2(session *Session, self DRTaskRef) (retval TaskRef, err error) {
	method := "Async.DR_task.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Create a disaster recovery task which will query the supplied list of devices
// Version: boston
func (dRTask) Create(session *Session, typeKey string, deviceConfig map[string]string, whitelist []string) (retval DRTaskRef, err error) {
	method := "DR_task.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	whitelistArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "whitelist"), whitelist)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, typeKeyArg, deviceConfigArg, whitelistArg)
	if err != nil {
		return
	}
	retval, err = deserializeDRTaskRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a disaster recovery task which will query the supplied list of devices
// Version: boston
func (dRTask) AsyncCreate(session *Session, typeKey string, deviceConfig map[string]string, whitelist []string) (retval TaskRef, err error) {
	method := "Async.DR_task.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	whitelistArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "whitelist"), whitelist)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, typeKeyArg, deviceConfigArg, whitelistArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Create4: Create a disaster recovery task which will query the supplied list of devices
// Version: boston
func (dRTask) Create4(session *Session, typeKey string, deviceConfig map[string]string, whitelist []string) (retval DRTaskRef, err error) {
	method := "DR_task.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	whitelistArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "whitelist"), whitelist)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, typeKeyArg, deviceConfigArg, whitelistArg)
	if err != nil {
		return
	}
	retval, err = deserializeDRTaskRef(method+" -> ", result)
	return
}

// AsyncCreate4: Create a disaster recovery task which will query the supplied list of devices
// Version: boston
func (dRTask) AsyncCreate4(session *Session, typeKey string, deviceConfig map[string]string, whitelist []string) (retval TaskRef, err error) {
	method := "Async.DR_task.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
	if err != nil {
		return
	}
	deviceConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "device_config"), deviceConfig)
	if err != nil {
		return
	}
	whitelistArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "whitelist"), whitelist)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, typeKeyArg, deviceConfigArg, whitelistArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetIntroducedSRs: Get the introduced_SRs field of the given DR_task.
// Version: boston
func (dRTask) GetIntroducedSRs(session *Session, self DRTaskRef) (retval []SRRef, err error) {
	method := "DR_task.get_introduced_SRs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRefSet(method+" -> ", result)
	return
}

// GetIntroducedSRs2: Get the introduced_SRs field of the given DR_task.
// Version: boston
func (dRTask) GetIntroducedSRs2(session *Session, self DRTaskRef) (retval []SRRef, err error) {
	method := "DR_task.get_introduced_SRs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given DR_task.
// Version: boston
func (dRTask) GetUUID(session *Session, self DRTaskRef) (retval string, err error) {
	method := "DR_task.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given DR_task.
// Version: boston
func (dRTask) GetUUID2(session *Session, self DRTaskRef) (retval string, err error) {
	method := "DR_task.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByUUID: Get a reference to the DR_task instance with the specified UUID.
// Version: boston
func (dRTask) GetByUUID(session *Session, uUID string) (retval DRTaskRef, err error) {
	method := "DR_task.get_by_uuid"
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
	retval, err = deserializeDRTaskRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the DR_task instance with the specified UUID.
// Version: boston
func (dRTask) GetByUUID2(session *Session, uUID string) (retval DRTaskRef, err error) {
	method := "DR_task.get_by_uuid"
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
	retval, err = deserializeDRTaskRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given DR_task.
// Version: boston
func (dRTask) GetRecord(session *Session, self DRTaskRef) (retval DRTaskRecord, err error) {
	method := "DR_task.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeDRTaskRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given DR_task.
// Version: boston
func (dRTask) GetRecord2(session *Session, self DRTaskRef) (retval DRTaskRecord, err error) {
	method := "DR_task.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeDRTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeDRTaskRecord(method+" -> ", result)
	return
}
