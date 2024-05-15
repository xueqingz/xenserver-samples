package xenapi

import (
	"fmt"
)

type PVSSiteRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// Unique identifier of the PVS site, as configured in PVS
	PVSUUID string
	// The SR used by PVS proxy for the cache
	CacheStorage []PVSCacheStorageRef
	// The set of PVS servers in the site
	Servers []PVSServerRef
	// The set of proxies associated with the site
	Proxies []PVSProxyRef
}

type PVSSiteRef string

// machines serving blocks of data for provisioning VMs
type pVSSite struct{}

var PVSSite pVSSite

// GetAllRecords: Return a map of PVS_site references to PVS_site records for all PVS_sites known to the system.
// Version: ely
func (pVSSite) GetAllRecords(session *Session) (retval map[PVSSiteRef]PVSSiteRecord, err error) {
	method := "PVS_site.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSSiteRefToPVSSiteRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of PVS_site references to PVS_site records for all PVS_sites known to the system.
// Version: ely
func (pVSSite) GetAllRecords1(session *Session) (retval map[PVSSiteRef]PVSSiteRecord, err error) {
	method := "PVS_site.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSSiteRefToPVSSiteRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the PVS_sites known to the system.
// Version: ely
func (pVSSite) GetAll(session *Session) (retval []PVSSiteRef, err error) {
	method := "PVS_site.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSSiteRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the PVS_sites known to the system.
// Version: ely
func (pVSSite) GetAll1(session *Session) (retval []PVSSiteRef, err error) {
	method := "PVS_site.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSSiteRefSet(method+" -> ", result)
	return
}

// SetPVSUUID: Update the PVS UUID of the PVS site
// Version: ely
func (pVSSite) SetPVSUUID(session *Session, self PVSSiteRef, value string) (err error) {
	method := "PVS_site.set_PVS_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetPVSUUID: Update the PVS UUID of the PVS site
// Version: ely
func (pVSSite) AsyncSetPVSUUID(session *Session, self PVSSiteRef, value string) (retval TaskRef, err error) {
	method := "Async.PVS_site.set_PVS_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetPVSUUID3: Update the PVS UUID of the PVS site
// Version: ely
func (pVSSite) SetPVSUUID3(session *Session, self PVSSiteRef, value string) (err error) {
	method := "PVS_site.set_PVS_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetPVSUUID3: Update the PVS UUID of the PVS site
// Version: ely
func (pVSSite) AsyncSetPVSUUID3(session *Session, self PVSSiteRef, value string) (retval TaskRef, err error) {
	method := "Async.PVS_site.set_PVS_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Forget: Remove a site&apos;s meta data
// Version: ely
//
// Errors:
// PVS_SITE_CONTAINS_RUNNING_PROXIES - The PVS site contains running proxies.
// PVS_SITE_CONTAINS_SERVERS - The PVS site contains servers and cannot be forgotten.
func (pVSSite) Forget(session *Session, self PVSSiteRef) (err error) {
	method := "PVS_site.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncForget: Remove a site&apos;s meta data
// Version: ely
//
// Errors:
// PVS_SITE_CONTAINS_RUNNING_PROXIES - The PVS site contains running proxies.
// PVS_SITE_CONTAINS_SERVERS - The PVS site contains servers and cannot be forgotten.
func (pVSSite) AsyncForget(session *Session, self PVSSiteRef) (retval TaskRef, err error) {
	method := "Async.PVS_site.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Forget2: Remove a site&apos;s meta data
// Version: ely
//
// Errors:
// PVS_SITE_CONTAINS_RUNNING_PROXIES - The PVS site contains running proxies.
// PVS_SITE_CONTAINS_SERVERS - The PVS site contains servers and cannot be forgotten.
func (pVSSite) Forget2(session *Session, self PVSSiteRef) (err error) {
	method := "PVS_site.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncForget2: Remove a site&apos;s meta data
// Version: ely
//
// Errors:
// PVS_SITE_CONTAINS_RUNNING_PROXIES - The PVS site contains running proxies.
// PVS_SITE_CONTAINS_SERVERS - The PVS site contains servers and cannot be forgotten.
func (pVSSite) AsyncForget2(session *Session, self PVSSiteRef) (retval TaskRef, err error) {
	method := "Async.PVS_site.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Introduce: Introduce new PVS site
// Version: ely
func (pVSSite) Introduce(session *Session, nameLabel string, nameDescription string, pVSUUID string) (retval PVSSiteRef, err error) {
	method := "PVS_site.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	pVSUUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "PVS_uuid"), pVSUUID)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, pVSUUIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSSiteRef(method+" -> ", result)
	return
}

// AsyncIntroduce: Introduce new PVS site
// Version: ely
func (pVSSite) AsyncIntroduce(session *Session, nameLabel string, nameDescription string, pVSUUID string) (retval TaskRef, err error) {
	method := "Async.PVS_site.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	pVSUUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "PVS_uuid"), pVSUUID)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, pVSUUIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Introduce4: Introduce new PVS site
// Version: ely
func (pVSSite) Introduce4(session *Session, nameLabel string, nameDescription string, pVSUUID string) (retval PVSSiteRef, err error) {
	method := "PVS_site.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	pVSUUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "PVS_uuid"), pVSUUID)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, pVSUUIDArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSSiteRef(method+" -> ", result)
	return
}

// AsyncIntroduce4: Introduce new PVS site
// Version: ely
func (pVSSite) AsyncIntroduce4(session *Session, nameLabel string, nameDescription string, pVSUUID string) (retval TaskRef, err error) {
	method := "Async.PVS_site.introduce"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameLabelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_label"), nameLabel)
	if err != nil {
		return
	}
	nameDescriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name_description"), nameDescription)
	if err != nil {
		return
	}
	pVSUUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "PVS_uuid"), pVSUUID)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, pVSUUIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetNameDescription: Set the name/description field of the given PVS_site.
// Version: ely
func (pVSSite) SetNameDescription(session *Session, self PVSSiteRef, value string) (err error) {
	method := "PVS_site.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription3: Set the name/description field of the given PVS_site.
// Version: ely
func (pVSSite) SetNameDescription3(session *Session, self PVSSiteRef, value string) (err error) {
	method := "PVS_site.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel: Set the name/label field of the given PVS_site.
// Version: ely
func (pVSSite) SetNameLabel(session *Session, self PVSSiteRef, value string) (err error) {
	method := "PVS_site.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel3: Set the name/label field of the given PVS_site.
// Version: ely
func (pVSSite) SetNameLabel3(session *Session, self PVSSiteRef, value string) (err error) {
	method := "PVS_site.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetProxies: Get the proxies field of the given PVS_site.
// Version: ely
func (pVSSite) GetProxies(session *Session, self PVSSiteRef) (retval []PVSProxyRef, err error) {
	method := "PVS_site.get_proxies"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSProxyRefSet(method+" -> ", result)
	return
}

// GetProxies2: Get the proxies field of the given PVS_site.
// Version: ely
func (pVSSite) GetProxies2(session *Session, self PVSSiteRef) (retval []PVSProxyRef, err error) {
	method := "PVS_site.get_proxies"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSProxyRefSet(method+" -> ", result)
	return
}

// GetServers: Get the servers field of the given PVS_site.
// Version: ely
func (pVSSite) GetServers(session *Session, self PVSSiteRef) (retval []PVSServerRef, err error) {
	method := "PVS_site.get_servers"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRefSet(method+" -> ", result)
	return
}

// GetServers2: Get the servers field of the given PVS_site.
// Version: ely
func (pVSSite) GetServers2(session *Session, self PVSSiteRef) (retval []PVSServerRef, err error) {
	method := "PVS_site.get_servers"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSServerRefSet(method+" -> ", result)
	return
}

// GetCacheStorage: Get the cache_storage field of the given PVS_site.
// Version: ely
func (pVSSite) GetCacheStorage(session *Session, self PVSSiteRef) (retval []PVSCacheStorageRef, err error) {
	method := "PVS_site.get_cache_storage"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSCacheStorageRefSet(method+" -> ", result)
	return
}

// GetCacheStorage2: Get the cache_storage field of the given PVS_site.
// Version: ely
func (pVSSite) GetCacheStorage2(session *Session, self PVSSiteRef) (retval []PVSCacheStorageRef, err error) {
	method := "PVS_site.get_cache_storage"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSCacheStorageRefSet(method+" -> ", result)
	return
}

// GetPVSUUID: Get the PVS_uuid field of the given PVS_site.
// Version: ely
func (pVSSite) GetPVSUUID(session *Session, self PVSSiteRef) (retval string, err error) {
	method := "PVS_site.get_PVS_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPVSUUID2: Get the PVS_uuid field of the given PVS_site.
// Version: ely
func (pVSSite) GetPVSUUID2(session *Session, self PVSSiteRef) (retval string, err error) {
	method := "PVS_site.get_PVS_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given PVS_site.
// Version: ely
func (pVSSite) GetNameDescription(session *Session, self PVSSiteRef) (retval string, err error) {
	method := "PVS_site.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription2: Get the name/description field of the given PVS_site.
// Version: ely
func (pVSSite) GetNameDescription2(session *Session, self PVSSiteRef) (retval string, err error) {
	method := "PVS_site.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given PVS_site.
// Version: ely
func (pVSSite) GetNameLabel(session *Session, self PVSSiteRef) (retval string, err error) {
	method := "PVS_site.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel2: Get the name/label field of the given PVS_site.
// Version: ely
func (pVSSite) GetNameLabel2(session *Session, self PVSSiteRef) (retval string, err error) {
	method := "PVS_site.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given PVS_site.
// Version: ely
func (pVSSite) GetUUID(session *Session, self PVSSiteRef) (retval string, err error) {
	method := "PVS_site.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given PVS_site.
// Version: ely
func (pVSSite) GetUUID2(session *Session, self PVSSiteRef) (retval string, err error) {
	method := "PVS_site.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByNameLabel: Get all the PVS_site instances with the given label.
// Version: ely
func (pVSSite) GetByNameLabel(session *Session, label string) (retval []PVSSiteRef, err error) {
	method := "PVS_site.get_by_name_label"
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
	retval, err = deserializePVSSiteRefSet(method+" -> ", result)
	return
}

// GetByNameLabel2: Get all the PVS_site instances with the given label.
// Version: ely
func (pVSSite) GetByNameLabel2(session *Session, label string) (retval []PVSSiteRef, err error) {
	method := "PVS_site.get_by_name_label"
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
	retval, err = deserializePVSSiteRefSet(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the PVS_site instance with the specified UUID.
// Version: ely
func (pVSSite) GetByUUID(session *Session, uUID string) (retval PVSSiteRef, err error) {
	method := "PVS_site.get_by_uuid"
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
	retval, err = deserializePVSSiteRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the PVS_site instance with the specified UUID.
// Version: ely
func (pVSSite) GetByUUID2(session *Session, uUID string) (retval PVSSiteRef, err error) {
	method := "PVS_site.get_by_uuid"
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
	retval, err = deserializePVSSiteRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given PVS_site.
// Version: ely
func (pVSSite) GetRecord(session *Session, self PVSSiteRef) (retval PVSSiteRecord, err error) {
	method := "PVS_site.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSSiteRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given PVS_site.
// Version: ely
func (pVSSite) GetRecord2(session *Session, self PVSSiteRef) (retval PVSSiteRecord, err error) {
	method := "PVS_site.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializePVSSiteRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializePVSSiteRecord(method+" -> ", result)
	return
}
