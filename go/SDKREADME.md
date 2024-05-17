# Go

The CitrixHypervisor-SDK directory contains the following folders that are relevant to Go programmers:

- `XenServerGo`
  The Citrix Hypervisor SDK for Go

  - `XenServerGo/src`
    The Go source files which can be used as the local module in a Go project. Every API object is associated with one Go file. For example the functions implementing the VM operations are contained within the file vm.go.

# Go SDK dependencies

The Go SDK supports the JSON-RPC v2.0 protocol. The Go SDK is backwards compatible and can be used to communicate with hosts running an older version of Citrix Hypervisor or XenServer. 

///
NOTE:

This module is backwards compatible by using different method names for a same API object method on different XenServer versions. To communicate with hosts running different versions of Citrix Hypervisor or XenServer, it is advisable to check both the method and XAPI version first. Like using "Host.Evacuate3" on version "1.297.0" and "Host.Evacuate4" on version "23.27.0".
///

## Platform supported:

- Linux
- Windows

## Library:

- None

## Dependencies:

- This library requires Go 1.22 or greater.

## Examples:

