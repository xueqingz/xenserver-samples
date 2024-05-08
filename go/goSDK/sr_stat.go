package xenapi

type OptionString *string

type SrStatRecord struct {
	// Uuid that uniquely identifies this SR, if one is available.
	UUID OptionString
	// Short, human-readable label for the SR.
	NameLabel string
	// Longer, human-readable description of the SR. Descriptions are generally only displayed by clients when the user is examining SRs in detail.
	NameDescription string
	// Number of bytes free on the backing storage (in bytes)
	FreeSpace int
	// Total physical size of the backing storage (in bytes)
	TotalSpace int
	// Indicates whether the SR uses clustered local storage.
	Clustered bool
	// The health status of the SR.
	Health SrHealth
}

type SrStatRef string

// A set of high-level properties associated with an SR.
type srStat struct{}

var SrStat srStat
