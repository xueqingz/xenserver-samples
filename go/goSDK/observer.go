package xenapi

import (
	"fmt"
)

type ObserverRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// The list of hosts the observer is active on. An empty list means all hosts
	Hosts []HostRef
	// Attributes that observer will add to the data they produce
	Attributes map[string]string
	// The list of endpoints where data is exported to. Each endpoint is a URL or the string &apos;bugtool&apos; refering to the internal logs
	Endpoints []string
	// The list of xenserver components the observer will broadcast. An empty list means all components
	Components []string
	// This denotes if the observer is enabled. true if it is enabled and false if it is disabled
	Enabled bool
}

type ObserverRef string

// Describes a observer which will control observability activity in the Toolstack
type observer struct{}

var Observer observer

// GetRecord: Get a record containing the current state of the given Observer.
func (observer) GetRecord(session *Session, self ObserverRef) (retval ObserverRecord, err error) {
	method := "Observer.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeObserverRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the Observer instance with the specified UUID.
func (observer) GetByUUID(session *Session, uUID string) (retval ObserverRef, err error) {
	method := "Observer.get_by_uuid"
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
	retval, err = deserializeObserverRef(method+" -> ", result)
	return
}

// Create: Create a new Observer instance, and return its handle. The constructor args are: name_label, name_description, hosts, attributes, endpoints, components, enabled (* = non-optional).
func (observer) Create(session *Session, args ObserverRecord) (retval ObserverRef, err error) {
	method := "Observer.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeObserverRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeObserverRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new Observer instance, and return its handle. The constructor args are: name_label, name_description, hosts, attributes, endpoints, components, enabled (* = non-optional).
func (observer) AsyncCreate(session *Session, args ObserverRecord) (retval TaskRef, err error) {
	method := "Async.Observer.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeObserverRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// Destroy: Destroy the specified Observer instance.
func (observer) Destroy(session *Session, self ObserverRef) (err error) {
	method := "Observer.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified Observer instance.
func (observer) AsyncDestroy(session *Session, self ObserverRef) (retval TaskRef, err error) {
	method := "Async.Observer.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByNameLabel: Get all the Observer instances with the given label.
func (observer) GetByNameLabel(session *Session, label string) (retval []ObserverRef, err error) {
	method := "Observer.get_by_name_label"
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
	retval, err = deserializeObserverRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given Observer.
func (observer) GetUUID(session *Session, self ObserverRef) (retval string, err error) {
	method := "Observer.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given Observer.
func (observer) GetNameLabel(session *Session, self ObserverRef) (retval string, err error) {
	method := "Observer.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given Observer.
func (observer) GetNameDescription(session *Session, self ObserverRef) (retval string, err error) {
	method := "Observer.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHosts: Get the hosts field of the given Observer.
func (observer) GetHosts(session *Session, self ObserverRef) (retval []HostRef, err error) {
	method := "Observer.get_hosts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRefSet(method+" -> ", result)
	return
}

// GetAttributes: Get the attributes field of the given Observer.
func (observer) GetAttributes(session *Session, self ObserverRef) (retval map[string]string, err error) {
	method := "Observer.get_attributes"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEndpoints: Get the endpoints field of the given Observer.
func (observer) GetEndpoints(session *Session, self ObserverRef) (retval []string, err error) {
	method := "Observer.get_endpoints"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetComponents: Get the components field of the given Observer.
func (observer) GetComponents(session *Session, self ObserverRef) (retval []string, err error) {
	method := "Observer.get_components"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnabled: Get the enabled field of the given Observer.
func (observer) GetEnabled(session *Session, self ObserverRef) (retval bool, err error) {
	method := "Observer.get_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel: Set the name/label field of the given Observer.
func (observer) SetNameLabel(session *Session, self ObserverRef, value string) (err error) {
	method := "Observer.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription: Set the name/description field of the given Observer.
func (observer) SetNameDescription(session *Session, self ObserverRef, value string) (err error) {
	method := "Observer.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetHosts: Sets the hosts that the observer is to be registered on
func (observer) SetHosts(session *Session, self ObserverRef, value []HostRef) (err error) {
	method := "Observer.set_hosts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeHostRefSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetHosts: Sets the hosts that the observer is to be registered on
func (observer) AsyncSetHosts(session *Session, self ObserverRef, value []HostRef) (retval TaskRef, err error) {
	method := "Async.Observer.set_hosts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeHostRefSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetEnabled: Enable / disable this observer which will stop the observer from producing observability information
func (observer) SetEnabled(session *Session, self ObserverRef, value bool) (err error) {
	method := "Observer.set_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetEnabled: Enable / disable this observer which will stop the observer from producing observability information
func (observer) AsyncSetEnabled(session *Session, self ObserverRef, value bool) (retval TaskRef, err error) {
	method := "Async.Observer.set_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetAttributes: Set the attributes of an observer. These are used to emit metadata by the observer
func (observer) SetAttributes(session *Session, self ObserverRef, value map[string]string) (err error) {
	method := "Observer.set_attributes"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetAttributes: Set the attributes of an observer. These are used to emit metadata by the observer
func (observer) AsyncSetAttributes(session *Session, self ObserverRef, value map[string]string) (retval TaskRef, err error) {
	method := "Async.Observer.set_attributes"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetEndpoints: Set the file/HTTP endpoints the observer sends data to
func (observer) SetEndpoints(session *Session, self ObserverRef, value []string) (err error) {
	method := "Observer.set_endpoints"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetEndpoints: Set the file/HTTP endpoints the observer sends data to
func (observer) AsyncSetEndpoints(session *Session, self ObserverRef, value []string) (retval TaskRef, err error) {
	method := "Async.Observer.set_endpoints"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetComponents: Set the components on which the observer will broadcast to. i.e. xapi, xenopsd, networkd, etc
func (observer) SetComponents(session *Session, self ObserverRef, value []string) (err error) {
	method := "Observer.set_components"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetComponents: Set the components on which the observer will broadcast to. i.e. xapi, xenopsd, networkd, etc
func (observer) AsyncSetComponents(session *Session, self ObserverRef, value []string) (retval TaskRef, err error) {
	method := "Async.Observer.set_components"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeObserverRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetAll: Return a list of all the Observers known to the system.
func (observer) GetAll(session *Session) (retval []ObserverRef, err error) {
	method := "Observer.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeObserverRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of Observer references to Observer records for all Observers known to the system.
func (observer) GetAllRecords(session *Session) (retval map[ObserverRef]ObserverRecord, err error) {
	method := "Observer.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeObserverRefToObserverRecordMap(method+" -> ", result)
	return
}

