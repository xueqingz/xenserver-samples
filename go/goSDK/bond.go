package xenapi

import (
	"fmt"
)

type BondRecord struct {
	// Unique identifier/object reference
	UUID string
	// The bonded interface
	Master PIFRef
	// The interfaces which are part of this bond
	Slaves []PIFRef
	// additional configuration
	OtherConfig map[string]string
	// The PIF of which the IP configuration and MAC were copied to the bond, and which will receive all configuration/VLANs/VIFs on the bond if the bond is destroyed
	PrimarySlave PIFRef
	// The algorithm used to distribute traffic among the bonded NICs
	Mode BondMode
	// Additional configuration properties specific to the bond mode.
	Properties map[string]string
	// Number of links up in this bond
	LinksUp int
	// true if the MAC was taken from the primary slave when the bond was created, and false if the client specified the MAC
	AutoUpdateMac bool
}

type BondRef string

// A Network bond that combines physical network interfaces, also known as link aggregation
type bond struct{}

var Bond bond

// GetRecord: Get a record containing the current state of the given Bond.
func (bond) GetRecord(session *Session, self BondRef) (retval BondRecord, err error) {
	method := "Bond.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeBondRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the Bond instance with the specified UUID.
func (bond) GetByUUID(session *Session, uUID string) (retval BondRef, err error) {
	method := "Bond.get_by_uuid"
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
	retval, err = deserializeBondRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given Bond.
func (bond) GetUUID(session *Session, self BondRef) (retval string, err error) {
	method := "Bond.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMaster: Get the master field of the given Bond.
func (bond) GetMaster(session *Session, self BondRef) (retval PIFRef, err error) {
	method := "Bond.get_master"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSlaves: Get the slaves field of the given Bond.
func (bond) GetSlaves(session *Session, self BondRef) (retval []PIFRef, err error) {
	method := "Bond.get_slaves"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePIFRefSet(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given Bond.
func (bond) GetOtherConfig(session *Session, self BondRef) (retval map[string]string, err error) {
	method := "Bond.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPrimarySlave: Get the primary_slave field of the given Bond.
func (bond) GetPrimarySlave(session *Session, self BondRef) (retval PIFRef, err error) {
	method := "Bond.get_primary_slave"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMode: Get the mode field of the given Bond.
func (bond) GetMode(session *Session, self BondRef) (retval BondMode, err error) {
	method := "Bond.get_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumBondMode(method+" -> ", result)
	return
}

// GetProperties: Get the properties field of the given Bond.
func (bond) GetProperties(session *Session, self BondRef) (retval map[string]string, err error) {
	method := "Bond.get_properties"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLinksUp: Get the links_up field of the given Bond.
func (bond) GetLinksUp(session *Session, self BondRef) (retval int, err error) {
	method := "Bond.get_links_up"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAutoUpdateMac: Get the auto_update_mac field of the given Bond.
func (bond) GetAutoUpdateMac(session *Session, self BondRef) (retval bool, err error) {
	method := "Bond.get_auto_update_mac"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given Bond.
func (bond) SetOtherConfig(session *Session, self BondRef, value map[string]string) (err error) {
	method := "Bond.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given Bond.
func (bond) AddToOtherConfig(session *Session, self BondRef, key string, value string) (err error) {
	method := "Bond.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given Bond.  If the key is not in that Map, then do nothing.
func (bond) RemoveFromOtherConfig(session *Session, self BondRef, key string) (err error) {
	method := "Bond.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Create an interface bond
func (bond) Create(session *Session, network NetworkRef, members []PIFRef, mAC string, mode BondMode, properties map[string]string) (retval BondRef, err error) {
	method := "Bond.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	membersArg, err := serializePIFRefSet(fmt.Sprintf("%s(%s)", method, "members"), members)
	if err != nil {
		return
	}
	mACArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "MAC"), mAC)
	if err != nil {
		return
	}
	modeArg, err := serializeEnumBondMode(fmt.Sprintf("%s(%s)", method, "mode"), mode)
	if err != nil {
		return
	}
	propertiesArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "properties"), properties)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, networkArg, membersArg, mACArg, modeArg, propertiesArg)
	if err != nil {
		return
	}
	retval, err = deserializeBondRef(method+" -> ", result)
	return
}

// AsyncCreate: Create an interface bond
func (bond) AsyncCreate(session *Session, network NetworkRef, members []PIFRef, mAC string, mode BondMode, properties map[string]string) (retval TaskRef, err error) {
	method := "Async.Bond.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	membersArg, err := serializePIFRefSet(fmt.Sprintf("%s(%s)", method, "members"), members)
	if err != nil {
		return
	}
	mACArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "MAC"), mAC)
	if err != nil {
		return
	}
	modeArg, err := serializeEnumBondMode(fmt.Sprintf("%s(%s)", method, "mode"), mode)
	if err != nil {
		return
	}
	propertiesArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "properties"), properties)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, networkArg, membersArg, mACArg, modeArg, propertiesArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Destroy an interface bond
func (bond) Destroy(session *Session, self BondRef) (err error) {
	method := "Bond.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy an interface bond
func (bond) AsyncDestroy(session *Session, self BondRef) (retval TaskRef, err error) {
	method := "Async.Bond.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetMode: Change the bond mode
func (bond) SetMode(session *Session, self BondRef, value BondMode) (err error) {
	method := "Bond.set_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumBondMode(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetMode: Change the bond mode
func (bond) AsyncSetMode(session *Session, self BondRef, value BondMode) (retval TaskRef, err error) {
	method := "Async.Bond.set_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumBondMode(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetProperty: Set the value of a property of the bond
func (bond) SetProperty(session *Session, self BondRef, name string, value string) (err error) {
	method := "Bond.set_property"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, nameArg, valueArg)
	return
}

// AsyncSetProperty: Set the value of a property of the bond
func (bond) AsyncSetProperty(session *Session, self BondRef, name string, value string) (retval TaskRef, err error) {
	method := "Async.Bond.set_property"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeBondRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, nameArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetAll: Return a list of all the Bonds known to the system.
func (bond) GetAll(session *Session) (retval []BondRef, err error) {
	method := "Bond.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeBondRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of Bond references to Bond records for all Bonds known to the system.
func (bond) GetAllRecords(session *Session) (retval map[BondRef]BondRecord, err error) {
	method := "Bond.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeBondRefToBondRecordMap(method+" -> ", result)
	return
}

