package xenapi

import (
	"fmt"
	"time"
)

type PoolRecord struct {
	// Unique identifier/object reference
	UUID string
	// Short name
	NameLabel string
	// Description
	NameDescription string
	// The host that is pool master
	Master HostRef
	// Default SR for VDIs
	DefaultSR SRRef
	// The SR in which VDIs for suspend images are created
	SuspendImageSR SRRef
	// The SR in which VDIs for crash dumps are created
	CrashDumpSR SRRef
	// additional configuration
	OtherConfig map[string]string
	// true if HA is enabled on the pool, false otherwise
	HaEnabled bool
	// The current HA configuration
	HaConfiguration map[string]string
	// HA statefile VDIs in use
	HaStatefiles []string
	// Number of host failures to tolerate before the Pool is declared to be overcommitted
	HaHostFailuresToTolerate int
	// Number of future host failures we have managed to find a plan for. Once this reaches zero any future host failures will cause the failure of protected VMs.
	HaPlanExistsFor int
	// If set to false then operations which would cause the Pool to become overcommitted will be blocked.
	HaAllowOvercommit bool
	// True if the Pool is considered to be overcommitted i.e. if there exist insufficient physical resources to tolerate the configured number of host failures
	HaOvercommitted bool
	// Binary blobs associated with this pool
	Blobs map[string]BlobRef
	// user-specified tags for categorization purposes
	Tags []string
	// gui-specific configuration for pool
	GuiConfig map[string]string
	// Configuration for the automatic health check feature
	HealthCheckConfig map[string]string
	// Url for the configured workload balancing host
	WlbURL string
	// Username for accessing the workload balancing host
	WlbUsername string
	// true if workload balancing is enabled on the pool, false otherwise
	WlbEnabled bool
	// true if communication with the WLB server should enforce TLS certificate verification.
	WlbVerifyCert bool
	// true a redo-log is to be used other than when HA is enabled, false otherwise
	RedoLogEnabled bool
	// indicates the VDI to use for the redo-log other than when HA is enabled
	RedoLogVdi VDIRef
	// address of the vswitch controller
	VswitchController string
	// Pool-wide restrictions currently in effect
	Restrictions map[string]string
	// The set of currently known metadata VDIs for this pool
	MetadataVDIs []VDIRef
	// The HA cluster stack that is currently in use. Only valid when HA is enabled.
	HaClusterStack string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []PoolAllowedOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]PoolAllowedOperations
	// Pool-wide guest agent configuration information
	GuestAgentConfig map[string]string
	// Details about the physical CPUs on the pool
	CPUInfo map[string]string
	// The pool-wide policy for clients on whether to use the vendor device or not on newly created VMs. This field will also be consulted if the &apos;has_vendor_device&apos; field is not specified in the VM.create call.
	PolicyNoVendorDevice bool
	// The pool-wide flag to show if the live patching feauture is disabled or not.
	LivePatchingDisabled bool
	// true if IGMP snooping is enabled in the pool, false otherwise.
	IgmpSnoopingEnabled bool
	// The UEFI certificates allowing Secure Boot
	UefiCertificates string
	// Custom UEFI certificates allowing Secure Boot
	CustomUefiCertificates string
	// True if either a PSR is running or we are waiting for a PSR to be re-run
	IsPsrPending bool
	// True iff TLS certificate verification is enabled
	TLSVerificationEnabled bool
	// The set of currently enabled repositories
	Repositories []RepositoryRef
	// True if authentication by TLS client certificates is enabled
	ClientCertificateAuthEnabled bool
	// The name (CN/SAN) that an incoming client certificate must have to allow authentication
	ClientCertificateAuthName string
	// Url of the proxy used in syncing with the enabled repositories
	RepositoryProxyURL string
	// Username for the authentication of the proxy used in syncing with the enabled repositories
	RepositoryProxyUsername string
	// Password for the authentication of the proxy used in syncing with the enabled repositories
	RepositoryProxyPassword SecretRef
	// Default behaviour during migration, True if stream compression should be used
	MigrationCompression bool
	// true if bias against pool master when scheduling vms is enabled, false otherwise
	CoordinatorBias bool
	// Maximum number of threads to use for PAM authentication
	LocalAuthMaxThreads int
	// Maximum number of threads to use for external (AD) authentication
	ExtAuthMaxThreads int
	// The UUID of the pool for identification of telemetry data
	TelemetryUUID SecretRef
	// How often the telemetry collection will be carried out
	TelemetryFrequency TelemetryFrequency
	// The earliest timestamp (in UTC) when the next round of telemetry collection can be carried out
	TelemetryNextCollection time.Time
	// time of the last update sychronization
	LastUpdateSync time.Time
	// The frequency at which updates are synchronized from a remote CDN: daily or weekly.
	UpdateSyncFrequency UpdateSyncFrequency
	// The day of the week the update synchronizations will be scheduled, based on pool&apos;s local timezone. Ignored when update_sync_frequency is daily
	UpdateSyncDay int
	// Whether periodic update synchronization is enabled or not
	UpdateSyncEnabled bool
}

type PoolRef string

// Pool-wide information
type pool struct{}

var Pool pool

// GetRecord: Get a record containing the current state of the given pool.
func (pool) GetRecord(session *Session, self PoolRef) (retval PoolRecord, err error) {
	method := "pool.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the pool instance with the specified UUID.
func (pool) GetByUUID(session *Session, uUID string) (retval PoolRef, err error) {
	method := "pool.get_by_uuid"
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
	retval, err = deserializePoolRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given pool.
func (pool) GetUUID(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name_label field of the given pool.
func (pool) GetNameLabel(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name_description field of the given pool.
func (pool) GetNameDescription(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMaster: Get the master field of the given pool.
func (pool) GetMaster(session *Session, self PoolRef) (retval HostRef, err error) {
	method := "pool.get_master"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDefaultSR: Get the default_SR field of the given pool.
func (pool) GetDefaultSR(session *Session, self PoolRef) (retval SRRef, err error) {
	method := "pool.get_default_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRef(method+" -> ", result)
	return
}

// GetSuspendImageSR: Get the suspend_image_SR field of the given pool.
func (pool) GetSuspendImageSR(session *Session, self PoolRef) (retval SRRef, err error) {
	method := "pool.get_suspend_image_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRef(method+" -> ", result)
	return
}

// GetCrashDumpSR: Get the crash_dump_SR field of the given pool.
func (pool) GetCrashDumpSR(session *Session, self PoolRef) (retval SRRef, err error) {
	method := "pool.get_crash_dump_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRef(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given pool.
func (pool) GetOtherConfig(session *Session, self PoolRef) (retval map[string]string, err error) {
	method := "pool.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaEnabled: Get the ha_enabled field of the given pool.
func (pool) GetHaEnabled(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_ha_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaConfiguration: Get the ha_configuration field of the given pool.
func (pool) GetHaConfiguration(session *Session, self PoolRef) (retval map[string]string, err error) {
	method := "pool.get_ha_configuration"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaStatefiles: Get the ha_statefiles field of the given pool.
func (pool) GetHaStatefiles(session *Session, self PoolRef) (retval []string, err error) {
	method := "pool.get_ha_statefiles"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaHostFailuresToTolerate: Get the ha_host_failures_to_tolerate field of the given pool.
func (pool) GetHaHostFailuresToTolerate(session *Session, self PoolRef) (retval int, err error) {
	method := "pool.get_ha_host_failures_to_tolerate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaPlanExistsFor: Get the ha_plan_exists_for field of the given pool.
func (pool) GetHaPlanExistsFor(session *Session, self PoolRef) (retval int, err error) {
	method := "pool.get_ha_plan_exists_for"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaAllowOvercommit: Get the ha_allow_overcommit field of the given pool.
func (pool) GetHaAllowOvercommit(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_ha_allow_overcommit"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaOvercommitted: Get the ha_overcommitted field of the given pool.
func (pool) GetHaOvercommitted(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_ha_overcommitted"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBlobs: Get the blobs field of the given pool.
func (pool) GetBlobs(session *Session, self PoolRef) (retval map[string]BlobRef, err error) {
	method := "pool.get_blobs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTags: Get the tags field of the given pool.
func (pool) GetTags(session *Session, self PoolRef) (retval []string, err error) {
	method := "pool.get_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetGuiConfig: Get the gui_config field of the given pool.
func (pool) GetGuiConfig(session *Session, self PoolRef) (retval map[string]string, err error) {
	method := "pool.get_gui_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHealthCheckConfig: Get the health_check_config field of the given pool.
func (pool) GetHealthCheckConfig(session *Session, self PoolRef) (retval map[string]string, err error) {
	method := "pool.get_health_check_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetWlbURL: Get the wlb_url field of the given pool.
func (pool) GetWlbURL(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_wlb_url"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetWlbUsername: Get the wlb_username field of the given pool.
func (pool) GetWlbUsername(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_wlb_username"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetWlbEnabled: Get the wlb_enabled field of the given pool.
func (pool) GetWlbEnabled(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_wlb_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetWlbVerifyCert: Get the wlb_verify_cert field of the given pool.
func (pool) GetWlbVerifyCert(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_wlb_verify_cert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRedoLogEnabled: Get the redo_log_enabled field of the given pool.
func (pool) GetRedoLogEnabled(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_redo_log_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRedoLogVdi: Get the redo_log_vdi field of the given pool.
func (pool) GetRedoLogVdi(session *Session, self PoolRef) (retval VDIRef, err error) {
	method := "pool.get_redo_log_vdi"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVswitchController: Get the vswitch_controller field of the given pool.
func (pool) GetVswitchController(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_vswitch_controller"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRestrictions: Get the restrictions field of the given pool.
func (pool) GetRestrictions(session *Session, self PoolRef) (retval map[string]string, err error) {
	method := "pool.get_restrictions"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMetadataVDIs: Get the metadata_VDIs field of the given pool.
func (pool) GetMetadataVDIs(session *Session, self PoolRef) (retval []VDIRef, err error) {
	method := "pool.get_metadata_VDIs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVDIRefSet(method+" -> ", result)
	return
}

// GetHaClusterStack: Get the ha_cluster_stack field of the given pool.
func (pool) GetHaClusterStack(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_ha_cluster_stack"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given pool.
func (pool) GetAllowedOperations(session *Session, self PoolRef) (retval []PoolAllowedOperations, err error) {
	method := "pool.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumPoolAllowedOperationsSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given pool.
func (pool) GetCurrentOperations(session *Session, self PoolRef) (retval map[string]PoolAllowedOperations, err error) {
	method := "pool.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumPoolAllowedOperationsMap(method+" -> ", result)
	return
}

// GetGuestAgentConfig: Get the guest_agent_config field of the given pool.
func (pool) GetGuestAgentConfig(session *Session, self PoolRef) (retval map[string]string, err error) {
	method := "pool.get_guest_agent_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCPUInfo: Get the cpu_info field of the given pool.
func (pool) GetCPUInfo(session *Session, self PoolRef) (retval map[string]string, err error) {
	method := "pool.get_cpu_info"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPolicyNoVendorDevice: Get the policy_no_vendor_device field of the given pool.
func (pool) GetPolicyNoVendorDevice(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_policy_no_vendor_device"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLivePatchingDisabled: Get the live_patching_disabled field of the given pool.
func (pool) GetLivePatchingDisabled(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_live_patching_disabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIgmpSnoopingEnabled: Get the igmp_snooping_enabled field of the given pool.
func (pool) GetIgmpSnoopingEnabled(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_igmp_snooping_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUefiCertificates: Get the uefi_certificates field of the given pool.
func (pool) GetUefiCertificates(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_uefi_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCustomUefiCertificates: Get the custom_uefi_certificates field of the given pool.
func (pool) GetCustomUefiCertificates(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_custom_uefi_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsPsrPending: Get the is_psr_pending field of the given pool.
func (pool) GetIsPsrPending(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_is_psr_pending"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTLSVerificationEnabled: Get the tls_verification_enabled field of the given pool.
func (pool) GetTLSVerificationEnabled(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_tls_verification_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRepositories: Get the repositories field of the given pool.
func (pool) GetRepositories(session *Session, self PoolRef) (retval []RepositoryRef, err error) {
	method := "pool.get_repositories"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeRepositoryRefSet(method+" -> ", result)
	return
}

// GetClientCertificateAuthEnabled: Get the client_certificate_auth_enabled field of the given pool.
func (pool) GetClientCertificateAuthEnabled(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_client_certificate_auth_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetClientCertificateAuthName: Get the client_certificate_auth_name field of the given pool.
func (pool) GetClientCertificateAuthName(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_client_certificate_auth_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRepositoryProxyURL: Get the repository_proxy_url field of the given pool.
func (pool) GetRepositoryProxyURL(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_repository_proxy_url"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRepositoryProxyUsername: Get the repository_proxy_username field of the given pool.
func (pool) GetRepositoryProxyUsername(session *Session, self PoolRef) (retval string, err error) {
	method := "pool.get_repository_proxy_username"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRepositoryProxyPassword: Get the repository_proxy_password field of the given pool.
func (pool) GetRepositoryProxyPassword(session *Session, self PoolRef) (retval SecretRef, err error) {
	method := "pool.get_repository_proxy_password"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRef(method+" -> ", result)
	return
}

// GetMigrationCompression: Get the migration_compression field of the given pool.
func (pool) GetMigrationCompression(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_migration_compression"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCoordinatorBias: Get the coordinator_bias field of the given pool.
func (pool) GetCoordinatorBias(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_coordinator_bias"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLocalAuthMaxThreads: Get the local_auth_max_threads field of the given pool.
func (pool) GetLocalAuthMaxThreads(session *Session, self PoolRef) (retval int, err error) {
	method := "pool.get_local_auth_max_threads"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetExtAuthMaxThreads: Get the ext_auth_max_threads field of the given pool.
func (pool) GetExtAuthMaxThreads(session *Session, self PoolRef) (retval int, err error) {
	method := "pool.get_ext_auth_max_threads"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTelemetryUUID: Get the telemetry_uuid field of the given pool.
func (pool) GetTelemetryUUID(session *Session, self PoolRef) (retval SecretRef, err error) {
	method := "pool.get_telemetry_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRef(method+" -> ", result)
	return
}

// GetTelemetryFrequency: Get the telemetry_frequency field of the given pool.
func (pool) GetTelemetryFrequency(session *Session, self PoolRef) (retval TelemetryFrequency, err error) {
	method := "pool.get_telemetry_frequency"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumTelemetryFrequency(method+" -> ", result)
	return
}

// GetTelemetryNextCollection: Get the telemetry_next_collection field of the given pool.
func (pool) GetTelemetryNextCollection(session *Session, self PoolRef) (retval time.Time, err error) {
	method := "pool.get_telemetry_next_collection"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastUpdateSync: Get the last_update_sync field of the given pool.
func (pool) GetLastUpdateSync(session *Session, self PoolRef) (retval time.Time, err error) {
	method := "pool.get_last_update_sync"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUpdateSyncFrequency: Get the update_sync_frequency field of the given pool.
func (pool) GetUpdateSyncFrequency(session *Session, self PoolRef) (retval UpdateSyncFrequency, err error) {
	method := "pool.get_update_sync_frequency"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumUpdateSyncFrequency(method+" -> ", result)
	return
}

// GetUpdateSyncDay: Get the update_sync_day field of the given pool.
func (pool) GetUpdateSyncDay(session *Session, self PoolRef) (retval int, err error) {
	method := "pool.get_update_sync_day"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUpdateSyncEnabled: Get the update_sync_enabled field of the given pool.
func (pool) GetUpdateSyncEnabled(session *Session, self PoolRef) (retval bool, err error) {
	method := "pool.get_update_sync_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel: Set the name_label field of the given pool.
func (pool) SetNameLabel(session *Session, self PoolRef, value string) (err error) {
	method := "pool.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription: Set the name_description field of the given pool.
func (pool) SetNameDescription(session *Session, self PoolRef, value string) (err error) {
	method := "pool.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetDefaultSR: Set the default_SR field of the given pool.
func (pool) SetDefaultSR(session *Session, self PoolRef, value SRRef) (err error) {
	method := "pool.set_default_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetSuspendImageSR: Set the suspend_image_SR field of the given pool.
func (pool) SetSuspendImageSR(session *Session, self PoolRef, value SRRef) (err error) {
	method := "pool.set_suspend_image_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetCrashDumpSR: Set the crash_dump_SR field of the given pool.
func (pool) SetCrashDumpSR(session *Session, self PoolRef, value SRRef) (err error) {
	method := "pool.set_crash_dump_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetOtherConfig: Set the other_config field of the given pool.
func (pool) SetOtherConfig(session *Session, self PoolRef, value map[string]string) (err error) {
	method := "pool.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given pool.
func (pool) AddToOtherConfig(session *Session, self PoolRef, key string, value string) (err error) {
	method := "pool.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given pool.  If the key is not in that Map, then do nothing.
func (pool) RemoveFromOtherConfig(session *Session, self PoolRef, key string) (err error) {
	method := "pool.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetHaAllowOvercommit: Set the ha_allow_overcommit field of the given pool.
func (pool) SetHaAllowOvercommit(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_ha_allow_overcommit"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetTags: Set the tags field of the given pool.
func (pool) SetTags(session *Session, self PoolRef, value []string) (err error) {
	method := "pool.set_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddTags: Add the given value to the tags field of the given pool.  If the value is already in that Set, then do nothing.
func (pool) AddTags(session *Session, self PoolRef, value string) (err error) {
	method := "pool.add_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveTags: Remove the given value from the tags field of the given pool.  If the value is not in that Set, then do nothing.
func (pool) RemoveTags(session *Session, self PoolRef, value string) (err error) {
	method := "pool.remove_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetGuiConfig: Set the gui_config field of the given pool.
func (pool) SetGuiConfig(session *Session, self PoolRef, value map[string]string) (err error) {
	method := "pool.set_gui_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToGuiConfig: Add the given key-value pair to the gui_config field of the given pool.
func (pool) AddToGuiConfig(session *Session, self PoolRef, key string, value string) (err error) {
	method := "pool.add_to_gui_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromGuiConfig: Remove the given key and its corresponding value from the gui_config field of the given pool.  If the key is not in that Map, then do nothing.
func (pool) RemoveFromGuiConfig(session *Session, self PoolRef, key string) (err error) {
	method := "pool.remove_from_gui_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetHealthCheckConfig: Set the health_check_config field of the given pool.
func (pool) SetHealthCheckConfig(session *Session, self PoolRef, value map[string]string) (err error) {
	method := "pool.set_health_check_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToHealthCheckConfig: Add the given key-value pair to the health_check_config field of the given pool.
func (pool) AddToHealthCheckConfig(session *Session, self PoolRef, key string, value string) (err error) {
	method := "pool.add_to_health_check_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromHealthCheckConfig: Remove the given key and its corresponding value from the health_check_config field of the given pool.  If the key is not in that Map, then do nothing.
func (pool) RemoveFromHealthCheckConfig(session *Session, self PoolRef, key string) (err error) {
	method := "pool.remove_from_health_check_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetWlbEnabled: Set the wlb_enabled field of the given pool.
func (pool) SetWlbEnabled(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_wlb_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetWlbVerifyCert: Set the wlb_verify_cert field of the given pool.
func (pool) SetWlbVerifyCert(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_wlb_verify_cert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetPolicyNoVendorDevice: Set the policy_no_vendor_device field of the given pool.
func (pool) SetPolicyNoVendorDevice(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_policy_no_vendor_device"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetLivePatchingDisabled: Set the live_patching_disabled field of the given pool.
func (pool) SetLivePatchingDisabled(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_live_patching_disabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetIsPsrPending: Set the is_psr_pending field of the given pool.
func (pool) SetIsPsrPending(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_is_psr_pending"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetMigrationCompression: Set the migration_compression field of the given pool.
func (pool) SetMigrationCompression(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_migration_compression"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetCoordinatorBias: Set the coordinator_bias field of the given pool.
func (pool) SetCoordinatorBias(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_coordinator_bias"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Join: Instruct host to join a new pool
//
// Errors:
// JOINING_HOST_CANNOT_CONTAIN_SHARED_SRS - The server joining the pool cannot contain any shared storage.
func (pool) Join(session *Session, masterAddress string, masterUsername string, masterPassword string) (err error) {
	method := "pool.join"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	masterAddressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_address"), masterAddress)
	if err != nil {
		return
	}
	masterUsernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_username"), masterUsername)
	if err != nil {
		return
	}
	masterPasswordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_password"), masterPassword)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, masterAddressArg, masterUsernameArg, masterPasswordArg)
	return
}

// AsyncJoin: Instruct host to join a new pool
//
// Errors:
// JOINING_HOST_CANNOT_CONTAIN_SHARED_SRS - The server joining the pool cannot contain any shared storage.
func (pool) AsyncJoin(session *Session, masterAddress string, masterUsername string, masterPassword string) (retval TaskRef, err error) {
	method := "Async.pool.join"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	masterAddressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_address"), masterAddress)
	if err != nil {
		return
	}
	masterUsernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_username"), masterUsername)
	if err != nil {
		return
	}
	masterPasswordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_password"), masterPassword)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, masterAddressArg, masterUsernameArg, masterPasswordArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// JoinForce: Instruct host to join a new pool
func (pool) JoinForce(session *Session, masterAddress string, masterUsername string, masterPassword string) (err error) {
	method := "pool.join_force"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	masterAddressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_address"), masterAddress)
	if err != nil {
		return
	}
	masterUsernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_username"), masterUsername)
	if err != nil {
		return
	}
	masterPasswordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_password"), masterPassword)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, masterAddressArg, masterUsernameArg, masterPasswordArg)
	return
}

// AsyncJoinForce: Instruct host to join a new pool
func (pool) AsyncJoinForce(session *Session, masterAddress string, masterUsername string, masterPassword string) (retval TaskRef, err error) {
	method := "Async.pool.join_force"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	masterAddressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_address"), masterAddress)
	if err != nil {
		return
	}
	masterUsernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_username"), masterUsername)
	if err != nil {
		return
	}
	masterPasswordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_password"), masterPassword)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, masterAddressArg, masterUsernameArg, masterPasswordArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Eject: Instruct a pool master to eject a host from the pool
func (pool) Eject(session *Session, host HostRef) (err error) {
	method := "pool.eject"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg)
	return
}

// AsyncEject: Instruct a pool master to eject a host from the pool
func (pool) AsyncEject(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.pool.eject"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EmergencyTransitionToMaster: Instruct host that&apos;s currently a slave to transition to being master
func (pool) EmergencyTransitionToMaster(session *Session) (err error) {
	method := "pool.emergency_transition_to_master"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// EmergencyResetMaster: Instruct a slave already in a pool that the master has changed
func (pool) EmergencyResetMaster(session *Session, masterAddress string) (err error) {
	method := "pool.emergency_reset_master"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	masterAddressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "master_address"), masterAddress)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, masterAddressArg)
	return
}

// RecoverSlaves: Instruct a pool master, M, to try and contact its slaves and, if slaves are in emergency mode, reset their master address to M.
func (pool) RecoverSlaves(session *Session) (retval []HostRef, err error) {
	method := "pool.recover_slaves"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRefSet(method+" -> ", result)
	return
}

// AsyncRecoverSlaves: Instruct a pool master, M, to try and contact its slaves and, if slaves are in emergency mode, reset their master address to M.
func (pool) AsyncRecoverSlaves(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.recover_slaves"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CreateVLAN: Create PIFs, mapping a network to the same physical interface/VLAN on each host. This call is deprecated: use Pool.create_VLAN_from_PIF instead.
//
// Errors:
// VLAN_TAG_INVALID - You tried to create a VLAN, but the tag you gave was invalid -- it must be between 0 and 4094. The parameter echoes the VLAN tag you gave.
func (pool) CreateVLAN(session *Session, device string, network NetworkRef, vLAN int) (retval []PIFRef, err error) {
	method := "pool.create_VLAN"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	deviceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "device"), device)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	vLANArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "VLAN"), vLAN)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, deviceArg, networkArg, vLANArg)
	if err != nil {
		return
	}
	retval, err = deserializePIFRefSet(method+" -> ", result)
	return
}

// AsyncCreateVLAN: Create PIFs, mapping a network to the same physical interface/VLAN on each host. This call is deprecated: use Pool.create_VLAN_from_PIF instead.
//
// Errors:
// VLAN_TAG_INVALID - You tried to create a VLAN, but the tag you gave was invalid -- it must be between 0 and 4094. The parameter echoes the VLAN tag you gave.
func (pool) AsyncCreateVLAN(session *Session, device string, network NetworkRef, vLAN int) (retval TaskRef, err error) {
	method := "Async.pool.create_VLAN"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	deviceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "device"), device)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	vLANArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "VLAN"), vLAN)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, deviceArg, networkArg, vLANArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// ManagementReconfigure: Reconfigure the management network interface for all Hosts in the Pool
//
// Errors:
// HA_IS_ENABLED - The operation could not be performed because HA is enabled on the Pool
// PIF_NOT_PRESENT - This host has no PIF on the given network.
// CANNOT_PLUG_BOND_SLAVE - This PIF is a bond member and cannot be plugged.
// PIF_INCOMPATIBLE_PRIMARY_ADDRESS_TYPE - The primary address types are not compatible
// PIF_HAS_NO_NETWORK_CONFIGURATION - PIF has no IP configuration (mode currently set to &apos;none&apos;)
// PIF_HAS_NO_V6_NETWORK_CONFIGURATION - PIF has no IPv6 configuration (mode currently set to &apos;none&apos;)
func (pool) ManagementReconfigure(session *Session, network NetworkRef) (err error) {
	method := "pool.management_reconfigure"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, networkArg)
	return
}

// AsyncManagementReconfigure: Reconfigure the management network interface for all Hosts in the Pool
//
// Errors:
// HA_IS_ENABLED - The operation could not be performed because HA is enabled on the Pool
// PIF_NOT_PRESENT - This host has no PIF on the given network.
// CANNOT_PLUG_BOND_SLAVE - This PIF is a bond member and cannot be plugged.
// PIF_INCOMPATIBLE_PRIMARY_ADDRESS_TYPE - The primary address types are not compatible
// PIF_HAS_NO_NETWORK_CONFIGURATION - PIF has no IP configuration (mode currently set to &apos;none&apos;)
// PIF_HAS_NO_V6_NETWORK_CONFIGURATION - PIF has no IPv6 configuration (mode currently set to &apos;none&apos;)
func (pool) AsyncManagementReconfigure(session *Session, network NetworkRef) (retval TaskRef, err error) {
	method := "Async.pool.management_reconfigure"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, networkArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CreateVLANFromPIF: Create a pool-wide VLAN by taking the PIF.
//
// Errors:
// VLAN_TAG_INVALID - You tried to create a VLAN, but the tag you gave was invalid -- it must be between 0 and 4094. The parameter echoes the VLAN tag you gave.
func (pool) CreateVLANFromPIF(session *Session, pif PIFRef, network NetworkRef, vLAN int) (retval []PIFRef, err error) {
	method := "pool.create_VLAN_from_PIF"
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
	vLANArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "VLAN"), vLAN)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, pifArg, networkArg, vLANArg)
	if err != nil {
		return
	}
	retval, err = deserializePIFRefSet(method+" -> ", result)
	return
}

// AsyncCreateVLANFromPIF: Create a pool-wide VLAN by taking the PIF.
//
// Errors:
// VLAN_TAG_INVALID - You tried to create a VLAN, but the tag you gave was invalid -- it must be between 0 and 4094. The parameter echoes the VLAN tag you gave.
func (pool) AsyncCreateVLANFromPIF(session *Session, pif PIFRef, network NetworkRef, vLAN int) (retval TaskRef, err error) {
	method := "Async.pool.create_VLAN_from_PIF"
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
	vLANArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "VLAN"), vLAN)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, pifArg, networkArg, vLANArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EnableHa: Turn on High Availability mode
func (pool) EnableHa(session *Session, heartbeatSrs []SRRef, configuration map[string]string) (err error) {
	method := "pool.enable_ha"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	heartbeatSrsArg, err := serializeSRRefSet(fmt.Sprintf("%s(%s)", method, "heartbeat_srs"), heartbeatSrs)
	if err != nil {
		return
	}
	configurationArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "configuration"), configuration)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, heartbeatSrsArg, configurationArg)
	return
}

// AsyncEnableHa: Turn on High Availability mode
func (pool) AsyncEnableHa(session *Session, heartbeatSrs []SRRef, configuration map[string]string) (retval TaskRef, err error) {
	method := "Async.pool.enable_ha"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	heartbeatSrsArg, err := serializeSRRefSet(fmt.Sprintf("%s(%s)", method, "heartbeat_srs"), heartbeatSrs)
	if err != nil {
		return
	}
	configurationArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "configuration"), configuration)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, heartbeatSrsArg, configurationArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// DisableHa: Turn off High Availability mode
func (pool) DisableHa(session *Session) (err error) {
	method := "pool.disable_ha"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// AsyncDisableHa: Turn off High Availability mode
func (pool) AsyncDisableHa(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.disable_ha"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SyncDatabase: Forcibly synchronise the database now
func (pool) SyncDatabase(session *Session) (err error) {
	method := "pool.sync_database"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// AsyncSyncDatabase: Forcibly synchronise the database now
func (pool) AsyncSyncDatabase(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.sync_database"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// DesignateNewMaster: Perform an orderly handover of the role of master to the referenced host.
func (pool) DesignateNewMaster(session *Session, host HostRef) (err error) {
	method := "pool.designate_new_master"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg)
	return
}

// AsyncDesignateNewMaster: Perform an orderly handover of the role of master to the referenced host.
func (pool) AsyncDesignateNewMaster(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.pool.designate_new_master"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// HaPreventRestartsFor: When this call returns the VM restart logic will not run for the requested number of seconds. If the argument is zero then the restart thread is immediately unblocked
func (pool) HaPreventRestartsFor(session *Session, seconds int) (err error) {
	method := "pool.ha_prevent_restarts_for"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	secondsArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "seconds"), seconds)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, secondsArg)
	return
}

// HaFailoverPlanExists: Returns true if a VM failover plan exists for up to &apos;n&apos; host failures
func (pool) HaFailoverPlanExists(session *Session, n int) (retval bool, err error) {
	method := "pool.ha_failover_plan_exists"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "n"), n)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nArg)
	if err != nil {
		return
	}
	retval, err = deserializeBool(method+" -> ", result)
	return
}

// HaComputeMaxHostFailuresToTolerate: Returns the maximum number of host failures we could tolerate before we would be unable to restart configured VMs
func (pool) HaComputeMaxHostFailuresToTolerate(session *Session) (retval int, err error) {
	method := "pool.ha_compute_max_host_failures_to_tolerate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// HaComputeHypotheticalMaxHostFailuresToTolerate: Returns the maximum number of host failures we could tolerate before we would be unable to restart the provided VMs
func (pool) HaComputeHypotheticalMaxHostFailuresToTolerate(session *Session, configuration map[VMRef]string) (retval int, err error) {
	method := "pool.ha_compute_hypothetical_max_host_failures_to_tolerate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	configurationArg, err := serializeVMRefToStringMap(fmt.Sprintf("%s(%s)", method, "configuration"), configuration)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, configurationArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// HaComputeVMFailoverPlan: Return a VM failover plan assuming a given subset of hosts fail
func (pool) HaComputeVMFailoverPlan(session *Session, failedHosts []HostRef, failedVms []VMRef) (retval map[VMRef]map[string]string, err error) {
	method := "pool.ha_compute_vm_failover_plan"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	failedHostsArg, err := serializeHostRefSet(fmt.Sprintf("%s(%s)", method, "failed_hosts"), failedHosts)
	if err != nil {
		return
	}
	failedVmsArg, err := serializeVMRefSet(fmt.Sprintf("%s(%s)", method, "failed_vms"), failedVms)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, failedHostsArg, failedVmsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefToStringToStringMapMap(method+" -> ", result)
	return
}

// SetHaHostFailuresToTolerate: Set the maximum number of host failures to consider in the HA VM restart planner
func (pool) SetHaHostFailuresToTolerate(session *Session, self PoolRef, value int) (err error) {
	method := "pool.set_ha_host_failures_to_tolerate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetHaHostFailuresToTolerate: Set the maximum number of host failures to consider in the HA VM restart planner
func (pool) AsyncSetHaHostFailuresToTolerate(session *Session, self PoolRef, value int) (retval TaskRef, err error) {
	method := "Async.pool.set_ha_host_failures_to_tolerate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this pool
func (pool) CreateNewBlob(session *Session, pool PoolRef, name string, mimeType string, public bool) (retval BlobRef, err error) {
	method := "pool.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	poolArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "pool"), pool)
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
	result, err := session.client.sendCall(method, sessionIDArg, poolArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeBlobRef(method+" -> ", result)
	return
}

// AsyncCreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this pool
func (pool) AsyncCreateNewBlob(session *Session, pool PoolRef, name string, mimeType string, public bool) (retval TaskRef, err error) {
	method := "Async.pool.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	poolArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "pool"), pool)
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
	result, err := session.client.sendCall(method, sessionIDArg, poolArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EnableExternalAuth: This call enables external authentication on all the hosts of the pool
func (pool) EnableExternalAuth(session *Session, pool PoolRef, config map[string]string, serviceName string, authType string) (err error) {
	method := "pool.enable_external_auth"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	poolArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "pool"), pool)
	if err != nil {
		return
	}
	configArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "config"), config)
	if err != nil {
		return
	}
	serviceNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "service_name"), serviceName)
	if err != nil {
		return
	}
	authTypeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "auth_type"), authType)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, poolArg, configArg, serviceNameArg, authTypeArg)
	return
}

// DisableExternalAuth: This call disables external authentication on all the hosts of the pool
func (pool) DisableExternalAuth(session *Session, pool PoolRef, config map[string]string) (err error) {
	method := "pool.disable_external_auth"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	poolArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "pool"), pool)
	if err != nil {
		return
	}
	configArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "config"), config)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, poolArg, configArg)
	return
}

// DetectNonhomogeneousExternalAuth: This call asynchronously detects if the external authentication configuration in any slave is different from that in the master and raises appropriate alerts
func (pool) DetectNonhomogeneousExternalAuth(session *Session, pool PoolRef) (err error) {
	method := "pool.detect_nonhomogeneous_external_auth"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	poolArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "pool"), pool)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, poolArg)
	return
}

// InitializeWlb: Initializes workload balancing monitoring on this pool with the specified wlb server
func (pool) InitializeWlb(session *Session, wlbURL string, wlbUsername string, wlbPassword string, xenserverUsername string, xenserverPassword string) (err error) {
	method := "pool.initialize_wlb"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	wlbURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "wlb_url"), wlbURL)
	if err != nil {
		return
	}
	wlbUsernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "wlb_username"), wlbUsername)
	if err != nil {
		return
	}
	wlbPasswordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "wlb_password"), wlbPassword)
	if err != nil {
		return
	}
	xenserverUsernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "xenserver_username"), xenserverUsername)
	if err != nil {
		return
	}
	xenserverPasswordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "xenserver_password"), xenserverPassword)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, wlbURLArg, wlbUsernameArg, wlbPasswordArg, xenserverUsernameArg, xenserverPasswordArg)
	return
}

// AsyncInitializeWlb: Initializes workload balancing monitoring on this pool with the specified wlb server
func (pool) AsyncInitializeWlb(session *Session, wlbURL string, wlbUsername string, wlbPassword string, xenserverUsername string, xenserverPassword string) (retval TaskRef, err error) {
	method := "Async.pool.initialize_wlb"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	wlbURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "wlb_url"), wlbURL)
	if err != nil {
		return
	}
	wlbUsernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "wlb_username"), wlbUsername)
	if err != nil {
		return
	}
	wlbPasswordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "wlb_password"), wlbPassword)
	if err != nil {
		return
	}
	xenserverUsernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "xenserver_username"), xenserverUsername)
	if err != nil {
		return
	}
	xenserverPasswordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "xenserver_password"), xenserverPassword)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, wlbURLArg, wlbUsernameArg, wlbPasswordArg, xenserverUsernameArg, xenserverPasswordArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// DeconfigureWlb: Permanently deconfigures workload balancing monitoring on this pool
func (pool) DeconfigureWlb(session *Session) (err error) {
	method := "pool.deconfigure_wlb"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// AsyncDeconfigureWlb: Permanently deconfigures workload balancing monitoring on this pool
func (pool) AsyncDeconfigureWlb(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.deconfigure_wlb"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SendWlbConfiguration: Sets the pool optimization criteria for the workload balancing server
func (pool) SendWlbConfiguration(session *Session, config map[string]string) (err error) {
	method := "pool.send_wlb_configuration"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	configArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "config"), config)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, configArg)
	return
}

// AsyncSendWlbConfiguration: Sets the pool optimization criteria for the workload balancing server
func (pool) AsyncSendWlbConfiguration(session *Session, config map[string]string) (retval TaskRef, err error) {
	method := "Async.pool.send_wlb_configuration"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	configArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "config"), config)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, configArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// RetrieveWlbConfiguration: Retrieves the pool optimization criteria from the workload balancing server
func (pool) RetrieveWlbConfiguration(session *Session) (retval map[string]string, err error) {
	method := "pool.retrieve_wlb_configuration"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToStringMap(method+" -> ", result)
	return
}

// AsyncRetrieveWlbConfiguration: Retrieves the pool optimization criteria from the workload balancing server
func (pool) AsyncRetrieveWlbConfiguration(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.retrieve_wlb_configuration"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// RetrieveWlbRecommendations: Retrieves vm migrate recommendations for the pool from the workload balancing server
func (pool) RetrieveWlbRecommendations(session *Session) (retval map[VMRef][]string, err error) {
	method := "pool.retrieve_wlb_recommendations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefToStringSetMap(method+" -> ", result)
	return
}

// AsyncRetrieveWlbRecommendations: Retrieves vm migrate recommendations for the pool from the workload balancing server
func (pool) AsyncRetrieveWlbRecommendations(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.retrieve_wlb_recommendations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SendTestPost: Send the given body to the given host and port, using HTTPS, and print the response.  This is used for debugging the SSL layer.
func (pool) SendTestPost(session *Session, host string, port int, body string) (retval string, err error) {
	method := "pool.send_test_post"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	portArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "port"), port)
	if err != nil {
		return
	}
	bodyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "body"), body)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, portArg, bodyArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncSendTestPost: Send the given body to the given host and port, using HTTPS, and print the response.  This is used for debugging the SSL layer.
func (pool) AsyncSendTestPost(session *Session, host string, port int, body string) (retval TaskRef, err error) {
	method := "Async.pool.send_test_post"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	portArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "port"), port)
	if err != nil {
		return
	}
	bodyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "body"), body)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, portArg, bodyArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CertificateInstall: Install a TLS CA certificate, pool-wide.
func (pool) CertificateInstall(session *Session, name string, cert string) (err error) {
	method := "pool.certificate_install"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	certArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cert"), cert)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, nameArg, certArg)
	return
}

// AsyncCertificateInstall: Install a TLS CA certificate, pool-wide.
func (pool) AsyncCertificateInstall(session *Session, name string, cert string) (retval TaskRef, err error) {
	method := "Async.pool.certificate_install"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	certArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cert"), cert)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameArg, certArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CertificateUninstall: Remove a pool-wide TLS CA certificate.
func (pool) CertificateUninstall(session *Session, name string) (err error) {
	method := "pool.certificate_uninstall"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, nameArg)
	return
}

// AsyncCertificateUninstall: Remove a pool-wide TLS CA certificate.
func (pool) AsyncCertificateUninstall(session *Session, name string) (retval TaskRef, err error) {
	method := "Async.pool.certificate_uninstall"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CertificateList: List the names of all installed TLS CA certificates.
func (pool) CertificateList(session *Session) (retval []string, err error) {
	method := "pool.certificate_list"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// AsyncCertificateList: List the names of all installed TLS CA certificates.
func (pool) AsyncCertificateList(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.certificate_list"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// InstallCaCertificate: Install a TLS CA certificate, pool-wide.
func (pool) InstallCaCertificate(session *Session, name string, cert string) (err error) {
	method := "pool.install_ca_certificate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	certArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cert"), cert)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, nameArg, certArg)
	return
}

// AsyncInstallCaCertificate: Install a TLS CA certificate, pool-wide.
func (pool) AsyncInstallCaCertificate(session *Session, name string, cert string) (retval TaskRef, err error) {
	method := "Async.pool.install_ca_certificate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	certArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cert"), cert)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameArg, certArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// UninstallCaCertificate: Remove a pool-wide TLS CA certificate.
func (pool) UninstallCaCertificate(session *Session, name string) (err error) {
	method := "pool.uninstall_ca_certificate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, nameArg)
	return
}

// AsyncUninstallCaCertificate: Remove a pool-wide TLS CA certificate.
func (pool) AsyncUninstallCaCertificate(session *Session, name string) (retval TaskRef, err error) {
	method := "Async.pool.uninstall_ca_certificate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CrlInstall: Install a TLS CA-issued Certificate Revocation List, pool-wide.
func (pool) CrlInstall(session *Session, name string, cert string) (err error) {
	method := "pool.crl_install"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	certArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cert"), cert)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, nameArg, certArg)
	return
}

// AsyncCrlInstall: Install a TLS CA-issued Certificate Revocation List, pool-wide.
func (pool) AsyncCrlInstall(session *Session, name string, cert string) (retval TaskRef, err error) {
	method := "Async.pool.crl_install"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	certArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "cert"), cert)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameArg, certArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CrlUninstall: Remove a pool-wide TLS CA-issued Certificate Revocation List.
func (pool) CrlUninstall(session *Session, name string) (err error) {
	method := "pool.crl_uninstall"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, nameArg)
	return
}

// AsyncCrlUninstall: Remove a pool-wide TLS CA-issued Certificate Revocation List.
func (pool) AsyncCrlUninstall(session *Session, name string) (retval TaskRef, err error) {
	method := "Async.pool.crl_uninstall"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CrlList: List the names of all installed TLS CA-issued Certificate Revocation Lists.
func (pool) CrlList(session *Session) (retval []string, err error) {
	method := "pool.crl_list"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// AsyncCrlList: List the names of all installed TLS CA-issued Certificate Revocation Lists.
func (pool) AsyncCrlList(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.crl_list"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CertificateSync: Copy the TLS CA certificates and CRLs of the master to all slaves.
func (pool) CertificateSync(session *Session) (err error) {
	method := "pool.certificate_sync"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// AsyncCertificateSync: Copy the TLS CA certificates and CRLs of the master to all slaves.
func (pool) AsyncCertificateSync(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.certificate_sync"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EnableTLSVerification: Enable TLS server certificate verification
func (pool) EnableTLSVerification(session *Session) (err error) {
	method := "pool.enable_tls_verification"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// EnableRedoLog: Enable the redo log on the given SR and start using it, unless HA is enabled.
func (pool) EnableRedoLog(session *Session, sr SRRef) (err error) {
	method := "pool.enable_redo_log"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, srArg)
	return
}

// AsyncEnableRedoLog: Enable the redo log on the given SR and start using it, unless HA is enabled.
func (pool) AsyncEnableRedoLog(session *Session, sr SRRef) (retval TaskRef, err error) {
	method := "Async.pool.enable_redo_log"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// DisableRedoLog: Disable the redo log if in use, unless HA is enabled.
func (pool) DisableRedoLog(session *Session) (err error) {
	method := "pool.disable_redo_log"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// AsyncDisableRedoLog: Disable the redo log if in use, unless HA is enabled.
func (pool) AsyncDisableRedoLog(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.disable_redo_log"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetVswitchController: Set the IP address of the vswitch controller.
func (pool) SetVswitchController(session *Session, address string) (err error) {
	method := "pool.set_vswitch_controller"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	addressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "address"), address)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, addressArg)
	return
}

// AsyncSetVswitchController: Set the IP address of the vswitch controller.
func (pool) AsyncSetVswitchController(session *Session, address string) (retval TaskRef, err error) {
	method := "Async.pool.set_vswitch_controller"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	addressArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "address"), address)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, addressArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// TestArchiveTarget: This call tests if a location is valid
func (pool) TestArchiveTarget(session *Session, self PoolRef, config map[string]string) (retval string, err error) {
	method := "pool.test_archive_target"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	configArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "config"), config)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, configArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// EnableLocalStorageCaching: This call attempts to enable pool-wide local storage caching
func (pool) EnableLocalStorageCaching(session *Session, self PoolRef) (err error) {
	method := "pool.enable_local_storage_caching"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncEnableLocalStorageCaching: This call attempts to enable pool-wide local storage caching
func (pool) AsyncEnableLocalStorageCaching(session *Session, self PoolRef) (retval TaskRef, err error) {
	method := "Async.pool.enable_local_storage_caching"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// DisableLocalStorageCaching: This call disables pool-wide local storage caching
func (pool) DisableLocalStorageCaching(session *Session, self PoolRef) (err error) {
	method := "pool.disable_local_storage_caching"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDisableLocalStorageCaching: This call disables pool-wide local storage caching
func (pool) AsyncDisableLocalStorageCaching(session *Session, self PoolRef) (retval TaskRef, err error) {
	method := "Async.pool.disable_local_storage_caching"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLicenseState: This call returns the license state for the pool
func (pool) GetLicenseState(session *Session, self PoolRef) (retval map[string]string, err error) {
	method := "pool.get_license_state"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncGetLicenseState: This call returns the license state for the pool
func (pool) AsyncGetLicenseState(session *Session, self PoolRef) (retval TaskRef, err error) {
	method := "Async.pool.get_license_state"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// ApplyEdition: Apply an edition to all hosts in the pool
func (pool) ApplyEdition(session *Session, self PoolRef, edition string) (err error) {
	method := "pool.apply_edition"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	editionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "edition"), edition)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, editionArg)
	return
}

// AsyncApplyEdition: Apply an edition to all hosts in the pool
func (pool) AsyncApplyEdition(session *Session, self PoolRef, edition string) (retval TaskRef, err error) {
	method := "Async.pool.apply_edition"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	editionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "edition"), edition)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, editionArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EnableSslLegacy: Sets ssl_legacy true on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy.
func (pool) EnableSslLegacy(session *Session, self PoolRef) (err error) {
	method := "pool.enable_ssl_legacy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncEnableSslLegacy: Sets ssl_legacy true on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy.
func (pool) AsyncEnableSslLegacy(session *Session, self PoolRef) (retval TaskRef, err error) {
	method := "Async.pool.enable_ssl_legacy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// DisableSslLegacy: Sets ssl_legacy false on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy.
func (pool) DisableSslLegacy(session *Session, self PoolRef) (err error) {
	method := "pool.disable_ssl_legacy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDisableSslLegacy: Sets ssl_legacy false on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy.
func (pool) AsyncDisableSslLegacy(session *Session, self PoolRef) (retval TaskRef, err error) {
	method := "Async.pool.disable_ssl_legacy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetIgmpSnoopingEnabled: Enable or disable IGMP Snooping on the pool.
func (pool) SetIgmpSnoopingEnabled(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_igmp_snooping_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetIgmpSnoopingEnabled: Enable or disable IGMP Snooping on the pool.
func (pool) AsyncSetIgmpSnoopingEnabled(session *Session, self PoolRef, value bool) (retval TaskRef, err error) {
	method := "Async.pool.set_igmp_snooping_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// HasExtension: Return true if the extension is available on the pool
func (pool) HasExtension(session *Session, self PoolRef, name string) (retval bool, err error) {
	method := "pool.has_extension"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, nameArg)
	if err != nil {
		return
	}
	retval, err = deserializeBool(method+" -> ", result)
	return
}

// AsyncHasExtension: Return true if the extension is available on the pool
func (pool) AsyncHasExtension(session *Session, self PoolRef, name string) (retval TaskRef, err error) {
	method := "Async.pool.has_extension"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, nameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// AddToGuestAgentConfig: Add a key-value pair to the pool-wide guest agent configuration
func (pool) AddToGuestAgentConfig(session *Session, self PoolRef, key string, value string) (err error) {
	method := "pool.add_to_guest_agent_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncAddToGuestAgentConfig: Add a key-value pair to the pool-wide guest agent configuration
func (pool) AsyncAddToGuestAgentConfig(session *Session, self PoolRef, key string, value string) (retval TaskRef, err error) {
	method := "Async.pool.add_to_guest_agent_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, keyArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// RemoveFromGuestAgentConfig: Remove a key-value pair from the pool-wide guest agent configuration
func (pool) RemoveFromGuestAgentConfig(session *Session, self PoolRef, key string) (err error) {
	method := "pool.remove_from_guest_agent_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncRemoveFromGuestAgentConfig: Remove a key-value pair from the pool-wide guest agent configuration
func (pool) AsyncRemoveFromGuestAgentConfig(session *Session, self PoolRef, key string) (retval TaskRef, err error) {
	method := "Async.pool.remove_from_guest_agent_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	keyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "key"), key)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, keyArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// RotateSecret:
//
// Errors:
// INTERNAL_ERROR - The server failed to handle your request, due to an internal error. The given message may give details useful for debugging the problem.
// HOST_IS_SLAVE - You cannot make regular API calls directly on a supporter. Please pass API calls via the coordinator host.
// CANNOT_CONTACT_HOST - Cannot forward messages because the server cannot be contacted. The server may be switched off or there may be network connectivity problems.
// HA_IS_ENABLED - The operation could not be performed because HA is enabled on the Pool
// NOT_SUPPORTED_DURING_UPGRADE - This operation is not supported during an upgrade.
func (pool) RotateSecret(session *Session) (err error) {
	method := "pool.rotate_secret"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// AsyncRotateSecret:
//
// Errors:
// INTERNAL_ERROR - The server failed to handle your request, due to an internal error. The given message may give details useful for debugging the problem.
// HOST_IS_SLAVE - You cannot make regular API calls directly on a supporter. Please pass API calls via the coordinator host.
// CANNOT_CONTACT_HOST - Cannot forward messages because the server cannot be contacted. The server may be switched off or there may be network connectivity problems.
// HA_IS_ENABLED - The operation could not be performed because HA is enabled on the Pool
// NOT_SUPPORTED_DURING_UPGRADE - This operation is not supported during an upgrade.
func (pool) AsyncRotateSecret(session *Session) (retval TaskRef, err error) {
	method := "Async.pool.rotate_secret"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetRepositories: Set enabled set of repositories
func (pool) SetRepositories(session *Session, self PoolRef, value []RepositoryRef) (err error) {
	method := "pool.set_repositories"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeRepositoryRefSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetRepositories: Set enabled set of repositories
func (pool) AsyncSetRepositories(session *Session, self PoolRef, value []RepositoryRef) (retval TaskRef, err error) {
	method := "Async.pool.set_repositories"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeRepositoryRefSet(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// AddRepository: Add a repository to the enabled set
func (pool) AddRepository(session *Session, self PoolRef, value RepositoryRef) (err error) {
	method := "pool.add_repository"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncAddRepository: Add a repository to the enabled set
func (pool) AsyncAddRepository(session *Session, self PoolRef, value RepositoryRef) (retval TaskRef, err error) {
	method := "Async.pool.add_repository"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// RemoveRepository: Remove a repository from the enabled set
func (pool) RemoveRepository(session *Session, self PoolRef, value RepositoryRef) (err error) {
	method := "pool.remove_repository"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncRemoveRepository: Remove a repository from the enabled set
func (pool) AsyncRemoveRepository(session *Session, self PoolRef, value RepositoryRef) (retval TaskRef, err error) {
	method := "Async.pool.remove_repository"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SyncUpdates: Sync with the enabled repository
func (pool) SyncUpdates(session *Session, self PoolRef, force bool, token string, tokenID string) (retval string, err error) {
	method := "pool.sync_updates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	tokenArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "token"), token)
	if err != nil {
		return
	}
	tokenIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "token_id"), tokenID)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, forceArg, tokenArg, tokenIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncSyncUpdates: Sync with the enabled repository
func (pool) AsyncSyncUpdates(session *Session, self PoolRef, force bool, token string, tokenID string) (retval TaskRef, err error) {
	method := "Async.pool.sync_updates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	tokenArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "token"), token)
	if err != nil {
		return
	}
	tokenIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "token_id"), tokenID)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, forceArg, tokenArg, tokenIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CheckUpdateReadiness: Check if the pool is ready to be updated. If not, report the reasons.
func (pool) CheckUpdateReadiness(session *Session, self PoolRef, requiresReboot bool) (retval [][]string, err error) {
	method := "pool.check_update_readiness"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	requiresRebootArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "requires_reboot"), requiresReboot)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, requiresRebootArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSetSet(method+" -> ", result)
	return
}

// AsyncCheckUpdateReadiness: Check if the pool is ready to be updated. If not, report the reasons.
func (pool) AsyncCheckUpdateReadiness(session *Session, self PoolRef, requiresReboot bool) (retval TaskRef, err error) {
	method := "Async.pool.check_update_readiness"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	requiresRebootArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "requires_reboot"), requiresReboot)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, requiresRebootArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EnableClientCertificateAuth: Enable client certificate authentication on the pool
func (pool) EnableClientCertificateAuth(session *Session, self PoolRef, name string) (err error) {
	method := "pool.enable_client_certificate_auth"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, nameArg)
	return
}

// AsyncEnableClientCertificateAuth: Enable client certificate authentication on the pool
func (pool) AsyncEnableClientCertificateAuth(session *Session, self PoolRef, name string) (retval TaskRef, err error) {
	method := "Async.pool.enable_client_certificate_auth"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, nameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// DisableClientCertificateAuth: Disable client certificate authentication on the pool
func (pool) DisableClientCertificateAuth(session *Session, self PoolRef) (err error) {
	method := "pool.disable_client_certificate_auth"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDisableClientCertificateAuth: Disable client certificate authentication on the pool
func (pool) AsyncDisableClientCertificateAuth(session *Session, self PoolRef) (retval TaskRef, err error) {
	method := "Async.pool.disable_client_certificate_auth"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// ConfigureRepositoryProxy: Configure proxy for RPM package repositories.
func (pool) ConfigureRepositoryProxy(session *Session, self PoolRef, uRL string, username string, password string) (err error) {
	method := "pool.configure_repository_proxy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	uRLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "url"), uRL)
	if err != nil {
		return
	}
	usernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "username"), username)
	if err != nil {
		return
	}
	passwordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "password"), password)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, uRLArg, usernameArg, passwordArg)
	return
}

// AsyncConfigureRepositoryProxy: Configure proxy for RPM package repositories.
func (pool) AsyncConfigureRepositoryProxy(session *Session, self PoolRef, uRL string, username string, password string) (retval TaskRef, err error) {
	method := "Async.pool.configure_repository_proxy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	uRLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "url"), uRL)
	if err != nil {
		return
	}
	usernameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "username"), username)
	if err != nil {
		return
	}
	passwordArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "password"), password)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, uRLArg, usernameArg, passwordArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// DisableRepositoryProxy: Disable the proxy for RPM package repositories.
func (pool) DisableRepositoryProxy(session *Session, self PoolRef) (err error) {
	method := "pool.disable_repository_proxy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDisableRepositoryProxy: Disable the proxy for RPM package repositories.
func (pool) AsyncDisableRepositoryProxy(session *Session, self PoolRef) (retval TaskRef, err error) {
	method := "Async.pool.disable_repository_proxy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetUefiCertificates: Set the UEFI certificates for a pool and all its hosts. Deprecated: use set_custom_uefi_certificates instead
func (pool) SetUefiCertificates(session *Session, self PoolRef, value string) (err error) {
	method := "pool.set_uefi_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetUefiCertificates: Set the UEFI certificates for a pool and all its hosts. Deprecated: use set_custom_uefi_certificates instead
func (pool) AsyncSetUefiCertificates(session *Session, self PoolRef, value string) (retval TaskRef, err error) {
	method := "Async.pool.set_uefi_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetCustomUefiCertificates: Set custom UEFI certificates for a pool and all its hosts. Need `allow-custom-uefi-certs` set to true in conf. If empty: default back to Pool.uefi_certificates
func (pool) SetCustomUefiCertificates(session *Session, self PoolRef, value string) (err error) {
	method := "pool.set_custom_uefi_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetCustomUefiCertificates: Set custom UEFI certificates for a pool and all its hosts. Need `allow-custom-uefi-certs` set to true in conf. If empty: default back to Pool.uefi_certificates
func (pool) AsyncSetCustomUefiCertificates(session *Session, self PoolRef, value string) (retval TaskRef, err error) {
	method := "Async.pool.set_custom_uefi_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetHTTPSOnly: updates all the host firewalls in the pool to open or close port 80 depending on the value
func (pool) SetHTTPSOnly(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_https_only"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetHTTPSOnly: updates all the host firewalls in the pool to open or close port 80 depending on the value
func (pool) AsyncSetHTTPSOnly(session *Session, self PoolRef, value bool) (retval TaskRef, err error) {
	method := "Async.pool.set_https_only"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetTelemetryNextCollection: Set the timestamp for the next telemetry data collection.
func (pool) SetTelemetryNextCollection(session *Session, self PoolRef, value time.Time) (err error) {
	method := "pool.set_telemetry_next_collection"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeTime(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetTelemetryNextCollection: Set the timestamp for the next telemetry data collection.
func (pool) AsyncSetTelemetryNextCollection(session *Session, self PoolRef, value time.Time) (retval TaskRef, err error) {
	method := "Async.pool.set_telemetry_next_collection"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeTime(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// ResetTelemetryUUID: Assign a new UUID to telemetry data.
func (pool) ResetTelemetryUUID(session *Session, self PoolRef) (err error) {
	method := "pool.reset_telemetry_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncResetTelemetryUUID: Assign a new UUID to telemetry data.
func (pool) AsyncResetTelemetryUUID(session *Session, self PoolRef) (retval TaskRef, err error) {
	method := "Async.pool.reset_telemetry_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// ConfigureUpdateSync: Configure periodic update synchronization to sync updates from a remote CDN
func (pool) ConfigureUpdateSync(session *Session, self PoolRef, updateSyncFrequency UpdateSyncFrequency, updateSyncDay int) (err error) {
	method := "pool.configure_update_sync"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	updateSyncFrequencyArg, err := serializeEnumUpdateSyncFrequency(fmt.Sprintf("%s(%s)", method, "update_sync_frequency"), updateSyncFrequency)
	if err != nil {
		return
	}
	updateSyncDayArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "update_sync_day"), updateSyncDay)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, updateSyncFrequencyArg, updateSyncDayArg)
	return
}

// AsyncConfigureUpdateSync: Configure periodic update synchronization to sync updates from a remote CDN
func (pool) AsyncConfigureUpdateSync(session *Session, self PoolRef, updateSyncFrequency UpdateSyncFrequency, updateSyncDay int) (retval TaskRef, err error) {
	method := "Async.pool.configure_update_sync"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	updateSyncFrequencyArg, err := serializeEnumUpdateSyncFrequency(fmt.Sprintf("%s(%s)", method, "update_sync_frequency"), updateSyncFrequency)
	if err != nil {
		return
	}
	updateSyncDayArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "update_sync_day"), updateSyncDay)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, updateSyncFrequencyArg, updateSyncDayArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetUpdateSyncEnabled: enable or disable periodic update synchronization depending on the value
func (pool) SetUpdateSyncEnabled(session *Session, self PoolRef, value bool) (err error) {
	method := "pool.set_update_sync_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetUpdateSyncEnabled: enable or disable periodic update synchronization depending on the value
func (pool) AsyncSetUpdateSyncEnabled(session *Session, self PoolRef, value bool) (retval TaskRef, err error) {
	method := "Async.pool.set_update_sync_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetLocalAuthMaxThreads:
func (pool) SetLocalAuthMaxThreads(session *Session, self PoolRef, value int) (err error) {
	method := "pool.set_local_auth_max_threads"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetExtAuthMaxThreads:
func (pool) SetExtAuthMaxThreads(session *Session, self PoolRef, value int) (err error) {
	method := "pool.set_ext_auth_max_threads"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePoolRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the pools known to the system.
func (pool) GetAll(session *Session) (retval []PoolRef, err error) {
	method := "pool.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of pool references to pool records for all pools known to the system.
func (pool) GetAllRecords(session *Session) (retval map[PoolRef]PoolRecord, err error) {
	method := "pool.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolRefToPoolRecordMap(method+" -> ", result)
	return
}
