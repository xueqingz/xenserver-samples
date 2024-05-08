package xenapi

import (
	"fmt"
	"time"
)

type ClusterHostRecord struct {
	// Unique identifier/object reference
	UUID string
	// Reference to the Cluster object
	Cluster ClusterRef
	// Reference to the Host object
	Host HostRef
	// Whether the cluster host believes that clustering should be enabled on this host. This field can be altered by calling the enable/disable message on a cluster host. Only enabled members run the underlying cluster stack. Disabled members are still considered a member of the cluster (see joined), and can be re-enabled by the user.
	Enabled bool
	// Reference to the PIF object
	PIF PIFRef
	// Whether the cluster host has joined the cluster. Contrary to enabled, a host that is not joined is not considered a member of the cluster, and hence enable and disable operations cannot be performed on this host.
	Joined bool
	// Whether the underlying cluster stack thinks we are live. This field is set automatically based on updates from the cluster stack and cannot be altered by the user.
	Live bool
	// Time when the live field was last updated based on information from the cluster stack
	LastUpdateLive time.Time
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []ClusterHostOperation
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]ClusterHostOperation
	// Additional configuration
	OtherConfig map[string]string
}

type ClusterHostRef string

// Cluster member metadata
type clusterHost struct{}

var ClusterHost clusterHost

// GetRecord: Get a record containing the current state of the given Cluster_host.
func (clusterHost) GetRecord(session *Session, self ClusterHostRef) (retval ClusterHostRecord, err error) {
	method := "Cluster_host.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterHostRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the Cluster_host instance with the specified UUID.
func (clusterHost) GetByUUID(session *Session, uUID string) (retval ClusterHostRef, err error) {
	method := "Cluster_host.get_by_uuid"
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
	retval, err = deserializeClusterHostRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given Cluster_host.
func (clusterHost) GetUUID(session *Session, self ClusterHostRef) (retval string, err error) {
	method := "Cluster_host.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCluster: Get the cluster field of the given Cluster_host.
func (clusterHost) GetCluster(session *Session, self ClusterHostRef) (retval ClusterRef, err error) {
	method := "Cluster_host.get_cluster"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterRef(method+" -> ", result)
	return
}

// GetHost: Get the host field of the given Cluster_host.
func (clusterHost) GetHost(session *Session, self ClusterHostRef) (retval HostRef, err error) {
	method := "Cluster_host.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnabled: Get the enabled field of the given Cluster_host.
func (clusterHost) GetEnabled(session *Session, self ClusterHostRef) (retval bool, err error) {
	method := "Cluster_host.get_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPIF: Get the PIF field of the given Cluster_host.
func (clusterHost) GetPIF(session *Session, self ClusterHostRef) (retval PIFRef, err error) {
	method := "Cluster_host.get_PIF"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetJoined: Get the joined field of the given Cluster_host.
func (clusterHost) GetJoined(session *Session, self ClusterHostRef) (retval bool, err error) {
	method := "Cluster_host.get_joined"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLive: Get the live field of the given Cluster_host.
func (clusterHost) GetLive(session *Session, self ClusterHostRef) (retval bool, err error) {
	method := "Cluster_host.get_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastUpdateLive: Get the last_update_live field of the given Cluster_host.
func (clusterHost) GetLastUpdateLive(session *Session, self ClusterHostRef) (retval time.Time, err error) {
	method := "Cluster_host.get_last_update_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTime(method+" -> ", result)
	return
}

// GetAllowedOperations: Get the allowed_operations field of the given Cluster_host.
func (clusterHost) GetAllowedOperations(session *Session, self ClusterHostRef) (retval []ClusterHostOperation, err error) {
	method := "Cluster_host.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumClusterHostOperationSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given Cluster_host.
func (clusterHost) GetCurrentOperations(session *Session, self ClusterHostRef) (retval map[string]ClusterHostOperation, err error) {
	method := "Cluster_host.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumClusterHostOperationMap(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given Cluster_host.
func (clusterHost) GetOtherConfig(session *Session, self ClusterHostRef) (retval map[string]string, err error) {
	method := "Cluster_host.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Add a new host to an existing cluster.
//
// Errors:
// PIF_NOT_ATTACHED_TO_HOST - Cluster_host creation failed as the PIF provided is not attached to the host.
// NO_CLUSTER_HOSTS_REACHABLE - No other cluster host was reachable when joining
func (clusterHost) Create(session *Session, cluster ClusterRef, host HostRef, pif PIFRef) (retval ClusterHostRef, err error) {
	method := "Cluster_host.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	clusterArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "cluster"), cluster)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	pifArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "pif"), pif)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, clusterArg, hostArg, pifArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterHostRef(method+" -> ", result)
	return
}

// AsyncCreate: Add a new host to an existing cluster.
//
// Errors:
// PIF_NOT_ATTACHED_TO_HOST - Cluster_host creation failed as the PIF provided is not attached to the host.
// NO_CLUSTER_HOSTS_REACHABLE - No other cluster host was reachable when joining
func (clusterHost) AsyncCreate(session *Session, cluster ClusterRef, host HostRef, pif PIFRef) (retval TaskRef, err error) {
	method := "Async.Cluster_host.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	clusterArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "cluster"), cluster)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	pifArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "pif"), pif)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, clusterArg, hostArg, pifArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Remove the host from an existing cluster. This operation is allowed even if a cluster host is not enabled.
//
// Errors:
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
// CLUSTERING_DISABLED - An operation was attempted while clustering was disabled on the cluster_host.
// CLUSTER_HOST_IS_LAST - The last cluster host cannot be destroyed. Destroy the cluster instead
func (clusterHost) Destroy(session *Session, self ClusterHostRef) (err error) {
	method := "Cluster_host.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Remove the host from an existing cluster. This operation is allowed even if a cluster host is not enabled.
//
// Errors:
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
// CLUSTERING_DISABLED - An operation was attempted while clustering was disabled on the cluster_host.
// CLUSTER_HOST_IS_LAST - The last cluster host cannot be destroyed. Destroy the cluster instead
func (clusterHost) AsyncDestroy(session *Session, self ClusterHostRef) (retval TaskRef, err error) {
	method := "Async.Cluster_host.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Enable: Enable cluster membership for a disabled cluster host.
//
// Errors:
// PIF_ALLOWS_UNPLUG - The operation you requested cannot be performed because the specified PIF allows unplug.
// REQUIRED_PIF_IS_UNPLUGGED - The operation you requested cannot be performed because the specified PIF is currently unplugged.
func (clusterHost) Enable(session *Session, self ClusterHostRef) (err error) {
	method := "Cluster_host.enable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncEnable: Enable cluster membership for a disabled cluster host.
//
// Errors:
// PIF_ALLOWS_UNPLUG - The operation you requested cannot be performed because the specified PIF allows unplug.
// REQUIRED_PIF_IS_UNPLUGGED - The operation you requested cannot be performed because the specified PIF is currently unplugged.
func (clusterHost) AsyncEnable(session *Session, self ClusterHostRef) (retval TaskRef, err error) {
	method := "Async.Cluster_host.enable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// ForceDestroy: Remove a host from an existing cluster forcefully.
//
// Errors:
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
func (clusterHost) ForceDestroy(session *Session, self ClusterHostRef) (err error) {
	method := "Cluster_host.force_destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncForceDestroy: Remove a host from an existing cluster forcefully.
//
// Errors:
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
func (clusterHost) AsyncForceDestroy(session *Session, self ClusterHostRef) (retval TaskRef, err error) {
	method := "Async.Cluster_host.force_destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Disable: Disable cluster membership for an enabled cluster host.
//
// Errors:
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
func (clusterHost) Disable(session *Session, self ClusterHostRef) (err error) {
	method := "Cluster_host.disable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDisable: Disable cluster membership for an enabled cluster host.
//
// Errors:
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
func (clusterHost) AsyncDisable(session *Session, self ClusterHostRef) (retval TaskRef, err error) {
	method := "Async.Cluster_host.disable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the Cluster_hosts known to the system.
func (clusterHost) GetAll(session *Session) (retval []ClusterHostRef, err error) {
	method := "Cluster_host.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterHostRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of Cluster_host references to Cluster_host records for all Cluster_hosts known to the system.
func (clusterHost) GetAllRecords(session *Session) (retval map[ClusterHostRef]ClusterHostRecord, err error) {
	method := "Cluster_host.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterHostRefToClusterHostRecordMap(method+" -> ", result)
	return
}

