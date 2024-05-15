package xenapi

import (
	"fmt"
)

type PCIRecord struct {
	// Unique identifier/object reference
	UUID string
	// PCI class name
	ClassName string
	// Vendor name
	VendorName string
	// Device name
	DeviceName string
	// Physical machine that owns the PCI device
	Host HostRef
	// PCI ID of the physical device
	PciID string
	// List of dependent PCI devices
	Dependencies []PCIRef
	// Additional configuration
	OtherConfig map[string]string
	// Subsystem vendor name
	SubsystemVendorName string
	// Subsystem device name
	SubsystemDeviceName string
	// Driver name
	DriverName string
}

type PCIRef string

// A PCI device
type pCI struct{}

var PCI pCI

// GetAllRecords: Return a map of PCI references to PCI records for all PCIs known to the system.
// Version: boston
func (pCI) GetAllRecords(session *Session) (retval map[PCIRef]PCIRecord, err error) {
	method := "PCI.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePCIRefToPCIRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of PCI references to PCI records for all PCIs known to the system.
// Version: boston
func (pCI) GetAllRecords1(session *Session) (retval map[PCIRef]PCIRecord, err error) {
	method := "PCI.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePCIRefToPCIRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the PCIs known to the system.
// Version: boston
func (pCI) GetAll(session *Session) (retval []PCIRef, err error) {
	method := "PCI.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePCIRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the PCIs known to the system.
// Version: boston
func (pCI) GetAll1(session *Session) (retval []PCIRef, err error) {
	method := "PCI.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePCIRefSet(method+" -> ", result)
	return
}

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PCI.  If the key is not in that Map, then do nothing.
// Version: boston
func (pCI) RemoveFromOtherConfig(session *Session, self PCIRef, key string) (err error) {
	method := "PCI.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig3: Remove the given key and its corresponding value from the other_config field of the given PCI.  If the key is not in that Map, then do nothing.
// Version: boston
func (pCI) RemoveFromOtherConfig3(session *Session, self PCIRef, key string) (err error) {
	method := "PCI.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given PCI.
// Version: boston
func (pCI) AddToOtherConfig(session *Session, self PCIRef, key string, value string) (err error) {
	method := "PCI.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig4: Add the given key-value pair to the other_config field of the given PCI.
// Version: boston
func (pCI) AddToOtherConfig4(session *Session, self PCIRef, key string, value string) (err error) {
	method := "PCI.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given PCI.
// Version: boston
func (pCI) SetOtherConfig(session *Session, self PCIRef, value map[string]string) (err error) {
	method := "PCI.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig3: Set the other_config field of the given PCI.
// Version: boston
func (pCI) SetOtherConfig3(session *Session, self PCIRef, value map[string]string) (err error) {
	method := "PCI.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDriverName: Get the driver_name field of the given PCI.
// Version: boston
func (pCI) GetDriverName(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_driver_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDriverName2: Get the driver_name field of the given PCI.
// Version: boston
func (pCI) GetDriverName2(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_driver_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSubsystemDeviceName: Get the subsystem_device_name field of the given PCI.
// Version: boston
func (pCI) GetSubsystemDeviceName(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_subsystem_device_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSubsystemDeviceName2: Get the subsystem_device_name field of the given PCI.
// Version: boston
func (pCI) GetSubsystemDeviceName2(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_subsystem_device_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSubsystemVendorName: Get the subsystem_vendor_name field of the given PCI.
// Version: boston
func (pCI) GetSubsystemVendorName(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_subsystem_vendor_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSubsystemVendorName2: Get the subsystem_vendor_name field of the given PCI.
// Version: boston
func (pCI) GetSubsystemVendorName2(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_subsystem_vendor_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given PCI.
// Version: boston
func (pCI) GetOtherConfig(session *Session, self PCIRef) (retval map[string]string, err error) {
	method := "PCI.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig2: Get the other_config field of the given PCI.
// Version: boston
func (pCI) GetOtherConfig2(session *Session, self PCIRef) (retval map[string]string, err error) {
	method := "PCI.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDependencies: Get the dependencies field of the given PCI.
// Version: boston
func (pCI) GetDependencies(session *Session, self PCIRef) (retval []PCIRef, err error) {
	method := "PCI.get_dependencies"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePCIRefSet(method+" -> ", result)
	return
}

// GetDependencies2: Get the dependencies field of the given PCI.
// Version: boston
func (pCI) GetDependencies2(session *Session, self PCIRef) (retval []PCIRef, err error) {
	method := "PCI.get_dependencies"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePCIRefSet(method+" -> ", result)
	return
}

// GetPciID: Get the pci_id field of the given PCI.
// Version: boston
func (pCI) GetPciID(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_pci_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPciID2: Get the pci_id field of the given PCI.
// Version: boston
func (pCI) GetPciID2(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_pci_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost: Get the host field of the given PCI.
// Version: boston
func (pCI) GetHost(session *Session, self PCIRef) (retval HostRef, err error) {
	method := "PCI.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost2: Get the host field of the given PCI.
// Version: boston
func (pCI) GetHost2(session *Session, self PCIRef) (retval HostRef, err error) {
	method := "PCI.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDeviceName: Get the device_name field of the given PCI.
// Version: boston
func (pCI) GetDeviceName(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_device_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDeviceName2: Get the device_name field of the given PCI.
// Version: boston
func (pCI) GetDeviceName2(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_device_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVendorName: Get the vendor_name field of the given PCI.
// Version: boston
func (pCI) GetVendorName(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_vendor_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVendorName2: Get the vendor_name field of the given PCI.
// Version: boston
func (pCI) GetVendorName2(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_vendor_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetClassName: Get the class_name field of the given PCI.
// Version: boston
func (pCI) GetClassName(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_class_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetClassName2: Get the class_name field of the given PCI.
// Version: boston
func (pCI) GetClassName2(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_class_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given PCI.
// Version: boston
func (pCI) GetUUID(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given PCI.
// Version: boston
func (pCI) GetUUID2(session *Session, self PCIRef) (retval string, err error) {
	method := "PCI.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByUUID: Get a reference to the PCI instance with the specified UUID.
// Version: boston
func (pCI) GetByUUID(session *Session, uUID string) (retval PCIRef, err error) {
	method := "PCI.get_by_uuid"
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
	retval, err = deserializePCIRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the PCI instance with the specified UUID.
// Version: boston
func (pCI) GetByUUID2(session *Session, uUID string) (retval PCIRef, err error) {
	method := "PCI.get_by_uuid"
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
	retval, err = deserializePCIRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given PCI.
// Version: boston
func (pCI) GetRecord(session *Session, self PCIRef) (retval PCIRecord, err error) {
	method := "PCI.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePCIRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given PCI.
// Version: boston
func (pCI) GetRecord2(session *Session, self PCIRef) (retval PCIRecord, err error) {
	method := "PCI.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePCIRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePCIRecord(method+" -> ", result)
	return
}
