package xenapi

import (
	"fmt"
	"time"
)

type HostRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// Virtualization memory overhead (bytes).
	MemoryOverhead int
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []HostAllowedOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]HostAllowedOperations
	// major version number
	APIVersionMajor int
	// minor version number
	APIVersionMinor int
	// identification of vendor
	APIVersionVendor string
	// details of vendor implementation
	APIVersionVendorImplementation map[string]string
	// True if the host is currently enabled
	Enabled bool
	// version strings
	SoftwareVersion map[string]string
	// additional configuration
	OtherConfig map[string]string
	// Xen capabilities
	Capabilities []string
	// The CPU configuration on this host.  May contain keys such as &quot;nr_nodes&quot;, &quot;sockets_per_node&quot;, &quot;cores_per_socket&quot;, or &quot;threads_per_core&quot;
	CPUConfiguration map[string]string
	// Scheduler policy currently in force on this host
	SchedPolicy string
	// a list of the bootloaders installed on the machine
	SupportedBootloaders []string
	// list of VMs currently resident on host
	ResidentVMs []VMRef
	// logging configuration
	Logging map[string]string
	// physical network interfaces
	PIFs []PIFRef
	// The SR in which VDIs for suspend images are created
	SuspendImageSr SRRef
	// The SR in which VDIs for crash dumps are created
	CrashDumpSr SRRef
	// Set of host crash dumps
	Crashdumps []HostCrashdumpRef
	// Set of host patches
	Patches []HostPatchRef
	// Set of updates
	Updates []PoolUpdateRef
	// physical blockdevices
	PBDs []PBDRef
	// The physical CPUs on this host
	HostCPUs []HostCPURef
	// Details about the physical CPUs on this host
	CPUInfo map[string]string
	// The hostname of this host
	Hostname string
	// The address by which this host can be contacted from any other host in the pool
	Address string
	// metrics associated with this host
	Metrics HostMetricsRef
	// State of the current license
	LicenseParams map[string]string
	// The set of statefiles accessible from this host
	HaStatefiles []string
	// The set of hosts visible via the network from this host
	HaNetworkPeers []string
	// Binary blobs associated with this host
	Blobs map[string]BlobRef
	// user-specified tags for categorization purposes
	Tags []string
	// type of external authentication service configured; empty if none configured.
	ExternalAuthType string
	// name of external authentication service configured; empty if none configured.
	ExternalAuthServiceName string
	// configuration specific to external authentication service
	ExternalAuthConfiguration map[string]string
	// Product edition
	Edition string
	// Contact information of the license server
	LicenseServer map[string]string
	// BIOS strings
	BiosStrings map[string]string
	// The power on mode
	PowerOnMode string
	// The power on config
	PowerOnConfig map[string]string
	// The SR that is used as a local cache
	LocalCacheSr SRRef
	// Information about chipset features
	ChipsetInfo map[string]string
	// List of PCI devices in the host
	PCIs []PCIRef
	// List of physical GPUs in the host
	PGPUs []PGPURef
	// List of physical USBs in the host
	PUSBs []PUSBRef
	// Allow SSLv3 protocol and ciphersuites as used by older server versions. This controls both incoming and outgoing connections. When this is set to a different value, the host immediately restarts its SSL/TLS listening service; typically this takes less than a second but existing connections to it will be broken. API login sessions will remain valid.
	SslLegacy bool
	// VCPUs params to apply to all resident guests
	GuestVCPUsParams map[string]string
	// indicates whether the host is configured to output its console to a physical display device
	Display HostDisplay
	// The set of versions of the virtual hardware platform that the host can offer to its guests
	VirtualHardwarePlatformVersions []int
	// The control domain (domain 0)
	ControlDomain VMRef
	// List of updates which require reboot
	UpdatesRequiringReboot []PoolUpdateRef
	// List of features available on this host
	Features []FeatureRef
	// The initiator IQN for the host
	IscsiIqn string
	// Specifies whether multipathing is enabled
	Multipathing bool
	// The UEFI certificates allowing Secure Boot
	UefiCertificates string
	// List of certificates installed in the host
	Certificates []CertificateRef
	// List of all available product editions
	Editions []string
	// The set of pending mandatory guidances after applying updates, which must be applied, as otherwise there may be e.g. VM failures
	PendingGuidances []UpdateGuidances
	// True if this host has TLS verifcation enabled
	TLSVerificationEnabled bool
	// Date and time when the last software update was applied
	LastSoftwareUpdate time.Time
	// Reflects whether port 80 is open (false) or not (true)
	HTTPSOnly bool
	// Default as &apos;unknown&apos;, &apos;yes&apos; if the host is up to date with updates synced from remote CDN, otherwise &apos;no&apos;
	LatestSyncedUpdatesApplied LatestSyncedUpdatesAppliedState
	// NUMA-aware VM memory and vCPU placement policy
	NumaAffinityPolicy HostNumaAffinityPolicy
	// The set of pending recommended guidances after applying updates, which most users should follow to make the updates effective, but if not followed, will not cause a failure
	PendingGuidancesRecommended []UpdateGuidances
	// The set of pending full guidances after applying updates, which a user should follow to make some updates, e.g. specific hardware drivers or CPU features, fully effective, but the &apos;average user&apos; doesn&apos;t need to
	PendingGuidancesFull []UpdateGuidances
	// The SHA256 checksum of updateinfo of the most recently applied update on the host
	LastUpdateHash string
}

type HostRef string

// A physical host
type host struct{}

var Host host

// GetRecord: Get a record containing the current state of the given host.
func (host) GetRecord(session *Session, self HostRef) (retval HostRecord, err error) {
	method := "host.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the host instance with the specified UUID.
func (host) GetByUUID(session *Session, uUID string) (retval HostRef, err error) {
	method := "host.get_by_uuid"
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
	retval, err = deserializeHostRef(method+" -> ", result)
	return
}

// GetByNameLabel: Get all the host instances with the given label.
func (host) GetByNameLabel(session *Session, label string) (retval []HostRef, err error) {
	method := "host.get_by_name_label"
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
	retval, err = deserializeHostRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given host.
func (host) GetUUID(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given host.
func (host) GetNameLabel(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given host.
func (host) GetNameDescription(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryOverhead: Get the memory/overhead field of the given host.
func (host) GetMemoryOverhead(session *Session, self HostRef) (retval int, err error) {
	method := "host.get_memory_overhead"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given host.
func (host) GetAllowedOperations(session *Session, self HostRef) (retval []HostAllowedOperations, err error) {
	method := "host.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumHostAllowedOperationsSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given host.
func (host) GetCurrentOperations(session *Session, self HostRef) (retval map[string]HostAllowedOperations, err error) {
	method := "host.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumHostAllowedOperationsMap(method+" -> ", result)
	return
}

// GetAPIVersionMajor: Get the API_version/major field of the given host.
func (host) GetAPIVersionMajor(session *Session, self HostRef) (retval int, err error) {
	method := "host.get_API_version_major"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAPIVersionMinor: Get the API_version/minor field of the given host.
func (host) GetAPIVersionMinor(session *Session, self HostRef) (retval int, err error) {
	method := "host.get_API_version_minor"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAPIVersionVendor: Get the API_version/vendor field of the given host.
func (host) GetAPIVersionVendor(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_API_version_vendor"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAPIVersionVendorImplementation: Get the API_version/vendor_implementation field of the given host.
func (host) GetAPIVersionVendorImplementation(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_API_version_vendor_implementation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEnabled: Get the enabled field of the given host.
func (host) GetEnabled(session *Session, self HostRef) (retval bool, err error) {
	method := "host.get_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSoftwareVersion: Get the software_version field of the given host.
func (host) GetSoftwareVersion(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_software_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given host.
func (host) GetOtherConfig(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCapabilities: Get the capabilities field of the given host.
func (host) GetCapabilities(session *Session, self HostRef) (retval []string, err error) {
	method := "host.get_capabilities"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCPUConfiguration: Get the cpu_configuration field of the given host.
func (host) GetCPUConfiguration(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_cpu_configuration"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSchedPolicy: Get the sched_policy field of the given host.
func (host) GetSchedPolicy(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_sched_policy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSupportedBootloaders: Get the supported_bootloaders field of the given host.
func (host) GetSupportedBootloaders(session *Session, self HostRef) (retval []string, err error) {
	method := "host.get_supported_bootloaders"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetResidentVMs: Get the resident_VMs field of the given host.
func (host) GetResidentVMs(session *Session, self HostRef) (retval []VMRef, err error) {
	method := "host.get_resident_VMs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefSet(method+" -> ", result)
	return
}

// GetLogging: Get the logging field of the given host.
func (host) GetLogging(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_logging"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPIFs: Get the PIFs field of the given host.
func (host) GetPIFs(session *Session, self HostRef) (retval []PIFRef, err error) {
	method := "host.get_PIFs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSuspendImageSr: Get the suspend_image_sr field of the given host.
func (host) GetSuspendImageSr(session *Session, self HostRef) (retval SRRef, err error) {
	method := "host.get_suspend_image_sr"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCrashDumpSr: Get the crash_dump_sr field of the given host.
func (host) GetCrashDumpSr(session *Session, self HostRef) (retval SRRef, err error) {
	method := "host.get_crash_dump_sr"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCrashdumps: Get the crashdumps field of the given host.
func (host) GetCrashdumps(session *Session, self HostRef) (retval []HostCrashdumpRef, err error) {
	method := "host.get_crashdumps"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostCrashdumpRefSet(method+" -> ", result)
	return
}

// GetPatches: Get the patches field of the given host.
func (host) GetPatches(session *Session, self HostRef) (retval []HostPatchRef, err error) {
	method := "host.get_patches"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostPatchRefSet(method+" -> ", result)
	return
}

// GetUpdates: Get the updates field of the given host.
func (host) GetUpdates(session *Session, self HostRef) (retval []PoolUpdateRef, err error) {
	method := "host.get_updates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolUpdateRefSet(method+" -> ", result)
	return
}

// GetPBDs: Get the PBDs field of the given host.
func (host) GetPBDs(session *Session, self HostRef) (retval []PBDRef, err error) {
	method := "host.get_PBDs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePBDRefSet(method+" -> ", result)
	return
}

// GetHostCPUs: Get the host_CPUs field of the given host.
func (host) GetHostCPUs(session *Session, self HostRef) (retval []HostCPURef, err error) {
	method := "host.get_host_CPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostCPURefSet(method+" -> ", result)
	return
}

// GetCPUInfo: Get the cpu_info field of the given host.
func (host) GetCPUInfo(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_cpu_info"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHostname: Get the hostname field of the given host.
func (host) GetHostname(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_hostname"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAddress: Get the address field of the given host.
func (host) GetAddress(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_address"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMetrics: Get the metrics field of the given host.
func (host) GetMetrics(session *Session, self HostRef) (retval HostMetricsRef, err error) {
	method := "host.get_metrics"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostMetricsRef(method+" -> ", result)
	return
}

// GetLicenseParams: Get the license_params field of the given host.
func (host) GetLicenseParams(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_license_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaStatefiles: Get the ha_statefiles field of the given host.
func (host) GetHaStatefiles(session *Session, self HostRef) (retval []string, err error) {
	method := "host.get_ha_statefiles"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaNetworkPeers: Get the ha_network_peers field of the given host.
func (host) GetHaNetworkPeers(session *Session, self HostRef) (retval []string, err error) {
	method := "host.get_ha_network_peers"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBlobs: Get the blobs field of the given host.
func (host) GetBlobs(session *Session, self HostRef) (retval map[string]BlobRef, err error) {
	method := "host.get_blobs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTags: Get the tags field of the given host.
func (host) GetTags(session *Session, self HostRef) (retval []string, err error) {
	method := "host.get_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetExternalAuthType: Get the external_auth_type field of the given host.
func (host) GetExternalAuthType(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_external_auth_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetExternalAuthServiceName: Get the external_auth_service_name field of the given host.
func (host) GetExternalAuthServiceName(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_external_auth_service_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetExternalAuthConfiguration: Get the external_auth_configuration field of the given host.
func (host) GetExternalAuthConfiguration(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_external_auth_configuration"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetEdition: Get the edition field of the given host.
func (host) GetEdition(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_edition"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLicenseServer: Get the license_server field of the given host.
func (host) GetLicenseServer(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_license_server"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBiosStrings: Get the bios_strings field of the given host.
func (host) GetBiosStrings(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_bios_strings"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPowerOnMode: Get the power_on_mode field of the given host.
func (host) GetPowerOnMode(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_power_on_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPowerOnConfig: Get the power_on_config field of the given host.
func (host) GetPowerOnConfig(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_power_on_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLocalCacheSr: Get the local_cache_sr field of the given host.
func (host) GetLocalCacheSr(session *Session, self HostRef) (retval SRRef, err error) {
	method := "host.get_local_cache_sr"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetChipsetInfo: Get the chipset_info field of the given host.
func (host) GetChipsetInfo(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_chipset_info"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPCIs: Get the PCIs field of the given host.
func (host) GetPCIs(session *Session, self HostRef) (retval []PCIRef, err error) {
	method := "host.get_PCIs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPGPUs: Get the PGPUs field of the given host.
func (host) GetPGPUs(session *Session, self HostRef) (retval []PGPURef, err error) {
	method := "host.get_PGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePGPURefSet(method+" -> ", result)
	return
}

// GetPUSBs: Get the PUSBs field of the given host.
func (host) GetPUSBs(session *Session, self HostRef) (retval []PUSBRef, err error) {
	method := "host.get_PUSBs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePUSBRefSet(method+" -> ", result)
	return
}

// GetSslLegacy: Get the ssl_legacy field of the given host.
func (host) GetSslLegacy(session *Session, self HostRef) (retval bool, err error) {
	method := "host.get_ssl_legacy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetGuestVCPUsParams: Get the guest_VCPUs_params field of the given host.
func (host) GetGuestVCPUsParams(session *Session, self HostRef) (retval map[string]string, err error) {
	method := "host.get_guest_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDisplay: Get the display field of the given host.
func (host) GetDisplay(session *Session, self HostRef) (retval HostDisplay, err error) {
	method := "host.get_display"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumHostDisplay(method+" -> ", result)
	return
}

// GetVirtualHardwarePlatformVersions: Get the virtual_hardware_platform_versions field of the given host.
func (host) GetVirtualHardwarePlatformVersions(session *Session, self HostRef) (retval []int, err error) {
	method := "host.get_virtual_hardware_platform_versions"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeIntSet(method+" -> ", result)
	return
}

// GetControlDomain: Get the control_domain field of the given host.
func (host) GetControlDomain(session *Session, self HostRef) (retval VMRef, err error) {
	method := "host.get_control_domain"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// GetUpdatesRequiringReboot: Get the updates_requiring_reboot field of the given host.
func (host) GetUpdatesRequiringReboot(session *Session, self HostRef) (retval []PoolUpdateRef, err error) {
	method := "host.get_updates_requiring_reboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePoolUpdateRefSet(method+" -> ", result)
	return
}

// GetFeatures: Get the features field of the given host.
func (host) GetFeatures(session *Session, self HostRef) (retval []FeatureRef, err error) {
	method := "host.get_features"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeFeatureRefSet(method+" -> ", result)
	return
}

// GetIscsiIqn: Get the iscsi_iqn field of the given host.
func (host) GetIscsiIqn(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_iscsi_iqn"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMultipathing: Get the multipathing field of the given host.
func (host) GetMultipathing(session *Session, self HostRef) (retval bool, err error) {
	method := "host.get_multipathing"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUefiCertificates: Get the uefi_certificates field of the given host.
func (host) GetUefiCertificates(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_uefi_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCertificates: Get the certificates field of the given host.
func (host) GetCertificates(session *Session, self HostRef) (retval []CertificateRef, err error) {
	method := "host.get_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeCertificateRefSet(method+" -> ", result)
	return
}

// GetEditions: Get the editions field of the given host.
func (host) GetEditions(session *Session, self HostRef) (retval []string, err error) {
	method := "host.get_editions"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPendingGuidances: Get the pending_guidances field of the given host.
func (host) GetPendingGuidances(session *Session, self HostRef) (retval []UpdateGuidances, err error) {
	method := "host.get_pending_guidances"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumUpdateGuidancesSet(method+" -> ", result)
	return
}

// GetTLSVerificationEnabled: Get the tls_verification_enabled field of the given host.
func (host) GetTLSVerificationEnabled(session *Session, self HostRef) (retval bool, err error) {
	method := "host.get_tls_verification_enabled"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastSoftwareUpdate: Get the last_software_update field of the given host.
func (host) GetLastSoftwareUpdate(session *Session, self HostRef) (retval time.Time, err error) {
	method := "host.get_last_software_update"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHTTPSOnly: Get the https_only field of the given host.
func (host) GetHTTPSOnly(session *Session, self HostRef) (retval bool, err error) {
	method := "host.get_https_only"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLatestSyncedUpdatesApplied: Get the latest_synced_updates_applied field of the given host.
func (host) GetLatestSyncedUpdatesApplied(session *Session, self HostRef) (retval LatestSyncedUpdatesAppliedState, err error) {
	method := "host.get_latest_synced_updates_applied"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumLatestSyncedUpdatesAppliedState(method+" -> ", result)
	return
}

// GetNumaAffinityPolicy: Get the numa_affinity_policy field of the given host.
func (host) GetNumaAffinityPolicy(session *Session, self HostRef) (retval HostNumaAffinityPolicy, err error) {
	method := "host.get_numa_affinity_policy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumHostNumaAffinityPolicy(method+" -> ", result)
	return
}

// GetPendingGuidancesRecommended: Get the pending_guidances_recommended field of the given host.
func (host) GetPendingGuidancesRecommended(session *Session, self HostRef) (retval []UpdateGuidances, err error) {
	method := "host.get_pending_guidances_recommended"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumUpdateGuidancesSet(method+" -> ", result)
	return
}

// GetPendingGuidancesFull: Get the pending_guidances_full field of the given host.
func (host) GetPendingGuidancesFull(session *Session, self HostRef) (retval []UpdateGuidances, err error) {
	method := "host.get_pending_guidances_full"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumUpdateGuidancesSet(method+" -> ", result)
	return
}

// GetLastUpdateHash: Get the last_update_hash field of the given host.
func (host) GetLastUpdateHash(session *Session, self HostRef) (retval string, err error) {
	method := "host.get_last_update_hash"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel: Set the name/label field of the given host.
func (host) SetNameLabel(session *Session, self HostRef, value string) (err error) {
	method := "host.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription: Set the name/description field of the given host.
func (host) SetNameDescription(session *Session, self HostRef, value string) (err error) {
	method := "host.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given host.
func (host) SetOtherConfig(session *Session, self HostRef, value map[string]string) (err error) {
	method := "host.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given host.
func (host) AddToOtherConfig(session *Session, self HostRef, key string, value string) (err error) {
	method := "host.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given host.  If the key is not in that Map, then do nothing.
func (host) RemoveFromOtherConfig(session *Session, self HostRef, key string) (err error) {
	method := "host.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetLogging: Set the logging field of the given host.
func (host) SetLogging(session *Session, self HostRef, value map[string]string) (err error) {
	method := "host.set_logging"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToLogging: Add the given key-value pair to the logging field of the given host.
func (host) AddToLogging(session *Session, self HostRef, key string, value string) (err error) {
	method := "host.add_to_logging"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromLogging: Remove the given key and its corresponding value from the logging field of the given host.  If the key is not in that Map, then do nothing.
func (host) RemoveFromLogging(session *Session, self HostRef, key string) (err error) {
	method := "host.remove_from_logging"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetSuspendImageSr: Set the suspend_image_sr field of the given host.
func (host) SetSuspendImageSr(session *Session, self HostRef, value SRRef) (err error) {
	method := "host.set_suspend_image_sr"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetCrashDumpSr: Set the crash_dump_sr field of the given host.
func (host) SetCrashDumpSr(session *Session, self HostRef, value SRRef) (err error) {
	method := "host.set_crash_dump_sr"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetHostname: Set the hostname field of the given host.
func (host) SetHostname(session *Session, self HostRef, value string) (err error) {
	method := "host.set_hostname"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetAddress: Set the address field of the given host.
func (host) SetAddress(session *Session, self HostRef, value string) (err error) {
	method := "host.set_address"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetTags: Set the tags field of the given host.
func (host) SetTags(session *Session, self HostRef, value []string) (err error) {
	method := "host.set_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddTags: Add the given value to the tags field of the given host.  If the value is already in that Set, then do nothing.
func (host) AddTags(session *Session, self HostRef, value string) (err error) {
	method := "host.add_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveTags: Remove the given value from the tags field of the given host.  If the value is not in that Set, then do nothing.
func (host) RemoveTags(session *Session, self HostRef, value string) (err error) {
	method := "host.remove_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetLicenseServer: Set the license_server field of the given host.
func (host) SetLicenseServer(session *Session, self HostRef, value map[string]string) (err error) {
	method := "host.set_license_server"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToLicenseServer: Add the given key-value pair to the license_server field of the given host.
func (host) AddToLicenseServer(session *Session, self HostRef, key string, value string) (err error) {
	method := "host.add_to_license_server"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromLicenseServer: Remove the given key and its corresponding value from the license_server field of the given host.  If the key is not in that Map, then do nothing.
func (host) RemoveFromLicenseServer(session *Session, self HostRef, key string) (err error) {
	method := "host.remove_from_license_server"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetGuestVCPUsParams: Set the guest_VCPUs_params field of the given host.
func (host) SetGuestVCPUsParams(session *Session, self HostRef, value map[string]string) (err error) {
	method := "host.set_guest_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToGuestVCPUsParams: Add the given key-value pair to the guest_VCPUs_params field of the given host.
func (host) AddToGuestVCPUsParams(session *Session, self HostRef, key string, value string) (err error) {
	method := "host.add_to_guest_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromGuestVCPUsParams: Remove the given key and its corresponding value from the guest_VCPUs_params field of the given host.  If the key is not in that Map, then do nothing.
func (host) RemoveFromGuestVCPUsParams(session *Session, self HostRef, key string) (err error) {
	method := "host.remove_from_guest_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetDisplay: Set the display field of the given host.
func (host) SetDisplay(session *Session, self HostRef, value HostDisplay) (err error) {
	method := "host.set_display"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumHostDisplay(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// Disable: Puts the host into a state in which no new VMs can be started. Currently active VMs on the host continue to execute.
func (host) Disable(session *Session, host HostRef) (err error) {
	method := "host.disable"
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

// AsyncDisable: Puts the host into a state in which no new VMs can be started. Currently active VMs on the host continue to execute.
func (host) AsyncDisable(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.disable"
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

// Enable: Puts the host into a state in which new VMs can be started.
func (host) Enable(session *Session, host HostRef) (err error) {
	method := "host.enable"
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

// AsyncEnable: Puts the host into a state in which new VMs can be started.
func (host) AsyncEnable(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.enable"
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

// Shutdown: Shutdown the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
func (host) Shutdown(session *Session, host HostRef) (err error) {
	method := "host.shutdown"
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

// AsyncShutdown: Shutdown the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
func (host) AsyncShutdown(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.shutdown"
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

// Reboot: Reboot the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
func (host) Reboot(session *Session, host HostRef) (err error) {
	method := "host.reboot"
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

// AsyncReboot: Reboot the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
func (host) AsyncReboot(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.reboot"
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

// Dmesg: Get the host xen dmesg.
func (host) Dmesg(session *Session, host HostRef) (retval string, err error) {
	method := "host.dmesg"
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
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncDmesg: Get the host xen dmesg.
func (host) AsyncDmesg(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.dmesg"
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

// DmesgClear: Get the host xen dmesg, and clear the buffer.
func (host) DmesgClear(session *Session, host HostRef) (retval string, err error) {
	method := "host.dmesg_clear"
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
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncDmesgClear: Get the host xen dmesg, and clear the buffer.
func (host) AsyncDmesgClear(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.dmesg_clear"
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

// GetLog: Get the host&apos;s log file
func (host) GetLog(session *Session, host HostRef) (retval string, err error) {
	method := "host.get_log"
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
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncGetLog: Get the host&apos;s log file
func (host) AsyncGetLog(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.get_log"
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

// SendDebugKeys: Inject the given string as debugging keys into Xen
func (host) SendDebugKeys(session *Session, host HostRef, keys string) (err error) {
	method := "host.send_debug_keys"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	keysArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "keys"), keys)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, keysArg)
	return
}

// AsyncSendDebugKeys: Inject the given string as debugging keys into Xen
func (host) AsyncSendDebugKeys(session *Session, host HostRef, keys string) (retval TaskRef, err error) {
	method := "Async.host.send_debug_keys"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	keysArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "keys"), keys)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, keysArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// BugreportUpload: Run xen-bugtool --yestoall and upload the output to support
func (host) BugreportUpload(session *Session, host HostRef, uRL string, options map[string]string) (err error) {
	method := "host.bugreport_upload"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	uRLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "url"), uRL)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, uRLArg, optionsArg)
	return
}

// AsyncBugreportUpload: Run xen-bugtool --yestoall and upload the output to support
func (host) AsyncBugreportUpload(session *Session, host HostRef, uRL string, options map[string]string) (retval TaskRef, err error) {
	method := "Async.host.bugreport_upload"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	uRLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "url"), uRL)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, uRLArg, optionsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// ListMethods: List all supported methods
func (host) ListMethods(session *Session) (retval []string, err error) {
	method := "host.list_methods"
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

// LicenseApply: Apply a new license to a host
//
// Errors:
// LICENSE_PROCESSING_ERROR - There was an error processing your license. Please contact your support representative.
func (host) LicenseApply(session *Session, host HostRef, contents string) (err error) {
	method := "host.license_apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	contentsArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "contents"), contents)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, contentsArg)
	return
}

// AsyncLicenseApply: Apply a new license to a host
//
// Errors:
// LICENSE_PROCESSING_ERROR - There was an error processing your license. Please contact your support representative.
func (host) AsyncLicenseApply(session *Session, host HostRef, contents string) (retval TaskRef, err error) {
	method := "Async.host.license_apply"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	contentsArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "contents"), contents)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, contentsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// LicenseAdd: Apply a new license to a host
//
// Errors:
// LICENSE_PROCESSING_ERROR - There was an error processing your license. Please contact your support representative.
func (host) LicenseAdd(session *Session, host HostRef, contents string) (err error) {
	method := "host.license_add"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	contentsArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "contents"), contents)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, contentsArg)
	return
}

// AsyncLicenseAdd: Apply a new license to a host
//
// Errors:
// LICENSE_PROCESSING_ERROR - There was an error processing your license. Please contact your support representative.
func (host) AsyncLicenseAdd(session *Session, host HostRef, contents string) (retval TaskRef, err error) {
	method := "Async.host.license_add"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	contentsArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "contents"), contents)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, contentsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// LicenseRemove: Remove any license file from the specified host, and switch that host to the unlicensed edition
func (host) LicenseRemove(session *Session, host HostRef) (err error) {
	method := "host.license_remove"
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

// AsyncLicenseRemove: Remove any license file from the specified host, and switch that host to the unlicensed edition
func (host) AsyncLicenseRemove(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.license_remove"
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

// Destroy: Destroy specified host record in database
func (host) Destroy(session *Session, self HostRef) (err error) {
	method := "host.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy specified host record in database
func (host) AsyncDestroy(session *Session, self HostRef) (retval TaskRef, err error) {
	method := "Async.host.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// PowerOn: Attempt to power-on the host (if the capability exists).
func (host) PowerOn(session *Session, host HostRef) (err error) {
	method := "host.power_on"
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

// AsyncPowerOn: Attempt to power-on the host (if the capability exists).
func (host) AsyncPowerOn(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.power_on"
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

// EmergencyHaDisable: This call disables HA on the local host. This should only be used with extreme care.
func (host) EmergencyHaDisable(session *Session, soft bool) (err error) {
	method := "host.emergency_ha_disable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	softArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "soft"), soft)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, softArg)
	return
}

// GetDataSources:
func (host) GetDataSources(session *Session, host HostRef) (retval []DataSourceRecord, err error) {
	method := "host.get_data_sources"
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
	retval, err = deserializeDataSourceRecordSet(method+" -> ", result)
	return
}

// RecordDataSource: Start recording the specified data source
func (host) RecordDataSource(session *Session, host HostRef, dataSource string) (err error) {
	method := "host.record_data_source"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	dataSourceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "data_source"), dataSource)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, dataSourceArg)
	return
}

// QueryDataSource: Query the latest value of the specified data source
func (host) QueryDataSource(session *Session, host HostRef, dataSource string) (retval float64, err error) {
	method := "host.query_data_source"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	dataSourceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "data_source"), dataSource)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, dataSourceArg)
	if err != nil {
		return
	}
	retval, err = deserializeFloat(method+" -> ", result)
	return
}

// ForgetDataSourceArchives: Forget the recorded statistics related to the specified data source
func (host) ForgetDataSourceArchives(session *Session, host HostRef, dataSource string) (err error) {
	method := "host.forget_data_source_archives"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	dataSourceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "data_source"), dataSource)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, dataSourceArg)
	return
}

// AssertCanEvacuate: Check this host can be evacuated.
func (host) AssertCanEvacuate(session *Session, host HostRef) (err error) {
	method := "host.assert_can_evacuate"
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

// AsyncAssertCanEvacuate: Check this host can be evacuated.
func (host) AsyncAssertCanEvacuate(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.assert_can_evacuate"
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

// GetVmsWhichPreventEvacuation: Return a set of VMs which prevent the host being evacuated, with per-VM error codes
func (host) GetVmsWhichPreventEvacuation(session *Session, self HostRef) (retval map[VMRef][]string, err error) {
	method := "host.get_vms_which_prevent_evacuation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefToStringSetMap(method+" -> ", result)
	return
}

// AsyncGetVmsWhichPreventEvacuation: Return a set of VMs which prevent the host being evacuated, with per-VM error codes
func (host) AsyncGetVmsWhichPreventEvacuation(session *Session, self HostRef) (retval TaskRef, err error) {
	method := "Async.host.get_vms_which_prevent_evacuation"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUncooperativeResidentVMs: Return a set of VMs which are not co-operating with the host&apos;s memory control system
func (host) GetUncooperativeResidentVMs(session *Session, self HostRef) (retval []VMRef, err error) {
	method := "host.get_uncooperative_resident_VMs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefSet(method+" -> ", result)
	return
}

// AsyncGetUncooperativeResidentVMs: Return a set of VMs which are not co-operating with the host&apos;s memory control system
func (host) AsyncGetUncooperativeResidentVMs(session *Session, self HostRef) (retval TaskRef, err error) {
	method := "Async.host.get_uncooperative_resident_VMs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Evacuate: Migrate all VMs off of this host, where possible.
func (host) Evacuate(session *Session, host HostRef, network NetworkRef, evacuateBatchSize int) (err error) {
	method := "host.evacuate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	evacuateBatchSizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "evacuate_batch_size"), evacuateBatchSize)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, networkArg, evacuateBatchSizeArg)
	return
}

// AsyncEvacuate: Migrate all VMs off of this host, where possible.
func (host) AsyncEvacuate(session *Session, host HostRef, network NetworkRef, evacuateBatchSize int) (retval TaskRef, err error) {
	method := "Async.host.evacuate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	evacuateBatchSizeArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "evacuate_batch_size"), evacuateBatchSize)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, networkArg, evacuateBatchSizeArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SyslogReconfigure: Re-configure syslog logging
func (host) SyslogReconfigure(session *Session, host HostRef) (err error) {
	method := "host.syslog_reconfigure"
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

// AsyncSyslogReconfigure: Re-configure syslog logging
func (host) AsyncSyslogReconfigure(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.syslog_reconfigure"
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

// ManagementReconfigure: Reconfigure the management network interface
func (host) ManagementReconfigure(session *Session, pif PIFRef) (err error) {
	method := "host.management_reconfigure"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	pifArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "pif"), pif)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, pifArg)
	return
}

// AsyncManagementReconfigure: Reconfigure the management network interface
func (host) AsyncManagementReconfigure(session *Session, pif PIFRef) (retval TaskRef, err error) {
	method := "Async.host.management_reconfigure"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	pifArg, err := serializePIFRef(fmt.Sprintf("%s(%s)", method, "pif"), pif)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, pifArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// LocalManagementReconfigure: Reconfigure the management network interface. Should only be used if Host.management_reconfigure is impossible because the network configuration is broken.
func (host) LocalManagementReconfigure(session *Session, inter string) (err error) {
	method := "host.local_management_reconfigure"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	interArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "interface"), inter)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, interArg)
	return
}

// ManagementDisable: Disable the management network interface
func (host) ManagementDisable(session *Session) (err error) {
	method := "host.management_disable"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// GetManagementInterface: Returns the management interface for the specified host
func (host) GetManagementInterface(session *Session, host HostRef) (retval PIFRef, err error) {
	method := "host.get_management_interface"
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
	retval, err = deserializePIFRef(method+" -> ", result)
	return
}

// AsyncGetManagementInterface: Returns the management interface for the specified host
func (host) AsyncGetManagementInterface(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.get_management_interface"
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

// GetSystemStatusCapabilities:
func (host) GetSystemStatusCapabilities(session *Session, host HostRef) (retval string, err error) {
	method := "host.get_system_status_capabilities"
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
	retval, err = deserializeString(method+" -> ", result)
	return
}

// RestartAgent: Restarts the agent after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called.
func (host) RestartAgent(session *Session, host HostRef) (err error) {
	method := "host.restart_agent"
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

// AsyncRestartAgent: Restarts the agent after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called.
func (host) AsyncRestartAgent(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.restart_agent"
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

// ShutdownAgent: Shuts the agent down after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called.
func (host) ShutdownAgent(session *Session) (err error) {
	method := "host.shutdown_agent"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// SetHostnameLive: Sets the host name to the specified string.  Both the API and lower-level system hostname are changed immediately.
//
// Errors:
// HOST_NAME_INVALID - The server name is invalid.
func (host) SetHostnameLive(session *Session, host HostRef, hostname string) (err error) {
	method := "host.set_hostname_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	hostnameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "hostname"), hostname)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, hostnameArg)
	return
}

// ComputeFreeMemory: Computes the amount of free memory on the host.
func (host) ComputeFreeMemory(session *Session, host HostRef) (retval int, err error) {
	method := "host.compute_free_memory"
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
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// AsyncComputeFreeMemory: Computes the amount of free memory on the host.
func (host) AsyncComputeFreeMemory(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.compute_free_memory"
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

// ComputeMemoryOverhead: Computes the virtualization memory overhead of a host.
func (host) ComputeMemoryOverhead(session *Session, host HostRef) (retval int, err error) {
	method := "host.compute_memory_overhead"
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
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// AsyncComputeMemoryOverhead: Computes the virtualization memory overhead of a host.
func (host) AsyncComputeMemoryOverhead(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.compute_memory_overhead"
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

// SyncData: This causes the synchronisation of the non-database data (messages, RRDs and so on) stored on the master to be synchronised with the host
func (host) SyncData(session *Session, host HostRef) (err error) {
	method := "host.sync_data"
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

// BackupRrds: This causes the RRDs to be backed up to the master
func (host) BackupRrds(session *Session, host HostRef, delay float64) (err error) {
	method := "host.backup_rrds"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	delayArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "delay"), delay)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, delayArg)
	return
}

// CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this host
func (host) CreateNewBlob(session *Session, host HostRef, name string, mimeType string, public bool) (retval BlobRef, err error) {
	method := "host.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
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
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeBlobRef(method+" -> ", result)
	return
}

// AsyncCreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this host
func (host) AsyncCreateNewBlob(session *Session, host HostRef, name string, mimeType string, public bool) (retval TaskRef, err error) {
	method := "Async.host.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
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
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CallPlugin: Call an API plugin on this host
func (host) CallPlugin(session *Session, host HostRef, plugin string, fn string, args map[string]string) (retval string, err error) {
	method := "host.call_plugin"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	pluginArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "plugin"), plugin)
	if err != nil {
		return
	}
	fnArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "fn"), fn)
	if err != nil {
		return
	}
	argsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, pluginArg, fnArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncCallPlugin: Call an API plugin on this host
func (host) AsyncCallPlugin(session *Session, host HostRef, plugin string, fn string, args map[string]string) (retval TaskRef, err error) {
	method := "Async.host.call_plugin"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	pluginArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "plugin"), plugin)
	if err != nil {
		return
	}
	fnArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "fn"), fn)
	if err != nil {
		return
	}
	argsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, pluginArg, fnArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// HasExtension: Return true if the extension is available on the host
func (host) HasExtension(session *Session, host HostRef, name string) (retval bool, err error) {
	method := "host.has_extension"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, nameArg)
	if err != nil {
		return
	}
	retval, err = deserializeBool(method+" -> ", result)
	return
}

// AsyncHasExtension: Return true if the extension is available on the host
func (host) AsyncHasExtension(session *Session, host HostRef, name string) (retval TaskRef, err error) {
	method := "Async.host.has_extension"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, nameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CallExtension: Call an API extension on this host
func (host) CallExtension(session *Session, host HostRef, call string) (retval string, err error) {
	method := "host.call_extension"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	callArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "call"), call)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, callArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// GetServertime: This call queries the host&apos;s clock for the current time
func (host) GetServertime(session *Session, host HostRef) (retval time.Time, err error) {
	method := "host.get_servertime"
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
	retval, err = deserializeTime(method+" -> ", result)
	return
}

// GetServerLocaltime: This call queries the host&apos;s clock for the current time in the host&apos;s local timezone
func (host) GetServerLocaltime(session *Session, host HostRef) (retval time.Time, err error) {
	method := "host.get_server_localtime"
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
	retval, err = deserializeTime(method+" -> ", result)
	return
}

// EnableExternalAuth: This call enables external authentication on a host
func (host) EnableExternalAuth(session *Session, host HostRef, config map[string]string, serviceName string, authType string) (err error) {
	method := "host.enable_external_auth"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
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
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, configArg, serviceNameArg, authTypeArg)
	return
}

// DisableExternalAuth: This call disables external authentication on the local host
func (host) DisableExternalAuth(session *Session, host HostRef, config map[string]string) (err error) {
	method := "host.disable_external_auth"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	configArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "config"), config)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, configArg)
	return
}

// RetrieveWlbEvacuateRecommendations: Retrieves recommended host migrations to perform when evacuating the host from the wlb server. If a VM cannot be migrated from the host the reason is listed instead of a recommendation.
func (host) RetrieveWlbEvacuateRecommendations(session *Session, self HostRef) (retval map[VMRef][]string, err error) {
	method := "host.retrieve_wlb_evacuate_recommendations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefToStringSetMap(method+" -> ", result)
	return
}

// AsyncRetrieveWlbEvacuateRecommendations: Retrieves recommended host migrations to perform when evacuating the host from the wlb server. If a VM cannot be migrated from the host the reason is listed instead of a recommendation.
func (host) AsyncRetrieveWlbEvacuateRecommendations(session *Session, self HostRef) (retval TaskRef, err error) {
	method := "Async.host.retrieve_wlb_evacuate_recommendations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetServerCertificate: Get the installed server public TLS certificate.
func (host) GetServerCertificate(session *Session, host HostRef) (retval string, err error) {
	method := "host.get_server_certificate"
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
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncGetServerCertificate: Get the installed server public TLS certificate.
func (host) AsyncGetServerCertificate(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.get_server_certificate"
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

// RefreshServerCertificate: Replace the internal self-signed host certficate with a new one.
func (host) RefreshServerCertificate(session *Session, host HostRef) (err error) {
	method := "host.refresh_server_certificate"
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

// AsyncRefreshServerCertificate: Replace the internal self-signed host certficate with a new one.
func (host) AsyncRefreshServerCertificate(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.refresh_server_certificate"
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

// InstallServerCertificate: Install the TLS server certificate.
func (host) InstallServerCertificate(session *Session, host HostRef, certificate string, privateKey string, certificateChain string) (err error) {
	method := "host.install_server_certificate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	certificateArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "certificate"), certificate)
	if err != nil {
		return
	}
	privateKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "private_key"), privateKey)
	if err != nil {
		return
	}
	certificateChainArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "certificate_chain"), certificateChain)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, certificateArg, privateKeyArg, certificateChainArg)
	return
}

// AsyncInstallServerCertificate: Install the TLS server certificate.
func (host) AsyncInstallServerCertificate(session *Session, host HostRef, certificate string, privateKey string, certificateChain string) (retval TaskRef, err error) {
	method := "Async.host.install_server_certificate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	certificateArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "certificate"), certificate)
	if err != nil {
		return
	}
	privateKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "private_key"), privateKey)
	if err != nil {
		return
	}
	certificateChainArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "certificate_chain"), certificateChain)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, certificateArg, privateKeyArg, certificateChainArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// EmergencyResetServerCertificate: Delete the current TLS server certificate and replace by a new, self-signed one. This should only be used with extreme care.
func (host) EmergencyResetServerCertificate(session *Session) (err error) {
	method := "host.emergency_reset_server_certificate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// ResetServerCertificate: Delete the current TLS server certificate and replace by a new, self-signed one. This should only be used with extreme care.
func (host) ResetServerCertificate(session *Session, host HostRef) (err error) {
	method := "host.reset_server_certificate"
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

// AsyncResetServerCertificate: Delete the current TLS server certificate and replace by a new, self-signed one. This should only be used with extreme care.
func (host) AsyncResetServerCertificate(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.reset_server_certificate"
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

// ApplyEdition: Change to another edition, or reactivate the current edition after a license has expired. This may be subject to the successful checkout of an appropriate license.
func (host) ApplyEdition(session *Session, host HostRef, edition string, force bool) (err error) {
	method := "host.apply_edition"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	editionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "edition"), edition)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, editionArg, forceArg)
	return
}

// RefreshPackInfo: Refresh the list of installed Supplemental Packs.
func (host) RefreshPackInfo(session *Session, host HostRef) (err error) {
	method := "host.refresh_pack_info"
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

// AsyncRefreshPackInfo: Refresh the list of installed Supplemental Packs.
func (host) AsyncRefreshPackInfo(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.refresh_pack_info"
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

// SetPowerOnMode: Set the power-on-mode, host, user and password
func (host) SetPowerOnMode(session *Session, self HostRef, powerOnMode string, powerOnConfig map[string]string) (err error) {
	method := "host.set_power_on_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	powerOnModeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "power_on_mode"), powerOnMode)
	if err != nil {
		return
	}
	powerOnConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "power_on_config"), powerOnConfig)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, powerOnModeArg, powerOnConfigArg)
	return
}

// AsyncSetPowerOnMode: Set the power-on-mode, host, user and password
func (host) AsyncSetPowerOnMode(session *Session, self HostRef, powerOnMode string, powerOnConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.host.set_power_on_mode"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	powerOnModeArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "power_on_mode"), powerOnMode)
	if err != nil {
		return
	}
	powerOnConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "power_on_config"), powerOnConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, powerOnModeArg, powerOnConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetCPUFeatures: Set the CPU features to be used after a reboot, if the given features string is valid.
func (host) SetCPUFeatures(session *Session, host HostRef, features string) (err error) {
	method := "host.set_cpu_features"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	featuresArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "features"), features)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, featuresArg)
	return
}

// ResetCPUFeatures: Remove the feature mask, such that after a reboot all features of the CPU are enabled.
func (host) ResetCPUFeatures(session *Session, host HostRef) (err error) {
	method := "host.reset_cpu_features"
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

// EnableLocalStorageCaching: Enable the use of a local SR for caching purposes
func (host) EnableLocalStorageCaching(session *Session, host HostRef, sr SRRef) (err error) {
	method := "host.enable_local_storage_caching"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, srArg)
	return
}

// DisableLocalStorageCaching: Disable the use of a local SR for caching purposes
func (host) DisableLocalStorageCaching(session *Session, host HostRef) (err error) {
	method := "host.disable_local_storage_caching"
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

// MigrateReceive: Prepare to receive a VM, returning a token which can be passed to VM.migrate.
func (host) MigrateReceive(session *Session, host HostRef, network NetworkRef, options map[string]string) (retval map[string]string, err error) {
	method := "host.migrate_receive"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, networkArg, optionsArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToStringMap(method+" -> ", result)
	return
}

// AsyncMigrateReceive: Prepare to receive a VM, returning a token which can be passed to VM.migrate.
func (host) AsyncMigrateReceive(session *Session, host HostRef, network NetworkRef, options map[string]string) (retval TaskRef, err error) {
	method := "Async.host.migrate_receive"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	networkArg, err := serializeNetworkRef(fmt.Sprintf("%s(%s)", method, "network"), network)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, networkArg, optionsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// DeclareDead: Declare that a host is dead. This is a dangerous operation, and should only be called if the administrator is absolutely sure the host is definitely dead
func (host) DeclareDead(session *Session, host HostRef) (err error) {
	method := "host.declare_dead"
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

// AsyncDeclareDead: Declare that a host is dead. This is a dangerous operation, and should only be called if the administrator is absolutely sure the host is definitely dead
func (host) AsyncDeclareDead(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.declare_dead"
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

// EnableDisplay: Enable console output to the physical display device next time this host boots
func (host) EnableDisplay(session *Session, host HostRef) (retval HostDisplay, err error) {
	method := "host.enable_display"
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
	retval, err = deserializeEnumHostDisplay(method+" -> ", result)
	return
}

// AsyncEnableDisplay: Enable console output to the physical display device next time this host boots
func (host) AsyncEnableDisplay(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.enable_display"
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

// DisableDisplay: Disable console output to the physical display device next time this host boots
func (host) DisableDisplay(session *Session, host HostRef) (retval HostDisplay, err error) {
	method := "host.disable_display"
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
	retval, err = deserializeEnumHostDisplay(method+" -> ", result)
	return
}

// AsyncDisableDisplay: Disable console output to the physical display device next time this host boots
func (host) AsyncDisableDisplay(session *Session, host HostRef) (retval TaskRef, err error) {
	method := "Async.host.disable_display"
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

// SetSslLegacy: Enable/disable SSLv3 for interoperability with older server versions. When this is set to a different value, the host immediately restarts its SSL/TLS listening service; typically this takes less than a second but existing connections to it will be broken. API login sessions will remain valid.
func (host) SetSslLegacy(session *Session, self HostRef, value bool) (err error) {
	method := "host.set_ssl_legacy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetSslLegacy: Enable/disable SSLv3 for interoperability with older server versions. When this is set to a different value, the host immediately restarts its SSL/TLS listening service; typically this takes less than a second but existing connections to it will be broken. API login sessions will remain valid.
func (host) AsyncSetSslLegacy(session *Session, self HostRef, value bool) (retval TaskRef, err error) {
	method := "Async.host.set_ssl_legacy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetIscsiIqn: Sets the initiator IQN for the host
func (host) SetIscsiIqn(session *Session, host HostRef, value string) (err error) {
	method := "host.set_iscsi_iqn"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, valueArg)
	return
}

// AsyncSetIscsiIqn: Sets the initiator IQN for the host
func (host) AsyncSetIscsiIqn(session *Session, host HostRef, value string) (retval TaskRef, err error) {
	method := "Async.host.set_iscsi_iqn"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetMultipathing: Specifies whether multipathing is enabled
func (host) SetMultipathing(session *Session, host HostRef, value bool) (err error) {
	method := "host.set_multipathing"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, valueArg)
	return
}

// AsyncSetMultipathing: Specifies whether multipathing is enabled
func (host) AsyncSetMultipathing(session *Session, host HostRef, value bool) (retval TaskRef, err error) {
	method := "Async.host.set_multipathing"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	valueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetUefiCertificates: Sets the UEFI certificates on a host
func (host) SetUefiCertificates(session *Session, host HostRef, value string) (err error) {
	method := "host.set_uefi_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, hostArg, valueArg)
	return
}

// AsyncSetUefiCertificates: Sets the UEFI certificates on a host
func (host) AsyncSetUefiCertificates(session *Session, host HostRef, value string) (retval TaskRef, err error) {
	method := "Async.host.set_uefi_certificates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, hostArg, valueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetSchedGran: Sets xen&apos;s sched-gran on a host. See: https://xenbits.xen.org/docs/unstable/misc/xen-command-line.html#sched-gran-x86
func (host) SetSchedGran(session *Session, self HostRef, value HostSchedGran) (err error) {
	method := "host.set_sched_gran"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumHostSchedGran(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetSchedGran: Sets xen&apos;s sched-gran on a host. See: https://xenbits.xen.org/docs/unstable/misc/xen-command-line.html#sched-gran-x86
func (host) AsyncSetSchedGran(session *Session, self HostRef, value HostSchedGran) (retval TaskRef, err error) {
	method := "Async.host.set_sched_gran"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumHostSchedGran(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// GetSchedGran: Gets xen&apos;s sched-gran on a host
func (host) GetSchedGran(session *Session, self HostRef) (retval HostSchedGran, err error) {
	method := "host.get_sched_gran"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumHostSchedGran(method+" -> ", result)
	return
}

// AsyncGetSchedGran: Gets xen&apos;s sched-gran on a host
func (host) AsyncGetSchedGran(session *Session, self HostRef) (retval TaskRef, err error) {
	method := "Async.host.get_sched_gran"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNumaAffinityPolicy: Set VM placement NUMA affinity policy
func (host) SetNumaAffinityPolicy(session *Session, self HostRef, value HostNumaAffinityPolicy) (err error) {
	method := "host.set_numa_affinity_policy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumHostNumaAffinityPolicy(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetNumaAffinityPolicy: Set VM placement NUMA affinity policy
func (host) AsyncSetNumaAffinityPolicy(session *Session, self HostRef, value HostNumaAffinityPolicy) (retval TaskRef, err error) {
	method := "Async.host.set_numa_affinity_policy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumHostNumaAffinityPolicy(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// EmergencyDisableTLSVerification: Disable TLS verification for this host only
func (host) EmergencyDisableTLSVerification(session *Session) (err error) {
	method := "host.emergency_disable_tls_verification"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// EmergencyReenableTLSVerification: Reenable TLS verification for this host only
func (host) EmergencyReenableTLSVerification(session *Session) (err error) {
	method := "host.emergency_reenable_tls_verification"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// ApplyUpdates: apply updates from current enabled repository on a host
func (host) ApplyUpdates(session *Session, self HostRef, hash string) (retval [][]string, err error) {
	method := "host.apply_updates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hashArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "hash"), hash)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, hashArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSetSet(method+" -> ", result)
	return
}

// AsyncApplyUpdates: apply updates from current enabled repository on a host
func (host) AsyncApplyUpdates(session *Session, self HostRef, hash string) (retval TaskRef, err error) {
	method := "Async.host.apply_updates"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hashArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "hash"), hash)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, hashArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetHTTPSOnly: updates the host firewall to open or close port 80 depending on the value
func (host) SetHTTPSOnly(session *Session, self HostRef, value bool) (err error) {
	method := "host.set_https_only"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetHTTPSOnly: updates the host firewall to open or close port 80 depending on the value
func (host) AsyncSetHTTPSOnly(session *Session, self HostRef, value bool) (retval TaskRef, err error) {
	method := "Async.host.set_https_only"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// ApplyRecommendedGuidances: apply all recommended guidances both on the host and on all HVM VMs on the host after updates are applied on the host
func (host) ApplyRecommendedGuidances(session *Session, self HostRef) (err error) {
	method := "host.apply_recommended_guidances"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncApplyRecommendedGuidances: apply all recommended guidances both on the host and on all HVM VMs on the host after updates are applied on the host
func (host) AsyncApplyRecommendedGuidances(session *Session, self HostRef) (retval TaskRef, err error) {
	method := "Async.host.apply_recommended_guidances"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// EmergencyClearMandatoryGuidance: Clear the pending mandatory guidance on this host
func (host) EmergencyClearMandatoryGuidance(session *Session) (err error) {
	method := "host.emergency_clear_mandatory_guidance"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg)
	return
}

// GetAll: Return a list of all the hosts known to the system.
func (host) GetAll(session *Session) (retval []HostRef, err error) {
	method := "host.get_all"
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

// GetAllRecords: Return a map of host references to host records for all hosts known to the system.
func (host) GetAllRecords(session *Session) (retval map[HostRef]HostRecord, err error) {
	method := "host.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRefToHostRecordMap(method+" -> ", result)
	return
}
