package xenapi

import (
	"fmt"
)

type LVHDRecord struct {
	// Unique identifier/object reference
	UUID string
}

type LVHDRef string

// LVHD SR specific operations
type lVHD struct{}

var LVHD lVHD

// GetRecord: Get a record containing the current state of the given LVHD.
func (lVHD) GetRecord(session *Session, self LVHDRef) (retval LVHDRecord, err error) {
	method := "LVHD.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeLVHDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeLVHDRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the LVHD instance with the specified UUID.
func (lVHD) GetByUUID(session *Session, uUID string) (retval LVHDRef, err error) {
	method := "LVHD.get_by_uuid"
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
	retval, err = deserializeLVHDRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given LVHD.
func (lVHD) GetUUID(session *Session, self LVHDRef) (retval string, err error) {
	method := "LVHD.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeLVHDRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// EnableThinProvisioning: Upgrades an LVHD SR to enable thin-provisioning. Future VDIs created in this SR will be thinly-provisioned, although existing VDIs will be left alone. Note that the SR must be attached to the SRmaster for upgrade to work.
func (lVHD) EnableThinProvisioning(session *Session, host HostRef, sR SRRef, initialAllocation int, allocationQuantum int) (retval string, err error) {
	method := "LVHD.enable_thin_provisioning"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	sRArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "SR"), sR)
	if err != nil {
		return
	}
	initialAllocationArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "initial_allocation"), initialAllocation)
	if err != nil {
		return
	}
	allocationQuantumArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "allocation_quantum"), allocationQuantum)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, sRArg, initialAllocationArg, allocationQuantumArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncEnableThinProvisioning: Upgrades an LVHD SR to enable thin-provisioning. Future VDIs created in this SR will be thinly-provisioned, although existing VDIs will be left alone. Note that the SR must be attached to the SRmaster for upgrade to work.
func (lVHD) AsyncEnableThinProvisioning(session *Session, host HostRef, sR SRRef, initialAllocation int, allocationQuantum int) (retval TaskRef, err error) {
	method := "Async.LVHD.enable_thin_provisioning"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	sRArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "SR"), sR)
	if err != nil {
		return
	}
	initialAllocationArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "initial_allocation"), initialAllocation)
	if err != nil {
		return
	}
	allocationQuantumArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "allocation_quantum"), allocationQuantum)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, sRArg, initialAllocationArg, allocationQuantumArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

