package xenapi

import (
	"fmt"
)

type NetworkSriovRecord struct {
	// Unique identifier/object reference
	UUID string
	// The PIF that has SR-IOV enabled
	PhysicalPIF PIFRef
	// The logical PIF to connect to the SR-IOV network after enable SR-IOV on the physical PIF
	LogicalPIF PIFRef
	// Indicates whether the host need to be rebooted before SR-IOV is enabled on the physical PIF
	RequiresReboot bool
	// The mode for configure network sriov
	ConfigurationMode SriovConfigurationMode
}

type NetworkSriovRef string

// network-sriov which connects logical pif and physical pif
type networkSriov struct{}

var NetworkSriov networkSriov

// GetRecord: Get a record containing the current state of the given network_sriov.
func (networkSriov) GetRecord(session *Session, self NetworkSriovRef) (retval NetworkSriovRecord, err error) {
	method := "network_sriov.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeNetworkSriovRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the network_sriov instance with the specified UUID.
func (networkSriov) GetByUUID(session *Session, uUID string) (retval NetworkSriovRef, err error) {
	method := "network_sriov.get_by_uuid"
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
	retval, err = deserializeNetworkSriovRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given network_sriov.
func (networkSriov) GetUUID(session *Session, self NetworkSriovRef) (retval string, err error) {
	method := "network_sriov.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPhysicalPIF: Get the physical_PIF field of the given network_sriov.
func (networkSriov) GetPhysicalPIF(session *Session, self NetworkSriovRef) (retval PIFRef, err error) {
	method := "network_sriov.get_physical_PIF"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLogicalPIF: Get the logical_PIF field of the given network_sriov.
func (networkSriov) GetLogicalPIF(session *Session, self NetworkSriovRef) (retval PIFRef, err error) {
	method := "network_sriov.get_logical_PIF"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRequiresReboot: Get the requires_reboot field of the given network_sriov.
func (networkSriov) GetRequiresReboot(session *Session, self NetworkSriovRef) (retval bool, err error) {
	method := "network_sriov.get_requires_reboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetConfigurationMode: Get the configuration_mode field of the given network_sriov.
func (networkSriov) GetConfigurationMode(session *Session, self NetworkSriovRef) (retval SriovConfigurationMode, err error) {
	method := "network_sriov.get_configuration_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumSriovConfigurationMode(method+" -> ", result)
	return
}

// Create: Enable SR-IOV on the specific PIF. It will create a network-sriov based on the specific PIF and automatically create a logical PIF to connect the specific network.
func (networkSriov) Create(session *Session, pif PIFRef, network NetworkRef) (retval NetworkSriovRef, err error) {
	method := "network_sriov.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	pifArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "pif"), pif)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, pifArg, networkArg)
	if err != nil {
		return
	}
	retval, err = deserializeNetworkSriovRef(method+" -> ", result)
	return
}

// AsyncCreate: Enable SR-IOV on the specific PIF. It will create a network-sriov based on the specific PIF and automatically create a logical PIF to connect the specific network.
func (networkSriov) AsyncCreate(session *Session, pif PIFRef, network NetworkRef) (retval TaskRef, err error) {
	method := "Async.network_sriov.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	pifArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "pif"), pif)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, pifArg, networkArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Disable SR-IOV on the specific PIF. It will destroy the network-sriov and the logical PIF accordingly.
func (networkSriov) Destroy(session *Session, self NetworkSriovRef) (err error) {
	method := "network_sriov.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Disable SR-IOV on the specific PIF. It will destroy the network-sriov and the logical PIF accordingly.
func (networkSriov) AsyncDestroy(session *Session, self NetworkSriovRef) (retval TaskRef, err error) {
	method := "Async.network_sriov.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRemainingCapacity: Get the number of free SR-IOV VFs on the associated PIF
func (networkSriov) GetRemainingCapacity(session *Session, self NetworkSriovRef) (retval int, err error) {
	method := "network_sriov.get_remaining_capacity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncGetRemainingCapacity: Get the number of free SR-IOV VFs on the associated PIF
func (networkSriov) AsyncGetRemainingCapacity(session *Session, self NetworkSriovRef) (retval TaskRef, err error) {
	method := "Async.network_sriov.get_remaining_capacity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkSriovRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the network_sriovs known to the system.
func (networkSriov) GetAll(session *Session) (retval []NetworkSriovRef, err error) {
	method := "network_sriov.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeNetworkSriovRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of network_sriov references to network_sriov records for all network_sriovs known to the system.
func (networkSriov) GetAllRecords(session *Session) (retval map[NetworkSriovRef]NetworkSriovRecord, err error) {
	method := "network_sriov.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeNetworkSriovRefToNetworkSriovRecordMap(method+" -> ", result)
	return
}

