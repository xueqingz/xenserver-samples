package xenapi

import (
	"fmt"
)

type VLANRecord struct {
	// Unique identifier/object reference
	UUID string
	// interface on which traffic is tagged
	TaggedPIF PIFRef
	// interface on which traffic is untagged
	UntaggedPIF PIFRef
	// VLAN tag in use
	Tag int
	// additional configuration
	OtherConfig map[string]string
}

type VLANRef string

// A VLAN mux/demux
type vLAN struct{}

var VLAN vLAN

// GetRecord: Get a record containing the current state of the given VLAN.
func (vLAN) GetRecord(session *Session, self VLANRef) (retval VLANRecord, err error) {
	method := "VLAN.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVLANRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VLAN instance with the specified UUID.
func (vLAN) GetByUUID(session *Session, uUID string) (retval VLANRef, err error) {
	method := "VLAN.get_by_uuid"
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
	retval, err = deserializeVLANRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VLAN.
func (vLAN) GetUUID(session *Session, self VLANRef) (retval string, err error) {
	method := "VLAN.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTaggedPIF: Get the tagged_PIF field of the given VLAN.
func (vLAN) GetTaggedPIF(session *Session, self VLANRef) (retval PIFRef, err error) {
	method := "VLAN.get_tagged_PIF"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePIFRef(method+" -> ", result)
	return
}

// GetUntaggedPIF: Get the untagged_PIF field of the given VLAN.
func (vLAN) GetUntaggedPIF(session *Session, self VLANRef) (retval PIFRef, err error) {
	method := "VLAN.get_untagged_PIF"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePIFRef(method+" -> ", result)
	return
}

// GetTag: Get the tag field of the given VLAN.
func (vLAN) GetTag(session *Session, self VLANRef) (retval int, err error) {
	method := "VLAN.get_tag"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given VLAN.
func (vLAN) GetOtherConfig(session *Session, self VLANRef) (retval map[string]string, err error) {
	method := "VLAN.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given VLAN.
func (vLAN) SetOtherConfig(session *Session, self VLANRef, value map[string]string) (err error) {
	method := "VLAN.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VLAN.
func (vLAN) AddToOtherConfig(session *Session, self VLANRef, key string, value string) (err error) {
	method := "VLAN.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VLAN.  If the key is not in that Map, then do nothing.
func (vLAN) RemoveFromOtherConfig(session *Session, self VLANRef, key string) (err error) {
	method := "VLAN.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Create a VLAN mux/demuxer
func (vLAN) Create(session *Session, taggedPIF PIFRef, tag int, network NetworkRef) (retval VLANRef, err error) {
	method := "VLAN.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	taggedPIFArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "tagged_PIF"), taggedPIF)
	if err != nil {
		return
	}
	tagArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "tag"), tag)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, taggedPIFArg, tagArg, networkArg)
	if err != nil {
		return
	}
	retval, err = deserializeVLANRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a VLAN mux/demuxer
func (vLAN) AsyncCreate(session *Session, taggedPIF PIFRef, tag int, network NetworkRef) (retval TaskRef, err error) {
	method := "Async.VLAN.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	taggedPIFArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "tagged_PIF"), taggedPIF)
	if err != nil {
		return
	}
	tagArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "tag"), tag)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, taggedPIFArg, tagArg, networkArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Destroy a VLAN mux/demuxer
func (vLAN) Destroy(session *Session, self VLANRef) (err error) {
	method := "VLAN.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy a VLAN mux/demuxer
func (vLAN) AsyncDestroy(session *Session, self VLANRef) (retval TaskRef, err error) {
	method := "Async.VLAN.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVLANRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the VLANs known to the system.
func (vLAN) GetAll(session *Session) (retval []VLANRef, err error) {
	method := "VLAN.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVLANRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VLAN references to VLAN records for all VLANs known to the system.
func (vLAN) GetAllRecords(session *Session) (retval map[VLANRef]VLANRecord, err error) {
	method := "VLAN.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVLANRefToVLANRecordMap(method+" -> ", result)
	return
}

