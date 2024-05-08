package xenapi

import (
	"fmt"
)

type VUSBRecord struct {
	// Unique identifier/object reference
	UUID string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VusbOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VusbOperations
	// VM that owns the VUSB
	VM VMRef
	// USB group used by the VUSB
	USBGroup USBGroupRef
	// Additional configuration
	OtherConfig map[string]string
	// is the device currently attached
	CurrentlyAttached bool
}

type VUSBRef string

// Describes the vusb device
type vUSB struct{}

var VUSB vUSB

// GetRecord: Get a record containing the current state of the given VUSB.
func (vUSB) GetRecord(session *Session, self VUSBRef) (retval VUSBRecord, err error) {
	method := "VUSB.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVUSBRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VUSB instance with the specified UUID.
func (vUSB) GetByUUID(session *Session, uUID string) (retval VUSBRef, err error) {
	method := "VUSB.get_by_uuid"
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
	retval, err = deserializeVUSBRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VUSB.
func (vUSB) GetUUID(session *Session, self VUSBRef) (retval string, err error) {
	method := "VUSB.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given VUSB.
func (vUSB) GetAllowedOperations(session *Session, self VUSBRef) (retval []VusbOperations, err error) {
	method := "VUSB.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVusbOperationsSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given VUSB.
func (vUSB) GetCurrentOperations(session *Session, self VUSBRef) (retval map[string]VusbOperations, err error) {
	method := "VUSB.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumVusbOperationsMap(method+" -> ", result)
	return
}

// GetVM: Get the VM field of the given VUSB.
func (vUSB) GetVM(session *Session, self VUSBRef) (retval VMRef, err error) {
	method := "VUSB.get_VM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// GetUSBGroup: Get the USB_group field of the given VUSB.
func (vUSB) GetUSBGroup(session *Session, self VUSBRef) (retval USBGroupRef, err error) {
	method := "VUSB.get_USB_group"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeUSBGroupRef(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given VUSB.
func (vUSB) GetOtherConfig(session *Session, self VUSBRef) (retval map[string]string, err error) {
	method := "VUSB.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCurrentlyAttached: Get the currently_attached field of the given VUSB.
func (vUSB) GetCurrentlyAttached(session *Session, self VUSBRef) (retval bool, err error) {
	method := "VUSB.get_currently_attached"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given VUSB.
func (vUSB) SetOtherConfig(session *Session, self VUSBRef, value map[string]string) (err error) {
	method := "VUSB.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VUSB.
func (vUSB) AddToOtherConfig(session *Session, self VUSBRef, key string, value string) (err error) {
	method := "VUSB.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VUSB.  If the key is not in that Map, then do nothing.
func (vUSB) RemoveFromOtherConfig(session *Session, self VUSBRef, key string) (err error) {
	method := "VUSB.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Create a new VUSB record in the database only
func (vUSB) Create(session *Session, vM VMRef, uSBGroup USBGroupRef, otherConfig map[string]string) (retval VUSBRef, err error) {
	method := "VUSB.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "VM"), vM)
	if err != nil {
		return
	}
	uSBGroupArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "USB_group"), uSBGroup)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, uSBGroupArg, otherConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeVUSBRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new VUSB record in the database only
func (vUSB) AsyncCreate(session *Session, vM VMRef, uSBGroup USBGroupRef, otherConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.VUSB.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "VM"), vM)
	if err != nil {
		return
	}
	uSBGroupArg, err := serializeUSBGroupRef(fmt.Sprintf("%s(%s)", method, "USB_group"), uSBGroup)
	if err != nil {
		return
	}
	otherConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "other_config"), otherConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, uSBGroupArg, otherConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Unplug: Unplug the vusb device from the vm.
func (vUSB) Unplug(session *Session, self VUSBRef) (err error) {
	method := "VUSB.unplug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncUnplug: Unplug the vusb device from the vm.
func (vUSB) AsyncUnplug(session *Session, self VUSBRef) (retval TaskRef, err error) {
	method := "Async.VUSB.unplug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy: Removes a VUSB record from the database
func (vUSB) Destroy(session *Session, self VUSBRef) (err error) {
	method := "VUSB.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Removes a VUSB record from the database
func (vUSB) AsyncDestroy(session *Session, self VUSBRef) (retval TaskRef, err error) {
	method := "Async.VUSB.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the VUSBs known to the system.
func (vUSB) GetAll(session *Session) (retval []VUSBRef, err error) {
	method := "VUSB.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVUSBRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VUSB references to VUSB records for all VUSBs known to the system.
func (vUSB) GetAllRecords(session *Session) (retval map[VUSBRef]VUSBRecord, err error) {
	method := "VUSB.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVUSBRefToVUSBRecordMap(method+" -> ", result)
	return
}

