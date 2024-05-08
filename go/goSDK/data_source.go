package xenapi

type DataSourceRecord struct {
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// true if the data source is being logged
	Enabled bool
	// true if the data source is enabled by default. Non-default data sources cannot be disabled
	Standard bool
	// the units of the value
	Units string
	// the minimum value of the data source
	Min float64
	// the maximum value of the data source
	Max float64
	// current value of the data source
	Value float64
}

type DataSourceRef string

// Data sources for logging in RRDs
type dataSource struct{}

var DataSource dataSource
