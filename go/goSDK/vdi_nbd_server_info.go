package xenapi

type VdiNbdServerInfoRecord struct {
	// The exportname to request over NBD. This holds details including an authentication token, so it must be protected appropriately. Clients should regard the exportname as an opaque string or token.
	Exportname string
	// An address on which the server can be reached; this can be IPv4, IPv6, or a DNS name.
	Address string
	// The TCP port
	Port int
	// The TLS certificate of the server
	Cert string
	// For convenience, this redundant field holds a DNS (hostname) subject of the certificate. This can be a wildcard, but only for a certificate that has a wildcard subject and no concrete hostname subjects.
	Subject string
}

type VdiNbdServerInfoRef string

// Details for connecting to a VDI using the Network Block Device protocol
type vdiNbdServerInfo struct{}

var VdiNbdServerInfo vdiNbdServerInfo
