package xenapi

import (
	"fmt"
)

type USBGroupRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// List of PUSBs in the group
	PUSBs []PUSBRef
	// List of VUSBs using the group
	VUSBs []VUSBRef
	// Additional configuration
	OtherConfig map[string]string
}

type USBGroupRef string

// A group of compatible USBs across the resource pool
type uSBGroup struct{}

var USBGroup uSBGroup

// GetRecord: Get a record containing the current state of the given USB_group.
func (uSBGroup) GetRecord(session *Session, self USBGroupRef) (retval USBGroupRecord, err error) {
	method := "USB_group.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeUSBGroupRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the USB_group instance with the specified UUID.
func (uSBGroup) GetByUUID(session *Session, uUID string) (retval USBGroupRef, err error) {
	method := "USB_group.get_by_uuid"
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
	retval, err = deserializeUSBGroupRef(method+" -> ", result)
	return
}

// GetByNameLabel: Get all the USB_group instances with the given label.
func (uSBGroup) GetByNameLabel(session *Session, label string) (retval []USBGroupRef, err error) {
	method := "USB_group.get_by_name_label"
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
	retval, err = deserializeUSBGroupRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given USB_group.
func (uSBGroup) GetUUID(session *Session, self USBGroupRef) (retval string, err error) {
	method := "USB_group.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given USB_group.
func (uSBGroup) GetNameLabel(session *Session, self USBGroupRef) (retval string, err error) {
	method := "USB_group.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given USB_group.
func (uSBGroup) GetNameDescription(session *Session, self USBGroupRef) (retval string, err error) {
	method := "USB_group.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPUSBs: Get the PUSBs field of the given USB_group.
func (uSBGroup) GetPUSBs(session *Session, self USBGroupRef) (retval []PUSBRef, err error) {
	method := "USB_group.get_PUSBs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePUSBRefSet(method+" -> ", result)
	return
}

// GetVUSBs: Get the VUSBs field of the given USB_group.
func (uSBGroup) GetVUSBs(session *Session, self USBGroupRef) (retval []VUSBRef, err error) {
	method := "USB_group.get_VUSBs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVUSBRefSet(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given USB_group.
func (uSBGroup) GetOtherConfig(session *Session, self USBGroupRef) (retval map[string]string, err error) {
	method := "USB_group.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel: Set the name/label field of the given USB_group.
func (uSBGroup) SetNameLabel(session *Session, self USBGroupRef, value string) (err error) {
	method := "USB_group.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription: Set the name/description field of the given USB_group.
func (uSBGroup) SetNameDescription(session *Session, self USBGroupRef, value string) (err error) {
	method := "USB_group.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given USB_group.
func (uSBGroup) SetOtherConfig(session *Session, self USBGroupRef, value map[string]string) (err error) {
	method := "USB_group.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given USB_group.
func (uSBGroup) AddToOtherConfig(session *Session, self USBGroupRef, key string, value string) (err error) {
	method := "USB_group.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given USB_group.  If the key is not in that Map, then do nothing.
func (uSBGroup) RemoveFromOtherConfig(session *Session, self USBGroupRef, key string) (err error) {
	method := "USB_group.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create:
func (uSBGroup) Create(session *Session, nameLabel string, nameDescription string, otherConfig map[string]string) (retval USBGroupRef, err error) {
	method := "USB_group.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, otherConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeUSBGroupRef(method+" -> ", result)
	return
}

// AsyncCreate:
func (uSBGroup) AsyncCreate(session *Session, nameLabel string, nameDescription string, otherConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.USB_group.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, otherConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy:
func (uSBGroup) Destroy(session *Session, self USBGroupRef) (err error) {
	method := "USB_group.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy:
func (uSBGroup) AsyncDestroy(session *Session, self USBGroupRef) (retval TaskRef, err error) {
	method := "Async.USB_group.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the USB_groups known to the system.
func (uSBGroup) GetAll(session *Session) (retval []USBGroupRef, err error) {
	method := "USB_group.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeUSBGroupRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of USB_group references to USB_group records for all USB_groups known to the system.
func (uSBGroup) GetAllRecords(session *Session) (retval map[USBGroupRef]USBGroupRecord, err error) {
	method := "USB_group.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeUSBGroupRefToUSBGroupRecordMap(method+" -> ", result)
	return
}
