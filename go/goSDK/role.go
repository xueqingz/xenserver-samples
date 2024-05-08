package xenapi

import (
	"fmt"
)

type RoleRecord struct {
	// Unique identifier/object reference
	UUID string
	// a short user-friendly name for the role
	NameLabel string
	// what this role is for
	NameDescription string
	// a list of pointers to other roles or permissions
	Subroles []RoleRef
	// Indicates whether the role is only to be assigned internally by xapi, or can be used by clients
	IsInternal bool
}

type RoleRef string

// A set of permissions associated with a subject
type role struct{}

var Role role

// GetRecord: Get a record containing the current state of the given role.
func (role) GetRecord(session *Session, self RoleRef) (retval RoleRecord, err error) {
	method := "role.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeRoleRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the role instance with the specified UUID.
func (role) GetByUUID(session *Session, uUID string) (retval RoleRef, err error) {
	method := "role.get_by_uuid"
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
	retval, err = deserializeRoleRef(method+" -> ", result)
	return
}

// GetByNameLabel: Get all the role instances with the given label.
func (role) GetByNameLabel(session *Session, label string) (retval []RoleRef, err error) {
	method := "role.get_by_name_label"
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
	retval, err = deserializeRoleRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given role.
func (role) GetUUID(session *Session, self RoleRef) (retval string, err error) {
	method := "role.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given role.
func (role) GetNameLabel(session *Session, self RoleRef) (retval string, err error) {
	method := "role.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given role.
func (role) GetNameDescription(session *Session, self RoleRef) (retval string, err error) {
	method := "role.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSubroles: Get the subroles field of the given role.
func (role) GetSubroles(session *Session, self RoleRef) (retval []RoleRef, err error) {
	method := "role.get_subroles"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsInternal: Get the is_internal field of the given role.
func (role) GetIsInternal(session *Session, self RoleRef) (retval bool, err error) {
	method := "role.get_is_internal"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPermissions: This call returns a list of permissions given a role
func (role) GetPermissions(session *Session, self RoleRef) (retval []RoleRef, err error) {
	method := "role.get_permissions"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPermissionsNameLabel: This call returns a list of permission names given a role
func (role) GetPermissionsNameLabel(session *Session, self RoleRef) (retval []string, err error) {
	method := "role.get_permissions_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByPermission: This call returns a list of roles given a permission
func (role) GetByPermission(session *Session, permission RoleRef) (retval []RoleRef, err error) {
	method := "role.get_by_permission"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	permissionArg, err := serializeRoleRef(fmt.Sprintf("%s(%s)", method, "permission"), permission)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, permissionArg)
	if err != nil {
		return
	}
	retval, err = deserializeRoleRefSet(method+" -> ", result)
	return
}

// GetByPermissionNameLabel: This call returns a list of roles given a permission name
func (role) GetByPermissionNameLabel(session *Session, label string) (retval []RoleRef, err error) {
	method := "role.get_by_permission_name_label"
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
	retval, err = deserializeRoleRefSet(method+" -> ", result)
	return
}

// GetAll: Return a list of all the roles known to the system.
func (role) GetAll(session *Session) (retval []RoleRef, err error) {
	method := "role.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeRoleRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of role references to role records for all roles known to the system.
func (role) GetAllRecords(session *Session) (retval map[RoleRef]RoleRecord, err error) {
	method := "role.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeRoleRefToRoleRecordMap(method+" -> ", result)
	return
}
