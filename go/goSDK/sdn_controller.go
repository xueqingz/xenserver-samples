package xenapi

import (
	"fmt"
)

type SDNControllerRecord struct {
	// Unique identifier/object reference
	UUID string
	// Protocol to connect with SDN controller
	Protocol SdnControllerProtocol
	// IP address of the controller
	Address string
	// TCP port of the controller
	Port int
}

type SDNControllerRef string

// Describes the SDN controller that is to connect with the pool
type sDNController struct{}

var SDNController sDNController

// GetAllRecords: Return a map of SDN_controller references to SDN_controller records for all SDN_controllers known to the system.
// Version: falcon
func (sDNController) GetAllRecords(session *Session) (retval map[SDNControllerRef]SDNControllerRecord, err error) {
	method := "SDN_controller.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSDNControllerRefToSDNControllerRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of SDN_controller references to SDN_controller records for all SDN_controllers known to the system.
// Version: falcon
func (sDNController) GetAllRecords1(session *Session) (retval map[SDNControllerRef]SDNControllerRecord, err error) {
	method := "SDN_controller.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSDNControllerRefToSDNControllerRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the SDN_controllers known to the system.
// Version: falcon
func (sDNController) GetAll(session *Session) (retval []SDNControllerRef, err error) {
	method := "SDN_controller.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSDNControllerRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the SDN_controllers known to the system.
// Version: falcon
func (sDNController) GetAll1(session *Session) (retval []SDNControllerRef, err error) {
	method := "SDN_controller.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSDNControllerRefSet(method+" -> ", result)
	return
}

// Forget: Remove the OVS manager of the pool and destroy the db record.
// Version: falcon
func (sDNController) Forget(session *Session, self SDNControllerRef) (err error) {
	method := "SDN_controller.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncForget: Remove the OVS manager of the pool and destroy the db record.
// Version: falcon
func (sDNController) AsyncForget(session *Session, self SDNControllerRef) (retval TaskRef, err error) {
	method := "Async.SDN_controller.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Forget2: Remove the OVS manager of the pool and destroy the db record.
// Version: falcon
func (sDNController) Forget2(session *Session, self SDNControllerRef) (err error) {
	method := "SDN_controller.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncForget2: Remove the OVS manager of the pool and destroy the db record.
// Version: falcon
func (sDNController) AsyncForget2(session *Session, self SDNControllerRef) (retval TaskRef, err error) {
	method := "Async.SDN_controller.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Introduce: Introduce an SDN controller to the pool.
// Version: falcon
func (sDNController) Introduce(session *Session, protocol SdnControllerProtocol, address string, port int) (retval SDNControllerRef, err error) {
	method := "SDN_controller.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	protocolArg, err := serializeEnumSdnControllerProtocol(fmt.Sprintf("%s(%s)", method, "protocol"), protocol)
	if err != nil {
		return
	}
	addressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "address"), address)
	if err != nil {
		return
	}
	portArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "port"), port)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, protocolArg, addressArg, portArg)
	if err != nil {
		return
	}
	retval, err = deserializeSDNControllerRef(method+" -> ", result)
	return
}

// AsyncIntroduce: Introduce an SDN controller to the pool.
// Version: falcon
func (sDNController) AsyncIntroduce(session *Session, protocol SdnControllerProtocol, address string, port int) (retval TaskRef, err error) {
	method := "Async.SDN_controller.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	protocolArg, err := serializeEnumSdnControllerProtocol(fmt.Sprintf("%s(%s)", method, "protocol"), protocol)
	if err != nil {
		return
	}
	addressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "address"), address)
	if err != nil {
		return
	}
	portArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "port"), port)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, protocolArg, addressArg, portArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Introduce4: Introduce an SDN controller to the pool.
// Version: falcon
func (sDNController) Introduce4(session *Session, protocol SdnControllerProtocol, address string, port int) (retval SDNControllerRef, err error) {
	method := "SDN_controller.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	protocolArg, err := serializeEnumSdnControllerProtocol(fmt.Sprintf("%s(%s)", method, "protocol"), protocol)
	if err != nil {
		return
	}
	addressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "address"), address)
	if err != nil {
		return
	}
	portArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "port"), port)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, protocolArg, addressArg, portArg)
	if err != nil {
		return
	}
	retval, err = deserializeSDNControllerRef(method+" -> ", result)
	return
}

// AsyncIntroduce4: Introduce an SDN controller to the pool.
// Version: falcon
func (sDNController) AsyncIntroduce4(session *Session, protocol SdnControllerProtocol, address string, port int) (retval TaskRef, err error) {
	method := "Async.SDN_controller.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	protocolArg, err := serializeEnumSdnControllerProtocol(fmt.Sprintf("%s(%s)", method, "protocol"), protocol)
	if err != nil {
		return
	}
	addressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "address"), address)
	if err != nil {
		return
	}
	portArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "port"), port)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, protocolArg, addressArg, portArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetPort: Get the port field of the given SDN_controller.
// Version: falcon
func (sDNController) GetPort(session *Session, self SDNControllerRef) (retval int, err error) {
	method := "SDN_controller.get_port"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPort2: Get the port field of the given SDN_controller.
// Version: falcon
func (sDNController) GetPort2(session *Session, self SDNControllerRef) (retval int, err error) {
	method := "SDN_controller.get_port"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAddress: Get the address field of the given SDN_controller.
// Version: falcon
func (sDNController) GetAddress(session *Session, self SDNControllerRef) (retval string, err error) {
	method := "SDN_controller.get_address"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAddress2: Get the address field of the given SDN_controller.
// Version: falcon
func (sDNController) GetAddress2(session *Session, self SDNControllerRef) (retval string, err error) {
	method := "SDN_controller.get_address"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetProtocol: Get the protocol field of the given SDN_controller.
// Version: falcon
func (sDNController) GetProtocol(session *Session, self SDNControllerRef) (retval SdnControllerProtocol, err error) {
	method := "SDN_controller.get_protocol"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumSdnControllerProtocol(method+" -> ", result)
	return
}

// GetProtocol2: Get the protocol field of the given SDN_controller.
// Version: falcon
func (sDNController) GetProtocol2(session *Session, self SDNControllerRef) (retval SdnControllerProtocol, err error) {
	method := "SDN_controller.get_protocol"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumSdnControllerProtocol(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given SDN_controller.
// Version: falcon
func (sDNController) GetUUID(session *Session, self SDNControllerRef) (retval string, err error) {
	method := "SDN_controller.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given SDN_controller.
// Version: falcon
func (sDNController) GetUUID2(session *Session, self SDNControllerRef) (retval string, err error) {
	method := "SDN_controller.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByUUID: Get a reference to the SDN_controller instance with the specified UUID.
// Version: falcon
func (sDNController) GetByUUID(session *Session, uUID string) (retval SDNControllerRef, err error) {
	method := "SDN_controller.get_by_uuid"
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
	retval, err = deserializeSDNControllerRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the SDN_controller instance with the specified UUID.
// Version: falcon
func (sDNController) GetByUUID2(session *Session, uUID string) (retval SDNControllerRef, err error) {
	method := "SDN_controller.get_by_uuid"
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
	retval, err = deserializeSDNControllerRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given SDN_controller.
// Version: falcon
func (sDNController) GetRecord(session *Session, self SDNControllerRef) (retval SDNControllerRecord, err error) {
	method := "SDN_controller.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSDNControllerRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given SDN_controller.
// Version: falcon
func (sDNController) GetRecord2(session *Session, self SDNControllerRef) (retval SDNControllerRecord, err error) {
	method := "SDN_controller.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSDNControllerRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSDNControllerRecord(method+" -> ", result)
	return
}
