package xenapi

type OptionSrStatRecord *SrStatRecord

type ProbeResultRecord struct {
	// Plugin-specific configuration which describes where and how to locate the storage repository. This may include the physical block device name, a remote NFS server and path or an RBD storage pool.
	Configuration map[string]string
	// True if this configuration is complete and can be used to call SR.create. False if it requires further iterative calls to SR.probe, to potentially narrow down on a configuration that can be used.
	Complete bool
	// Existing SR found for this configuration
	Sr OptionSrStatRecord
	// Additional plugin-specific information about this configuration, that might be of use for an API user. This can for example include the LUN or the WWPN.
	ExtraInfo map[string]string
}

type ProbeResultRef string

// A set of properties that describe one result element of SR.probe. Result elements and properties can change dynamically based on changes to the the SR.probe input-parameters or the target.
type probeResult struct{}

var ProbeResult probeResult
