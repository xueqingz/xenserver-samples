package xenapi

import (
	"fmt"
	"time"
)

type TaskRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []TaskAllowedOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]TaskAllowedOperations
	// Time task was created
	Created time.Time
	// Time task finished (i.e. succeeded or failed). If task-status is pending, then the value of this field has no meaning
	Finished time.Time
	// current status of the task
	Status TaskStatusType
	// the host on which the task is running
	ResidentOn HostRef
	// This field contains the estimated fraction of the task which is complete. This field should not be used to determine whether the task is complete - for this the status field of the task should be used.
	Progress float64
	// if the task has completed successfully, this field contains the type of the encoded result (i.e. name of the class whose reference is in the result field). Undefined otherwise.
	Type string
	// if the task has completed successfully, this field contains the result value (either Void or an object reference). Undefined otherwise.
	Result string
	// if the task has failed, this field contains the set of associated error strings. Undefined otherwise.
	ErrorInfo []string
	// additional configuration
	OtherConfig map[string]string
	// Ref pointing to the task this is a substask of.
	SubtaskOf TaskRef
	// List pointing to all the substasks.
	Subtasks []TaskRef
	// Function call trace for debugging.
	Backtrace string
}

type TaskRef string

// A long-running asynchronous task
type task struct{}

var Task task

// GetRecord: Get a record containing the current state of the given task.
func (task) GetRecord(session *Session, self TaskRef) (retval TaskRecord, err error) {
	method := "task.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRecord(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the task instance with the specified UUID.
func (task) GetByUUID(session *Session, uUID string) (retval TaskRef, err error) {
	method := "task.get_by_uuid"
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
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// GetByNameLabel: Get all the task instances with the given label.
func (task) GetByNameLabel(session *Session, label string) (retval []TaskRef, err error) {
	method := "task.get_by_name_label"
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
	retval, err = deserializeTaskRefSet(method+" -> ", result)
	return
}

// GetUUID: Get the uuid field of the given task.
func (task) GetUUID(session *Session, self TaskRef) (retval string, err error) {
	method := "task.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given task.
func (task) GetNameLabel(session *Session, self TaskRef) (retval string, err error) {
	method := "task.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given task.
func (task) GetNameDescription(session *Session, self TaskRef) (retval string, err error) {
	method := "task.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetAllowedOperations: Get the allowed_operations field of the given task.
func (task) GetAllowedOperations(session *Session, self TaskRef) (retval []TaskAllowedOperations, err error) {
	method := "task.get_allowed_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumTaskAllowedOperationsSet(method+" -> ", result)
	return
}

// GetCurrentOperations: Get the current_operations field of the given task.
func (task) GetCurrentOperations(session *Session, self TaskRef) (retval map[string]TaskAllowedOperations, err error) {
	method := "task.get_current_operations"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringToEnumTaskAllowedOperationsMap(method+" -> ", result)
	return
}

// GetCreated: Get the created field of the given task.
func (task) GetCreated(session *Session, self TaskRef) (retval time.Time, err error) {
	method := "task.get_created"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetFinished: Get the finished field of the given task.
func (task) GetFinished(session *Session, self TaskRef) (retval time.Time, err error) {
	method := "task.get_finished"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetStatus: Get the status field of the given task.
func (task) GetStatus(session *Session, self TaskRef) (retval TaskStatusType, err error) {
	method := "task.get_status"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeEnumTaskStatusType(method+" -> ", result)
	return
}

// GetResidentOn: Get the resident_on field of the given task.
func (task) GetResidentOn(session *Session, self TaskRef) (retval HostRef, err error) {
	method := "task.get_resident_on"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetProgress: Get the progress field of the given task.
func (task) GetProgress(session *Session, self TaskRef) (retval float64, err error) {
	method := "task.get_progress"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeFloat(method+" -> ", result)
	return
}

// GetType: Get the type field of the given task.
func (task) GetType(session *Session, self TaskRef) (retval string, err error) {
	method := "task.get_type"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetResult: Get the result field of the given task.
func (task) GetResult(session *Session, self TaskRef) (retval string, err error) {
	method := "task.get_result"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetErrorInfo: Get the error_info field of the given task.
func (task) GetErrorInfo(session *Session, self TaskRef) (retval []string, err error) {
	method := "task.get_error_info"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeStringSet(method+" -> ", result)
	return
}

// GetOtherConfig: Get the other_config field of the given task.
func (task) GetOtherConfig(session *Session, self TaskRef) (retval map[string]string, err error) {
	method := "task.get_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSubtaskOf: Get the subtask_of field of the given task.
func (task) GetSubtaskOf(session *Session, self TaskRef) (retval TaskRef, err error) {
	method := "task.get_subtask_of"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSubtasks: Get the subtasks field of the given task.
func (task) GetSubtasks(session *Session, self TaskRef) (retval []TaskRef, err error) {
	method := "task.get_subtasks"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRefSet(method+" -> ", result)
	return
}

// GetBacktrace: Get the backtrace field of the given task.
func (task) GetBacktrace(session *Session, self TaskRef) (retval string, err error) {
	method := "task.get_backtrace"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetOtherConfig: Set the other_config field of the given task.
func (task) SetOtherConfig(session *Session, self TaskRef, value map[string]string) (err error) {
	method := "task.set_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AddToOtherConfig: Add the given key-value pair to the other_config field of the given task.
func (task) AddToOtherConfig(session *Session, self TaskRef, key string, value string) (err error) {
	method := "task.add_to_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given task.  If the key is not in that Map, then do nothing.
func (task) RemoveFromOtherConfig(session *Session, self TaskRef, key string) (err error) {
	method := "task.remove_from_other_config"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Create: Create a new task object which must be manually destroyed.
func (task) Create(session *Session, label string, description string) (retval TaskRef, err error) {
	method := "task.create"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	labelArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "label"), label)
	if err != nil {
		return
	}
	descriptionArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "description"), description)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, labelArg, descriptionArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Destroy: Destroy the task object
func (task) Destroy(session *Session, self TaskRef) (err error) {
	method := "task.destroy"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// Cancel: Request that a task be cancelled. Note that a task may fail to be cancelled and may complete or fail normally and note that, even when a task does cancel, it might take an arbitrary amount of time.
//
// Errors:
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (task) Cancel(session *Session, task TaskRef) (err error) {
	method := "task.cancel"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	taskArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "task"), task)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, taskArg)
	return
}

// AsyncCancel: Request that a task be cancelled. Note that a task may fail to be cancelled and may complete or fail normally and note that, even when a task does cancel, it might take an arbitrary amount of time.
//
// Errors:
// OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (task) AsyncCancel(session *Session, task TaskRef) (retval TaskRef, err error) {
	method := "Async.task.cancel"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	taskArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "task"), task)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, taskArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetStatus: Set the task status
func (task) SetStatus(session *Session, self TaskRef, value TaskStatusType) (err error) {
	method := "task.set_status"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeEnumTaskStatusType(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetProgress: Set the task progress
func (task) SetProgress(session *Session, self TaskRef, value float64) (err error) {
	method := "task.set_progress"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// SetResult: Set the task result
func (task) SetResult(session *Session, self TaskRef, value string) (err error) {
	method := "task.set_result"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetErrorInfo: Set the task error info
func (task) SetErrorInfo(session *Session, self TaskRef, value []string) (err error) {
	method := "task.set_error_info"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeTaskRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	valueArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "value"), value)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg, valueArg)
	return
}

// GetAll: Return a list of all the tasks known to the system.
func (task) GetAll(session *Session) (retval []TaskRef, err error) {
	method := "task.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRefSet(method+" -> ", result)
	return
}

// GetAllRecords: Return a map of task references to task records for all tasks known to the system.
func (task) GetAllRecords(session *Session) (retval map[TaskRef]TaskRecord, err error) {
	method := "task.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRefToTaskRecordMap(method+" -> ", result)
	return
}
