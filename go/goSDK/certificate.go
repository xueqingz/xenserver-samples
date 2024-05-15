package xenapi

import (
	"fmt"
	"time"
)

type CertificateRecord struct {
	// Unique identifier/object reference
	UUID string
	// The name of the certificate, only present on certificates of type &apos;ca&apos;
	Name string
	// The type of the certificate, either &apos;ca&apos;, &apos;host&apos; or &apos;host_internal&apos;
	Type CertificateType
	// The host where the certificate is installed
	Host HostRef
	// Date after which the certificate is valid
	NotBefore time.Time
	// Date before which the certificate is valid
	NotAfter time.Time
	// The certificate&apos;s SHA256 fingerprint / hash
	Fingerprint string
}

type CertificateRef string

// An X509 certificate used for TLS connections
type certificate struct{}

var Certificate certificate

// GetAllRecords: Return a map of Certificate references to Certificate records for all Certificates known to the system.
// Version: stockholm
func (certificate) GetAllRecords(session *Session) (retval map[CertificateRef]CertificateRecord, err error) {
	method := "Certificate.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeCertificateRefToCertificateRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of Certificate references to Certificate records for all Certificates known to the system.
// Version: stockholm
func (certificate) GetAllRecords1(session *Session) (retval map[CertificateRef]CertificateRecord, err error) {
	method := "Certificate.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeCertificateRefToCertificateRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the Certificates known to the system.
// Version: stockholm
func (certificate) GetAll(session *Session) (retval []CertificateRef, err error) {
	method := "Certificate.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeCertificateRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the Certificates known to the system.
// Version: stockholm
func (certificate) GetAll1(session *Session) (retval []CertificateRef, err error) {
	method := "Certificate.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeCertificateRefSet(method+" -> ", result)
	return
}

// GetFingerprint: Get the fingerprint field of the given Certificate.
// Version: stockholm
func (certificate) GetFingerprint(session *Session, self CertificateRef) (retval string, err error) {
	method := "Certificate.get_fingerprint"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetFingerprint2: Get the fingerprint field of the given Certificate.
// Version: stockholm
func (certificate) GetFingerprint2(session *Session, self CertificateRef) (retval string, err error) {
	method := "Certificate.get_fingerprint"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNotAfter: Get the not_after field of the given Certificate.
// Version: stockholm
func (certificate) GetNotAfter(session *Session, self CertificateRef) (retval time.Time, err error) {
	method := "Certificate.get_not_after"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNotAfter2: Get the not_after field of the given Certificate.
// Version: stockholm
func (certificate) GetNotAfter2(session *Session, self CertificateRef) (retval time.Time, err error) {
	method := "Certificate.get_not_after"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNotBefore: Get the not_before field of the given Certificate.
// Version: stockholm
func (certificate) GetNotBefore(session *Session, self CertificateRef) (retval time.Time, err error) {
	method := "Certificate.get_not_before"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNotBefore2: Get the not_before field of the given Certificate.
// Version: stockholm
func (certificate) GetNotBefore2(session *Session, self CertificateRef) (retval time.Time, err error) {
	method := "Certificate.get_not_before"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost: Get the host field of the given Certificate.
// Version: stockholm
func (certificate) GetHost(session *Session, self CertificateRef) (retval HostRef, err error) {
	method := "Certificate.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHost2: Get the host field of the given Certificate.
// Version: stockholm
func (certificate) GetHost2(session *Session, self CertificateRef) (retval HostRef, err error) {
	method := "Certificate.get_host"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetType: Get the type field of the given Certificate.
// Version: stockholm
func (certificate) GetType(session *Session, self CertificateRef) (retval CertificateType, err error) {
	method := "Certificate.get_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumCertificateType(method+" -> ", result)
	return
}

// GetType2: Get the type field of the given Certificate.
// Version: stockholm
func (certificate) GetType2(session *Session, self CertificateRef) (retval CertificateType, err error) {
	method := "Certificate.get_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumCertificateType(method+" -> ", result)
	return
}

// GetName: Get the name field of the given Certificate.
// Version: stockholm
func (certificate) GetName(session *Session, self CertificateRef) (retval string, err error) {
	method := "Certificate.get_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetName2: Get the name field of the given Certificate.
// Version: stockholm
func (certificate) GetName2(session *Session, self CertificateRef) (retval string, err error) {
	method := "Certificate.get_name"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given Certificate.
// Version: stockholm
func (certificate) GetUUID(session *Session, self CertificateRef) (retval string, err error) {
	method := "Certificate.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given Certificate.
// Version: stockholm
func (certificate) GetUUID2(session *Session, self CertificateRef) (retval string, err error) {
	method := "Certificate.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByUUID: Get a reference to the Certificate instance with the specified UUID.
// Version: stockholm
func (certificate) GetByUUID(session *Session, uUID string) (retval CertificateRef, err error) {
	method := "Certificate.get_by_uuid"
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
	retval, err = deserializeCertificateRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the Certificate instance with the specified UUID.
// Version: stockholm
func (certificate) GetByUUID2(session *Session, uUID string) (retval CertificateRef, err error) {
	method := "Certificate.get_by_uuid"
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
	retval, err = deserializeCertificateRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given Certificate.
// Version: stockholm
func (certificate) GetRecord(session *Session, self CertificateRef) (retval CertificateRecord, err error) {
	method := "Certificate.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeCertificateRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given Certificate.
// Version: stockholm
func (certificate) GetRecord2(session *Session, self CertificateRef) (retval CertificateRecord, err error) {
	method := "Certificate.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeCertificateRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeCertificateRecord(method+" -> ", result)
	return
}
