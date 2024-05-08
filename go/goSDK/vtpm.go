package xenapi

import (
	"fmt"
)

type VTPMRecord struct {
	// Unique identifier/object reference
	UUID string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VtpmOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VtpmOperations
	// The virtual machine the TPM is attached to
	VM VMRef
	// The domain where the backend is located (unused)
	Backend VMRef
	// The backend where the vTPM is persisted
	PersistenceBackend PersistenceBackend
	// Whether the contents are never copied, satisfying the TPM spec
	IsUnique bool
	// Whether the contents of the VTPM are secured according to the TPM spec
	IsProtected bool
}

type VTPMRef string

// A virtual TPM device
type vTPM struct{}

var VTPM vTPM

// GetRecord: Get a record containing the current state of the given VTPM.
func (vTPM) GetRecord(session *Session, self VTPMRef) (retval VTPMRecord, err error) {
	method := "VTPM.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeVTPMRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the VTPM instance with the specified UUID.
func (vTPM) GetByUUID(session *Session, uUID string) (retval VTPMRef, err error) {
	method := "VTPM.get_by_uuid"
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
	retval, err = deserializeVTPMRef(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given VTPM.
func (vTPM) GetUUID(session *Session, self VTPMRef) (retval string, err error) {
	method := "VTPM.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given VTPM.
func (vTPM) GetAllowedOperations(session *Session, self VTPMRef) (retval []VtpmOperations, err error) {
	method := "VTPM.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumVtpmOperationsSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given VTPM.
func (vTPM) GetCurrentOperations(session *Session, self VTPMRef) (retval map[string]VtpmOperations, err error) {
	method := "VTPM.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumVtpmOperationsMap(method+" -> ", result)
	return
}

// GetVM: Get the VM field of the given VTPM.
func (vTPM) GetVM(session *Session, self VTPMRef) (retval VMRef, err error) {
	method := "VTPM.get_VM"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBackend: Get the backend field of the given VTPM.
func (vTPM) GetBackend(session *Session, self VTPMRef) (retval VMRef, err error) {
	method := "VTPM.get_backend"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetPersistenceBackend: Get the persistence_backend field of the given VTPM.
func (vTPM) GetPersistenceBackend(session *Session, self VTPMRef) (retval PersistenceBackend, err error) {
	method := "VTPM.get_persistence_backend"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumPersistenceBackend(method+" -> ", result)
	return
}

// GetIsUnique: Get the is_unique field of the given VTPM.
func (vTPM) GetIsUnique(session *Session, self VTPMRef) (retval bool, err error) {
	method := "VTPM.get_is_unique"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetIsProtected: Get the is_protected field of the given VTPM.
func (vTPM) GetIsProtected(session *Session, self VTPMRef) (retval bool, err error) {
	method := "VTPM.get_is_protected"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Create a new VTPM instance, and return its handle.
func (vTPM) Create(session *Session, vM VMRef, isUnique bool) (retval VTPMRef, err error) {
	method := "VTPM.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vM"), vM)
	if err != nil {
		return
	}
	isUniqueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "is_unique"), isUnique)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, isUniqueArg)
	if err != nil {
		return
	}
	retval, err = deserializeVTPMRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new VTPM instance, and return its handle.
func (vTPM) AsyncCreate(session *Session, vM VMRef, isUnique bool) (retval TaskRef, err error) {
	method := "Async.VTPM.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	vMArg, err := serializeVMRef(fmt.Sprintf("%s(%s)", method, "vM"), vM)
	if err != nil {
		return
	}
	isUniqueArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "is_unique"), isUnique)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, vMArg, isUniqueArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Destroy the specified VTPM instance, along with its state.
func (vTPM) Destroy(session *Session, self VTPMRef) (err error) {
	method := "VTPM.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified VTPM instance, along with its state.
func (vTPM) AsyncDestroy(session *Session, self VTPMRef) (retval TaskRef, err error) {
	method := "Async.VTPM.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeVTPMRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAll: Return a list of all the VTPMs known to the system.
func (vTPM) GetAll(session *Session) (retval []VTPMRef, err error) {
	method := "VTPM.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVTPMRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of VTPM references to VTPM records for all VTPMs known to the system.
func (vTPM) GetAllRecords(session *Session) (retval map[VTPMRef]VTPMRecord, err error) {
	method := "VTPM.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeVTPMRefToVTPMRecordMap(method+" -> ", result)
	return
}

