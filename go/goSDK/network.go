package xenapi

import (
	"fmt"
)

type NetworkRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []NetworkOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]NetworkOperations
	// list of connected vifs
	VIFs []VIFRef
	// list of connected pifs
	PIFs []PIFRef
	// MTU in octets
	MTU int
	// additional configuration
	OtherConfig map[string]string
	// name of the bridge corresponding to this network on the local host
	Bridge string
	// true if the bridge is managed by xapi
	Managed bool
	// Binary blobs associated with this network
	Blobs map[string]BlobRef
	// user-specified tags for categorization purposes
	Tags []string
	// The network will use this value to determine the behaviour of all VIFs where locking_mode = default
	DefaultLockingMode NetworkDefaultLockingMode
	// The IP addresses assigned to VIFs on networks that have active xapi-managed DHCP
	AssignedIps map[VIFRef]string
	// Set of purposes for which the server will use this network
	Purpose []NetworkPurpose
}

type NetworkRef string

// A virtual network
type network struct{}

var Network network

// GetRecord: Get a record containing the current state of the given network.
func (network) GetRecord(session *Session, self NetworkRef) (retval NetworkRecord, err error) {
	method := "network.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeNetworkRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the network instance with the specified UUID.
func (network) GetByUUID(session *Session, uUID string) (retval NetworkRef, err error) {
	method := "network.get_by_uuid"
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
	retval, err = deserializeNetworkRef(method+" -> ", result)
	return
}

// Create: Create a new network instance, and return its handle. The constructor args are: name_label, name_description, MTU, other_config*, bridge, managed, tags (* = non-optional).
func (network) Create(session *Session, args NetworkRecord) (retval NetworkRef, err error) {
	method := "network.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeNetworkRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeNetworkRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new network instance, and return its handle. The constructor args are: name_label, name_description, MTU, other_config*, bridge, managed, tags (* = non-optional).
func (network) AsyncCreate(session *Session, args NetworkRecord) (retval TaskRef, err error) {
	method := "Async.network.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeNetworkRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// Destroy: Destroy the specified network instance.
func (network) Destroy(session *Session, self NetworkRef) (err error) {
	method := "network.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified network instance.
func (network) AsyncDestroy(session *Session, self NetworkRef) (retval TaskRef, err error) {
	method := "Async.network.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByNameLabel: Get all the network instances with the given label.
func (network) GetByNameLabel(session *Session, label string) (retval []NetworkRef, err error) {
	method := "network.get_by_name_label"
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
	retval, err = deserializeNetworkRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given network.
func (network) GetUUID(session *Session, self NetworkRef) (retval string, err error) {
	method := "network.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given network.
func (network) GetNameLabel(session *Session, self NetworkRef) (retval string, err error) {
	method := "network.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given network.
func (network) GetNameDescription(session *Session, self NetworkRef) (retval string, err error) {
	method := "network.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given network.
func (network) GetAllowedOperations(session *Session, self NetworkRef) (retval []NetworkOperations, err error) {
	method := "network.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumNetworkOperationsSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given network.
func (network) GetCurrentOperations(session *Session, self NetworkRef) (retval map[string]NetworkOperations, err error) {
	method := "network.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumNetworkOperationsMap(method+" -> ", result)
	return
}

// GetVIFs: Get the VIFs field of the given network.
func (network) GetVIFs(session *Session, self NetworkRef) (retval []VIFRef, err error) {
	method := "network.get_VIFs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVIFRefSet(method+" -> ", result)
	return
}

// GetPIFs: Get the PIFs field of the given network.
func (network) GetPIFs(session *Session, self NetworkRef) (retval []PIFRef, err error) {
	method := "network.get_PIFs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePIFRefSet(method+" -> ", result)
	return
}

// GetMTU: Get the MTU field of the given network.
func (network) GetMTU(session *Session, self NetworkRef) (retval int, err error) {
	method := "network.get_MTU"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given network.
func (network) GetOtherConfig(session *Session, self NetworkRef) (retval map[string]string, err error) {
	method := "network.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBridge: Get the bridge field of the given network.
func (network) GetBridge(session *Session, self NetworkRef) (retval string, err error) {
	method := "network.get_bridge"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetManaged: Get the managed field of the given network.
func (network) GetManaged(session *Session, self NetworkRef) (retval bool, err error) {
	method := "network.get_managed"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBlobs: Get the blobs field of the given network.
func (network) GetBlobs(session *Session, self NetworkRef) (retval map[string]BlobRef, err error) {
	method := "network.get_blobs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToBlobRefMap(method+" -> ", result)
	return
}

// GetTags: Get the tags field of the given network.
func (network) GetTags(session *Session, self NetworkRef) (retval []string, err error) {
	method := "network.get_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDefaultLockingMode: Get the default_locking_mode field of the given network.
func (network) GetDefaultLockingMode(session *Session, self NetworkRef) (retval NetworkDefaultLockingMode, err error) {
	method := "network.get_default_locking_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumNetworkDefaultLockingMode(method+" -> ", result)
	return
}

// GetAssignedIps: Get the assigned_ips field of the given network.
func (network) GetAssignedIps(session *Session, self NetworkRef) (retval map[VIFRef]string, err error) {
	method := "network.get_assigned_ips"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVIFRefToStringMap(method+" -> ", result)
	return
}

// GetPurpose: Get the purpose field of the given network.
func (network) GetPurpose(session *Session, self NetworkRef) (retval []NetworkPurpose, err error) {
	method := "network.get_purpose"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumNetworkPurposeSet(method+" -> ", result)
	return
}

// SetNameLabel: Set the name/label field of the given network.
func (network) SetNameLabel(session *Session, self NetworkRef, value string) (err error) {
	method := "network.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription: Set the name/description field of the given network.
func (network) SetNameDescription(session *Session, self NetworkRef, value string) (err error) {
	method := "network.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetMTU: Set the MTU field of the given network.
func (network) SetMTU(session *Session, self NetworkRef, value int) (err error) {
	method := "network.set_MTU"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetOtherConfig: Set the other_config field of the given network.
func (network) SetOtherConfig(session *Session, self NetworkRef, value map[string]string) (err error) {
	method := "network.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given network.
func (network) AddToOtherConfig(session *Session, self NetworkRef, key string, value string) (err error) {
	method := "network.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given network.  If the key is not in that Map, then do nothing.
func (network) RemoveFromOtherConfig(session *Session, self NetworkRef, key string) (err error) {
	method := "network.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetTags: Set the tags field of the given network.
func (network) SetTags(session *Session, self NetworkRef, value []string) (err error) {
	method := "network.set_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AddTags: Add the given value to the tags field of the given network.  If the value is already in that Set, then do nothing.
func (network) AddTags(session *Session, self NetworkRef, value string) (err error) {
	method := "network.add_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveTags: Remove the given value from the tags field of the given network.  If the value is not in that Set, then do nothing.
func (network) RemoveTags(session *Session, self NetworkRef, value string) (err error) {
	method := "network.remove_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this pool
func (network) CreateNewBlob(session *Session, network NetworkRef, name string, mimeType string, public bool) (retval BlobRef, err error) {
	method := "network.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	mimeTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "mime_type"), mimeType)
	if err != nil {
		return
	}
	publicArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "public"), public)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, networkArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeBlobRef(method+" -> ", result)
	return
}

// AsyncCreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this pool
func (network) AsyncCreateNewBlob(session *Session, network NetworkRef, name string, mimeType string, public bool) (retval TaskRef, err error) {
	method := "Async.network.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	mimeTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "mime_type"), mimeType)
	if err != nil {
		return
	}
	publicArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "public"), public)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, networkArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetDefaultLockingMode: Set the default locking mode for VIFs attached to this network
func (network) SetDefaultLockingMode(session *Session, network NetworkRef, value NetworkDefaultLockingMode) (err error) {
	method := "network.set_default_locking_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumNetworkDefaultLockingMode(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, networkArg, valueArg)
	return
}

// AsyncSetDefaultLockingMode: Set the default locking mode for VIFs attached to this network
func (network) AsyncSetDefaultLockingMode(session *Session, network NetworkRef, value NetworkDefaultLockingMode) (retval TaskRef, err error) {
	method := "Async.network.set_default_locking_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumNetworkDefaultLockingMode(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, networkArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// AddPurpose: Give a network a new purpose (if not present already)
//
// Errors:
// NETWORK_INCOMPATIBLE_PURPOSES - You tried to add a purpose to a network but the new purpose is not compatible with an existing purpose of the network or other networks.
func (network) AddPurpose(session *Session, self NetworkRef, value NetworkPurpose) (err error) {
	method := "network.add_purpose"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumNetworkPurpose(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncAddPurpose: Give a network a new purpose (if not present already)
//
// Errors:
// NETWORK_INCOMPATIBLE_PURPOSES - You tried to add a purpose to a network but the new purpose is not compatible with an existing purpose of the network or other networks.
func (network) AsyncAddPurpose(session *Session, self NetworkRef, value NetworkPurpose) (retval TaskRef, err error) {
	method := "Async.network.add_purpose"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumNetworkPurpose(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// RemovePurpose: Remove a purpose from a network (if present)
func (network) RemovePurpose(session *Session, self NetworkRef, value NetworkPurpose) (err error) {
	method := "network.remove_purpose"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumNetworkPurpose(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncRemovePurpose: Remove a purpose from a network (if present)
func (network) AsyncRemovePurpose(session *Session, self NetworkRef, value NetworkPurpose) (retval TaskRef, err error) {
	method := "Async.network.remove_purpose"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumNetworkPurpose(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// GetAll: Return a list of all the networks known to the system.
func (network) GetAll(session *Session) (retval []NetworkRef, err error) {
	method := "network.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeNetworkRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of network references to network records for all networks known to the system.
func (network) GetAllRecords(session *Session) (retval map[NetworkRef]NetworkRecord, err error) {
	method := "network.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeNetworkRefToNetworkRecordMap(method+" -> ", result)
	return
}

