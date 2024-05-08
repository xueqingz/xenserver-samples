package xenapi

import (
	"fmt"
)

type TunnelRecord struct {
	// Unique identifier/object reference
	UUID string
	// The interface through which the tunnel is accessed
	AccessPIF PIFRef
	// The interface used by the tunnel
	TransportPIF PIFRef
	// Status information about the tunnel
	Status map[string]string
	// Additional configuration
	OtherConfig map[string]string
	// The protocol used for tunneling (either GRE or VxLAN)
	Protocol TunnelProtocol
}

type TunnelRef string

// A tunnel for network traffic
type tunnel struct{}

var Tunnel tunnel

// GetRecord: Get a record containing the current state of the given tunnel.
func (tunnel) GetRecord(session *Session, self TunnelRef) (retval TunnelRecord, err error) {
	method := "tunnel.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTunnelRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the tunnel instance with the specified UUID.
func (tunnel) GetByUUID(session *Session, uUID string) (retval TunnelRef, err error) {
	method := "tunnel.get_by_uuid"
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
	retval, err = deserializeTunnelRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given tunnel.
func (tunnel) GetUUID(session *Session, self TunnelRef) (retval string, err error) {
	method := "tunnel.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAccessPIF: Get the access_PIF field of the given tunnel.
func (tunnel) GetAccessPIF(session *Session, self TunnelRef) (retval PIFRef, err error) {
	method := "tunnel.get_access_PIF"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTransportPIF: Get the transport_PIF field of the given tunnel.
func (tunnel) GetTransportPIF(session *Session, self TunnelRef) (retval PIFRef, err error) {
	method := "tunnel.get_transport_PIF"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStatus: Get the status field of the given tunnel.
func (tunnel) GetStatus(session *Session, self TunnelRef) (retval map[string]string, err error) {
	method := "tunnel.get_status"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given tunnel.
func (tunnel) GetOtherConfig(session *Session, self TunnelRef) (retval map[string]string, err error) {
	method := "tunnel.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetProtocol: Get the protocol field of the given tunnel.
func (tunnel) GetProtocol(session *Session, self TunnelRef) (retval TunnelProtocol, err error) {
	method := "tunnel.get_protocol"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumTunnelProtocol(method+" -> ", result)
	return
}

// SetStatus: Set the status field of the given tunnel.
func (tunnel) SetStatus(session *Session, self TunnelRef, value map[string]string) (err error) {
	method := "tunnel.set_status"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToStatus: Add the given key-value pair to the status field of the given tunnel.
func (tunnel) AddToStatus(session *Session, self TunnelRef, key string, value string) (err error) {
	method := "tunnel.add_to_status"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromStatus: Remove the given key and its corresponding value from the status field of the given tunnel.  If the key is not in that Map, then do nothing.
func (tunnel) RemoveFromStatus(session *Session, self TunnelRef, key string) (err error) {
	method := "tunnel.remove_from_status"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given tunnel.
func (tunnel) SetOtherConfig(session *Session, self TunnelRef, value map[string]string) (err error) {
	method := "tunnel.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given tunnel.
func (tunnel) AddToOtherConfig(session *Session, self TunnelRef, key string, value string) (err error) {
	method := "tunnel.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given tunnel.  If the key is not in that Map, then do nothing.
func (tunnel) RemoveFromOtherConfig(session *Session, self TunnelRef, key string) (err error) {
	method := "tunnel.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetProtocol: Set the protocol field of the given tunnel.
func (tunnel) SetProtocol(session *Session, self TunnelRef, value TunnelProtocol) (err error) {
	method := "tunnel.set_protocol"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumTunnelProtocol(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// Create: Create a tunnel
//
// Errors:
// OPENVSWITCH_NOT_ACTIVE - This operation needs the OpenVSwitch networking backend to be enabled on all hosts in the pool.
// TRANSPORT_PIF_NOT_CONFIGURED - The tunnel transport PIF has no IP configuration set.
// IS_TUNNEL_ACCESS_PIF - Cannot create a VLAN or tunnel on top of a tunnel access PIF - use the underlying transport PIF instead.
func (tunnel) Create(session *Session, transportPIF PIFRef, network NetworkRef, protocol TunnelProtocol) (retval TunnelRef, err error) {
	method := "tunnel.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	transportPIFArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "transport_PIF"), transportPIF)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	protocolArg, err := serializeEnumTunnelProtocol(fmt.Sprintf("%s(%s)", method, "protocol"), protocol)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, transportPIFArg, networkArg, protocolArg)
	if err != nil {
		return
	}
	retval, err = deserializeTunnelRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a tunnel
//
// Errors:
// OPENVSWITCH_NOT_ACTIVE - This operation needs the OpenVSwitch networking backend to be enabled on all hosts in the pool.
// TRANSPORT_PIF_NOT_CONFIGURED - The tunnel transport PIF has no IP configuration set.
// IS_TUNNEL_ACCESS_PIF - Cannot create a VLAN or tunnel on top of a tunnel access PIF - use the underlying transport PIF instead.
func (tunnel) AsyncCreate(session *Session, transportPIF PIFRef, network NetworkRef, protocol TunnelProtocol) (retval TaskRef, err error) {
	method := "Async.tunnel.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	transportPIFArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "transport_PIF"), transportPIF)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	protocolArg, err := serializeEnumTunnelProtocol(fmt.Sprintf("%s(%s)", method, "protocol"), protocol)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, transportPIFArg, networkArg, protocolArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Destroy a tunnel
func (tunnel) Destroy(session *Session, self TunnelRef) (err error) {
	method := "tunnel.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy a tunnel
func (tunnel) AsyncDestroy(session *Session, self TunnelRef) (retval TaskRef, err error) {
	method := "Async.tunnel.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTunnelRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the tunnels known to the system.
func (tunnel) GetAll(session *Session) (retval []TunnelRef, err error) {
	method := "tunnel.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTunnelRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of tunnel references to tunnel records for all tunnels known to the system.
func (tunnel) GetAllRecords(session *Session) (retval map[TunnelRef]TunnelRecord, err error) {
	method := "tunnel.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTunnelRefToTunnelRecordMap(method+" -> ", result)
	return
}

