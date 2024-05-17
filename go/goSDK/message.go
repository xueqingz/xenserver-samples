package xenapi

import (
	"fmt"
	"time"
)

type MessageRecord struct {
	// Unique identifier/object reference
	UUID string
	// The name of the message
	Name string
	// The message priority, 0 being low priority
	Priority int
	// The class of the object this message is associated with
	Cls Cls
	// The uuid of the object this message is associated with
	ObjUUID string
	// The time at which the message was created
	Timestamp time.Time
	// The body of the message
	Body string
}

type MessageRef string

// An message for the attention of the administrator
type message struct{}

var Message message

// GetAllRecordsWhere:
// Version: orlando
func (message) GetAllRecordsWhere(session *Session, expr string) (retval map[MessageRef]MessageRecord, err error) {
	method := "message.get_all_records_where"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	exprArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "expr"), expr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, exprArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefToMessageRecordMap(method+" -> ", result)
	return
}

// GetAllRecordsWhere2:
// Version: orlando
func (message) GetAllRecordsWhere2(session *Session, expr string) (retval map[MessageRef]MessageRecord, err error) {
	method := "message.get_all_records_where"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	exprArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "expr"), expr)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, exprArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefToMessageRecordMap(method+" -> ", result)
	return
}

// GetAllRecords:
// Version: orlando
func (message) GetAllRecords(session *Session) (retval map[MessageRef]MessageRecord, err error) {
	method := "message.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefToMessageRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1:
// Version: orlando
func (message) GetAllRecords1(session *Session) (retval map[MessageRef]MessageRecord, err error) {
	method := "message.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefToMessageRecordMap(method+" -> ", result)
	return
}

// GetByUUID:
// Version: orlando
func (message) GetByUUID(session *Session, uuid string) (retval MessageRef, err error) {
	method := "message.get_by_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	uuidArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uuid"), uuid)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uuidArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRef(method+" -> ", result)
	return
}

// GetByUUID2:
// Version: orlando
func (message) GetByUUID2(session *Session, uuid string) (retval MessageRef, err error) {
	method := "message.get_by_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	uuidArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "uuid"), uuid)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, uuidArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRef(method+" -> ", result)
	return
}

// GetRecord:
// Version: orlando
func (message) GetRecord(session *Session, self MessageRef) (retval MessageRecord, err error) {
	method := "message.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeMessageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRecord(method+" -> ", result)
	return
}

// GetRecord2:
// Version: orlando
func (message) GetRecord2(session *Session, self MessageRef) (retval MessageRecord, err error) {
	method := "message.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeMessageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRecord(method+" -> ", result)
	return
}

// GetSince:
// Version: orlando
func (message) GetSince(session *Session, since time.Time) (retval map[MessageRef]MessageRecord, err error) {
	method := "message.get_since"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	sinceArg, err := serializeTime(fmt.Sprintf("%s(%s)", method, "since"), since)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, sinceArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefToMessageRecordMap(method+" -> ", result)
	return
}

// GetSince2:
// Version: orlando
func (message) GetSince2(session *Session, since time.Time) (retval map[MessageRef]MessageRecord, err error) {
	method := "message.get_since"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	sinceArg, err := serializeTime(fmt.Sprintf("%s(%s)", method, "since"), since)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, sinceArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefToMessageRecordMap(method+" -> ", result)
	return
}

// GetAll:
// Version: orlando
func (message) GetAll(session *Session) (retval []MessageRef, err error) {
	method := "message.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefSet(method+" -> ", result)
	return
}

// GetAll1:
// Version: orlando
func (message) GetAll1(session *Session) (retval []MessageRef, err error) {
	method := "message.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefSet(method+" -> ", result)
	return
}

// Get:
// Version: orlando
func (message) Get(session *Session, cls Cls, objUUID string, since time.Time) (retval map[MessageRef]MessageRecord, err error) {
	method := "message.get"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	clsArg, err := serializeEnumCls(fmt.Sprintf("%s(%s)", method, "cls"), cls)
	if err != nil {
		return
	}
	objUUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "obj_uuid"), objUUID)
	if err != nil {
		return
	}
	sinceArg, err := serializeTime(fmt.Sprintf("%s(%s)", method, "since"), since)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, clsArg, objUUIDArg, sinceArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefToMessageRecordMap(method+" -> ", result)
	return
}

// Get4:
// Version: orlando
func (message) Get4(session *Session, cls Cls, objUUID string, since time.Time) (retval map[MessageRef]MessageRecord, err error) {
	method := "message.get"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	clsArg, err := serializeEnumCls(fmt.Sprintf("%s(%s)", method, "cls"), cls)
	if err != nil {
		return
	}
	objUUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "obj_uuid"), objUUID)
	if err != nil {
		return
	}
	sinceArg, err := serializeTime(fmt.Sprintf("%s(%s)", method, "since"), since)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, clsArg, objUUIDArg, sinceArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRefToMessageRecordMap(method+" -> ", result)
	return
}

// DestroyMany:
// Version: closed
func (message) DestroyMany(session *Session, messages []MessageRef) (err error) {
	method := "message.destroy_many"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	messagesArg, err := serializeMessageRefSet(fmt.Sprintf("%s(%s)", method, "messages"), messages)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, messagesArg)
	return
}

// AsyncDestroyMany:
// Version: closed
func (message) AsyncDestroyMany(session *Session, messages []MessageRef) (retval TaskRef, err error) {
	method := "Async.message.destroy_many"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	messagesArg, err := serializeMessageRefSet(fmt.Sprintf("%s(%s)", method, "messages"), messages)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, messagesArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// DestroyMany2:
// Version: closed
func (message) DestroyMany2(session *Session, messages []MessageRef) (err error) {
	method := "message.destroy_many"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	messagesArg, err := serializeMessageRefSet(fmt.Sprintf("%s(%s)", method, "messages"), messages)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, messagesArg)
	return
}

// AsyncDestroyMany2:
// Version: closed
func (message) AsyncDestroyMany2(session *Session, messages []MessageRef) (retval TaskRef, err error) {
	method := "Async.message.destroy_many"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	messagesArg, err := serializeMessageRefSet(fmt.Sprintf("%s(%s)", method, "messages"), messages)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, messagesArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy:
// Version: orlando
func (message) Destroy(session *Session, self MessageRef) (err error) {
	method := "message.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeMessageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// Destroy2:
// Version: orlando
func (message) Destroy2(session *Session, self MessageRef) (err error) {
	method := "message.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeMessageRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// Create:
// Version: orlando
func (message) Create(session *Session, name string, priority int, cls Cls, objUUID string, body string) (retval MessageRef, err error) {
	method := "message.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	priorityArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "priority"), priority)
	if err != nil {
		return
	}
	clsArg, err := serializeEnumCls(fmt.Sprintf("%s(%s)", method, "cls"), cls)
	if err != nil {
		return
	}
	objUUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "obj_uuid"), objUUID)
	if err != nil {
		return
	}
	bodyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "body"), body)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameArg, priorityArg, clsArg, objUUIDArg, bodyArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRef(method+" -> ", result)
	return
}

// Create6:
// Version: orlando
func (message) Create6(session *Session, name string, priority int, cls Cls, objUUID string, body string) (retval MessageRef, err error) {
	method := "message.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	nameArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "name"), name)
	if err != nil {
		return
	}
	priorityArg, err := serializeInt(fmt.Sprintf("%s(%s)", method, "priority"), priority)
	if err != nil {
		return
	}
	clsArg, err := serializeEnumCls(fmt.Sprintf("%s(%s)", method, "cls"), cls)
	if err != nil {
		return
	}
	objUUIDArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "obj_uuid"), objUUID)
	if err != nil {
		return
	}
	bodyArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "body"), body)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameArg, priorityArg, clsArg, objUUIDArg, bodyArg)
	if err != nil {
		return
	}
	retval, err = deserializeMessageRef(method+" -> ", result)
	return
}
