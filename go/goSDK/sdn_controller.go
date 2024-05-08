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

// GetRecord: Get a record containing the current state of the given SDN_controller.
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

// GetByUUID: Get a reference to the SDN_controller instance with the specified UUID.
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

// GetUUID: Get the uuid field of the given SDN_controller.
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

// GetProtocol: Get the protocol field of the given SDN_controller.
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

// GetAddress: Get the address field of the given SDN_controller.
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

// GetPort: Get the port field of the given SDN_controller.
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

// Introduce: Introduce an SDN controller to the pool.
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

// Forget: Remove the OVS manager of the pool and destroy the db record.
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

// GetAll: Return a list of all the SDN_controllers known to the system.
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

// GetAllRecords: Return a map of SDN_controller references to SDN_controller records for all SDN_controllers known to the system.
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
