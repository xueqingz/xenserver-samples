package xenapi

import (
	"fmt"
)

type SubjectRecord struct {
	// Unique identifier/object reference
	UUID string
	// the subject identifier, unique in the external directory service
	SubjectIdentifier string
	// additional configuration
	OtherConfig map[string]string
	// the roles associated with this subject
	Roles []RoleRef
}

type SubjectRef string

// A user or group that can log in xapi
type subject struct{}

var Subject subject

// GetRecord: Get a record containing the current state of the given subject.
func (subject) GetRecord(session *Session, self SubjectRef) (retval SubjectRecord, err error) {
	method := "subject.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSubjectRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the subject instance with the specified UUID.
func (subject) GetByUUID(session *Session, uUID string) (retval SubjectRef, err error) {
	method := "subject.get_by_uuid"
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
	retval, err = deserializeSubjectRef(method+" -> ", result)
	return
}

// Create: Create a new subject instance, and return its handle. The constructor args are: subject_identifier, other_config (* = non-optional).
func (subject) Create(session *Session, args SubjectRecord) (retval SubjectRef, err error) {
	method := "subject.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeSubjectRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeSubjectRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new subject instance, and return its handle. The constructor args are: subject_identifier, other_config (* = non-optional).
func (subject) AsyncCreate(session *Session, args SubjectRecord) (retval TaskRef, err error) {
	method := "Async.subject.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeSubjectRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// Destroy: Destroy the specified subject instance.
func (subject) Destroy(session *Session, self SubjectRef) (err error) {
	method := "subject.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified subject instance.
func (subject) AsyncDestroy(session *Session, self SubjectRef) (retval TaskRef, err error) {
	method := "Async.subject.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given subject.
func (subject) GetUUID(session *Session, self SubjectRef) (retval string, err error) {
	method := "subject.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSubjectIdentifier: Get the subject_identifier field of the given subject.
func (subject) GetSubjectIdentifier(session *Session, self SubjectRef) (retval string, err error) {
	method := "subject.get_subject_identifier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given subject.
func (subject) GetOtherConfig(session *Session, self SubjectRef) (retval map[string]string, err error) {
	method := "subject.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRoles: Get the roles field of the given subject.
func (subject) GetRoles(session *Session, self SubjectRef) (retval []RoleRef, err error) {
	method := "subject.get_roles"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeRoleRefSet(method+" -> ", result)
	return
}

// AddToRoles: This call adds a new role to a subject
func (subject) AddToRoles(session *Session, self SubjectRef, role RoleRef) (err error) {
	method := "subject.add_to_roles"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	roleArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "role"), role)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, roleArg)
	return
}

// RemoveFromRoles: This call removes a role from a subject
func (subject) RemoveFromRoles(session *Session, self SubjectRef, role RoleRef) (err error) {
	method := "subject.remove_from_roles"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	roleArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "role"), role)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, roleArg)
	return
}

// GetPermissionsNameLabel: This call returns a list of permission names given a subject
func (subject) GetPermissionsNameLabel(session *Session, self SubjectRef) (retval []string, err error) {
	method := "subject.get_permissions_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSubjectRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the subjects known to the system.
func (subject) GetAll(session *Session) (retval []SubjectRef, err error) {
	method := "subject.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSubjectRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of subject references to subject records for all subjects known to the system.
func (subject) GetAllRecords(session *Session) (retval map[SubjectRef]SubjectRecord, err error) {
	method := "subject.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSubjectRefToSubjectRecordMap(method+" -> ", result)
	return
}
