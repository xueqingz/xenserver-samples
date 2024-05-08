package xenapi

import (
	"fmt"
)

type FeatureRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// Indicates whether the feature is enabled
	Enabled bool
	// Indicates whether the feature is experimental (as opposed to stable and fully supported)
	Experimental bool
	// The version of this feature
	Version string
	// The host where this feature is available
	Host HostRef
}

type FeatureRef string

// A new piece of functionality
type feature struct{}

var Feature feature

// GetRecord: Get a record containing the current state of the given Feature.
func (feature) GetRecord(session *Session, self FeatureRef) (retval FeatureRecord, err error) {
	method := "Feature.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeFeatureRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeFeatureRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the Feature instance with the specified UUID.
func (feature) GetByUUID(session *Session, uUID string) (retval FeatureRef, err error) {
	method := "Feature.get_by_uuid"
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
	retval, err = deserializeFeatureRef(method+" -> ", result)
	return
}

// GetByNameLabel: Get all the Feature instances with the given label.
func (feature) GetByNameLabel(session *Session, label string) (retval []FeatureRef, err error) {
	method := "Feature.get_by_name_label"
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
	retval, err = deserializeFeatureRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given Feature.
func (feature) GetUUID(session *Session, self FeatureRef) (retval string, err error) {
	method := "Feature.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeFeatureRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given Feature.
func (feature) GetNameLabel(session *Session, self FeatureRef) (retval string, err error) {
	method := "Feature.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeFeatureRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given Feature.
func (feature) GetNameDescription(session *Session, self FeatureRef) (retval string, err error) {
	method := "Feature.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeFeatureRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnabled: Get the enabled field of the given Feature.
func (feature) GetEnabled(session *Session, self FeatureRef) (retval bool, err error) {
	method := "Feature.get_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeFeatureRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetExperimental: Get the experimental field of the given Feature.
func (feature) GetExperimental(session *Session, self FeatureRef) (retval bool, err error) {
	method := "Feature.get_experimental"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeFeatureRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVersion: Get the version field of the given Feature.
func (feature) GetVersion(session *Session, self FeatureRef) (retval string, err error) {
	method := "Feature.get_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeFeatureRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost: Get the host field of the given Feature.
func (feature) GetHost(session *Session, self FeatureRef) (retval HostRef, err error) {
	method := "Feature.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeFeatureRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRef(method+" -> ", result)
	return
}

// GetAll: Return a list of all the Features known to the system.
func (feature) GetAll(session *Session) (retval []FeatureRef, err error) {
	method := "Feature.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeFeatureRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of Feature references to Feature records for all Features known to the system.
func (feature) GetAllRecords(session *Session) (retval map[FeatureRef]FeatureRecord, err error) {
	method := "Feature.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeFeatureRefToFeatureRecordMap(method+" -> ", result)
	return
}

