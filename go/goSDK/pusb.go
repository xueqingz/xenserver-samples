package xenapi

import (
	"fmt"
)

type PUSBRecord struct {
	// Unique identifier/object reference
	UUID string
	// USB group the PUSB is contained in
	USBGroup USBGroupRef
	// Physical machine that owns the USB device
	Host HostRef
	// port path of USB device
	Path string
	// vendor id of the USB device
	VendorID string
	// vendor description of the USB device
	VendorDesc string
	// product id of the USB device
	ProductID string
	// product description of the USB device
	ProductDesc string
	// serial of the USB device
	Serial string
	// USB device version
	Version string
	// USB device description
	Description string
	// enabled for passthrough
	PassthroughEnabled bool
	// additional configuration
	OtherConfig map[string]string
	// USB device speed
	Speed float64
}

type PUSBRef string

// A physical USB device
type pUSB struct{}

var PUSB pUSB

// GetRecord: Get a record containing the current state of the given PUSB.
func (pUSB) GetRecord(session *Session, self PUSBRef) (retval PUSBRecord, err error) {
	method := "PUSB.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePUSBRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the PUSB instance with the specified UUID.
func (pUSB) GetByUUID(session *Session, uUID string) (retval PUSBRef, err error) {
	method := "PUSB.get_by_uuid"
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
	retval, err = deserializePUSBRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given PUSB.
func (pUSB) GetUUID(session *Session, self PUSBRef) (retval string, err error) {
	method := "PUSB.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUSBGroup: Get the USB_group field of the given PUSB.
func (pUSB) GetUSBGroup(session *Session, self PUSBRef) (retval USBGroupRef, err error) {
	method := "PUSB.get_USB_group"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost: Get the host field of the given PUSB.
func (pUSB) GetHost(session *Session, self PUSBRef) (retval HostRef, err error) {
	method := "PUSB.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPath: Get the path field of the given PUSB.
func (pUSB) GetPath(session *Session, self PUSBRef) (retval string, err error) {
	method := "PUSB.get_path"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVendorID: Get the vendor_id field of the given PUSB.
func (pUSB) GetVendorID(session *Session, self PUSBRef) (retval string, err error) {
	method := "PUSB.get_vendor_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVendorDesc: Get the vendor_desc field of the given PUSB.
func (pUSB) GetVendorDesc(session *Session, self PUSBRef) (retval string, err error) {
	method := "PUSB.get_vendor_desc"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetProductID: Get the product_id field of the given PUSB.
func (pUSB) GetProductID(session *Session, self PUSBRef) (retval string, err error) {
	method := "PUSB.get_product_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetProductDesc: Get the product_desc field of the given PUSB.
func (pUSB) GetProductDesc(session *Session, self PUSBRef) (retval string, err error) {
	method := "PUSB.get_product_desc"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSerial: Get the serial field of the given PUSB.
func (pUSB) GetSerial(session *Session, self PUSBRef) (retval string, err error) {
	method := "PUSB.get_serial"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVersion: Get the version field of the given PUSB.
func (pUSB) GetVersion(session *Session, self PUSBRef) (retval string, err error) {
	method := "PUSB.get_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDescription: Get the description field of the given PUSB.
func (pUSB) GetDescription(session *Session, self PUSBRef) (retval string, err error) {
	method := "PUSB.get_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPassthroughEnabled: Get the passthrough_enabled field of the given PUSB.
func (pUSB) GetPassthroughEnabled(session *Session, self PUSBRef) (retval bool, err error) {
	method := "PUSB.get_passthrough_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given PUSB.
func (pUSB) GetOtherConfig(session *Session, self PUSBRef) (retval map[string]string, err error) {
	method := "PUSB.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSpeed: Get the speed field of the given PUSB.
func (pUSB) GetSpeed(session *Session, self PUSBRef) (retval float64, err error) {
	method := "PUSB.get_speed"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeFloat(method+" -> ", result)
	return
}

// SetOtherConfig: Set the other_config field of the given PUSB.
func (pUSB) SetOtherConfig(session *Session, self PUSBRef, value map[string]string) (err error) {
	method := "PUSB.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given PUSB.
func (pUSB) AddToOtherConfig(session *Session, self PUSBRef, key string, value string) (err error) {
	method := "PUSB.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PUSB.  If the key is not in that Map, then do nothing.
func (pUSB) RemoveFromOtherConfig(session *Session, self PUSBRef, key string) (err error) {
	method := "PUSB.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Scan:
func (pUSB) Scan(session *Session, host HostRef) (err error) {
	method := "PUSB.scan"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg)
	return
}

// AsyncScan:
func (pUSB) AsyncScan(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.PUSB.scan"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetPassthroughEnabled:
func (pUSB) SetPassthroughEnabled(session *Session, self PUSBRef, value bool) (err error) {
	method := "PUSB.set_passthrough_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetPassthroughEnabled:
func (pUSB) AsyncSetPassthroughEnabled(session *Session, self PUSBRef, value bool) (retval TaskRef, err error) {
	method := "Async.PUSB.set_passthrough_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePUSBRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the PUSBs known to the system.
func (pUSB) GetAll(session *Session) (retval []PUSBRef, err error) {
	method := "PUSB.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePUSBRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of PUSB references to PUSB records for all PUSBs known to the system.
func (pUSB) GetAllRecords(session *Session) (retval map[PUSBRef]PUSBRecord, err error) {
	method := "PUSB.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePUSBRefToPUSBRecordMap(method+" -> ", result)
	return
}

