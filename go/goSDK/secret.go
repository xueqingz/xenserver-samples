package xenapi

import (
	"fmt"
)

type SecretRecord struct {
	// Unique identifier/object reference
	UUID string
	// the secret
	Value string
	// other_config
	OtherConfig map[string]string
}

type SecretRef string

// A secret
type secret struct{}

var Secret secret

// GetAllRecords: Return a map of secret references to secret records for all secrets known to the system.
// Version: midnight-ride
func (secret) GetAllRecords(session *Session) (retval map[SecretRef]SecretRecord, err error) {
	method := "secret.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRefToSecretRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of secret references to secret records for all secrets known to the system.
// Version: midnight-ride
func (secret) GetAllRecords1(session *Session) (retval map[SecretRef]SecretRecord, err error) {
	method := "secret.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRefToSecretRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the secrets known to the system.
// Version: midnight-ride
func (secret) GetAll(session *Session) (retval []SecretRef, err error) {
	method := "secret.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the secrets known to the system.
// Version: midnight-ride
func (secret) GetAll1(session *Session) (retval []SecretRef, err error) {
	method := "secret.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRefSet(method+" -> ", result)
	return
}

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given secret.  If the key is not in that Map, then do nothing.
// Version: midnight-ride
func (secret) RemoveFromOtherConfig(session *Session, self SecretRef, key string) (err error) {
	method := "secret.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig3: Remove the given key and its corresponding value from the other_config field of the given secret.  If the key is not in that Map, then do nothing.
// Version: midnight-ride
func (secret) RemoveFromOtherConfig3(session *Session, self SecretRef, key string) (err error) {
	method := "secret.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig2: Remove the given key and its corresponding value from the other_config field of the given secret.  If the key is not in that Map, then do nothing.
// Version: rio
func (secret) RemoveFromOtherConfig2(session *Session, key string) (err error) {
	method := "secret.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	keyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "key"), key)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, keyArg)
	return
}

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given secret.
// Version: midnight-ride
func (secret) AddToOtherConfig(session *Session, self SecretRef, key string, value string) (err error) {
	method := "secret.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig4: Add the given key-value pair to the other_config field of the given secret.
// Version: midnight-ride
func (secret) AddToOtherConfig4(session *Session, self SecretRef, key string, value string) (err error) {
	method := "secret.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig3: Add the given key-value pair to the other_config field of the given secret.
// Version: rio
func (secret) AddToOtherConfig3(session *Session, key string, value string) (err error) {
	method := "secret.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
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
	_, err = session.client.sendCall(method, sessionIDArg, keyArg, valueArg)
	return
}

// SetOtherConfig: Set the other_config field of the given secret.
// Version: midnight-ride
func (secret) SetOtherConfig(session *Session, self SecretRef, value map[string]string) (err error) {
	method := "secret.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig3: Set the other_config field of the given secret.
// Version: midnight-ride
func (secret) SetOtherConfig3(session *Session, self SecretRef, value map[string]string) (err error) {
	method := "secret.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig2: Set the other_config field of the given secret.
// Version: rio
func (secret) SetOtherConfig2(session *Session, value map[string]string) (err error) {
	method := "secret.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	valueArg, err := serializeStringToStringMap(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, valueArg)
	return
}

// SetValue: Set the value field of the given secret.
// Version: midnight-ride
func (secret) SetValue(session *Session, self SecretRef, value string) (err error) {
	method := "secret.set_value"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetValue3: Set the value field of the given secret.
// Version: midnight-ride
func (secret) SetValue3(session *Session, self SecretRef, value string) (err error) {
	method := "secret.set_value"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetValue2: Set the value field of the given secret.
// Version: rio
func (secret) SetValue2(session *Session, value string) (err error) {
	method := "secret.set_value"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	valueArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, valueArg)
	return
}

// GetOtherConfig: Get the other_config field of the given secret.
// Version: midnight-ride
func (secret) GetOtherConfig(session *Session, self SecretRef) (retval map[string]string, err error) {
	method := "secret.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetOtherConfig2: Get the other_config field of the given secret.
// Version: midnight-ride
func (secret) GetOtherConfig2(session *Session, self SecretRef) (retval map[string]string, err error) {
	method := "secret.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetValue: Get the value field of the given secret.
// Version: midnight-ride
func (secret) GetValue(session *Session, self SecretRef) (retval string, err error) {
	method := "secret.get_value"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetValue2: Get the value field of the given secret.
// Version: midnight-ride
func (secret) GetValue2(session *Session, self SecretRef) (retval string, err error) {
	method := "secret.get_value"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given secret.
// Version: midnight-ride
func (secret) GetUUID(session *Session, self SecretRef) (retval string, err error) {
	method := "secret.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given secret.
// Version: midnight-ride
func (secret) GetUUID2(session *Session, self SecretRef) (retval string, err error) {
	method := "secret.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy: Destroy the specified secret instance.
// Version: midnight-ride
func (secret) Destroy(session *Session, self SecretRef) (err error) {
	method := "secret.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy: Destroy the specified secret instance.
// Version: midnight-ride
func (secret) AsyncDestroy(session *Session, self SecretRef) (retval TaskRef, err error) {
	method := "Async.secret.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Destroy2: Destroy the specified secret instance.
// Version: midnight-ride
func (secret) Destroy2(session *Session, self SecretRef) (err error) {
	method := "secret.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncDestroy2: Destroy the specified secret instance.
// Version: midnight-ride
func (secret) AsyncDestroy2(session *Session, self SecretRef) (retval TaskRef, err error) {
	method := "Async.secret.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Create a new secret instance, and return its handle. The constructor args are: value*, other_config (* = non-optional).
// Version: midnight-ride
func (secret) Create(session *Session, args SecretRecord) (retval SecretRef, err error) {
	method := "secret.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeSecretRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRef(method+" -> ", result)
	return
}

// AsyncCreate: Create a new secret instance, and return its handle. The constructor args are: value*, other_config (* = non-optional).
// Version: midnight-ride
func (secret) AsyncCreate(session *Session, args SecretRecord) (retval TaskRef, err error) {
	method := "Async.secret.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeSecretRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// Create2: Create a new secret instance, and return its handle. The constructor args are: value*, other_config (* = non-optional).
// Version: midnight-ride
func (secret) Create2(session *Session, args SecretRecord) (retval SecretRef, err error) {
	method := "secret.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeSecretRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, argsArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRef(method+" -> ", result)
	return
}

// AsyncCreate2: Create a new secret instance, and return its handle. The constructor args are: value*, other_config (* = non-optional).
// Version: midnight-ride
func (secret) AsyncCreate2(session *Session, args SecretRecord) (retval TaskRef, err error) {
	method := "Async.secret.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	argsArg, err := serializeSecretRecord(fmt.Sprintf("%s(%s)", method, "args"), args)
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

// GetByUUID: Get a reference to the secret instance with the specified UUID.
// Version: midnight-ride
func (secret) GetByUUID(session *Session, uUID string) (retval SecretRef, err error) {
	method := "secret.get_by_uuid"
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
	retval, err = deserializeSecretRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the secret instance with the specified UUID.
// Version: midnight-ride
func (secret) GetByUUID2(session *Session, uUID string) (retval SecretRef, err error) {
	method := "secret.get_by_uuid"
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
	retval, err = deserializeSecretRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given secret.
// Version: midnight-ride
func (secret) GetRecord(session *Session, self SecretRef) (retval SecretRecord, err error) {
	method := "secret.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given secret.
// Version: midnight-ride
func (secret) GetRecord2(session *Session, self SecretRef) (retval SecretRecord, err error) {
	method := "secret.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeSecretRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeSecretRecord(method+" -> ", result)
	return
}
