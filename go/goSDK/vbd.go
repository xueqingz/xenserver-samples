package xenapi

import (
	"fmt"
)

type VBDRecord struct {
	// Unique identifier/object reference
	UUID string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VbdOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VbdOperations
	// the virtual machine
	VM VMRef
	// the virtual disk
	VDI VDIRef
	// device seen by the guest e.g. hda1
	Device string
	// user-friendly device name e.g. 0,1,2,etc.
	Userdevice string
	// true if this VBD is bootable
	Bootable bool
	// the mode the VBD should be mounted with
	Mode VbdMode
	// how the VBD will appear to the guest (e.g. disk or CD)
	Type VbdType
	// true if this VBD will support hot-unplug
	Unpluggable bool
	// true if a storage level lock was acquired
	StorageLock bool
	// if true this represents an empty drive
	Empty bool
	// additional configuration
	OtherConfig map[string]string
	// is the device currently attached (erased on reboot)
	CurrentlyAttached bool
	// error/success code associated with last attach-operation (erased on reboot)
	StatusCode int
	// error/success information associated with last attach-operation status (erased on reboot)
	StatusDetail string
	// Device runtime properties
	RuntimeProperties map[string]string
	// QoS algorithm to use
	QosAlgorithmType string
	// parameters for chosen QoS algorithm
	QosAlgorithmParams map[string]string
	// supported QoS algorithms for this VBD
	QosSupportedAlgorithms []string
	// metrics associated with this VBD
	Metrics VBDMetricsRef
}

type VBDRef string

// A virtual block device
type vBD struct{}

var VBD vBD

// GetAllRecords: Return a map of VBD references to VBD records for all VBDs known to the system.
// Version: rio
func (vBD) GetAllRecords(session *Session) (retval map[VBDRef]VBDRecord, err error) {
	method := "VBD.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRefToVBDRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of VBD references to VBD records for all VBDs known to the system.
// Version: rio
func (vBD) GetAllRecords1(session *Session) (retval map[VBDRef]VBDRecord, err error) {
	method := "VBD.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRefToVBDRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the VBDs known to the system.
// Version: rio
func (vBD) GetAll(session *Session) (retval []VBDRef, err error) {
	method := "VBD.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the VBDs known to the system.
// Version: rio
func (vBD) GetAll1(session *Session) (retval []VBDRef, err error) {
	method := "VBD.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRefSet(method+" -> ", result)
	return
}

// SetMode: Sets the mode of the VBD. The power_state of the VM must be halted.
// Version: rio
func (vBD) SetMode(session *Session, self VBDRef, value VbdMode) (err error) {
	method := "VBD.set_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVbdMode(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetMode: Sets the mode of the VBD. The power_state of the VM must be halted.
// Version: rio
func (vBD) AsyncSetMode(session *Session, self VBDRef, value VbdMode) (retval TaskRef, err error) {
	method := "Async.VBD.set_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVbdMode(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetMode3: Sets the mode of the VBD. The power_state of the VM must be halted.
// Version: rio
func (vBD) SetMode3(session *Session, self VBDRef, value VbdMode) (err error) {
	method := "VBD.set_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVbdMode(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetMode3: Sets the mode of the VBD. The power_state of the VM must be halted.
// Version: rio
func (vBD) AsyncSetMode3(session *Session, self VBDRef, value VbdMode) (retval TaskRef, err error) {
	method := "Async.VBD.set_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVbdMode(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// AssertAttachable: Throws an error if this VBD could not be attached to this VM if the VM were running. Intended for debugging.
// Version: rio
func (vBD) AssertAttachable(session *Session, self VBDRef) (err error) {
	method := "VBD.assert_attachable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncAssertAttachable: Throws an error if this VBD could not be attached to this VM if the VM were running. Intended for debugging.
// Version: rio
func (vBD) AsyncAssertAttachable(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.assert_attachable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AssertAttachable2: Throws an error if this VBD could not be attached to this VM if the VM were running. Intended for debugging.
// Version: rio
func (vBD) AssertAttachable2(session *Session, self VBDRef) (err error) {
	method := "VBD.assert_attachable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncAssertAttachable2: Throws an error if this VBD could not be attached to this VM if the VM were running. Intended for debugging.
// Version: rio
func (vBD) AsyncAssertAttachable2(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.assert_attachable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// UnplugForce: Forcibly unplug the specified VBD
// Version: rio
func (vBD) UnplugForce(session *Session, self VBDRef) (err error) {
	method := "VBD.unplug_force"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncUnplugForce: Forcibly unplug the specified VBD
// Version: rio
func (vBD) AsyncUnplugForce(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.unplug_force"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// UnplugForce2: Forcibly unplug the specified VBD
// Version: rio
func (vBD) UnplugForce2(session *Session, self VBDRef) (err error) {
	method := "VBD.unplug_force"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncUnplugForce2: Forcibly unplug the specified VBD
// Version: rio
func (vBD) AsyncUnplugForce2(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.unplug_force"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Unplug: Hot-unplug the specified VBD, dynamically unattaching it from the running VM
// Version: rio
//
// Errors:
// DEVICE_DETACH_REJECTED - The VM rejected the attempt to detach the device.
// DEVICE_ALREADY_DETACHED - The device is not currently attached
func (vBD) Unplug(session *Session, self VBDRef) (err error) {
	method := "VBD.unplug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncUnplug: Hot-unplug the specified VBD, dynamically unattaching it from the running VM
// Version: rio
//
// Errors:
// DEVICE_DETACH_REJECTED - The VM rejected the attempt to detach the device.
// DEVICE_ALREADY_DETACHED - The device is not currently attached
func (vBD) AsyncUnplug(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.unplug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Unplug2: Hot-unplug the specified VBD, dynamically unattaching it from the running VM
// Version: rio
//
// Errors:
// DEVICE_DETACH_REJECTED - The VM rejected the attempt to detach the device.
// DEVICE_ALREADY_DETACHED - The device is not currently attached
func (vBD) Unplug2(session *Session, self VBDRef) (err error) {
	method := "VBD.unplug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncUnplug2: Hot-unplug the specified VBD, dynamically unattaching it from the running VM
// Version: rio
//
// Errors:
// DEVICE_DETACH_REJECTED - The VM rejected the attempt to detach the device.
// DEVICE_ALREADY_DETACHED - The device is not currently attached
func (vBD) AsyncUnplug2(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.unplug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Plug: Hotplug the specified VBD, dynamically attaching it to the running VM
// Version: rio
func (vBD) Plug(session *Session, self VBDRef) (err error) {
	method := "VBD.plug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPlug: Hotplug the specified VBD, dynamically attaching it to the running VM
// Version: rio
func (vBD) AsyncPlug(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.plug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Plug2: Hotplug the specified VBD, dynamically attaching it to the running VM
// Version: rio
func (vBD) Plug2(session *Session, self VBDRef) (err error) {
	method := "VBD.plug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPlug2: Hotplug the specified VBD, dynamically attaching it to the running VM
// Version: rio
func (vBD) AsyncPlug2(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.plug"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Insert: Insert new media into the device
// Version: rio
//
// Errors:
// VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
// VBD_NOT_EMPTY - Operation could not be performed because the drive is not empty
func (vBD) Insert(session *Session, vbd VBDRef, vdi VDIRef) (err error) {
	method := "VBD.insert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vbdArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "vbd"), vbd)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vbdArg, vdiArg)
	return
}

// AsyncInsert: Insert new media into the device
// Version: rio
//
// Errors:
// VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
// VBD_NOT_EMPTY - Operation could not be performed because the drive is not empty
func (vBD) AsyncInsert(session *Session, vbd VBDRef, vdi VDIRef) (retval TaskRef, err error) {
	method := "Async.VBD.insert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vbdArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "vbd"), vbd)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vbdArg, vdiArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Insert3: Insert new media into the device
// Version: rio
//
// Errors:
// VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
// VBD_NOT_EMPTY - Operation could not be performed because the drive is not empty
func (vBD) Insert3(session *Session, vbd VBDRef, vdi VDIRef) (err error) {
	method := "VBD.insert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vbdArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "vbd"), vbd)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vbdArg, vdiArg)
	return
}

// AsyncInsert3: Insert new media into the device
// Version: rio
//
// Errors:
// VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
// VBD_NOT_EMPTY - Operation could not be performed because the drive is not empty
func (vBD) AsyncInsert3(session *Session, vbd VBDRef, vdi VDIRef) (retval TaskRef, err error) {
	method := "Async.VBD.insert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vbdArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "vbd"), vbd)
	if err != nil {
		return
	}
	vdiArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "vdi"), vdi)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vbdArg, vdiArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Eject: Remove the media from the device and leave it empty
// Version: rio
//
// Errors:
// VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
// VBD_IS_EMPTY - Operation could not be performed because the drive is empty
func (vBD) Eject(session *Session, vbd VBDRef) (err error) {
	method := "VBD.eject"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vbdArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "vbd"), vbd)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vbdArg)
	return
}

// AsyncEject: Remove the media from the device and leave it empty
// Version: rio
//
// Errors:
// VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
// VBD_IS_EMPTY - Operation could not be performed because the drive is empty
func (vBD) AsyncEject(session *Session, vbd VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.eject"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vbdArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "vbd"), vbd)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vbdArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Eject2: Remove the media from the device and leave it empty
// Version: rio
//
// Errors:
// VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
// VBD_IS_EMPTY - Operation could not be performed because the drive is empty
func (vBD) Eject2(session *Session, vbd VBDRef) (err error) {
	method := "VBD.eject"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vbdArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "vbd"), vbd)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vbdArg)
	return
}

// AsyncEject2: Remove the media from the device and leave it empty
// Version: rio
//
// Errors:
// VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
// VBD_IS_EMPTY - Operation could not be performed because the drive is empty
func (vBD) AsyncEject2(session *Session, vbd VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.eject"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vbdArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "vbd"), vbd)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vbdArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// RemoveFromQosAlgorithmParams: Remove the given key and its corresponding value from the qos/algorithm_params field of the given VBD.  If the key is not in that Map, then do nothing.
// Version: rio
func (vBD) RemoveFromQosAlgorithmParams(session *Session, self VBDRef, key string) (err error) {
	method := "VBD.remove_from_qos_algorithm_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromQosAlgorithmParams3: Remove the given key and its corresponding value from the qos/algorithm_params field of the given VBD.  If the key is not in that Map, then do nothing.
// Version: rio
func (vBD) RemoveFromQosAlgorithmParams3(session *Session, self VBDRef, key string) (err error) {
	method := "VBD.remove_from_qos_algorithm_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToQosAlgorithmParams: Add the given key-value pair to the qos/algorithm_params field of the given VBD.
// Version: rio
func (vBD) AddToQosAlgorithmParams(session *Session, self VBDRef, key string, value string) (err error) {
	method := "VBD.add_to_qos_algorithm_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToQosAlgorithmParams4: Add the given key-value pair to the qos/algorithm_params field of the given VBD.
// Version: rio
func (vBD) AddToQosAlgorithmParams4(session *Session, self VBDRef, key string, value string) (err error) {
	method := "VBD.add_to_qos_algorithm_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetQosAlgorithmParams: Set the qos/algorithm_params field of the given VBD.
// Version: rio
func (vBD) SetQosAlgorithmParams(session *Session, self VBDRef, value map[string]string) (err error) {
	method := "VBD.set_qos_algorithm_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetQosAlgorithmParams3: Set the qos/algorithm_params field of the given VBD.
// Version: rio
func (vBD) SetQosAlgorithmParams3(session *Session, self VBDRef, value map[string]string) (err error) {
	method := "VBD.set_qos_algorithm_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetQosAlgorithmType: Set the qos/algorithm_type field of the given VBD.
// Version: rio
func (vBD) SetQosAlgorithmType(session *Session, self VBDRef, value string) (err error) {
	method := "VBD.set_qos_algorithm_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetQosAlgorithmType3: Set the qos/algorithm_type field of the given VBD.
// Version: rio
func (vBD) SetQosAlgorithmType3(session *Session, self VBDRef, value string) (err error) {
	method := "VBD.set_qos_algorithm_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VBD.  If the key is not in that Map, then do nothing.
// Version: rio
func (vBD) RemoveFromOtherConfig(session *Session, self VBDRef, key string) (err error) {
	method := "VBD.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig3: Remove the given key and its corresponding value from the other_config field of the given VBD.  If the key is not in that Map, then do nothing.
// Version: rio
func (vBD) RemoveFromOtherConfig3(session *Session, self VBDRef, key string) (err error) {
	method := "VBD.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VBD.
// Version: rio
func (vBD) AddToOtherConfig(session *Session, self VBDRef, key string, value string) (err error) {
	method := "VBD.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig4: Add the given key-value pair to the other_config field of the given VBD.
// Version: rio
func (vBD) AddToOtherConfig4(session *Session, self VBDRef, key string, value string) (err error) {
	method := "VBD.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given VBD.
// Version: rio
func (vBD) SetOtherConfig(session *Session, self VBDRef, value map[string]string) (err error) {
	method := "VBD.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig3: Set the other_config field of the given VBD.
// Version: rio
func (vBD) SetOtherConfig3(session *Session, self VBDRef, value map[string]string) (err error) {
	method := "VBD.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetUnpluggable: Set the unpluggable field of the given VBD.
// Version: miami
func (vBD) SetUnpluggable(session *Session, self VBDRef, value bool) (err error) {
	method := "VBD.set_unpluggable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetUnpluggable3: Set the unpluggable field of the given VBD.
// Version: miami
func (vBD) SetUnpluggable3(session *Session, self VBDRef, value bool) (err error) {
	method := "VBD.set_unpluggable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetUnpluggable2: Set the unpluggable field of the given VBD.
// Version: rio
func (vBD) SetUnpluggable2(session *Session, self VBDRef) (err error) {
	method := "VBD.set_unpluggable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// SetType: Set the type field of the given VBD.
// Version: rio
func (vBD) SetType(session *Session, self VBDRef, value VbdType) (err error) {
	method := "VBD.set_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVbdType(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetType3: Set the type field of the given VBD.
// Version: rio
func (vBD) SetType3(session *Session, self VBDRef, value VbdType) (err error) {
	method := "VBD.set_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVbdType(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetBootable: Set the bootable field of the given VBD.
// Version: rio
func (vBD) SetBootable(session *Session, self VBDRef, value bool) (err error) {
	method := "VBD.set_bootable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetBootable3: Set the bootable field of the given VBD.
// Version: rio
func (vBD) SetBootable3(session *Session, self VBDRef, value bool) (err error) {
	method := "VBD.set_bootable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetUserdevice: Set the userdevice field of the given VBD.
// Version: rio
func (vBD) SetUserdevice(session *Session, self VBDRef, value string) (err error) {
	method := "VBD.set_userdevice"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetUserdevice3: Set the userdevice field of the given VBD.
// Version: rio
func (vBD) SetUserdevice3(session *Session, self VBDRef, value string) (err error) {
	method := "VBD.set_userdevice"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMetrics: Get the metrics field of the given VBD.
// Version: rio
func (vBD) GetMetrics(session *Session, self VBDRef) (retval VBDMetricsRef, err error) {
	method := "VBD.get_metrics"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDMetricsRef(method+" -> ", result)
	return
}

// GetMetrics2: Get the metrics field of the given VBD.
// Version: rio
func (vBD) GetMetrics2(session *Session, self VBDRef) (retval VBDMetricsRef, err error) {
	method := "VBD.get_metrics"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDMetricsRef(method+" -> ", result)
	return
}

// GetQosSupportedAlgorithms: Get the qos/supported_algorithms field of the given VBD.
// Version: rio
func (vBD) GetQosSupportedAlgorithms(session *Session, self VBDRef) (retval []string, err error) {
	method := "VBD.get_qos_supported_algorithms"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetQosSupportedAlgorithms2: Get the qos/supported_algorithms field of the given VBD.
// Version: rio
func (vBD) GetQosSupportedAlgorithms2(session *Session, self VBDRef) (retval []string, err error) {
	method := "VBD.get_qos_supported_algorithms"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetQosAlgorithmParams: Get the qos/algorithm_params field of the given VBD.
// Version: rio
func (vBD) GetQosAlgorithmParams(session *Session, self VBDRef) (retval map[string]string, err error) {
	method := "VBD.get_qos_algorithm_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetQosAlgorithmParams2: Get the qos/algorithm_params field of the given VBD.
// Version: rio
func (vBD) GetQosAlgorithmParams2(session *Session, self VBDRef) (retval map[string]string, err error) {
	method := "VBD.get_qos_algorithm_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetQosAlgorithmType: Get the qos/algorithm_type field of the given VBD.
// Version: rio
func (vBD) GetQosAlgorithmType(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_qos_algorithm_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetQosAlgorithmType2: Get the qos/algorithm_type field of the given VBD.
// Version: rio
func (vBD) GetQosAlgorithmType2(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_qos_algorithm_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRuntimeProperties: Get the runtime_properties field of the given VBD.
// Version: rio
func (vBD) GetRuntimeProperties(session *Session, self VBDRef) (retval map[string]string, err error) {
	method := "VBD.get_runtime_properties"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRuntimeProperties2: Get the runtime_properties field of the given VBD.
// Version: rio
func (vBD) GetRuntimeProperties2(session *Session, self VBDRef) (retval map[string]string, err error) {
	method := "VBD.get_runtime_properties"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStatusDetail: Get the status_detail field of the given VBD.
// Version: rio
func (vBD) GetStatusDetail(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_status_detail"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStatusDetail2: Get the status_detail field of the given VBD.
// Version: rio
func (vBD) GetStatusDetail2(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_status_detail"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStatusCode: Get the status_code field of the given VBD.
// Version: rio
func (vBD) GetStatusCode(session *Session, self VBDRef) (retval int, err error) {
	method := "VBD.get_status_code"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStatusCode2: Get the status_code field of the given VBD.
// Version: rio
func (vBD) GetStatusCode2(session *Session, self VBDRef) (retval int, err error) {
	method := "VBD.get_status_code"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCurrentlyAttached: Get the currently_attached field of the given VBD.
// Version: rio
func (vBD) GetCurrentlyAttached(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_currently_attached"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCurrentlyAttached2: Get the currently_attached field of the given VBD.
// Version: rio
func (vBD) GetCurrentlyAttached2(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_currently_attached"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given VBD.
// Version: rio
func (vBD) GetOtherConfig(session *Session, self VBDRef) (retval map[string]string, err error) {
	method := "VBD.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig2: Get the other_config field of the given VBD.
// Version: rio
func (vBD) GetOtherConfig2(session *Session, self VBDRef) (retval map[string]string, err error) {
	method := "VBD.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEmpty: Get the empty field of the given VBD.
// Version: rio
func (vBD) GetEmpty(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_empty"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEmpty2: Get the empty field of the given VBD.
// Version: rio
func (vBD) GetEmpty2(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_empty"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStorageLock: Get the storage_lock field of the given VBD.
// Version: rio
func (vBD) GetStorageLock(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_storage_lock"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStorageLock2: Get the storage_lock field of the given VBD.
// Version: rio
func (vBD) GetStorageLock2(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_storage_lock"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUnpluggable: Get the unpluggable field of the given VBD.
// Version: rio
func (vBD) GetUnpluggable(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_unpluggable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUnpluggable2: Get the unpluggable field of the given VBD.
// Version: rio
func (vBD) GetUnpluggable2(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_unpluggable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetType: Get the type field of the given VBD.
// Version: rio
func (vBD) GetType(session *Session, self VBDRef) (retval VbdType, err error) {
	method := "VBD.get_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVbdType(method+" -> ", result)
	return
}

// GetType2: Get the type field of the given VBD.
// Version: rio
func (vBD) GetType2(session *Session, self VBDRef) (retval VbdType, err error) {
	method := "VBD.get_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVbdType(method+" -> ", result)
	return
}

// GetMode: Get the mode field of the given VBD.
// Version: rio
func (vBD) GetMode(session *Session, self VBDRef) (retval VbdMode, err error) {
	method := "VBD.get_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVbdMode(method+" -> ", result)
	return
}

// GetMode2: Get the mode field of the given VBD.
// Version: rio
func (vBD) GetMode2(session *Session, self VBDRef) (retval VbdMode, err error) {
	method := "VBD.get_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVbdMode(method+" -> ", result)
	return
}

// GetBootable: Get the bootable field of the given VBD.
// Version: rio
func (vBD) GetBootable(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_bootable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBootable2: Get the bootable field of the given VBD.
// Version: rio
func (vBD) GetBootable2(session *Session, self VBDRef) (retval bool, err error) {
	method := "VBD.get_bootable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUserdevice: Get the userdevice field of the given VBD.
// Version: rio
func (vBD) GetUserdevice(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_userdevice"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUserdevice2: Get the userdevice field of the given VBD.
// Version: rio
func (vBD) GetUserdevice2(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_userdevice"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDevice: Get the device field of the given VBD.
// Version: rio
func (vBD) GetDevice(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_device"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDevice2: Get the device field of the given VBD.
// Version: rio
func (vBD) GetDevice2(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_device"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVDI: Get the VDI field of the given VBD.
// Version: rio
func (vBD) GetVDI(session *Session, self VBDRef) (retval VDIRef, err error) {
	method := "VBD.get_VDI"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// GetVDI2: Get the VDI field of the given VBD.
// Version: rio
func (vBD) GetVDI2(session *Session, self VBDRef) (retval VDIRef, err error) {
	method := "VBD.get_VDI"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRef(method+" -> ", result)
	return
}

// GetVM: Get the VM field of the given VBD.
// Version: rio
func (vBD) GetVM(session *Session, self VBDRef) (retval VMRef, err error) {
	method := "VBD.get_VM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVM2: Get the VM field of the given VBD.
// Version: rio
func (vBD) GetVM2(session *Session, self VBDRef) (retval VMRef, err error) {
	method := "VBD.get_VM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCurrentOperations: Get the current_operations field of the given VBD.
// Version: rio
func (vBD) GetCurrentOperations(session *Session, self VBDRef) (retval map[string]VbdOperations, err error) {
	method := "VBD.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumVbdOperationsMap(method+" -> ", result)
	return
}

// GetCurrentOperations2: Get the current_operations field of the given VBD.
// Version: rio
func (vBD) GetCurrentOperations2(session *Session, self VBDRef) (retval map[string]VbdOperations, err error) {
	method := "VBD.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumVbdOperationsMap(method+" -> ", result)
	return
}

// GetAllowedOperations: Get the allowed_operations field of the given VBD.
// Version: rio
func (vBD) GetAllowedOperations(session *Session, self VBDRef) (retval []VbdOperations, err error) {
	method := "VBD.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVbdOperationsSet(method+" -> ", result)
	return
}

// GetAllowedOperations2: Get the allowed_operations field of the given VBD.
// Version: rio
func (vBD) GetAllowedOperations2(session *Session, self VBDRef) (retval []VbdOperations, err error) {
	method := "VBD.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVbdOperationsSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VBD.
// Version: rio
func (vBD) GetUUID(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given VBD.
// Version: rio
func (vBD) GetUUID2(session *Session, self VBDRef) (retval string, err error) {
	method := "VBD.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy: Destroy the specified VBD instance.
// Version: rio
func (vBD) Destroy(session *Session, self VBDRef) (err error) {
	method := "VBD.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified VBD instance.
// Version: rio
func (vBD) AsyncDestroy(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy2: Destroy the specified VBD instance.
// Version: rio
func (vBD) Destroy2(session *Session, self VBDRef) (err error) {
	method := "VBD.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy2: Destroy the specified VBD instance.
// Version: rio
func (vBD) AsyncDestroy2(session *Session, self VBDRef) (retval TaskRef, err error) {
	method := "Async.VBD.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Create a new VBD instance, and return its handle. The constructor args are: VM*, VDI*, device, userdevice*, bootable*, mode*, type*, unpluggable, empty*, other_config*, currently_attached, qos_algorithm_type*, qos_algorithm_params* (* = non-optional).
// Version: rio
func (vBD) Create(session *Session, args VBDRecord) (retval VBDRef, err error) {
	method := "VBD.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVBDRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new VBD instance, and return its handle. The constructor args are: VM*, VDI*, device, userdevice*, bootable*, mode*, type*, unpluggable, empty*, other_config*, currently_attached, qos_algorithm_type*, qos_algorithm_params* (* = non-optional).
// Version: rio
func (vBD) AsyncCreate(session *Session, args VBDRecord) (retval TaskRef, err error) {
	method := "Async.VBD.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVBDRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// Create2: Create a new VBD instance, and return its handle. The constructor args are: VM*, VDI*, device, userdevice*, bootable*, mode*, type*, unpluggable, empty*, other_config*, currently_attached, qos_algorithm_type*, qos_algorithm_params* (* = non-optional).
// Version: rio
func (vBD) Create2(session *Session, args VBDRecord) (retval VBDRef, err error) {
	method := "VBD.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVBDRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRef(method+" -> ", result)
	return
}

// AsyncCreate2: Create a new VBD instance, and return its handle. The constructor args are: VM*, VDI*, device, userdevice*, bootable*, mode*, type*, unpluggable, empty*, other_config*, currently_attached, qos_algorithm_type*, qos_algorithm_params* (* = non-optional).
// Version: rio
func (vBD) AsyncCreate2(session *Session, args VBDRecord) (retval TaskRef, err error) {
	method := "Async.VBD.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVBDRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// GetByUUID: Get a reference to the VBD instance with the specified UUID.
// Version: rio
func (vBD) GetByUUID(session *Session, uUID string) (retval VBDRef, err error) {
	method := "VBD.get_by_uuid"
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
	retval, err = deserializeVBDRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the VBD instance with the specified UUID.
// Version: rio
func (vBD) GetByUUID2(session *Session, uUID string) (retval VBDRef, err error) {
	method := "VBD.get_by_uuid"
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
	retval, err = deserializeVBDRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given VBD.
// Version: rio
func (vBD) GetRecord(session *Session, self VBDRef) (retval VBDRecord, err error) {
	method := "VBD.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given VBD.
// Version: rio
func (vBD) GetRecord2(session *Session, self VBDRef) (retval VBDRecord, err error) {
	method := "VBD.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVBDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRecord(method+" -> ", result)
	return
}
