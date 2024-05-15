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

// GetAllRecords: Return a map of Feature references to Feature records for all Features known to the system.
// Version: falcon
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

// GetAllRecords1: Return a map of Feature references to Feature records for all Features known to the system.
// Version: falcon
func (feature) GetAllRecords1(session *Session) (retval map[FeatureRef]FeatureRecord, err error) {
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

// GetAll: Return a list of all the Features known to the system.
// Version: falcon
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

// GetAll1: Return a list of all the Features known to the system.
// Version: falcon
func (feature) GetAll1(session *Session) (retval []FeatureRef, err error) {
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

// GetHost: Get the host field of the given Feature.
// Version: falcon
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

// GetHost2: Get the host field of the given Feature.
// Version: falcon
func (feature) GetHost2(session *Session, self FeatureRef) (retval HostRef, err error) {
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

// GetVersion: Get the version field of the given Feature.
// Version: falcon
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

// GetVersion2: Get the version field of the given Feature.
// Version: falcon
func (feature) GetVersion2(session *Session, self FeatureRef) (retval string, err error) {
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

// GetExperimental: Get the experimental field of the given Feature.
// Version: falcon
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

// GetExperimental2: Get the experimental field of the given Feature.
// Version: falcon
func (feature) GetExperimental2(session *Session, self FeatureRef) (retval bool, err error) {
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

// GetEnabled: Get the enabled field of the given Feature.
// Version: falcon
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

// GetEnabled2: Get the enabled field of the given Feature.
// Version: falcon
func (feature) GetEnabled2(session *Session, self FeatureRef) (retval bool, err error) {
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

// GetNameDescription: Get the name/description field of the given Feature.
// Version: falcon
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

// GetNameDescription2: Get the name/description field of the given Feature.
// Version: falcon
func (feature) GetNameDescription2(session *Session, self FeatureRef) (retval string, err error) {
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

// GetNameLabel: Get the name/label field of the given Feature.
// Version: falcon
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

// GetNameLabel2: Get the name/label field of the given Feature.
// Version: falcon
func (feature) GetNameLabel2(session *Session, self FeatureRef) (retval string, err error) {
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

// GetUUID: Get the uuid field of the given Feature.
// Version: falcon
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

// GetUUID2: Get the uuid field of the given Feature.
// Version: falcon
func (feature) GetUUID2(session *Session, self FeatureRef) (retval string, err error) {
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

// GetByNameLabel: Get all the Feature instances with the given label.
// Version: falcon
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

// GetByNameLabel2: Get all the Feature instances with the given label.
// Version: falcon
func (feature) GetByNameLabel2(session *Session, label string) (retval []FeatureRef, err error) {
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

// GetByUUID: Get a reference to the Feature instance with the specified UUID.
// Version: falcon
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

// GetByUUID2: Get a reference to the Feature instance with the specified UUID.
// Version: falcon
func (feature) GetByUUID2(session *Session, uUID string) (retval FeatureRef, err error) {
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

// GetRecord: Get a record containing the current state of the given Feature.
// Version: falcon
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

// GetRecord2: Get a record containing the current state of the given Feature.
// Version: falcon
func (feature) GetRecord2(session *Session, self FeatureRef) (retval FeatureRecord, err error) {
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
