package xenapi

import (
	"fmt"
)

type AuthRecord struct {
}

type AuthRef string

// Management of remote authentication services
type auth struct{}

var Auth auth

// GetSubjectIdentifier: This call queries the external directory service to obtain the subject_identifier as a string from the human-readable subject_name
func (auth) GetSubjectIdentifier(session *Session, subjectName string) (retval string, err error) {
	method := "auth.get_subject_identifier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	subjectNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "subject_name"), subjectName)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, subjectNameArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// GetSubjectInformationFromIdentifier: This call queries the external directory service to obtain the user information (e.g. username, organization etc) from the specified subject_identifier
func (auth) GetSubjectInformationFromIdentifier(session *Session, subjectIdentifier string) (retval map[string]string, err error) {
	method := "auth.get_subject_information_from_identifier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	subjectIdentifierArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "subject_identifier"), subjectIdentifier)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, subjectIdentifierArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToStringMap(method+" -> ", result)
	return
}

// GetGroupMembership: This calls queries the external directory service to obtain the transitively-closed set of groups that the the subject_identifier is member of.
func (auth) GetGroupMembership(session *Session, subjectIdentifier string) (retval []string, err error) {
	method := "auth.get_group_membership"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	subjectIdentifierArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "subject_identifier"), subjectIdentifier)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, subjectIdentifierArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}
