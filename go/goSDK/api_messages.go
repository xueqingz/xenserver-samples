package xenapi

const (
	//
	MessageXapiStartupBlockedAsVersionHigherThanCoordinator = "XAPI_STARTUP_BLOCKED_AS_VERSION_HIGHER_THAN_COORDINATOR"
	//
	MessagePeriodicUpdateSyncFailed = "PERIODIC_UPDATE_SYNC_FAILED"
	//
	MessageTLSVerificationEmergencyDisabled = "TLS_VERIFICATION_EMERGENCY_DISABLED"
	//
	MessageFailedLoginAttempts = "FAILED_LOGIN_ATTEMPTS"
	//
	MessageHostInternalCertificateExpiring07 = "HOST_INTERNAL_CERTIFICATE_EXPIRING_07"
	//
	MessageHostInternalCertificateExpiring14 = "HOST_INTERNAL_CERTIFICATE_EXPIRING_14"
	//
	MessageHostInternalCertificateExpiring30 = "HOST_INTERNAL_CERTIFICATE_EXPIRING_30"
	//
	MessagePoolCaCertificateExpiring07 = "POOL_CA_CERTIFICATE_EXPIRING_07"
	//
	MessagePoolCaCertificateExpiring14 = "POOL_CA_CERTIFICATE_EXPIRING_14"
	//
	MessagePoolCaCertificateExpiring30 = "POOL_CA_CERTIFICATE_EXPIRING_30"
	//
	MessageHostServerCertificateExpiring07 = "HOST_SERVER_CERTIFICATE_EXPIRING_07"
	//
	MessageHostServerCertificateExpiring14 = "HOST_SERVER_CERTIFICATE_EXPIRING_14"
	//
	MessageHostServerCertificateExpiring30 = "HOST_SERVER_CERTIFICATE_EXPIRING_30"
	//
	MessagePoolCaCertificateExpired = "POOL_CA_CERTIFICATE_EXPIRED"
	//
	MessageHostInternalCertificateExpired = "HOST_INTERNAL_CERTIFICATE_EXPIRED"
	//
	MessageHostServerCertificateExpired = "HOST_SERVER_CERTIFICATE_EXPIRED"
	//
	MessageClusterHostJoining = "CLUSTER_HOST_JOINING"
	//
	MessageClusterHostLeaving = "CLUSTER_HOST_LEAVING"
	//
	MessageClusterHostFencing = "CLUSTER_HOST_FENCING"
	//
	MessageClusterHostEnableFailed = "CLUSTER_HOST_ENABLE_FAILED"
	//
	MessageClusterQuorumApproachingLost = "CLUSTER_QUORUM_APPROACHING_LOST"
	//
	MessagePoolCPUFeaturesUp = "POOL_CPU_FEATURES_UP"
	//
	MessagePoolCPUFeaturesDown = "POOL_CPU_FEATURES_DOWN"
	//
	MessageHostCPUFeaturesUp = "HOST_CPU_FEATURES_UP"
	//
	MessageHostCPUFeaturesDown = "HOST_CPU_FEATURES_DOWN"
	//
	MessageBondStatusChanged = "BOND_STATUS_CHANGED"
	//
	MessageVdiCbtResizeFailed = "VDI_CBT_RESIZE_FAILED"
	//
	MessageVdiCbtSnapshotFailed = "VDI_CBT_SNAPSHOT_FAILED"
	//
	MessageVdiCbtMetadataInconsistent = "VDI_CBT_METADATA_INCONSISTENT"
	//
	MessageVmssSnapshotMissedEvent = "VMSS_SNAPSHOT_MISSED_EVENT"
	//
	MessageVmssXapiLogonFailure = "VMSS_XAPI_LOGON_FAILURE"
	//
	MessageVmssLicenseError = "VMSS_LICENSE_ERROR"
	//
	MessageVmssSnapshotFailed = "VMSS_SNAPSHOT_FAILED"
	//
	MessageVmssSnapshotSucceeded = "VMSS_SNAPSHOT_SUCCEEDED"
	//
	MessageVmssSnapshotLockFailed = "VMSS_SNAPSHOT_LOCK_FAILED"
	//
	MessageVmppSnapshotArchiveAlreadyExists = "VMPP_SNAPSHOT_ARCHIVE_ALREADY_EXISTS"
	//
	MessageVmppArchiveMissedEvent = "VMPP_ARCHIVE_MISSED_EVENT"
	//
	MessageVmppSnapshotMissedEvent = "VMPP_SNAPSHOT_MISSED_EVENT"
	//
	MessageVmppXapiLogonFailure = "VMPP_XAPI_LOGON_FAILURE"
	//
	MessageVmppLicenseError = "VMPP_LICENSE_ERROR"
	//
	MessageVmppArchiveTargetUnmountFailed = "VMPP_ARCHIVE_TARGET_UNMOUNT_FAILED"
	//
	MessageVmppArchiveTargetMountFailed = "VMPP_ARCHIVE_TARGET_MOUNT_FAILED"
	//
	MessageVmppArchiveSucceeded = "VMPP_ARCHIVE_SUCCEEDED"
	//
	MessageVmppArchiveFailed0 = "VMPP_ARCHIVE_FAILED_0"
	//
	MessageVmppArchiveLockFailed = "VMPP_ARCHIVE_LOCK_FAILED"
	//
	MessageVmppSnapshotFailed = "VMPP_SNAPSHOT_FAILED"
	//
	MessageVmppSnapshotSucceeded = "VMPP_SNAPSHOT_SUCCEEDED"
	//
	MessageVmppSnapshotLockFailed = "VMPP_SNAPSHOT_LOCK_FAILED"
	//
	MessagePvsProxySrOutOfSpace = "PVS_PROXY_SR_OUT_OF_SPACE"
	//
	MessagePvsProxyNoServerAvailable = "PVS_PROXY_NO_SERVER_AVAILABLE"
	//
	MessagePvsProxySetupFailed = "PVS_PROXY_SETUP_FAILED"
	//
	MessagePvsProxyNoCacheSrAvailable = "PVS_PROXY_NO_CACHE_SR_AVAILABLE"
	//
	MessageLicenseServerVersionObsolete = "LICENSE_SERVER_VERSION_OBSOLETE"
	//
	MessageLicenseServerUnreachable = "LICENSE_SERVER_UNREACHABLE"
	//
	MessageLicenseNotAvailable = "LICENSE_NOT_AVAILABLE"
	//
	MessageGraceLicense = "GRACE_LICENSE"
	//
	MessageLicenseServerUnavailable = "LICENSE_SERVER_UNAVAILABLE"
	//
	MessageLicenseServerConnected = "LICENSE_SERVER_CONNECTED"
	//
	MessageLicenseExpired = "LICENSE_EXPIRED"
	//
	MessageLicenseExpiresSoon = "LICENSE_EXPIRES_SOON"
	//
	MessageLicenseDoesNotSupportPooling = "LICENSE_DOES_NOT_SUPPORT_POOLING"
	//
	MessageMultipathPeriodicAlert = "MULTIPATH_PERIODIC_ALERT"
	//
	MessageExtauthInPoolIsNonHomogeneous = "EXTAUTH_IN_POOL_IS_NON_HOMOGENEOUS"
	//
	MessageExtauthInitInHostFailed = "EXTAUTH_INIT_IN_HOST_FAILED"
	//
	MessageWlbOptimizationAlert = "WLB_OPTIMIZATION_ALERT"
	//
	MessageWlbConsultationFailed = "WLB_CONSULTATION_FAILED"
	//
	MessageAlarm = "ALARM"
	//
	MessagePbdPlugFailedOnServerStart = "PBD_PLUG_FAILED_ON_SERVER_START"
	//
	MessagePoolMasterTransition = "POOL_MASTER_TRANSITION"
	//
	MessageHostClockWentBackwards = "HOST_CLOCK_WENT_BACKWARDS"
	//
	MessageHostClockSkewDetected = "HOST_CLOCK_SKEW_DETECTED"
	//
	MessageHostSyncDataFailed = "HOST_SYNC_DATA_FAILED"
	//
	MessageVMSecureBootFailed = "VM_SECURE_BOOT_FAILED"
	//
	MessageVMCloned = "VM_CLONED"
	//
	MessageVMCrashed = "VM_CRASHED"
	//
	MessageVMUnpaused = "VM_UNPAUSED"
	//
	MessageVMPaused = "VM_PAUSED"
	//
	MessageVMResumed = "VM_RESUMED"
	//
	MessageVMSuspended = "VM_SUSPENDED"
	//
	MessageVMCheckpointed = "VM_CHECKPOINTED"
	//
	MessageVMSnapshotReverted = "VM_SNAPSHOT_REVERTED"
	//
	MessageVMSnapshotted = "VM_SNAPSHOTTED"
	//
	MessageVMMigrated = "VM_MIGRATED"
	//
	MessageVMRebooted = "VM_REBOOTED"
	//
	MessageVMShutdown = "VM_SHUTDOWN"
	//
	MessageVMStarted = "VM_STARTED"
	//
	MessageVcpuQosFailed = "VCPU_QOS_FAILED"
	//
	MessageVbdQosFailed = "VBD_QOS_FAILED"
	//
	MessageVifQosFailed = "VIF_QOS_FAILED"
	//
	MessageIPConfiguredPifCanUnplug = "IP_CONFIGURED_PIF_CAN_UNPLUG"
	//
	MessageMetadataLunBroken = "METADATA_LUN_BROKEN"
	//
	MessageMetadataLunHealthy = "METADATA_LUN_HEALTHY"
	//
	MessageHaHostWasFenced = "HA_HOST_WAS_FENCED"
	//
	MessageHaHostFailed = "HA_HOST_FAILED"
	//
	MessageHaProtectedVMRestartFailed = "HA_PROTECTED_VM_RESTART_FAILED"
	//
	MessageHaPoolDropInPlanExistsFor = "HA_POOL_DROP_IN_PLAN_EXISTS_FOR"
	//
	MessageHaPoolOvercommitted = "HA_POOL_OVERCOMMITTED"
	//
	MessageHaNetworkBondingError = "HA_NETWORK_BONDING_ERROR"
	//
	MessageHaXapiHealthcheckApproachingTimeout = "HA_XAPI_HEALTHCHECK_APPROACHING_TIMEOUT"
	//
	MessageHaStatefileApproachingTimeout = "HA_STATEFILE_APPROACHING_TIMEOUT"
	//
	MessageHaHeartbeatApproachingTimeout = "HA_HEARTBEAT_APPROACHING_TIMEOUT"
	//
	MessageHaStatefileLost = "HA_STATEFILE_LOST"
)
