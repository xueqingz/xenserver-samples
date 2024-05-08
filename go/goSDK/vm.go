package xenapi

import (
	"fmt"
	"time"
)

type VMRecord struct {
	// Unique identifier/object reference
	UUID string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VMOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VMOperations
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// Current power state of the machine
	PowerState VMPowerState
	// Creators of VMs and templates may store version information here.
	UserVersion int
	// true if this is a template. Template VMs can never be started, they are used only for cloning other VMs
	IsATemplate bool
	// true if this is a default template. Default template VMs can never be started or migrated, they are used only for cloning other VMs
	IsDefaultTemplate bool
	// The VDI that a suspend image is stored on. (Only has meaning if VM is currently suspended)
	SuspendVDI VDIRef
	// the host the VM is currently resident on
	ResidentOn HostRef
	// the host on which the VM is due to be started/resumed/migrated. This acts as a memory reservation indicator
	ScheduledToBeResidentOn HostRef
	// A host which the VM has some affinity for (or NULL). This is used as a hint to the start call when it decides where to run the VM. Resource constraints may cause the VM to be started elsewhere.
	Affinity HostRef
	// Virtualization memory overhead (bytes).
	MemoryOverhead int
	// Dynamically-set memory target (bytes). The value of this field indicates the current target for memory available to this VM.
	MemoryTarget int
	// Statically-set (i.e. absolute) maximum (bytes). The value of this field at VM start time acts as a hard limit of the amount of memory a guest can use. New values only take effect on reboot.
	MemoryStaticMax int
	// Dynamic maximum (bytes)
	MemoryDynamicMax int
	// Dynamic minimum (bytes)
	MemoryDynamicMin int
	// Statically-set (i.e. absolute) mininum (bytes). The value of this field indicates the least amount of memory this VM can boot with without crashing.
	MemoryStaticMin int
	// configuration parameters for the selected VCPU policy
	VCPUsParams map[string]string
	// Max number of VCPUs
	VCPUsMax int
	// Boot number of VCPUs
	VCPUsAtStartup int
	// action to take after soft reboot
	ActionsAfterSoftreboot OnSoftrebootBehavior
	// action to take after the guest has shutdown itself
	ActionsAfterShutdown OnNormalExit
	// action to take after the guest has rebooted itself
	ActionsAfterReboot OnNormalExit
	// action to take if the guest crashes
	ActionsAfterCrash OnCrashBehaviour
	// virtual console devices
	Consoles []ConsoleRef
	// virtual network interfaces
	VIFs []VIFRef
	// virtual block devices
	VBDs []VBDRef
	// vitual usb devices
	VUSBs []VUSBRef
	// crash dumps associated with this VM
	CrashDumps []CrashdumpRef
	// virtual TPMs
	VTPMs []VTPMRef
	// name of or path to bootloader
	PVBootloader string
	// path to the kernel
	PVKernel string
	// path to the initrd
	PVRamdisk string
	// kernel command-line arguments
	PVArgs string
	// miscellaneous arguments for the bootloader
	PVBootloaderArgs string
	// to make Zurich guests boot
	PVLegacyArgs string
	// HVM boot policy
	HVMBootPolicy string
	// HVM boot params
	HVMBootParams map[string]string
	// multiplier applied to the amount of shadow that will be made available to the guest
	HVMShadowMultiplier float64
	// platform-specific configuration
	Platform map[string]string
	// PCI bus path for pass-through devices
	PCIBus string
	// additional configuration
	OtherConfig map[string]string
	// domain ID (if available, -1 otherwise)
	Domid int
	// Domain architecture (if available, null string otherwise)
	Domarch string
	// describes the CPU flags on which the VM was last booted
	LastBootCPUFlags map[string]string
	// true if this is a control domain (domain 0 or a driver domain)
	IsControlDomain bool
	// metrics associated with this VM
	Metrics VMMetricsRef
	// metrics associated with the running guest
	GuestMetrics VMGuestMetricsRef
	// marshalled value containing VM record at time of last boot
	LastBootedRecord string
	// An XML specification of recommended values and ranges for properties of this VM
	Recommendations string
	// data to be inserted into the xenstore tree (/local/domain/&lt;domid&gt;/vm-data) after the VM is created.
	XenstoreData map[string]string
	// if true then the system will attempt to keep the VM running as much as possible.
	HaAlwaysRun bool
	// has possible values: &quot;best-effort&quot; meaning &quot;try to restart this VM if possible but don&apos;t consider the Pool to be overcommitted if this is not possible&quot;; &quot;restart&quot; meaning &quot;this VM should be restarted&quot;; &quot;&quot; meaning &quot;do not try to restart this VM&quot;
	HaRestartPriority string
	// true if this is a snapshot. Snapshotted VMs can never be started, they are used only for cloning other VMs
	IsASnapshot bool
	// Ref pointing to the VM this snapshot is of.
	SnapshotOf VMRef
	// List pointing to all the VM snapshots.
	Snapshots []VMRef
	// Date/time when this snapshot was created.
	SnapshotTime time.Time
	// Transportable ID of the snapshot VM
	TransportableSnapshotID string
	// Binary blobs associated with this VM
	Blobs map[string]BlobRef
	// user-specified tags for categorization purposes
	Tags []string
	// List of operations which have been explicitly blocked and an error code
	BlockedOperations map[VMOperations]string
	// Human-readable information concerning this snapshot
	SnapshotInfo map[string]string
	// Encoded information about the VM&apos;s metadata this is a snapshot of
	SnapshotMetadata string
	// Ref pointing to the parent of this VM
	Parent VMRef
	// List pointing to all the children of this VM
	Children []VMRef
	// BIOS strings
	BiosStrings map[string]string
	// Ref pointing to a protection policy for this VM
	ProtectionPolicy VMPPRef
	// true if this snapshot was created by the protection policy
	IsSnapshotFromVmpp bool
	// Ref pointing to a snapshot schedule for this VM
	SnapshotSchedule VMSSRef
	// true if this snapshot was created by the snapshot schedule
	IsVmssSnapshot bool
	// the appliance to which this VM belongs
	Appliance VMApplianceRef
	// The delay to wait before proceeding to the next order in the startup sequence (seconds)
	StartDelay int
	// The delay to wait before proceeding to the next order in the shutdown sequence (seconds)
	ShutdownDelay int
	// The point in the startup or shutdown sequence at which this VM will be started
	Order int
	// Virtual GPUs
	VGPUs []VGPURef
	// Currently passed-through PCI devices
	AttachedPCIs []PCIRef
	// The SR on which a suspend image is stored
	SuspendSR SRRef
	// The number of times this VM has been recovered
	Version int
	// Generation ID of the VM
	GenerationID string
	// The host virtual hardware platform version the VM can run on
	HardwarePlatformVersion int
	// When an HVM guest starts, this controls the presence of the emulated C000 PCI device which triggers Windows Update to fetch or update PV drivers.
	HasVendorDevice bool
	// Indicates whether a VM requires a reboot in order to update its configuration, e.g. its memory allocation.
	RequiresReboot bool
	// Textual reference to the template used to create a VM. This can be used by clients in need of an immutable reference to the template since the latter&apos;s uuid and name_label may change, for example, after a package installation or upgrade.
	ReferenceLabel string
	// The type of domain that will be created when the VM is started
	DomainType DomainType
	// initial value for guest NVRAM (containing UEFI variables, etc). Cannot be changed while the VM is running
	NVRAM map[string]string
	// The set of pending mandatory guidances after applying updates, which must be applied, as otherwise there may be e.g. VM failures
	PendingGuidances []UpdateGuidances
	// The set of pending recommended guidances after applying updates, which most users should follow to make the updates effective, but if not followed, will not cause a failure
	PendingGuidancesRecommended []UpdateGuidances
	// The set of pending full guidances after applying updates, which a user should follow to make some updates, e.g. specific hardware drivers or CPU features, fully effective, but the &apos;average user&apos; doesn&apos;t need to
	PendingGuidancesFull []UpdateGuidances
}

type VMRef string

// A virtual machine (or &apos;guest&apos;).
type vM struct{}

var VM vM

// GetRecord: Get a record containing the current state of the given VM.
func (vM) GetRecord(session *Session, self VMRef) (retval VMRecord, err error) {
	method := "VM.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VM instance with the specified UUID.
func (vM) GetByUUID(session *Session, uUID string) (retval VMRef, err error) {
	method := "VM.get_by_uuid"
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
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// Create: NOT RECOMMENDED! VM.clone or VM.copy (or VM.import) is a better choice in almost all situations. The standard way to obtain a new VM is to call VM.clone on a template VM, then call VM.provision on the new clone. Caution: if VM.create is used and then the new VM is attached to a virtual disc that has an operating system already installed, then there is no guarantee that the operating system will boot and run. Any software that calls VM.create on a future version of this API may fail or give unexpected results. For example this could happen if an additional parameter were added to VM.create. VM.create is intended only for use in the automatic creation of the system VM templates. It creates a new VM instance, and returns its handle. The constructor args are: name_label, name_description, power_state, user_version*, is_a_template*, suspend_VDI, affinity*, memory_target, memory_static_max*, memory_dynamic_max*, memory_dynamic_min*, memory_static_min*, VCPUs_params*, VCPUs_max*, VCPUs_at_startup*, actions_after_softreboot, actions_after_shutdown*, actions_after_reboot*, actions_after_crash*, PV_bootloader*, PV_kernel*, PV_ramdisk*, PV_args*, PV_bootloader_args*, PV_legacy_args*, HVM_boot_policy*, HVM_boot_params*, HVM_shadow_multiplier, platform*, PCI_bus*, other_config*, last_boot_CPU_flags, last_booted_record, recommendations*, xenstore_data, ha_always_run, ha_restart_priority, tags, blocked_operations, protection_policy, is_snapshot_from_vmpp, snapshot_schedule, is_vmss_snapshot, appliance, start_delay, shutdown_delay, order, suspend_SR, version, generation_id, hardware_platform_version, has_vendor_device, reference_label, domain_type, NVRAM (* = non-optional).
func (vM) Create(session *Session, args VMRecord) (retval VMRef, err error) {
	method := "VM.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVMRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// AsyncCreate: NOT RECOMMENDED! VM.clone or VM.copy (or VM.import) is a better choice in almost all situations. The standard way to obtain a new VM is to call VM.clone on a template VM, then call VM.provision on the new clone. Caution: if VM.create is used and then the new VM is attached to a virtual disc that has an operating system already installed, then there is no guarantee that the operating system will boot and run. Any software that calls VM.create on a future version of this API may fail or give unexpected results. For example this could happen if an additional parameter were added to VM.create. VM.create is intended only for use in the automatic creation of the system VM templates. It creates a new VM instance, and returns its handle. The constructor args are: name_label, name_description, power_state, user_version*, is_a_template*, suspend_VDI, affinity*, memory_target, memory_static_max*, memory_dynamic_max*, memory_dynamic_min*, memory_static_min*, VCPUs_params*, VCPUs_max*, VCPUs_at_startup*, actions_after_softreboot, actions_after_shutdown*, actions_after_reboot*, actions_after_crash*, PV_bootloader*, PV_kernel*, PV_ramdisk*, PV_args*, PV_bootloader_args*, PV_legacy_args*, HVM_boot_policy*, HVM_boot_params*, HVM_shadow_multiplier, platform*, PCI_bus*, other_config*, last_boot_CPU_flags, last_booted_record, recommendations*, xenstore_data, ha_always_run, ha_restart_priority, tags, blocked_operations, protection_policy, is_snapshot_from_vmpp, snapshot_schedule, is_vmss_snapshot, appliance, start_delay, shutdown_delay, order, suspend_SR, version, generation_id, hardware_platform_version, has_vendor_device, reference_label, domain_type, NVRAM (* = non-optional).
func (vM) AsyncCreate(session *Session, args VMRecord) (retval TaskRef, err error) {
	method := "Async.VM.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeVMRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// Destroy: Destroy the specified VM.  The VM is completely removed from the system.  This function can only be called when the VM is in the Halted State.
func (vM) Destroy(session *Session, self VMRef) (err error) {
	method := "VM.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified VM.  The VM is completely removed from the system.  This function can only be called when the VM is in the Halted State.
func (vM) AsyncDestroy(session *Session, self VMRef) (retval TaskRef, err error) {
	method := "Async.VM.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByNameLabel: Get all the VM instances with the given label.
func (vM) GetByNameLabel(session *Session, label string) (retval []VMRef, err error) {
	method := "VM.get_by_name_label"
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
	retval, err = deserializeVMRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VM.
func (vM) GetUUID(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given VM.
func (vM) GetAllowedOperations(session *Session, self VMRef) (retval []VMOperations, err error) {
	method := "VM.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVMOperationsSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given VM.
func (vM) GetCurrentOperations(session *Session, self VMRef) (retval map[string]VMOperations, err error) {
	method := "VM.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumVMOperationsMap(method+" -> ", result)
	return
}

// GetNameLabel: Get the name/label field of the given VM.
func (vM) GetNameLabel(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given VM.
func (vM) GetNameDescription(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPowerState: Get the power_state field of the given VM.
func (vM) GetPowerState(session *Session, self VMRef) (retval VMPowerState, err error) {
	method := "VM.get_power_state"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVMPowerState(method+" -> ", result)
	return
}

// GetUserVersion: Get the user_version field of the given VM.
func (vM) GetUserVersion(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_user_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsATemplate: Get the is_a_template field of the given VM.
func (vM) GetIsATemplate(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_is_a_template"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsDefaultTemplate: Get the is_default_template field of the given VM.
func (vM) GetIsDefaultTemplate(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_is_default_template"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSuspendVDI: Get the suspend_VDI field of the given VM.
func (vM) GetSuspendVDI(session *Session, self VMRef) (retval VDIRef, err error) {
	method := "VM.get_suspend_VDI"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetResidentOn: Get the resident_on field of the given VM.
func (vM) GetResidentOn(session *Session, self VMRef) (retval HostRef, err error) {
	method := "VM.get_resident_on"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetScheduledToBeResidentOn: Get the scheduled_to_be_resident_on field of the given VM.
func (vM) GetScheduledToBeResidentOn(session *Session, self VMRef) (retval HostRef, err error) {
	method := "VM.get_scheduled_to_be_resident_on"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAffinity: Get the affinity field of the given VM.
func (vM) GetAffinity(session *Session, self VMRef) (retval HostRef, err error) {
	method := "VM.get_affinity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryOverhead: Get the memory/overhead field of the given VM.
func (vM) GetMemoryOverhead(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_memory_overhead"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryTarget: Get the memory/target field of the given VM.
func (vM) GetMemoryTarget(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_memory_target"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryStaticMax: Get the memory/static_max field of the given VM.
func (vM) GetMemoryStaticMax(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_memory_static_max"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryDynamicMax: Get the memory/dynamic_max field of the given VM.
func (vM) GetMemoryDynamicMax(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_memory_dynamic_max"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryDynamicMin: Get the memory/dynamic_min field of the given VM.
func (vM) GetMemoryDynamicMin(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_memory_dynamic_min"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMemoryStaticMin: Get the memory/static_min field of the given VM.
func (vM) GetMemoryStaticMin(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_memory_static_min"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVCPUsParams: Get the VCPUs/params field of the given VM.
func (vM) GetVCPUsParams(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.get_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVCPUsMax: Get the VCPUs/max field of the given VM.
func (vM) GetVCPUsMax(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_VCPUs_max"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVCPUsAtStartup: Get the VCPUs/at_startup field of the given VM.
func (vM) GetVCPUsAtStartup(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_VCPUs_at_startup"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetActionsAfterSoftreboot: Get the actions/after_softreboot field of the given VM.
func (vM) GetActionsAfterSoftreboot(session *Session, self VMRef) (retval OnSoftrebootBehavior, err error) {
	method := "VM.get_actions_after_softreboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumOnSoftrebootBehavior(method+" -> ", result)
	return
}

// GetActionsAfterShutdown: Get the actions/after_shutdown field of the given VM.
func (vM) GetActionsAfterShutdown(session *Session, self VMRef) (retval OnNormalExit, err error) {
	method := "VM.get_actions_after_shutdown"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumOnNormalExit(method+" -> ", result)
	return
}

// GetActionsAfterReboot: Get the actions/after_reboot field of the given VM.
func (vM) GetActionsAfterReboot(session *Session, self VMRef) (retval OnNormalExit, err error) {
	method := "VM.get_actions_after_reboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumOnNormalExit(method+" -> ", result)
	return
}

// GetActionsAfterCrash: Get the actions/after_crash field of the given VM.
func (vM) GetActionsAfterCrash(session *Session, self VMRef) (retval OnCrashBehaviour, err error) {
	method := "VM.get_actions_after_crash"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumOnCrashBehaviour(method+" -> ", result)
	return
}

// GetConsoles: Get the consoles field of the given VM.
func (vM) GetConsoles(session *Session, self VMRef) (retval []ConsoleRef, err error) {
	method := "VM.get_consoles"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeConsoleRefSet(method+" -> ", result)
	return
}

// GetVIFs: Get the VIFs field of the given VM.
func (vM) GetVIFs(session *Session, self VMRef) (retval []VIFRef, err error) {
	method := "VM.get_VIFs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVBDs: Get the VBDs field of the given VM.
func (vM) GetVBDs(session *Session, self VMRef) (retval []VBDRef, err error) {
	method := "VM.get_VBDs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVBDRefSet(method+" -> ", result)
	return
}

// GetVUSBs: Get the VUSBs field of the given VM.
func (vM) GetVUSBs(session *Session, self VMRef) (retval []VUSBRef, err error) {
	method := "VM.get_VUSBs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVUSBRefSet(method+" -> ", result)
	return
}

// GetCrashDumps: Get the crash_dumps field of the given VM.
func (vM) GetCrashDumps(session *Session, self VMRef) (retval []CrashdumpRef, err error) {
	method := "VM.get_crash_dumps"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeCrashdumpRefSet(method+" -> ", result)
	return
}

// GetVTPMs: Get the VTPMs field of the given VM.
func (vM) GetVTPMs(session *Session, self VMRef) (retval []VTPMRef, err error) {
	method := "VM.get_VTPMs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVTPMRefSet(method+" -> ", result)
	return
}

// GetPVBootloader: Get the PV/bootloader field of the given VM.
func (vM) GetPVBootloader(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_PV_bootloader"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPVKernel: Get the PV/kernel field of the given VM.
func (vM) GetPVKernel(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_PV_kernel"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPVRamdisk: Get the PV/ramdisk field of the given VM.
func (vM) GetPVRamdisk(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_PV_ramdisk"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPVArgs: Get the PV/args field of the given VM.
func (vM) GetPVArgs(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_PV_args"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPVBootloaderArgs: Get the PV/bootloader_args field of the given VM.
func (vM) GetPVBootloaderArgs(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_PV_bootloader_args"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPVLegacyArgs: Get the PV/legacy_args field of the given VM.
func (vM) GetPVLegacyArgs(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_PV_legacy_args"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHVMBootPolicy: Get the HVM/boot_policy field of the given VM.
func (vM) GetHVMBootPolicy(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_HVM_boot_policy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHVMBootParams: Get the HVM/boot_params field of the given VM.
func (vM) GetHVMBootParams(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.get_HVM_boot_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHVMShadowMultiplier: Get the HVM/shadow_multiplier field of the given VM.
func (vM) GetHVMShadowMultiplier(session *Session, self VMRef) (retval float64, err error) {
	method := "VM.get_HVM_shadow_multiplier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPlatform: Get the platform field of the given VM.
func (vM) GetPlatform(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.get_platform"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPCIBus: Get the PCI_bus field of the given VM.
func (vM) GetPCIBus(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_PCI_bus"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig: Get the other_config field of the given VM.
func (vM) GetOtherConfig(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDomid: Get the domid field of the given VM.
func (vM) GetDomid(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_domid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDomarch: Get the domarch field of the given VM.
func (vM) GetDomarch(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_domarch"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetLastBootCPUFlags: Get the last_boot_CPU_flags field of the given VM.
func (vM) GetLastBootCPUFlags(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.get_last_boot_CPU_flags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsControlDomain: Get the is_control_domain field of the given VM.
func (vM) GetIsControlDomain(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_is_control_domain"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetMetrics: Get the metrics field of the given VM.
func (vM) GetMetrics(session *Session, self VMRef) (retval VMMetricsRef, err error) {
	method := "VM.get_metrics"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMMetricsRef(method+" -> ", result)
	return
}

// GetGuestMetrics: Get the guest_metrics field of the given VM.
func (vM) GetGuestMetrics(session *Session, self VMRef) (retval VMGuestMetricsRef, err error) {
	method := "VM.get_guest_metrics"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMGuestMetricsRef(method+" -> ", result)
	return
}

// GetLastBootedRecord: Get the last_booted_record field of the given VM.
func (vM) GetLastBootedRecord(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_last_booted_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRecommendations: Get the recommendations field of the given VM.
func (vM) GetRecommendations(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_recommendations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetXenstoreData: Get the xenstore_data field of the given VM.
func (vM) GetXenstoreData(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.get_xenstore_data"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaAlwaysRun: Get the ha_always_run field of the given VM.
func (vM) GetHaAlwaysRun(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_ha_always_run"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHaRestartPriority: Get the ha_restart_priority field of the given VM.
func (vM) GetHaRestartPriority(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_ha_restart_priority"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsASnapshot: Get the is_a_snapshot field of the given VM.
func (vM) GetIsASnapshot(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_is_a_snapshot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSnapshotOf: Get the snapshot_of field of the given VM.
func (vM) GetSnapshotOf(session *Session, self VMRef) (retval VMRef, err error) {
	method := "VM.get_snapshot_of"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSnapshots: Get the snapshots field of the given VM.
func (vM) GetSnapshots(session *Session, self VMRef) (retval []VMRef, err error) {
	method := "VM.get_snapshots"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSnapshotTime: Get the snapshot_time field of the given VM.
func (vM) GetSnapshotTime(session *Session, self VMRef) (retval time.Time, err error) {
	method := "VM.get_snapshot_time"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTransportableSnapshotID: Get the transportable_snapshot_id field of the given VM.
func (vM) GetTransportableSnapshotID(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_transportable_snapshot_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBlobs: Get the blobs field of the given VM.
func (vM) GetBlobs(session *Session, self VMRef) (retval map[string]BlobRef, err error) {
	method := "VM.get_blobs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetTags: Get the tags field of the given VM.
func (vM) GetTags(session *Session, self VMRef) (retval []string, err error) {
	method := "VM.get_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBlockedOperations: Get the blocked_operations field of the given VM.
func (vM) GetBlockedOperations(session *Session, self VMRef) (retval map[VMOperations]string, err error) {
	method := "VM.get_blocked_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVMOperationsToStringMap(method+" -> ", result)
	return
}

// GetSnapshotInfo: Get the snapshot_info field of the given VM.
func (vM) GetSnapshotInfo(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.get_snapshot_info"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSnapshotMetadata: Get the snapshot_metadata field of the given VM.
func (vM) GetSnapshotMetadata(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_snapshot_metadata"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetParent: Get the parent field of the given VM.
func (vM) GetParent(session *Session, self VMRef) (retval VMRef, err error) {
	method := "VM.get_parent"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetChildren: Get the children field of the given VM.
func (vM) GetChildren(session *Session, self VMRef) (retval []VMRef, err error) {
	method := "VM.get_children"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBiosStrings: Get the bios_strings field of the given VM.
func (vM) GetBiosStrings(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.get_bios_strings"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetProtectionPolicy: Get the protection_policy field of the given VM.
func (vM) GetProtectionPolicy(session *Session, self VMRef) (retval VMPPRef, err error) {
	method := "VM.get_protection_policy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMPPRef(method+" -> ", result)
	return
}

// GetIsSnapshotFromVmpp: Get the is_snapshot_from_vmpp field of the given VM.
func (vM) GetIsSnapshotFromVmpp(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_is_snapshot_from_vmpp"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSnapshotSchedule: Get the snapshot_schedule field of the given VM.
func (vM) GetSnapshotSchedule(session *Session, self VMRef) (retval VMSSRef, err error) {
	method := "VM.get_snapshot_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMSSRef(method+" -> ", result)
	return
}

// GetIsVmssSnapshot: Get the is_vmss_snapshot field of the given VM.
func (vM) GetIsVmssSnapshot(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_is_vmss_snapshot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAppliance: Get the appliance field of the given VM.
func (vM) GetAppliance(session *Session, self VMRef) (retval VMApplianceRef, err error) {
	method := "VM.get_appliance"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMApplianceRef(method+" -> ", result)
	return
}

// GetStartDelay: Get the start_delay field of the given VM.
func (vM) GetStartDelay(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_start_delay"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetShutdownDelay: Get the shutdown_delay field of the given VM.
func (vM) GetShutdownDelay(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_shutdown_delay"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOrder: Get the order field of the given VM.
func (vM) GetOrder(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_order"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVGPUs: Get the VGPUs field of the given VM.
func (vM) GetVGPUs(session *Session, self VMRef) (retval []VGPURef, err error) {
	method := "VM.get_VGPUs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVGPURefSet(method+" -> ", result)
	return
}

// GetAttachedPCIs: Get the attached_PCIs field of the given VM.
func (vM) GetAttachedPCIs(session *Session, self VMRef) (retval []PCIRef, err error) {
	method := "VM.get_attached_PCIs"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSuspendSR: Get the suspend_SR field of the given VM.
func (vM) GetSuspendSR(session *Session, self VMRef) (retval SRRef, err error) {
	method := "VM.get_suspend_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetVersion: Get the version field of the given VM.
func (vM) GetVersion(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetGenerationID: Get the generation_id field of the given VM.
func (vM) GetGenerationID(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_generation_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHardwarePlatformVersion: Get the hardware_platform_version field of the given VM.
func (vM) GetHardwarePlatformVersion(session *Session, self VMRef) (retval int, err error) {
	method := "VM.get_hardware_platform_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHasVendorDevice: Get the has_vendor_device field of the given VM.
func (vM) GetHasVendorDevice(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_has_vendor_device"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetRequiresReboot: Get the requires_reboot field of the given VM.
func (vM) GetRequiresReboot(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_requires_reboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetReferenceLabel: Get the reference_label field of the given VM.
func (vM) GetReferenceLabel(session *Session, self VMRef) (retval string, err error) {
	method := "VM.get_reference_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetDomainType: Get the domain_type field of the given VM.
func (vM) GetDomainType(session *Session, self VMRef) (retval DomainType, err error) {
	method := "VM.get_domain_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumDomainType(method+" -> ", result)
	return
}

// GetNVRAM: Get the NVRAM field of the given VM.
func (vM) GetNVRAM(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.get_NVRAM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPendingGuidances: Get the pending_guidances field of the given VM.
func (vM) GetPendingGuidances(session *Session, self VMRef) (retval []UpdateGuidances, err error) {
	method := "VM.get_pending_guidances"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPendingGuidancesRecommended: Get the pending_guidances_recommended field of the given VM.
func (vM) GetPendingGuidancesRecommended(session *Session, self VMRef) (retval []UpdateGuidances, err error) {
	method := "VM.get_pending_guidances_recommended"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPendingGuidancesFull: Get the pending_guidances_full field of the given VM.
func (vM) GetPendingGuidancesFull(session *Session, self VMRef) (retval []UpdateGuidances, err error) {
	method := "VM.get_pending_guidances_full"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel: Set the name/label field of the given VM.
func (vM) SetNameLabel(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription: Set the name/description field of the given VM.
func (vM) SetNameDescription(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetUserVersion: Set the user_version field of the given VM.
func (vM) SetUserVersion(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_user_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetIsATemplate: Set the is_a_template field of the given VM.
func (vM) SetIsATemplate(session *Session, self VMRef, value bool) (err error) {
	method := "VM.set_is_a_template"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetAffinity: Set the affinity field of the given VM.
func (vM) SetAffinity(session *Session, self VMRef, value HostRef) (err error) {
	method := "VM.set_affinity"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetVCPUsParams: Set the VCPUs/params field of the given VM.
func (vM) SetVCPUsParams(session *Session, self VMRef, value map[string]string) (err error) {
	method := "VM.set_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToVCPUsParams: Add the given key-value pair to the VCPUs/params field of the given VM.
func (vM) AddToVCPUsParams(session *Session, self VMRef, key string, value string) (err error) {
	method := "VM.add_to_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromVCPUsParams: Remove the given key and its corresponding value from the VCPUs/params field of the given VM.  If the key is not in that Map, then do nothing.
func (vM) RemoveFromVCPUsParams(session *Session, self VMRef, key string) (err error) {
	method := "VM.remove_from_VCPUs_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetActionsAfterSoftreboot: Set the actions/after_softreboot field of the given VM.
func (vM) SetActionsAfterSoftreboot(session *Session, self VMRef, value OnSoftrebootBehavior) (err error) {
	method := "VM.set_actions_after_softreboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumOnSoftrebootBehavior(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetActionsAfterShutdown: Set the actions/after_shutdown field of the given VM.
func (vM) SetActionsAfterShutdown(session *Session, self VMRef, value OnNormalExit) (err error) {
	method := "VM.set_actions_after_shutdown"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumOnNormalExit(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetActionsAfterReboot: Set the actions/after_reboot field of the given VM.
func (vM) SetActionsAfterReboot(session *Session, self VMRef, value OnNormalExit) (err error) {
	method := "VM.set_actions_after_reboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumOnNormalExit(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetPVBootloader: Set the PV/bootloader field of the given VM.
func (vM) SetPVBootloader(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_PV_bootloader"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetPVKernel: Set the PV/kernel field of the given VM.
func (vM) SetPVKernel(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_PV_kernel"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetPVRamdisk: Set the PV/ramdisk field of the given VM.
func (vM) SetPVRamdisk(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_PV_ramdisk"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetPVArgs: Set the PV/args field of the given VM.
func (vM) SetPVArgs(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_PV_args"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetPVBootloaderArgs: Set the PV/bootloader_args field of the given VM.
func (vM) SetPVBootloaderArgs(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_PV_bootloader_args"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetPVLegacyArgs: Set the PV/legacy_args field of the given VM.
func (vM) SetPVLegacyArgs(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_PV_legacy_args"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetHVMBootParams: Set the HVM/boot_params field of the given VM.
func (vM) SetHVMBootParams(session *Session, self VMRef, value map[string]string) (err error) {
	method := "VM.set_HVM_boot_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToHVMBootParams: Add the given key-value pair to the HVM/boot_params field of the given VM.
func (vM) AddToHVMBootParams(session *Session, self VMRef, key string, value string) (err error) {
	method := "VM.add_to_HVM_boot_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromHVMBootParams: Remove the given key and its corresponding value from the HVM/boot_params field of the given VM.  If the key is not in that Map, then do nothing.
func (vM) RemoveFromHVMBootParams(session *Session, self VMRef, key string) (err error) {
	method := "VM.remove_from_HVM_boot_params"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetPlatform: Set the platform field of the given VM.
func (vM) SetPlatform(session *Session, self VMRef, value map[string]string) (err error) {
	method := "VM.set_platform"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToPlatform: Add the given key-value pair to the platform field of the given VM.
func (vM) AddToPlatform(session *Session, self VMRef, key string, value string) (err error) {
	method := "VM.add_to_platform"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromPlatform: Remove the given key and its corresponding value from the platform field of the given VM.  If the key is not in that Map, then do nothing.
func (vM) RemoveFromPlatform(session *Session, self VMRef, key string) (err error) {
	method := "VM.remove_from_platform"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetPCIBus: Set the PCI_bus field of the given VM.
func (vM) SetPCIBus(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_PCI_bus"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given VM.
func (vM) SetOtherConfig(session *Session, self VMRef, value map[string]string) (err error) {
	method := "VM.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given VM.
func (vM) AddToOtherConfig(session *Session, self VMRef, key string, value string) (err error) {
	method := "VM.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VM.  If the key is not in that Map, then do nothing.
func (vM) RemoveFromOtherConfig(session *Session, self VMRef, key string) (err error) {
	method := "VM.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetRecommendations: Set the recommendations field of the given VM.
func (vM) SetRecommendations(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_recommendations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetXenstoreData: Set the xenstore_data field of the given VM.
func (vM) SetXenstoreData(session *Session, self VMRef, value map[string]string) (err error) {
	method := "VM.set_xenstore_data"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToXenstoreData: Add the given key-value pair to the xenstore_data field of the given VM.
func (vM) AddToXenstoreData(session *Session, self VMRef, key string, value string) (err error) {
	method := "VM.add_to_xenstore_data"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromXenstoreData: Remove the given key and its corresponding value from the xenstore_data field of the given VM.  If the key is not in that Map, then do nothing.
func (vM) RemoveFromXenstoreData(session *Session, self VMRef, key string) (err error) {
	method := "VM.remove_from_xenstore_data"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetTags: Set the tags field of the given VM.
func (vM) SetTags(session *Session, self VMRef, value []string) (err error) {
	method := "VM.set_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddTags: Add the given value to the tags field of the given VM.  If the value is already in that Set, then do nothing.
func (vM) AddTags(session *Session, self VMRef, value string) (err error) {
	method := "VM.add_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveTags: Remove the given value from the tags field of the given VM.  If the value is not in that Set, then do nothing.
func (vM) RemoveTags(session *Session, self VMRef, value string) (err error) {
	method := "VM.remove_tags"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetBlockedOperations: Set the blocked_operations field of the given VM.
func (vM) SetBlockedOperations(session *Session, self VMRef, value map[VMOperations]string) (err error) {
	method := "VM.set_blocked_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumVMOperationsToStringMap(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AddToBlockedOperations: Add the given key-value pair to the blocked_operations field of the given VM.
func (vM) AddToBlockedOperations(session *Session, self VMRef, key VMOperations, value string) (err error) {
	method := "VM.add_to_blocked_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	keyArg, err := serializeEnumVMOperations(fmt.Sprintf("%s(%s)", method, "key"), key)
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

// RemoveFromBlockedOperations: Remove the given key and its corresponding value from the blocked_operations field of the given VM.  If the key is not in that Map, then do nothing.
func (vM) RemoveFromBlockedOperations(session *Session, self VMRef, key VMOperations) (err error) {
	method := "VM.remove_from_blocked_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	keyArg, err := serializeEnumVMOperations(fmt.Sprintf("%s(%s)", method, "key"), key)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, keyArg)
	return
}

// SetSuspendSR: Set the suspend_SR field of the given VM.
func (vM) SetSuspendSR(session *Session, self VMRef, value SRRef) (err error) {
	method := "VM.set_suspend_SR"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetHardwarePlatformVersion: Set the hardware_platform_version field of the given VM.
func (vM) SetHardwarePlatformVersion(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_hardware_platform_version"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Snapshot: Snapshots the specified VM, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored (e.g. Copy on Write).
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (vM) Snapshot(session *Session, vM VMRef, newName string, ignoreVdis []VDIRef) (retval VMRef, err error) {
	method := "VM.snapshot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	ignoreVdisArg, err := serializeVDIRefSet(fmt.Sprintf("%s(%s)", method, "ignore_vdis"), ignoreVdis)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg, ignoreVdisArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// AsyncSnapshot: Snapshots the specified VM, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored (e.g. Copy on Write).
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (vM) AsyncSnapshot(session *Session, vM VMRef, newName string, ignoreVdis []VDIRef) (retval TaskRef, err error) {
	method := "Async.VM.snapshot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	ignoreVdisArg, err := serializeVDIRefSet(fmt.Sprintf("%s(%s)", method, "ignore_vdis"), ignoreVdis)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg, ignoreVdisArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SnapshotWithQuiesce: Snapshots the specified VM with quiesce, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored (e.g. Copy on Write).
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_SNAPSHOT_WITH_QUIESCE_FAILED - The quiesced-snapshot operation failed for an unexpected reason
// VM_SNAPSHOT_WITH_QUIESCE_TIMEOUT - The VSS plug-in has timed out
// VM_SNAPSHOT_WITH_QUIESCE_PLUGIN_DEOS_NOT_RESPOND - The VSS plug-in cannot be contacted
// VM_SNAPSHOT_WITH_QUIESCE_NOT_SUPPORTED - The VSS plug-in is not installed on this virtual machine
func (vM) SnapshotWithQuiesce(session *Session, vM VMRef, newName string) (retval VMRef, err error) {
	method := "VM.snapshot_with_quiesce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// AsyncSnapshotWithQuiesce: Snapshots the specified VM with quiesce, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored (e.g. Copy on Write).
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_SNAPSHOT_WITH_QUIESCE_FAILED - The quiesced-snapshot operation failed for an unexpected reason
// VM_SNAPSHOT_WITH_QUIESCE_TIMEOUT - The VSS plug-in has timed out
// VM_SNAPSHOT_WITH_QUIESCE_PLUGIN_DEOS_NOT_RESPOND - The VSS plug-in cannot be contacted
// VM_SNAPSHOT_WITH_QUIESCE_NOT_SUPPORTED - The VSS plug-in is not installed on this virtual machine
func (vM) AsyncSnapshotWithQuiesce(session *Session, vM VMRef, newName string) (retval TaskRef, err error) {
	method := "Async.VM.snapshot_with_quiesce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Clone: Clones the specified VM, making a new VM. Clone automatically exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored (e.g. Copy on Write).   This function can only be called when the VM is in the Halted State.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) Clone(session *Session, vM VMRef, newName string) (retval VMRef, err error) {
	method := "VM.clone"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// AsyncClone: Clones the specified VM, making a new VM. Clone automatically exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored (e.g. Copy on Write).   This function can only be called when the VM is in the Halted State.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) AsyncClone(session *Session, vM VMRef, newName string) (retval TaskRef, err error) {
	method := "Async.VM.clone"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Copy: Copied the specified VM, making a new VM. Unlike clone, copy does not exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored. Instead, copy guarantees that the disk images of the newly created VM will be &apos;full disks&apos; - i.e. not part of a CoW chain.  This function can only be called when the VM is in the Halted State.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) Copy(session *Session, vM VMRef, newName string, sr SRRef) (retval VMRef, err error) {
	method := "VM.copy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// AsyncCopy: Copied the specified VM, making a new VM. Unlike clone, copy does not exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored. Instead, copy guarantees that the disk images of the newly created VM will be &apos;full disks&apos; - i.e. not part of a CoW chain.  This function can only be called when the VM is in the Halted State.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) AsyncCopy(session *Session, vM VMRef, newName string, sr SRRef) (retval TaskRef, err error) {
	method := "Async.VM.copy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg, srArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Revert: Reverts the specified VM to a previous state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// VM_REVERT_FAILED - An error occured while reverting the specified virtual machine to the specified snapshot
func (vM) Revert(session *Session, snapshot VMRef) (err error) {
	method := "VM.revert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	snapshotArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "snapshot"), snapshot)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, snapshotArg)
	return
}

// AsyncRevert: Reverts the specified VM to a previous state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// VM_REVERT_FAILED - An error occured while reverting the specified virtual machine to the specified snapshot
func (vM) AsyncRevert(session *Session, snapshot VMRef) (retval TaskRef, err error) {
	method := "Async.VM.revert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	snapshotArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "snapshot"), snapshot)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, snapshotArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Checkpoint: Checkpoints the specified VM, making a new VM. Checkpoint automatically exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored (e.g. Copy on Write) and saves the memory image as well.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_CHECKPOINT_SUSPEND_FAILED - An error occured while saving the memory image of the specified virtual machine
// VM_CHECKPOINT_RESUME_FAILED - An error occured while restoring the memory image of the specified virtual machine
func (vM) Checkpoint(session *Session, vM VMRef, newName string) (retval VMRef, err error) {
	method := "VM.checkpoint"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// AsyncCheckpoint: Checkpoints the specified VM, making a new VM. Checkpoint automatically exploits the capabilities of the underlying storage repository in which the VM&apos;s disk images are stored (e.g. Copy on Write) and saves the memory image as well.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_CHECKPOINT_SUSPEND_FAILED - An error occured while saving the memory image of the specified virtual machine
// VM_CHECKPOINT_RESUME_FAILED - An error occured while restoring the memory image of the specified virtual machine
func (vM) AsyncCheckpoint(session *Session, vM VMRef, newName string) (retval TaskRef, err error) {
	method := "Async.VM.checkpoint"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	newNameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "new_name"), newName)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, newNameArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Provision: Inspects the disk configuration contained within the VM&apos;s other_config, creates VDIs and VBDs and then executes any applicable post-install script.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) Provision(session *Session, vM VMRef) (err error) {
	method := "VM.provision"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncProvision: Inspects the disk configuration contained within the VM&apos;s other_config, creates VDIs and VBDs and then executes any applicable post-install script.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// SR_FULL - The SR is full. Requested new size exceeds the maximum size
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) AsyncProvision(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.provision"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Start: Start the specified VM.  This function can only be called with the VM is in the Halted State.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// VM_HVM_REQUIRED - HVM is required for this operation
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// BOOTLOADER_FAILED - The bootloader returned an error
// UNKNOWN_BOOTLOADER - The requested bootloader is unknown
// NO_HOSTS_AVAILABLE - There were no servers available to complete the specified operation.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) Start(session *Session, vM VMRef, startPaused bool, force bool) (err error) {
	method := "VM.start"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	startPausedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "start_paused"), startPaused)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg, startPausedArg, forceArg)
	return
}

// AsyncStart: Start the specified VM.  This function can only be called with the VM is in the Halted State.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// VM_HVM_REQUIRED - HVM is required for this operation
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// BOOTLOADER_FAILED - The bootloader returned an error
// UNKNOWN_BOOTLOADER - The requested bootloader is unknown
// NO_HOSTS_AVAILABLE - There were no servers available to complete the specified operation.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) AsyncStart(session *Session, vM VMRef, startPaused bool, force bool) (retval TaskRef, err error) {
	method := "Async.VM.start"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	startPausedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "start_paused"), startPaused)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, startPausedArg, forceArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// StartOn: Start the specified VM on a particular host.  This function can only be called with the VM is in the Halted State.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// BOOTLOADER_FAILED - The bootloader returned an error
// UNKNOWN_BOOTLOADER - The requested bootloader is unknown
func (vM) StartOn(session *Session, vM VMRef, host HostRef, startPaused bool, force bool) (err error) {
	method := "VM.start_on"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	startPausedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "start_paused"), startPaused)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg, hostArg, startPausedArg, forceArg)
	return
}

// AsyncStartOn: Start the specified VM on a particular host.  This function can only be called with the VM is in the Halted State.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// BOOTLOADER_FAILED - The bootloader returned an error
// UNKNOWN_BOOTLOADER - The requested bootloader is unknown
func (vM) AsyncStartOn(session *Session, vM VMRef, host HostRef, startPaused bool, force bool) (retval TaskRef, err error) {
	method := "Async.VM.start_on"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	startPausedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "start_paused"), startPaused)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, hostArg, startPausedArg, forceArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Pause: Pause the specified VM. This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) Pause(session *Session, vM VMRef) (err error) {
	method := "VM.pause"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncPause: Pause the specified VM. This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncPause(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.pause"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Unpause: Resume the specified VM. This can only be called when the specified VM is in the Paused state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) Unpause(session *Session, vM VMRef) (err error) {
	method := "VM.unpause"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncUnpause: Resume the specified VM. This can only be called when the specified VM is in the Paused state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncUnpause(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.unpause"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CleanShutdown: Attempt to cleanly shutdown the specified VM. (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) CleanShutdown(session *Session, vM VMRef) (err error) {
	method := "VM.clean_shutdown"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncCleanShutdown: Attempt to cleanly shutdown the specified VM. (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncCleanShutdown(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.clean_shutdown"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Shutdown: Attempts to first clean shutdown a VM and if it should fail then perform a hard shutdown on it.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) Shutdown(session *Session, vM VMRef) (err error) {
	method := "VM.shutdown"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncShutdown: Attempts to first clean shutdown a VM and if it should fail then perform a hard shutdown on it.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncShutdown(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.shutdown"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CleanReboot: Attempt to cleanly shutdown the specified VM (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) CleanReboot(session *Session, vM VMRef) (err error) {
	method := "VM.clean_reboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncCleanReboot: Attempt to cleanly shutdown the specified VM (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncCleanReboot(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.clean_reboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// HardShutdown: Stop executing the specified VM without attempting a clean shutdown.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) HardShutdown(session *Session, vM VMRef) (err error) {
	method := "VM.hard_shutdown"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncHardShutdown: Stop executing the specified VM without attempting a clean shutdown.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncHardShutdown(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.hard_shutdown"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// PowerStateReset: Reset the power-state of the VM to halted in the database only. (Used to recover from slave failures in pooling scenarios by resetting the power-states of VMs running on dead slaves to halted.) This is a potentially dangerous operation; use with care.
func (vM) PowerStateReset(session *Session, vM VMRef) (err error) {
	method := "VM.power_state_reset"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncPowerStateReset: Reset the power-state of the VM to halted in the database only. (Used to recover from slave failures in pooling scenarios by resetting the power-states of VMs running on dead slaves to halted.) This is a potentially dangerous operation; use with care.
func (vM) AsyncPowerStateReset(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.power_state_reset"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// HardReboot: Stop executing the specified VM without attempting a clean shutdown and immediately restart the VM.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) HardReboot(session *Session, vM VMRef) (err error) {
	method := "VM.hard_reboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncHardReboot: Stop executing the specified VM without attempting a clean shutdown and immediately restart the VM.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncHardReboot(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.hard_reboot"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Suspend: Suspend the specified VM to disk.  This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) Suspend(session *Session, vM VMRef) (err error) {
	method := "VM.suspend"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg)
	return
}

// AsyncSuspend: Suspend the specified VM to disk.  This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncSuspend(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.suspend"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Resume: Awaken the specified VM and resume it.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) Resume(session *Session, vM VMRef, startPaused bool, force bool) (err error) {
	method := "VM.resume"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	startPausedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "start_paused"), startPaused)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg, startPausedArg, forceArg)
	return
}

// AsyncResume: Awaken the specified VM and resume it.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncResume(session *Session, vM VMRef, startPaused bool, force bool) (retval TaskRef, err error) {
	method := "Async.VM.resume"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	startPausedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "start_paused"), startPaused)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, startPausedArg, forceArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// ResumeOn: Awaken the specified VM and resume it on a particular Host.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) ResumeOn(session *Session, vM VMRef, host HostRef, startPaused bool, force bool) (err error) {
	method := "VM.resume_on"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	startPausedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "start_paused"), startPaused)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg, hostArg, startPausedArg, forceArg)
	return
}

// AsyncResumeOn: Awaken the specified VM and resume it on a particular Host.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (vM) AsyncResumeOn(session *Session, vM VMRef, host HostRef, startPaused bool, force bool) (retval TaskRef, err error) {
	method := "Async.VM.resume_on"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	startPausedArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "start_paused"), startPaused)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, hostArg, startPausedArg, forceArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// PoolMigrate: Migrate a VM to another Host.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
func (vM) PoolMigrate(session *Session, vM VMRef, host HostRef, options map[string]string) (err error) {
	method := "VM.pool_migrate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg, hostArg, optionsArg)
	return
}

// AsyncPoolMigrate: Migrate a VM to another Host.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
func (vM) AsyncPoolMigrate(session *Session, vM VMRef, host HostRef, options map[string]string) (retval TaskRef, err error) {
	method := "Async.VM.pool_migrate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, hostArg, optionsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetVCPUsNumberLive: Set the number of VCPUs for a running VM
//
// Errors:
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) SetVCPUsNumberLive(session *Session, self VMRef, nvcpu int) (err error) {
	method := "VM.set_VCPUs_number_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	nvcpuArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "nvcpu"), nvcpu)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, nvcpuArg)
	return
}

// AsyncSetVCPUsNumberLive: Set the number of VCPUs for a running VM
//
// Errors:
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) AsyncSetVCPUsNumberLive(session *Session, self VMRef, nvcpu int) (retval TaskRef, err error) {
	method := "Async.VM.set_VCPUs_number_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	nvcpuArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "nvcpu"), nvcpu)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, nvcpuArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// AddToVCPUsParamsLive: Add the given key-value pair to VM.VCPUs_params, and apply that value on the running VM
func (vM) AddToVCPUsParamsLive(session *Session, self VMRef, key string, value string) (err error) {
	method := "VM.add_to_VCPUs_params_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncAddToVCPUsParamsLive: Add the given key-value pair to VM.VCPUs_params, and apply that value on the running VM
func (vM) AsyncAddToVCPUsParamsLive(session *Session, self VMRef, key string, value string) (retval TaskRef, err error) {
	method := "Async.VM.add_to_VCPUs_params_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNVRAM:
func (vM) SetNVRAM(session *Session, self VMRef, value map[string]string) (err error) {
	method := "VM.set_NVRAM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToNVRAM:
func (vM) AddToNVRAM(session *Session, self VMRef, key string, value string) (err error) {
	method := "VM.add_to_NVRAM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromNVRAM:
func (vM) RemoveFromNVRAM(session *Session, self VMRef, key string) (err error) {
	method := "VM.remove_from_NVRAM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetHaRestartPriority: Set the value of the ha_restart_priority field
func (vM) SetHaRestartPriority(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_ha_restart_priority"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetHaAlwaysRun: Set the value of the ha_always_run
func (vM) SetHaAlwaysRun(session *Session, self VMRef, value bool) (err error) {
	method := "VM.set_ha_always_run"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// ComputeMemoryOverhead: Computes the virtualization memory overhead of a VM.
func (vM) ComputeMemoryOverhead(session *Session, vM VMRef) (retval int, err error) {
	method := "VM.compute_memory_overhead"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// AsyncComputeMemoryOverhead: Computes the virtualization memory overhead of a VM.
func (vM) AsyncComputeMemoryOverhead(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.compute_memory_overhead"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetMemoryDynamicMax: Set the value of the memory_dynamic_max field
func (vM) SetMemoryDynamicMax(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_memory_dynamic_max"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetMemoryDynamicMin: Set the value of the memory_dynamic_min field
func (vM) SetMemoryDynamicMin(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_memory_dynamic_min"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetMemoryDynamicRange: Set the minimum and maximum amounts of physical memory the VM is allowed to use.
func (vM) SetMemoryDynamicRange(session *Session, self VMRef, min int, max int) (err error) {
	method := "VM.set_memory_dynamic_range"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	minArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "min"), min)
	if err != nil {
		return
	}
	maxArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "max"), max)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, minArg, maxArg)
	return
}

// AsyncSetMemoryDynamicRange: Set the minimum and maximum amounts of physical memory the VM is allowed to use.
func (vM) AsyncSetMemoryDynamicRange(session *Session, self VMRef, min int, max int) (retval TaskRef, err error) {
	method := "Async.VM.set_memory_dynamic_range"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	minArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "min"), min)
	if err != nil {
		return
	}
	maxArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "max"), max)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, minArg, maxArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetMemoryStaticMax: Set the value of the memory_static_max field
//
// Errors:
// HA_OPERATION_WOULD_BREAK_FAILOVER_PLAN - This operation cannot be performed because it would invalidate VM failover planning such that the system would be unable to guarantee to restart protected VMs after a Host failure.
func (vM) SetMemoryStaticMax(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_memory_static_max"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetMemoryStaticMin: Set the value of the memory_static_min field
func (vM) SetMemoryStaticMin(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_memory_static_min"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetMemoryStaticRange: Set the static (ie boot-time) range of virtual memory that the VM is allowed to use.
func (vM) SetMemoryStaticRange(session *Session, self VMRef, min int, max int) (err error) {
	method := "VM.set_memory_static_range"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	minArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "min"), min)
	if err != nil {
		return
	}
	maxArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "max"), max)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, minArg, maxArg)
	return
}

// AsyncSetMemoryStaticRange: Set the static (ie boot-time) range of virtual memory that the VM is allowed to use.
func (vM) AsyncSetMemoryStaticRange(session *Session, self VMRef, min int, max int) (retval TaskRef, err error) {
	method := "Async.VM.set_memory_static_range"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	minArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "min"), min)
	if err != nil {
		return
	}
	maxArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "max"), max)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, minArg, maxArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetMemoryLimits: Set the memory limits of this VM.
func (vM) SetMemoryLimits(session *Session, self VMRef, staticMin int, staticMax int, dynamicMin int, dynamicMax int) (err error) {
	method := "VM.set_memory_limits"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	staticMinArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "static_min"), staticMin)
	if err != nil {
		return
	}
	staticMaxArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "static_max"), staticMax)
	if err != nil {
		return
	}
	dynamicMinArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "dynamic_min"), dynamicMin)
	if err != nil {
		return
	}
	dynamicMaxArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "dynamic_max"), dynamicMax)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, staticMinArg, staticMaxArg, dynamicMinArg, dynamicMaxArg)
	return
}

// AsyncSetMemoryLimits: Set the memory limits of this VM.
func (vM) AsyncSetMemoryLimits(session *Session, self VMRef, staticMin int, staticMax int, dynamicMin int, dynamicMax int) (retval TaskRef, err error) {
	method := "Async.VM.set_memory_limits"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	staticMinArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "static_min"), staticMin)
	if err != nil {
		return
	}
	staticMaxArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "static_max"), staticMax)
	if err != nil {
		return
	}
	dynamicMinArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "dynamic_min"), dynamicMin)
	if err != nil {
		return
	}
	dynamicMaxArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "dynamic_max"), dynamicMax)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, staticMinArg, staticMaxArg, dynamicMinArg, dynamicMaxArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetMemory: Set the memory allocation of this VM. Sets all of memory_static_max, memory_dynamic_min, and memory_dynamic_max to the given value, and leaves memory_static_min untouched.
func (vM) SetMemory(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_memory"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetMemory: Set the memory allocation of this VM. Sets all of memory_static_max, memory_dynamic_min, and memory_dynamic_max to the given value, and leaves memory_static_min untouched.
func (vM) AsyncSetMemory(session *Session, self VMRef, value int) (retval TaskRef, err error) {
	method := "Async.VM.set_memory"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetMemoryTargetLive: Set the memory target for a running VM
func (vM) SetMemoryTargetLive(session *Session, self VMRef, target int) (err error) {
	method := "VM.set_memory_target_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	targetArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "target"), target)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, targetArg)
	return
}

// AsyncSetMemoryTargetLive: Set the memory target for a running VM
func (vM) AsyncSetMemoryTargetLive(session *Session, self VMRef, target int) (retval TaskRef, err error) {
	method := "Async.VM.set_memory_target_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	targetArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "target"), target)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, targetArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// WaitMemoryTargetLive: Wait for a running VM to reach its current memory target
func (vM) WaitMemoryTargetLive(session *Session, self VMRef) (err error) {
	method := "VM.wait_memory_target_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncWaitMemoryTargetLive: Wait for a running VM to reach its current memory target
func (vM) AsyncWaitMemoryTargetLive(session *Session, self VMRef) (retval TaskRef, err error) {
	method := "Async.VM.wait_memory_target_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetCooperative: Return true if the VM is currently &apos;co-operative&apos; i.e. is expected to reach a balloon target and actually has done
func (vM) GetCooperative(session *Session, self VMRef) (retval bool, err error) {
	method := "VM.get_cooperative"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncGetCooperative: Return true if the VM is currently &apos;co-operative&apos; i.e. is expected to reach a balloon target and actually has done
func (vM) AsyncGetCooperative(session *Session, self VMRef) (retval TaskRef, err error) {
	method := "Async.VM.get_cooperative"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetHVMShadowMultiplier: Set the shadow memory multiplier on a halted VM
func (vM) SetHVMShadowMultiplier(session *Session, self VMRef, value float64) (err error) {
	method := "VM.set_HVM_shadow_multiplier"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetShadowMultiplierLive: Set the shadow memory multiplier on a running VM
func (vM) SetShadowMultiplierLive(session *Session, self VMRef, multiplier float64) (err error) {
	method := "VM.set_shadow_multiplier_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	multiplierArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "multiplier"), multiplier)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, multiplierArg)
	return
}

// AsyncSetShadowMultiplierLive: Set the shadow memory multiplier on a running VM
func (vM) AsyncSetShadowMultiplierLive(session *Session, self VMRef, multiplier float64) (retval TaskRef, err error) {
	method := "Async.VM.set_shadow_multiplier_live"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	multiplierArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "multiplier"), multiplier)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, multiplierArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetVCPUsMax: Set the maximum number of VCPUs for a halted VM
func (vM) SetVCPUsMax(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_VCPUs_max"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetVCPUsAtStartup: Set the number of startup VCPUs for a halted VM
func (vM) SetVCPUsAtStartup(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_VCPUs_at_startup"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SendSysrq: Send the given key as a sysrq to this VM.  The key is specified as a single character (a String of length 1).  This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
func (vM) SendSysrq(session *Session, vM VMRef, key string) (err error) {
	method := "VM.send_sysrq"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	keyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "key"), key)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg, keyArg)
	return
}

// AsyncSendSysrq: Send the given key as a sysrq to this VM.  The key is specified as a single character (a String of length 1).  This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
func (vM) AsyncSendSysrq(session *Session, vM VMRef, key string) (retval TaskRef, err error) {
	method := "Async.VM.send_sysrq"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	keyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "key"), key)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, keyArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SendTrigger: Send the named trigger to this VM.  This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
func (vM) SendTrigger(session *Session, vM VMRef, trigger string) (err error) {
	method := "VM.send_trigger"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	triggerArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "trigger"), trigger)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg, triggerArg)
	return
}

// AsyncSendTrigger: Send the named trigger to this VM.  This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
func (vM) AsyncSendTrigger(session *Session, vM VMRef, trigger string) (retval TaskRef, err error) {
	method := "Async.VM.send_trigger"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	triggerArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "trigger"), trigger)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, triggerArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// MaximiseMemory: Returns the maximum amount of guest memory which will fit, together with overheads, in the supplied amount of physical memory. If &apos;exact&apos; is true then an exact calculation is performed using the VM&apos;s current settings. If &apos;exact&apos; is false then a more conservative approximation is used
func (vM) MaximiseMemory(session *Session, self VMRef, total int, approximate bool) (retval int, err error) {
	method := "VM.maximise_memory"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	totalArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "total"), total)
	if err != nil {
		return
	}
	approximateArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "approximate"), approximate)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, totalArg, approximateArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// AsyncMaximiseMemory: Returns the maximum amount of guest memory which will fit, together with overheads, in the supplied amount of physical memory. If &apos;exact&apos; is true then an exact calculation is performed using the VM&apos;s current settings. If &apos;exact&apos; is false then a more conservative approximation is used
func (vM) AsyncMaximiseMemory(session *Session, self VMRef, total int, approximate bool) (retval TaskRef, err error) {
	method := "Async.VM.maximise_memory"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	totalArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "total"), total)
	if err != nil {
		return
	}
	approximateArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "approximate"), approximate)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, totalArg, approximateArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// MigrateSend: Migrate the VM to another host.  This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) MigrateSend(session *Session, vM VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (retval VMRef, err error) {
	method := "VM.migrate_send"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	destArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "dest"), dest)
	if err != nil {
		return
	}
	liveArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "live"), live)
	if err != nil {
		return
	}
	vdiMapArg, err := serializeVDIRefToSRRefMap(fmt.Sprintf("%s(%s)", method, "vdi_map"), vdiMap)
	if err != nil {
		return
	}
	vifMapArg, err := serializeVIFRefToNetworkRefMap(fmt.Sprintf("%s(%s)", method, "vif_map"), vifMap)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	vgpuMapArg, err := serializeVGPURefToGPUGroupRefMap(fmt.Sprintf("%s(%s)", method, "vgpu_map"), vgpuMap)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, destArg, liveArg, vdiMapArg, vifMapArg, optionsArg, vgpuMapArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRef(method+" -> ", result)
	return
}

// AsyncMigrateSend: Migrate the VM to another host.  This can only be called when the specified VM is in the Running state.
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) AsyncMigrateSend(session *Session, vM VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (retval TaskRef, err error) {
	method := "Async.VM.migrate_send"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	destArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "dest"), dest)
	if err != nil {
		return
	}
	liveArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "live"), live)
	if err != nil {
		return
	}
	vdiMapArg, err := serializeVDIRefToSRRefMap(fmt.Sprintf("%s(%s)", method, "vdi_map"), vdiMap)
	if err != nil {
		return
	}
	vifMapArg, err := serializeVIFRefToNetworkRefMap(fmt.Sprintf("%s(%s)", method, "vif_map"), vifMap)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	vgpuMapArg, err := serializeVGPURefToGPUGroupRefMap(fmt.Sprintf("%s(%s)", method, "vgpu_map"), vgpuMap)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, destArg, liveArg, vdiMapArg, vifMapArg, optionsArg, vgpuMapArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// AssertCanMigrate: Assert whether a VM can be migrated to the specified destination.
//
// Errors:
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) AssertCanMigrate(session *Session, vM VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (err error) {
	method := "VM.assert_can_migrate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	destArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "dest"), dest)
	if err != nil {
		return
	}
	liveArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "live"), live)
	if err != nil {
		return
	}
	vdiMapArg, err := serializeVDIRefToSRRefMap(fmt.Sprintf("%s(%s)", method, "vdi_map"), vdiMap)
	if err != nil {
		return
	}
	vifMapArg, err := serializeVIFRefToNetworkRefMap(fmt.Sprintf("%s(%s)", method, "vif_map"), vifMap)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	vgpuMapArg, err := serializeVGPURefToGPUGroupRefMap(fmt.Sprintf("%s(%s)", method, "vgpu_map"), vgpuMap)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg, destArg, liveArg, vdiMapArg, vifMapArg, optionsArg, vgpuMapArg)
	return
}

// AsyncAssertCanMigrate: Assert whether a VM can be migrated to the specified destination.
//
// Errors:
// LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature. Please contact your support representative.
func (vM) AsyncAssertCanMigrate(session *Session, vM VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (retval TaskRef, err error) {
	method := "Async.VM.assert_can_migrate"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	destArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "dest"), dest)
	if err != nil {
		return
	}
	liveArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "live"), live)
	if err != nil {
		return
	}
	vdiMapArg, err := serializeVDIRefToSRRefMap(fmt.Sprintf("%s(%s)", method, "vdi_map"), vdiMap)
	if err != nil {
		return
	}
	vifMapArg, err := serializeVIFRefToNetworkRefMap(fmt.Sprintf("%s(%s)", method, "vif_map"), vifMap)
	if err != nil {
		return
	}
	optionsArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "options"), options)
	if err != nil {
		return
	}
	vgpuMapArg, err := serializeVGPURefToGPUGroupRefMap(fmt.Sprintf("%s(%s)", method, "vgpu_map"), vgpuMap)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, destArg, liveArg, vdiMapArg, vifMapArg, optionsArg, vgpuMapArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetBootRecord: Returns a record describing the VM&apos;s dynamic state, initialised when the VM boots and updated to reflect runtime configuration changes e.g. CPU hotplug
func (vM) GetBootRecord(session *Session, self VMRef) (retval VMRecord, err error) {
	method := "VM.get_boot_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRecord(method+" -> ", result)
	return
}

// GetDataSources:
func (vM) GetDataSources(session *Session, self VMRef) (retval []DataSourceRecord, err error) {
	method := "VM.get_data_sources"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeDataSourceRecordSet(method+" -> ", result)
	return
}

// RecordDataSource: Start recording the specified data source
func (vM) RecordDataSource(session *Session, self VMRef, dataSource string) (err error) {
	method := "VM.record_data_source"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	dataSourceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "data_source"), dataSource)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, dataSourceArg)
	return
}

// QueryDataSource: Query the latest value of the specified data source
func (vM) QueryDataSource(session *Session, self VMRef, dataSource string) (retval float64, err error) {
	method := "VM.query_data_source"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	dataSourceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "data_source"), dataSource)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, dataSourceArg)
	if err != nil {
		return
	}
	retval, err = deserializeFloat(method+" -> ", result)
	return
}

// ForgetDataSourceArchives: Forget the recorded statistics related to the specified data source
func (vM) ForgetDataSourceArchives(session *Session, self VMRef, dataSource string) (err error) {
	method := "VM.forget_data_source_archives"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	dataSourceArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "data_source"), dataSource)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, dataSourceArg)
	return
}

// AssertOperationValid: Check to see whether this operation is acceptable in the current state of the system, raising an error if the operation is invalid for some reason
func (vM) AssertOperationValid(session *Session, self VMRef, op VMOperations) (err error) {
	method := "VM.assert_operation_valid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	opArg, err := serializeEnumVMOperations(fmt.Sprintf("%s(%s)", method, "op"), op)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, opArg)
	return
}

// AsyncAssertOperationValid: Check to see whether this operation is acceptable in the current state of the system, raising an error if the operation is invalid for some reason
func (vM) AsyncAssertOperationValid(session *Session, self VMRef, op VMOperations) (retval TaskRef, err error) {
	method := "Async.VM.assert_operation_valid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	opArg, err := serializeEnumVMOperations(fmt.Sprintf("%s(%s)", method, "op"), op)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, opArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// UpdateAllowedOperations: Recomputes the list of acceptable operations
func (vM) UpdateAllowedOperations(session *Session, self VMRef) (err error) {
	method := "VM.update_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncUpdateAllowedOperations: Recomputes the list of acceptable operations
func (vM) AsyncUpdateAllowedOperations(session *Session, self VMRef) (retval TaskRef, err error) {
	method := "Async.VM.update_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedVBDDevices: Returns a list of the allowed values that a VBD device field can take
func (vM) GetAllowedVBDDevices(session *Session, vM VMRef) (retval []string, err error) {
	method := "VM.get_allowed_VBD_devices"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// GetAllowedVIFDevices: Returns a list of the allowed values that a VIF device field can take
func (vM) GetAllowedVIFDevices(session *Session, vM VMRef) (retval []string, err error) {
	method := "VM.get_allowed_VIF_devices"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// GetPossibleHosts: Return the list of hosts on which this VM may run.
func (vM) GetPossibleHosts(session *Session, vM VMRef) (retval []HostRef, err error) {
	method := "VM.get_possible_hosts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRefSet(method+" -> ", result)
	return
}

// AsyncGetPossibleHosts: Return the list of hosts on which this VM may run.
func (vM) AsyncGetPossibleHosts(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.get_possible_hosts"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// AssertCanBootHere: Returns an error if the VM could not boot on this host for some reason
//
// Errors:
// HOST_NOT_ENOUGH_FREE_MEMORY - Not enough server memory is available to perform this operation.
// HOST_NOT_ENOUGH_PCPUS - The host does not have enough pCPUs to run the VM. It needs at least as many as the VM has vCPUs.
// NETWORK_SRIOV_INSUFFICIENT_CAPACITY - There is insufficient capacity for VF reservation
// HOST_NOT_LIVE - This operation cannot be completed as the server is not live.
// HOST_DISABLED - The specified server is disabled.
// HOST_CANNOT_ATTACH_NETWORK - Server cannot attach network (in the case of NIC bonding, this may be because attaching the network on this server would require other networks - that are currently active - to be taken down).
// VM_HVM_REQUIRED - HVM is required for this operation
// VM_REQUIRES_GPU - You attempted to run a VM on a host which doesn&apos;t have a pGPU available in the GPU group needed by the VM. The VM has a vGPU attached to this GPU group.
// VM_REQUIRES_IOMMU - You attempted to run a VM on a host which doesn&apos;t have I/O virtualization (IOMMU/VT-d) enabled, which is needed by the VM.
// VM_REQUIRES_NETWORK - You attempted to run a VM on a host which doesn&apos;t have a PIF on a Network needed by the VM. The VM has at least one VIF attached to the Network.
// VM_REQUIRES_SR - You attempted to run a VM on a host which doesn&apos;t have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
// VM_REQUIRES_VGPU - You attempted to run a VM on a host on which the vGPU required by the VM cannot be allocated on any pGPUs in the GPU_group needed by the VM.
// VM_HOST_INCOMPATIBLE_VERSION - This VM operation cannot be performed on an older-versioned host during an upgrade.
// VM_HOST_INCOMPATIBLE_VIRTUAL_HARDWARE_PLATFORM_VERSION - You attempted to run a VM on a host that cannot provide the VM&apos;s required Virtual Hardware Platform version.
// INVALID_VALUE - The value given is invalid
// MEMORY_CONSTRAINT_VIOLATION - The dynamic memory range does not satisfy the following constraint.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VALUE_NOT_SUPPORTED - You attempted to set a value that is not supported by this implementation. The fully-qualified field name and the value that you tried to set are returned. Also returned is a developer-only diagnostic reason.
// VM_INCOMPATIBLE_WITH_THIS_HOST - The VM is incompatible with the CPU features of this host.
func (vM) AssertCanBootHere(session *Session, self VMRef, host HostRef) (err error) {
	method := "VM.assert_can_boot_here"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, hostArg)
	return
}

// AsyncAssertCanBootHere: Returns an error if the VM could not boot on this host for some reason
//
// Errors:
// HOST_NOT_ENOUGH_FREE_MEMORY - Not enough server memory is available to perform this operation.
// HOST_NOT_ENOUGH_PCPUS - The host does not have enough pCPUs to run the VM. It needs at least as many as the VM has vCPUs.
// NETWORK_SRIOV_INSUFFICIENT_CAPACITY - There is insufficient capacity for VF reservation
// HOST_NOT_LIVE - This operation cannot be completed as the server is not live.
// HOST_DISABLED - The specified server is disabled.
// HOST_CANNOT_ATTACH_NETWORK - Server cannot attach network (in the case of NIC bonding, this may be because attaching the network on this server would require other networks - that are currently active - to be taken down).
// VM_HVM_REQUIRED - HVM is required for this operation
// VM_REQUIRES_GPU - You attempted to run a VM on a host which doesn&apos;t have a pGPU available in the GPU group needed by the VM. The VM has a vGPU attached to this GPU group.
// VM_REQUIRES_IOMMU - You attempted to run a VM on a host which doesn&apos;t have I/O virtualization (IOMMU/VT-d) enabled, which is needed by the VM.
// VM_REQUIRES_NETWORK - You attempted to run a VM on a host which doesn&apos;t have a PIF on a Network needed by the VM. The VM has at least one VIF attached to the Network.
// VM_REQUIRES_SR - You attempted to run a VM on a host which doesn&apos;t have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
// VM_REQUIRES_VGPU - You attempted to run a VM on a host on which the vGPU required by the VM cannot be allocated on any pGPUs in the GPU_group needed by the VM.
// VM_HOST_INCOMPATIBLE_VERSION - This VM operation cannot be performed on an older-versioned host during an upgrade.
// VM_HOST_INCOMPATIBLE_VIRTUAL_HARDWARE_PLATFORM_VERSION - You attempted to run a VM on a host that cannot provide the VM&apos;s required Virtual Hardware Platform version.
// INVALID_VALUE - The value given is invalid
// MEMORY_CONSTRAINT_VIOLATION - The dynamic memory range does not satisfy the following constraint.
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VALUE_NOT_SUPPORTED - You attempted to set a value that is not supported by this implementation. The fully-qualified field name and the value that you tried to set are returned. Also returned is a developer-only diagnostic reason.
// VM_INCOMPATIBLE_WITH_THIS_HOST - The VM is incompatible with the CPU features of this host.
func (vM) AsyncAssertCanBootHere(session *Session, self VMRef, host HostRef) (retval TaskRef, err error) {
	method := "Async.VM.assert_can_boot_here"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this VM
func (vM) CreateNewBlob(session *Session, vM VMRef, name string, mimeType string, public bool) (retval BlobRef, err error) {
	method := "VM.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
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
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeBlobRef(method+" -> ", result)
	return
}

// AsyncCreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this VM
func (vM) AsyncCreateNewBlob(session *Session, vM VMRef, name string, mimeType string, public bool) (retval TaskRef, err error) {
	method := "Async.VM.create_new_blob"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
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
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, nameArg, mimeTypeArg, publicArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// AssertAgile: Returns an error if the VM is not considered agile e.g. because it is tied to a resource local to a host
func (vM) AssertAgile(session *Session, self VMRef) (err error) {
	method := "VM.assert_agile"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncAssertAgile: Returns an error if the VM is not considered agile e.g. because it is tied to a resource local to a host
func (vM) AsyncAssertAgile(session *Session, self VMRef) (retval TaskRef, err error) {
	method := "Async.VM.assert_agile"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RetrieveWlbRecommendations: Returns mapping of hosts to ratings, indicating the suitability of starting the VM at that location according to wlb. Rating is replaced with an error if the VM cannot boot there.
func (vM) RetrieveWlbRecommendations(session *Session, vM VMRef) (retval map[HostRef][]string, err error) {
	method := "VM.retrieve_wlb_recommendations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeHostRefToStringSetMap(method+" -> ", result)
	return
}

// AsyncRetrieveWlbRecommendations: Returns mapping of hosts to ratings, indicating the suitability of starting the VM at that location according to wlb. Rating is replaced with an error if the VM cannot boot there.
func (vM) AsyncRetrieveWlbRecommendations(session *Session, vM VMRef) (retval TaskRef, err error) {
	method := "Async.VM.retrieve_wlb_recommendations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetBiosStrings: Set custom BIOS strings to this VM. VM will be given a default set of BIOS strings, only some of which can be overridden by the supplied values. Allowed keys are: &apos;bios-vendor&apos;, &apos;bios-version&apos;, &apos;system-manufacturer&apos;, &apos;system-product-name&apos;, &apos;system-version&apos;, &apos;system-serial-number&apos;, &apos;enclosure-asset-tag&apos;, &apos;baseboard-manufacturer&apos;, &apos;baseboard-product-name&apos;, &apos;baseboard-version&apos;, &apos;baseboard-serial-number&apos;, &apos;baseboard-asset-tag&apos;, &apos;baseboard-location-in-chassis&apos;, &apos;enclosure-asset-tag&apos;
//
// Errors:
// VM_BIOS_STRINGS_ALREADY_SET - The BIOS strings for this VM have already been set and cannot be changed.
// INVALID_VALUE - The value given is invalid
func (vM) SetBiosStrings(session *Session, self VMRef, value map[string]string) (err error) {
	method := "VM.set_bios_strings"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetBiosStrings: Set custom BIOS strings to this VM. VM will be given a default set of BIOS strings, only some of which can be overridden by the supplied values. Allowed keys are: &apos;bios-vendor&apos;, &apos;bios-version&apos;, &apos;system-manufacturer&apos;, &apos;system-product-name&apos;, &apos;system-version&apos;, &apos;system-serial-number&apos;, &apos;enclosure-asset-tag&apos;, &apos;baseboard-manufacturer&apos;, &apos;baseboard-product-name&apos;, &apos;baseboard-version&apos;, &apos;baseboard-serial-number&apos;, &apos;baseboard-asset-tag&apos;, &apos;baseboard-location-in-chassis&apos;, &apos;enclosure-asset-tag&apos;
//
// Errors:
// VM_BIOS_STRINGS_ALREADY_SET - The BIOS strings for this VM have already been set and cannot be changed.
// INVALID_VALUE - The value given is invalid
func (vM) AsyncSetBiosStrings(session *Session, self VMRef, value map[string]string) (retval TaskRef, err error) {
	method := "Async.VM.set_bios_strings"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// CopyBiosStrings: Copy the BIOS strings from the given host to this VM
func (vM) CopyBiosStrings(session *Session, vM VMRef, host HostRef) (err error) {
	method := "VM.copy_bios_strings"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, vMArg, hostArg)
	return
}

// AsyncCopyBiosStrings: Copy the BIOS strings from the given host to this VM
func (vM) AsyncCopyBiosStrings(session *Session, vM VMRef, host HostRef) (retval TaskRef, err error) {
	method := "Async.VM.copy_bios_strings"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
	if err != nil {
		return
	}
	hostArg, err := serializeHostRef(fmt.Sprintf("%s(%s)", method, "host"), host)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, hostArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetProtectionPolicy: Set the value of the protection_policy field
func (vM) SetProtectionPolicy(session *Session, self VMRef, value VMPPRef) (err error) {
	method := "VM.set_protection_policy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVMPPRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetSnapshotSchedule: Set the value of the snapshot schedule field
func (vM) SetSnapshotSchedule(session *Session, self VMRef, value VMSSRef) (err error) {
	method := "VM.set_snapshot_schedule"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVMSSRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetStartDelay: Set this VM&apos;s start delay in seconds
func (vM) SetStartDelay(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_start_delay"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetStartDelay: Set this VM&apos;s start delay in seconds
func (vM) AsyncSetStartDelay(session *Session, self VMRef, value int) (retval TaskRef, err error) {
	method := "Async.VM.set_start_delay"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetShutdownDelay: Set this VM&apos;s shutdown delay in seconds
func (vM) SetShutdownDelay(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_shutdown_delay"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetShutdownDelay: Set this VM&apos;s shutdown delay in seconds
func (vM) AsyncSetShutdownDelay(session *Session, self VMRef, value int) (retval TaskRef, err error) {
	method := "Async.VM.set_shutdown_delay"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOrder: Set this VM&apos;s boot order
func (vM) SetOrder(session *Session, self VMRef, value int) (err error) {
	method := "VM.set_order"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetOrder: Set this VM&apos;s boot order
func (vM) AsyncSetOrder(session *Session, self VMRef, value int) (retval TaskRef, err error) {
	method := "Async.VM.set_order"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetSuspendVDI: Set this VM&apos;s suspend VDI, which must be indentical to its current one
func (vM) SetSuspendVDI(session *Session, self VMRef, value VDIRef) (err error) {
	method := "VM.set_suspend_VDI"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetSuspendVDI: Set this VM&apos;s suspend VDI, which must be indentical to its current one
func (vM) AsyncSetSuspendVDI(session *Session, self VMRef, value VDIRef) (retval TaskRef, err error) {
	method := "Async.VM.set_suspend_VDI"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVDIRef(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// AssertCanBeRecovered: Assert whether all SRs required to recover this VM are available.
//
// Errors:
// VM_IS_PART_OF_AN_APPLIANCE - This operation is not allowed as the VM is part of an appliance.
// VM_REQUIRES_SR - You attempted to run a VM on a host which doesn&apos;t have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
func (vM) AssertCanBeRecovered(session *Session, self VMRef, sessionTo SessionRef) (err error) {
	method := "VM.assert_can_be_recovered"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	sessionToArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_to"), sessionTo)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, sessionToArg)
	return
}

// AsyncAssertCanBeRecovered: Assert whether all SRs required to recover this VM are available.
//
// Errors:
// VM_IS_PART_OF_AN_APPLIANCE - This operation is not allowed as the VM is part of an appliance.
// VM_REQUIRES_SR - You attempted to run a VM on a host which doesn&apos;t have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
func (vM) AsyncAssertCanBeRecovered(session *Session, self VMRef, sessionTo SessionRef) (retval TaskRef, err error) {
	method := "Async.VM.assert_can_be_recovered"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	sessionToArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_to"), sessionTo)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, sessionToArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetSRsRequiredForRecovery: List all the SR&apos;s that are required for the VM to be recovered
func (vM) GetSRsRequiredForRecovery(session *Session, self VMRef, sessionTo SessionRef) (retval []SRRef, err error) {
	method := "VM.get_SRs_required_for_recovery"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	sessionToArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_to"), sessionTo)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, sessionToArg)
	if err != nil {
		return
	}
	retval, err = deserializeSRRefSet(method+" -> ", result)
	return
}

// AsyncGetSRsRequiredForRecovery: List all the SR&apos;s that are required for the VM to be recovered
func (vM) AsyncGetSRsRequiredForRecovery(session *Session, self VMRef, sessionTo SessionRef) (retval TaskRef, err error) {
	method := "Async.VM.get_SRs_required_for_recovery"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	sessionToArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_to"), sessionTo)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, sessionToArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Recover: Recover the VM
func (vM) Recover(session *Session, self VMRef, sessionTo SessionRef, force bool) (err error) {
	method := "VM.recover"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	sessionToArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_to"), sessionTo)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, sessionToArg, forceArg)
	return
}

// AsyncRecover: Recover the VM
func (vM) AsyncRecover(session *Session, self VMRef, sessionTo SessionRef, force bool) (retval TaskRef, err error) {
	method := "Async.VM.recover"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	sessionToArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_to"), sessionTo)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg, sessionToArg, forceArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// ImportConvert: Import using a conversion service.
func (vM) ImportConvert(session *Session, typeKey string, username string, password string, sr SRRef, remoteConfig map[string]string) (err error) {
	method := "VM.import_convert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
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
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	remoteConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "remote_config"), remoteConfig)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, typeKeyArg, usernameArg, passwordArg, srArg, remoteConfigArg)
	return
}

// AsyncImportConvert: Import using a conversion service.
func (vM) AsyncImportConvert(session *Session, typeKey string, username string, password string, sr SRRef, remoteConfig map[string]string) (retval TaskRef, err error) {
	method := "Async.VM.import_convert"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	typeKeyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "type"), typeKey)
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
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	remoteConfigArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "remote_config"), remoteConfig)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, typeKeyArg, usernameArg, passwordArg, srArg, remoteConfigArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetAppliance: Assign this VM to an appliance.
func (vM) SetAppliance(session *Session, self VMRef, value VMApplianceRef) (err error) {
	method := "VM.set_appliance"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVMApplianceRef(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetAppliance: Assign this VM to an appliance.
func (vM) AsyncSetAppliance(session *Session, self VMRef, value VMApplianceRef) (retval TaskRef, err error) {
	method := "Async.VM.set_appliance"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeVMApplianceRef(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// QueryServices: Query the system services advertised by this VM and register them. This can only be applied to a system domain.
func (vM) QueryServices(session *Session, self VMRef) (retval map[string]string, err error) {
	method := "VM.query_services"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncQueryServices: Query the system services advertised by this VM and register them. This can only be applied to a system domain.
func (vM) AsyncQueryServices(session *Session, self VMRef) (retval TaskRef, err error) {
	method := "Async.VM.query_services"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// CallPlugin: Call an API plugin on this vm
func (vM) CallPlugin(session *Session, vM VMRef, plugin string, fn string, args map[string]string) (retval string, err error) {
	method := "VM.call_plugin"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
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
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, pluginArg, fnArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// AsyncCallPlugin: Call an API plugin on this vm
func (vM) AsyncCallPlugin(session *Session, vM VMRef, plugin string, fn string, args map[string]string) (retval TaskRef, err error) {
	method := "Async.VM.call_plugin"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vm"), vM)
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
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, pluginArg, fnArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetHasVendorDevice: Controls whether, when the VM starts in HVM mode, its virtual hardware will include the emulated PCI device for which drivers may be available through Windows Update. Usually this should never be changed on a VM on which Windows has been installed: changing it on such a VM is likely to lead to a crash on next start.
func (vM) SetHasVendorDevice(session *Session, self VMRef, value bool) (err error) {
	method := "VM.set_has_vendor_device"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetHasVendorDevice: Controls whether, when the VM starts in HVM mode, its virtual hardware will include the emulated PCI device for which drivers may be available through Windows Update. Usually this should never be changed on a VM on which Windows has been installed: changing it on such a VM is likely to lead to a crash on next start.
func (vM) AsyncSetHasVendorDevice(session *Session, self VMRef, value bool) (retval TaskRef, err error) {
	method := "Async.VM.set_has_vendor_device"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Import: Import an XVA from a URI
func (vM) Import(session *Session, uRL string, sr SRRef, fullRestore bool, force bool) (retval []VMRef, err error) {
	method := "VM.import"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	uRLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "url"), uRL)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	fullRestoreArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "full_restore"), fullRestore)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uRLArg, srArg, fullRestoreArg, forceArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefSet(method+" -> ", result)
	return
}

// AsyncImport: Import an XVA from a URI
func (vM) AsyncImport(session *Session, uRL string, sr SRRef, fullRestore bool, force bool) (retval TaskRef, err error) {
	method := "Async.VM.import"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	uRLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "url"), uRL)
	if err != nil {
		return
	}
	srArg, err := serializeSRRef(fmt.Sprintf("%s(%s)", method, "sr"), sr)
	if err != nil {
		return
	}
	fullRestoreArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "full_restore"), fullRestore)
	if err != nil {
		return
	}
	forceArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "force"), force)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uRLArg, srArg, fullRestoreArg, forceArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetActionsAfterCrash: Sets the actions_after_crash parameter
func (vM) SetActionsAfterCrash(session *Session, self VMRef, value OnCrashBehaviour) (err error) {
	method := "VM.set_actions_after_crash"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumOnCrashBehaviour(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// AsyncSetActionsAfterCrash: Sets the actions_after_crash parameter
func (vM) AsyncSetActionsAfterCrash(session *Session, self VMRef, value OnCrashBehaviour) (retval TaskRef, err error) {
	method := "Async.VM.set_actions_after_crash"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumOnCrashBehaviour(fmt.Sprintf("%s(%s)", method, "value"), value)
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

// SetDomainType: Set the VM.domain_type field of the given VM, which will take effect when it is next started
func (vM) SetDomainType(session *Session, self VMRef, value DomainType) (err error) {
	method := "VM.set_domain_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumDomainType(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetHVMBootPolicy: Set the VM.HVM_boot_policy field of the given VM, which will take effect when it is next started
func (vM) SetHVMBootPolicy(session *Session, self VMRef, value string) (err error) {
	method := "VM.set_HVM_boot_policy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RestartDeviceModels: Restart device models of the VM
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
func (vM) RestartDeviceModels(session *Session, self VMRef) (err error) {
	method := "VM.restart_device_models"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncRestartDeviceModels: Restart device models of the VM
//
// Errors:
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
// OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
// VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
// VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running. The parameters returned are the VM&apos;s handle, and the expected and actual VM state at the time of the call.
func (vM) AsyncRestartDeviceModels(session *Session, self VMRef) (retval TaskRef, err error) {
	method := "Async.VM.restart_device_models"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the VMs known to the system.
func (vM) GetAll(session *Session) (retval []VMRef, err error) {
	method := "VM.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VM references to VM records for all VMs known to the system.
func (vM) GetAllRecords(session *Session) (retval map[VMRef]VMRecord, err error) {
	method := "VM.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVMRefToVMRecordMap(method+" -> ", result)
	return
}

