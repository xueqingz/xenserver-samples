package xenapi

import (
	"fmt"
	"time"
)

type SessionRecord struct {
	// Unique identifier/object reference
	UUID string
	// Currently connected host
	ThisHost HostRef
	// Currently connected user
	ThisUser UserRef
	// Timestamp for last time session was active
	LastActive time.Time
	// True if this session relates to a intra-pool login, false otherwise
	Pool bool
	// additional configuration
	OtherConfig map[string]string
	// true iff this session was created using local superuser credentials
	IsLocalSuperuser bool
	// references the subject instance that created the session. If a session instance has is_local_superuser set, then the value of this field is undefined.
	Subject SubjectRef
	// time when session was last validated
	ValidationTime time.Time
	// the subject identifier of the user that was externally authenticated. If a session instance has is_local_superuser set, then the value of this field is undefined.
	AuthUserSid string
	// the subject name of the user that was externally authenticated. If a session instance has is_local_superuser set, then the value of this field is undefined.
	AuthUserName string
	// list with all RBAC permissions for this session
	RbacPermissions []string
	// list of tasks created using the current session
	Tasks []TaskRef
	// references the parent session that created this session
	Parent SessionRef
	// a key string provided by a API user to distinguish itself from other users sharing the same login name
	Originator string
	// indicates whether this session was authenticated using a client certificate
	ClientCertificate bool
}

type SessionRef string

// A session
type Session struct {
	APIVersion  APIVersion
	client      *rpcClient
	ref         SessionRef
	XAPIVersion string
}

func NewSession(opts *ClientOpts) *Session {
	client := newJSONRPCClient(opts)
	var session Session
	session.client = client

	return &session
}

// GetRecord: Get a record containing the current state of the given session.
func (class *Session) GetRecord(self SessionRef) (retval SessionRecord, err error) {
	method := "session.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSessionRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the session instance with the specified UUID.
func (class *Session) GetByUUID(uUID string) (retval SessionRef, err error) {
	method := "session.get_by_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	uUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uuid"), uUID)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, uUIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSessionRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given session.
func (class *Session) GetUUID(self SessionRef) (retval string, err error) {
	method := "session.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// GetThisHost: Get the this_host field of the given session.
func (class *Session) GetThisHost(self SessionRef) (retval HostRef, err error) {
	method := "session.get_this_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRef(method+" -> ", result)
	return
}

// GetThisUser: Get the this_user field of the given session.
func (class *Session) GetThisUser(self SessionRef) (retval UserRef, err error) {
	method := "session.get_this_user"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeUserRef(method+" -> ", result)
	return
}

// GetLastActive: Get the last_active field of the given session.
func (class *Session) GetLastActive(self SessionRef) (retval time.Time, err error) {
	method := "session.get_last_active"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTime(method+" -> ", result)
	return
}

// GetPool: Get the pool field of the given session.
func (class *Session) GetPool(self SessionRef) (retval bool, err error) {
	method := "session.get_pool"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeBool(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given session.
func (class *Session) GetOtherConfig(self SessionRef) (retval map[string]string, err error) {
	method := "session.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToStringMap(method+" -> ", result)
	return
}

// GetIsLocalSuperuser: Get the is_local_superuser field of the given session.
func (class *Session) GetIsLocalSuperuser(self SessionRef) (retval bool, err error) {
	method := "session.get_is_local_superuser"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeBool(method+" -> ", result)
	return
}

// GetSubject: Get the subject field of the given session.
func (class *Session) GetSubject(self SessionRef) (retval SubjectRef, err error) {
	method := "session.get_subject"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSubjectRef(method+" -> ", result)
	return
}

// GetValidationTime: Get the validation_time field of the given session.
func (class *Session) GetValidationTime(self SessionRef) (retval time.Time, err error) {
	method := "session.get_validation_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTime(method+" -> ", result)
	return
}

// GetAuthUserSid: Get the auth_user_sid field of the given session.
func (class *Session) GetAuthUserSid(self SessionRef) (retval string, err error) {
	method := "session.get_auth_user_sid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// GetAuthUserName: Get the auth_user_name field of the given session.
func (class *Session) GetAuthUserName(self SessionRef) (retval string, err error) {
	method := "session.get_auth_user_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// GetRbacPermissions: Get the rbac_permissions field of the given session.
func (class *Session) GetRbacPermissions(self SessionRef) (retval []string, err error) {
	method := "session.get_rbac_permissions"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// GetTasks: Get the tasks field of the given session.
func (class *Session) GetTasks(self SessionRef) (retval []TaskRef, err error) {
	method := "session.get_tasks"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRefSet(method+" -> ", result)
	return
}

// GetParent: Get the parent field of the given session.
func (class *Session) GetParent(self SessionRef) (retval SessionRef, err error) {
	method := "session.get_parent"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSessionRef(method+" -> ", result)
	return
}

// GetOriginator: Get the originator field of the given session.
func (class *Session) GetOriginator(self SessionRef) (retval string, err error) {
	method := "session.get_originator"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// GetClientCertificate: Get the client_certificate field of the given session.
func (class *Session) GetClientCertificate(self SessionRef) (retval bool, err error) {
	method := "session.get_client_certificate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeBool(method+" -> ", result)
	return
}

// SetOtherConfig: Set the other_config field of the given session.
func (class *Session) SetOtherConfig(self SessionRef, value map[string]string) (err error) {
	method := "session.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = class.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given session.
func (class *Session) AddToOtherConfig(self SessionRef, key string, value string) (err error) {
	method := "session.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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
	_, err = class.client.sendCall(method, sessionIDArg, selfArg, keyArg, valueArg)
	return
}

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given session.  If the key is not in that Map, then do nothing.
func (class *Session) RemoveFromOtherConfig(self SessionRef, key string) (err error) {
	method := "session.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	keyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "key"), key)
	if err != nil {
		return
	}
	_, err = class.client.sendCall(method, sessionIDArg, selfArg, keyArg)
	return
}

// LoginWithPassword: Attempt to authenticate the user, returning a session reference if successful
//
// Errors:
// SESSION_AUTHENTICATION_FAILED - The credentials given by the user are incorrect, so access has been denied, and you have not been issued a session handle.
// HOST_IS_SLAVE - You cannot make regular API calls directly on a supporter. Please pass API calls via the coordinator host.
func (class *Session) LoginWithPassword(uname string, pwd string, version string, originator string) (retval SessionRef, err error) {
	method := "session.login_with_password"
	unameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uname"), uname)
	if err != nil {
		return
	}
	pwdArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "pwd"), pwd)
	if err != nil {
		return
	}
	versionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "version"), version)
	if err != nil {
		return
	}
	originatorArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "originator"), originator)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, unameArg, pwdArg, versionArg, originatorArg)
	if err != nil {
		return
	}
	retval, err = deserializeSessionRef(method+" -> ", result)
	if err != nil {
		return
	}
	class.ref = retval
	err = setSessionDetails(class)
	return
}

// Logout: Log out of a session
func (class *Session) Logout() (err error) {
	method := "session.logout"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	_, err = class.client.sendCall(method, sessionIDArg)
	class.ref = ""
	return
}

// ChangePassword: Change the account password; if your session is authenticated with root privileges then the old_pwd is validated and the new_pwd is set regardless
func (class *Session) ChangePassword(oldPwd string, newPwd string) (err error) {
	method := "session.change_password"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	oldPwdArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "old_pwd"), oldPwd)
	if err != nil {
		return
	}
	newPwdArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_pwd"), newPwd)
	if err != nil {
		return
	}
	_, err = class.client.sendCall(method, sessionIDArg, oldPwdArg, newPwdArg)
	return
}

// SlaveLocalLoginWithPassword: Authenticate locally against a slave in emergency mode. Note the resulting sessions are only good for use on this host.
func (class *Session) SlaveLocalLoginWithPassword(uname string, pwd string) (retval SessionRef, err error) {
	method := "session.slave_local_login_with_password"
	unameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uname"), uname)
	if err != nil {
		return
	}
	pwdArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "pwd"), pwd)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, unameArg, pwdArg)
	if err != nil {
		return
	}
	retval, err = deserializeSessionRef(method+" -> ", result)
	if err != nil {
		return
	}
	class.ref = retval
	err = setSessionDetails(class)
	return
}

// CreateFromDBFile:
func (class *Session) CreateFromDBFile(filename string) (retval SessionRef, err error) {
	method := "session.create_from_db_file"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	filenameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "filename"), filename)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, filenameArg)
	if err != nil {
		return
	}
	retval, err = deserializeSessionRef(method+" -> ", result)
	return
}

// AsyncCreateFromDBFile:
func (class *Session) AsyncCreateFromDBFile(filename string) (retval TaskRef, err error) {
	method := "Async.session.create_from_db_file"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	filenameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "filename"), filename)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, filenameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// LocalLogout: Log out of local session.
func (class *Session) LocalLogout() (err error) {
	method := "session.local_logout"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	_, err = class.client.sendCall(method, sessionIDArg)
	class.ref = ""
	return
}

// GetAllSubjectIdentifiers: Return a list of all the user subject-identifiers of all existing sessions
func (class *Session) GetAllSubjectIdentifiers() (retval []string, err error) {
	method := "session.get_all_subject_identifiers"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// AsyncGetAllSubjectIdentifiers: Return a list of all the user subject-identifiers of all existing sessions
func (class *Session) AsyncGetAllSubjectIdentifiers() (retval TaskRef, err error) {
	method := "Async.session.get_all_subject_identifiers"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// LogoutSubjectIdentifier: Log out all sessions associated to a user subject-identifier, except the session associated with the context calling this function
func (class *Session) LogoutSubjectIdentifier(subjectIdentifier string) (err error) {
	method := "session.logout_subject_identifier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	subjectIdentifierArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "subject_identifier"), subjectIdentifier)
	if err != nil {
		return
	}
	_, err = class.client.sendCall(method, sessionIDArg, subjectIdentifierArg)
	return
}

// AsyncLogoutSubjectIdentifier: Log out all sessions associated to a user subject-identifier, except the session associated with the context calling this function
func (class *Session) AsyncLogoutSubjectIdentifier(subjectIdentifier string) (retval TaskRef, err error) {
	method := "Async.session.logout_subject_identifier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), class.ref)
	if err != nil {
		return
	}
	subjectIdentifierArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "subject_identifier"), subjectIdentifier)
	if err != nil {
		return
	}
	result, err := class.client.sendCall(method, sessionIDArg, subjectIdentifierArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}
