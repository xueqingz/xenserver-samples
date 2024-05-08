package xenapi

import (
	"fmt"
)

type ClusterRecord struct {
	// Unique identifier/object reference
	UUID string
	// A list of the cluster_host objects associated with the Cluster
	ClusterHosts []ClusterHostRef
	// Internal field used by Host.destroy to store the IP of cluster members marked as permanently dead but not yet removed
	PendingForget []string
	// The secret key used by xapi-clusterd when it talks to itself on other hosts
	ClusterToken string
	// Simply the string &apos;corosync&apos;. No other cluster stacks are currently supported
	ClusterStack string
	// Whether the cluster stack thinks the cluster is quorate
	IsQuorate bool
	// Number of live hosts in order to be quorate
	Quorum int
	// Current number of live hosts, according to the cluster stack
	LiveHosts int
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []ClusterOperation
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]ClusterOperation
	// True if automatically joining new pool members to the cluster. This will be `true` in the first release
	PoolAutoJoin bool
	// The corosync token timeout in seconds
	TokenTimeout float64
	// The corosync token timeout coefficient in seconds
	TokenTimeoutCoefficient float64
	// Contains read-only settings for the cluster, such as timeouts and other options. It can only be set at cluster create time
	ClusterConfig map[string]string
	// Additional configuration
	OtherConfig map[string]string
}

type ClusterRef string

// Cluster-wide Cluster metadata
type cluster struct{}

var Cluster cluster

// GetRecord: Get a record containing the current state of the given Cluster.
func (cluster) GetRecord(session *Session, self ClusterRef) (retval ClusterRecord, err error) {
	method := "Cluster.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the Cluster instance with the specified UUID.
func (cluster) GetByUUID(session *Session, uUID string) (retval ClusterRef, err error) {
	method := "Cluster.get_by_uuid"
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
	retval, err = deserializeClusterRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given Cluster.
func (cluster) GetUUID(session *Session, self ClusterRef) (retval string, err error) {
	method := "Cluster.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetClusterHosts: Get the cluster_hosts field of the given Cluster.
func (cluster) GetClusterHosts(session *Session, self ClusterRef) (retval []ClusterHostRef, err error) {
	method := "Cluster.get_cluster_hosts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterHostRefSet(method+" -> ", result)
	return
}

// GetPendingForget: Get the pending_forget field of the given Cluster.
func (cluster) GetPendingForget(session *Session, self ClusterRef) (retval []string, err error) {
	method := "Cluster.get_pending_forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetClusterToken: Get the cluster_token field of the given Cluster.
func (cluster) GetClusterToken(session *Session, self ClusterRef) (retval string, err error) {
	method := "Cluster.get_cluster_token"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetClusterStack: Get the cluster_stack field of the given Cluster.
func (cluster) GetClusterStack(session *Session, self ClusterRef) (retval string, err error) {
	method := "Cluster.get_cluster_stack"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsQuorate: Get the is_quorate field of the given Cluster.
func (cluster) GetIsQuorate(session *Session, self ClusterRef) (retval bool, err error) {
	method := "Cluster.get_is_quorate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetQuorum: Get the quorum field of the given Cluster.
func (cluster) GetQuorum(session *Session, self ClusterRef) (retval int, err error) {
	method := "Cluster.get_quorum"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLiveHosts: Get the live_hosts field of the given Cluster.
func (cluster) GetLiveHosts(session *Session, self ClusterRef) (retval int, err error) {
	method := "Cluster.get_live_hosts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given Cluster.
func (cluster) GetAllowedOperations(session *Session, self ClusterRef) (retval []ClusterOperation, err error) {
	method := "Cluster.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumClusterOperationSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given Cluster.
func (cluster) GetCurrentOperations(session *Session, self ClusterRef) (retval map[string]ClusterOperation, err error) {
	method := "Cluster.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumClusterOperationMap(method+" -> ", result)
	return
}

// GetPoolAutoJoin: Get the pool_auto_join field of the given Cluster.
func (cluster) GetPoolAutoJoin(session *Session, self ClusterRef) (retval bool, err error) {
	method := "Cluster.get_pool_auto_join"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTokenTimeout: Get the token_timeout field of the given Cluster.
func (cluster) GetTokenTimeout(session *Session, self ClusterRef) (retval float64, err error) {
	method := "Cluster.get_token_timeout"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTokenTimeoutCoefficient: Get the token_timeout_coefficient field of the given Cluster.
func (cluster) GetTokenTimeoutCoefficient(session *Session, self ClusterRef) (retval float64, err error) {
	method := "Cluster.get_token_timeout_coefficient"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetClusterConfig: Get the cluster_config field of the given Cluster.
func (cluster) GetClusterConfig(session *Session, self ClusterRef) (retval map[string]string, err error) {
	method := "Cluster.get_cluster_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given Cluster.
func (cluster) GetOtherConfig(session *Session, self ClusterRef) (retval map[string]string, err error) {
	method := "Cluster.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given Cluster.
func (cluster) SetOtherConfig(session *Session, self ClusterRef, value map[string]string) (err error) {
	method := "Cluster.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given Cluster.
func (cluster) AddToOtherConfig(session *Session, self ClusterRef, key string, value string) (err error) {
	method := "Cluster.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given Cluster.  If the key is not in that Map, then do nothing.
func (cluster) RemoveFromOtherConfig(session *Session, self ClusterRef, key string) (err error) {
	method := "Cluster.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Creates a Cluster object and one Cluster_host object as its first member
//
// Errors:
// INVALID_CLUSTER_STACK - The cluster stack provided is not supported.
// INVALID_VALUE - The value given is invalid
// PIF_ALLOWS_UNPLUG - The operation you requested cannot be performed because the specified PIF allows unplug.
// REQUIRED_PIF_IS_UNPLUGGED - The operation you requested cannot be performed because the specified PIF is currently unplugged.
func (cluster) Create(session *Session, pIF PIFRef, clusterStack string, poolAutoJoin bool, tokenTimeout float64, tokenTimeoutCoefficient float64) (retval ClusterRef, err error) {
	method := "Cluster.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	pIFArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "PIF"), pIF)
	if err != nil {
		return
	}
	clusterStackArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cluster_stack"), clusterStack)
	if err != nil {
		return
	}
	poolAutoJoinArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "pool_auto_join"), poolAutoJoin)
	if err != nil {
		return
	}
	tokenTimeoutArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "token_timeout"), tokenTimeout)
	if err != nil {
		return
	}
	tokenTimeoutCoefficientArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "token_timeout_coefficient"), tokenTimeoutCoefficient)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, pIFArg, clusterStackArg, poolAutoJoinArg, tokenTimeoutArg, tokenTimeoutCoefficientArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterRef(method+" -> ", result)
	return
}

// AsyncCreate: Creates a Cluster object and one Cluster_host object as its first member
//
// Errors:
// INVALID_CLUSTER_STACK - The cluster stack provided is not supported.
// INVALID_VALUE - The value given is invalid
// PIF_ALLOWS_UNPLUG - The operation you requested cannot be performed because the specified PIF allows unplug.
// REQUIRED_PIF_IS_UNPLUGGED - The operation you requested cannot be performed because the specified PIF is currently unplugged.
func (cluster) AsyncCreate(session *Session, pIF PIFRef, clusterStack string, poolAutoJoin bool, tokenTimeout float64, tokenTimeoutCoefficient float64) (retval TaskRef, err error) {
	method := "Async.Cluster.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	pIFArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "PIF"), pIF)
	if err != nil {
		return
	}
	clusterStackArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cluster_stack"), clusterStack)
	if err != nil {
		return
	}
	poolAutoJoinArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "pool_auto_join"), poolAutoJoin)
	if err != nil {
		return
	}
	tokenTimeoutArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "token_timeout"), tokenTimeout)
	if err != nil {
		return
	}
	tokenTimeoutCoefficientArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "token_timeout_coefficient"), tokenTimeoutCoefficient)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, pIFArg, clusterStackArg, poolAutoJoinArg, tokenTimeoutArg, tokenTimeoutCoefficientArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Destroys a Cluster object and the one remaining Cluster_host member
//
// Errors:
// CLUSTER_DOES_NOT_HAVE_ONE_NODE - An operation failed as it expected the cluster to have only one node but found multiple cluster_hosts.
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
func (cluster) Destroy(session *Session, self ClusterRef) (err error) {
	method := "Cluster.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroys a Cluster object and the one remaining Cluster_host member
//
// Errors:
// CLUSTER_DOES_NOT_HAVE_ONE_NODE - An operation failed as it expected the cluster to have only one node but found multiple cluster_hosts.
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
func (cluster) AsyncDestroy(session *Session, self ClusterRef) (retval TaskRef, err error) {
	method := "Async.Cluster.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNetwork: Returns the network used by the cluster for inter-host communication, i.e. the network shared by all cluster host PIFs
func (cluster) GetNetwork(session *Session, self ClusterRef) (retval NetworkRef, err error) {
	method := "Cluster.get_network"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeNetworkRef(method+" -> ", result)
	return
}

// AsyncGetNetwork: Returns the network used by the cluster for inter-host communication, i.e. the network shared by all cluster host PIFs
func (cluster) AsyncGetNetwork(session *Session, self ClusterRef) (retval TaskRef, err error) {
	method := "Async.Cluster.get_network"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// PoolCreate: Attempt to create a Cluster from the entire pool
func (cluster) PoolCreate(session *Session, network NetworkRef, clusterStack string, tokenTimeout float64, tokenTimeoutCoefficient float64) (retval ClusterRef, err error) {
	method := "Cluster.pool_create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	clusterStackArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cluster_stack"), clusterStack)
	if err != nil {
		return
	}
	tokenTimeoutArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "token_timeout"), tokenTimeout)
	if err != nil {
		return
	}
	tokenTimeoutCoefficientArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "token_timeout_coefficient"), tokenTimeoutCoefficient)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, networkArg, clusterStackArg, tokenTimeoutArg, tokenTimeoutCoefficientArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterRef(method+" -> ", result)
	return
}

// AsyncPoolCreate: Attempt to create a Cluster from the entire pool
func (cluster) AsyncPoolCreate(session *Session, network NetworkRef, clusterStack string, tokenTimeout float64, tokenTimeoutCoefficient float64) (retval TaskRef, err error) {
	method := "Async.Cluster.pool_create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	clusterStackArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cluster_stack"), clusterStack)
	if err != nil {
		return
	}
	tokenTimeoutArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "token_timeout"), tokenTimeout)
	if err != nil {
		return
	}
	tokenTimeoutCoefficientArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "token_timeout_coefficient"), tokenTimeoutCoefficient)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, networkArg, clusterStackArg, tokenTimeoutArg, tokenTimeoutCoefficientArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// PoolForceDestroy: Attempt to force destroy the Cluster_host objects, and then destroy the Cluster.
//
// Errors:
// CLUSTER_FORCE_DESTROY_FAILED - Force destroy failed on a Cluster_host while force destroying the cluster.
func (cluster) PoolForceDestroy(session *Session, self ClusterRef) (err error) {
	method := "Cluster.pool_force_destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPoolForceDestroy: Attempt to force destroy the Cluster_host objects, and then destroy the Cluster.
//
// Errors:
// CLUSTER_FORCE_DESTROY_FAILED - Force destroy failed on a Cluster_host while force destroying the cluster.
func (cluster) AsyncPoolForceDestroy(session *Session, self ClusterRef) (retval TaskRef, err error) {
	method := "Async.Cluster.pool_force_destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// PoolDestroy: Attempt to destroy the Cluster_host objects for all hosts in the pool and then destroy the Cluster.
//
// Errors:
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
// CLUSTERING_DISABLED - An operation was attempted while clustering was disabled on the cluster_host.
// CLUSTER_HOST_IS_LAST - The last cluster host cannot be destroyed. Destroy the cluster instead
func (cluster) PoolDestroy(session *Session, self ClusterRef) (err error) {
	method := "Cluster.pool_destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPoolDestroy: Attempt to destroy the Cluster_host objects for all hosts in the pool and then destroy the Cluster.
//
// Errors:
// CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
// CLUSTERING_DISABLED - An operation was attempted while clustering was disabled on the cluster_host.
// CLUSTER_HOST_IS_LAST - The last cluster host cannot be destroyed. Destroy the cluster instead
func (cluster) AsyncPoolDestroy(session *Session, self ClusterRef) (retval TaskRef, err error) {
	method := "Async.Cluster.pool_destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// PoolResync: Resynchronise the cluster_host objects across the pool. Creates them where they need creating and then plugs them
func (cluster) PoolResync(session *Session, self ClusterRef) (err error) {
	method := "Cluster.pool_resync"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncPoolResync: Resynchronise the cluster_host objects across the pool. Creates them where they need creating and then plugs them
func (cluster) AsyncPoolResync(session *Session, self ClusterRef) (retval TaskRef, err error) {
	method := "Async.Cluster.pool_resync"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeClusterRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the Clusters known to the system.
func (cluster) GetAll(session *Session) (retval []ClusterRef, err error) {
	method := "Cluster.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of Cluster references to Cluster records for all Clusters known to the system.
func (cluster) GetAllRecords(session *Session) (retval map[ClusterRef]ClusterRecord, err error) {
	method := "Cluster.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeClusterRefToClusterRecordMap(method+" -> ", result)
	return
}

