package xenapi

import (
	"fmt"
)

type RepositoryRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// Base URL of binary packages in this repository
	BinaryURL string
	// Base URL of source packages in this repository
	SourceURL string
	// True if updateinfo.xml in this repository needs to be parsed
	Update bool
	// SHA256 checksum of latest updateinfo.xml.gz in this repository if its &apos;update&apos; is true
	Hash string
	// True if all hosts in pool is up to date with this repository
	UpToDate bool
	// The file name of the GPG public key of this repository
	GpgkeyPath string
}

type RepositoryRef string

// Repository for updates
type repository struct{}

var Repository repository

// GetAllRecords: Return a map of Repository references to Repository records for all Repositorys known to the system.
// Version: 1.301.0
func (repository) GetAllRecords(session *Session) (retval map[RepositoryRef]RepositoryRecord, err error) {
	method := "Repository.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeRepositoryRefToRepositoryRecordMap(method+" -> ", result)
	return
}

// GetAllRecords1: Return a map of Repository references to Repository records for all Repositorys known to the system.
// Version: 1.301.0
func (repository) GetAllRecords1(session *Session) (retval map[RepositoryRef]RepositoryRecord, err error) {
	method := "Repository.get_all_records"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeRepositoryRefToRepositoryRecordMap(method+" -> ", result)
	return
}

// GetAll: Return a list of all the Repositorys known to the system.
// Version: 1.301.0
func (repository) GetAll(session *Session) (retval []RepositoryRef, err error) {
	method := "Repository.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeRepositoryRefSet(method+" -> ", result)
	return
}

// GetAll1: Return a list of all the Repositorys known to the system.
// Version: 1.301.0
func (repository) GetAll1(session *Session) (retval []RepositoryRef, err error) {
	method := "Repository.get_all"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeRepositoryRefSet(method+" -> ", result)
	return
}

// SetGpgkeyPath: Set the file name of the GPG public key of the repository
// Version: closed
func (repository) SetGpgkeyPath(session *Session, self RepositoryRef, value string) (err error) {
	method := "Repository.set_gpgkey_path"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetGpgkeyPath: Set the file name of the GPG public key of the repository
// Version: closed
func (repository) AsyncSetGpgkeyPath(session *Session, self RepositoryRef, value string) (retval TaskRef, err error) {
	method := "Async.Repository.set_gpgkey_path"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetGpgkeyPath3: Set the file name of the GPG public key of the repository
// Version: closed
func (repository) SetGpgkeyPath3(session *Session, self RepositoryRef, value string) (err error) {
	method := "Repository.set_gpgkey_path"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// AsyncSetGpgkeyPath3: Set the file name of the GPG public key of the repository
// Version: closed
func (repository) AsyncSetGpgkeyPath3(session *Session, self RepositoryRef, value string) (retval TaskRef, err error) {
	method := "Async.Repository.set_gpgkey_path"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Forget: Remove the repository record from the database
// Version: 1.301.0
func (repository) Forget(session *Session, self RepositoryRef) (err error) {
	method := "Repository.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncForget: Remove the repository record from the database
// Version: 1.301.0
func (repository) AsyncForget(session *Session, self RepositoryRef) (retval TaskRef, err error) {
	method := "Async.Repository.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Forget2: Remove the repository record from the database
// Version: 1.301.0
func (repository) Forget2(session *Session, self RepositoryRef) (err error) {
	method := "Repository.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, selfArg)
	return
}

// AsyncForget2: Remove the repository record from the database
// Version: 1.301.0
func (repository) AsyncForget2(session *Session, self RepositoryRef) (retval TaskRef, err error) {
	method := "Async.Repository.forget"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// Introduce: Add the configuration for a new repository
// Version: 1.301.0
func (repository) Introduce(session *Session, nameLabel string, nameDescription string, binaryURL string, sourceURL string, update bool, gpgkeyPath string) (retval RepositoryRef, err error) {
	method := "Repository.introduce"
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
	binaryURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "binary_url"), binaryURL)
	if err != nil {
		return
	}
	sourceURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "source_url"), sourceURL)
	if err != nil {
		return
	}
	updateArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "update"), update)
	if err != nil {
		return
	}
	gpgkeyPathArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "gpgkey_path"), gpgkeyPath)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, binaryURLArg, sourceURLArg, updateArg, gpgkeyPathArg)
	if err != nil {
		return
	}
	retval, err = deserializeRepositoryRef(method+" -> ", result)
	return
}

// AsyncIntroduce: Add the configuration for a new repository
// Version: 1.301.0
func (repository) AsyncIntroduce(session *Session, nameLabel string, nameDescription string, binaryURL string, sourceURL string, update bool, gpgkeyPath string) (retval TaskRef, err error) {
	method := "Async.Repository.introduce"
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
	binaryURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "binary_url"), binaryURL)
	if err != nil {
		return
	}
	sourceURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "source_url"), sourceURL)
	if err != nil {
		return
	}
	updateArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "update"), update)
	if err != nil {
		return
	}
	gpgkeyPathArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "gpgkey_path"), gpgkeyPath)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, binaryURLArg, sourceURLArg, updateArg, gpgkeyPathArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Introduce7: Add the configuration for a new repository
// Version: 1.301.0
func (repository) Introduce7(session *Session, nameLabel string, nameDescription string, binaryURL string, sourceURL string, update bool, gpgkeyPath string) (retval RepositoryRef, err error) {
	method := "Repository.introduce"
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
	binaryURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "binary_url"), binaryURL)
	if err != nil {
		return
	}
	sourceURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "source_url"), sourceURL)
	if err != nil {
		return
	}
	updateArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "update"), update)
	if err != nil {
		return
	}
	gpgkeyPathArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "gpgkey_path"), gpgkeyPath)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, binaryURLArg, sourceURLArg, updateArg, gpgkeyPathArg)
	if err != nil {
		return
	}
	retval, err = deserializeRepositoryRef(method+" -> ", result)
	return
}

// AsyncIntroduce7: Add the configuration for a new repository
// Version: 1.301.0
func (repository) AsyncIntroduce7(session *Session, nameLabel string, nameDescription string, binaryURL string, sourceURL string, update bool, gpgkeyPath string) (retval TaskRef, err error) {
	method := "Async.Repository.introduce"
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
	binaryURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "binary_url"), binaryURL)
	if err != nil {
		return
	}
	sourceURLArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "source_url"), sourceURL)
	if err != nil {
		return
	}
	updateArg, err := serializeBool(fmt.Sprintf("%s(%s)", method, "update"), update)
	if err != nil {
		return
	}
	gpgkeyPathArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "gpgkey_path"), gpgkeyPath)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, nameLabelArg, nameDescriptionArg, binaryURLArg, sourceURLArg, updateArg, gpgkeyPathArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// SetNameDescription: Set the name/description field of the given Repository.
// Version: 1.301.0
func (repository) SetNameDescription(session *Session, self RepositoryRef, value string) (err error) {
	method := "Repository.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription3: Set the name/description field of the given Repository.
// Version: 1.301.0
func (repository) SetNameDescription3(session *Session, self RepositoryRef, value string) (err error) {
	method := "Repository.set_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameDescription2: Set the name/description field of the given Repository.
// Version: rio
func (repository) SetNameDescription2(session *Session, value string) (err error) {
	method := "Repository.set_name_description"
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

// SetNameLabel: Set the name/label field of the given Repository.
// Version: 1.301.0
func (repository) SetNameLabel(session *Session, self RepositoryRef, value string) (err error) {
	method := "Repository.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel3: Set the name/label field of the given Repository.
// Version: 1.301.0
func (repository) SetNameLabel3(session *Session, self RepositoryRef, value string) (err error) {
	method := "Repository.set_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// SetNameLabel2: Set the name/label field of the given Repository.
// Version: rio
func (repository) SetNameLabel2(session *Session, value string) (err error) {
	method := "Repository.set_name_label"
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

// GetGpgkeyPath: Get the gpgkey_path field of the given Repository.
// Version: 1.301.0
func (repository) GetGpgkeyPath(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_gpgkey_path"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetGpgkeyPath2: Get the gpgkey_path field of the given Repository.
// Version: 1.301.0
func (repository) GetGpgkeyPath2(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_gpgkey_path"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUpToDate: Get the up_to_date field of the given Repository.
// Version: 1.301.0
func (repository) GetUpToDate(session *Session, self RepositoryRef) (retval bool, err error) {
	method := "Repository.get_up_to_date"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUpToDate2: Get the up_to_date field of the given Repository.
// Version: 1.301.0
func (repository) GetUpToDate2(session *Session, self RepositoryRef) (retval bool, err error) {
	method := "Repository.get_up_to_date"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHash: Get the hash field of the given Repository.
// Version: 1.301.0
func (repository) GetHash(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_hash"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetHash2: Get the hash field of the given Repository.
// Version: 1.301.0
func (repository) GetHash2(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_hash"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUpdate: Get the update field of the given Repository.
// Version: 1.301.0
func (repository) GetUpdate(session *Session, self RepositoryRef) (retval bool, err error) {
	method := "Repository.get_update"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUpdate2: Get the update field of the given Repository.
// Version: 1.301.0
func (repository) GetUpdate2(session *Session, self RepositoryRef) (retval bool, err error) {
	method := "Repository.get_update"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSourceURL: Get the source_url field of the given Repository.
// Version: 1.301.0
func (repository) GetSourceURL(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_source_url"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetSourceURL2: Get the source_url field of the given Repository.
// Version: 1.301.0
func (repository) GetSourceURL2(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_source_url"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBinaryURL: Get the binary_url field of the given Repository.
// Version: 1.301.0
func (repository) GetBinaryURL(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_binary_url"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetBinaryURL2: Get the binary_url field of the given Repository.
// Version: 1.301.0
func (repository) GetBinaryURL2(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_binary_url"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription: Get the name/description field of the given Repository.
// Version: 1.301.0
func (repository) GetNameDescription(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameDescription2: Get the name/description field of the given Repository.
// Version: 1.301.0
func (repository) GetNameDescription2(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_name_description"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel: Get the name/label field of the given Repository.
// Version: 1.301.0
func (repository) GetNameLabel(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetNameLabel2: Get the name/label field of the given Repository.
// Version: 1.301.0
func (repository) GetNameLabel2(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_name_label"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID: Get the uuid field of the given Repository.
// Version: 1.301.0
func (repository) GetUUID(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetUUID2: Get the uuid field of the given Repository.
// Version: 1.301.0
func (repository) GetUUID2(session *Session, self RepositoryRef) (retval string, err error) {
	method := "Repository.get_uuid"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
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

// GetByNameLabel: Get all the Repository instances with the given label.
// Version: 1.301.0
func (repository) GetByNameLabel(session *Session, label string) (retval []RepositoryRef, err error) {
	method := "Repository.get_by_name_label"
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
	retval, err = deserializeRepositoryRefSet(method+" -> ", result)
	return
}

// GetByNameLabel2: Get all the Repository instances with the given label.
// Version: 1.301.0
func (repository) GetByNameLabel2(session *Session, label string) (retval []RepositoryRef, err error) {
	method := "Repository.get_by_name_label"
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
	retval, err = deserializeRepositoryRefSet(method+" -> ", result)
	return
}

// GetByUUID: Get a reference to the Repository instance with the specified UUID.
// Version: 1.301.0
func (repository) GetByUUID(session *Session, uUID string) (retval RepositoryRef, err error) {
	method := "Repository.get_by_uuid"
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
	retval, err = deserializeRepositoryRef(method+" -> ", result)
	return
}

// GetByUUID2: Get a reference to the Repository instance with the specified UUID.
// Version: 1.301.0
func (repository) GetByUUID2(session *Session, uUID string) (retval RepositoryRef, err error) {
	method := "Repository.get_by_uuid"
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
	retval, err = deserializeRepositoryRef(method+" -> ", result)
	return
}

// GetRecord: Get a record containing the current state of the given Repository.
// Version: 1.301.0
func (repository) GetRecord(session *Session, self RepositoryRef) (retval RepositoryRecord, err error) {
	method := "Repository.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeRepositoryRecord(method+" -> ", result)
	return
}

// GetRecord2: Get a record containing the current state of the given Repository.
// Version: 1.301.0
func (repository) GetRecord2(session *Session, self RepositoryRef) (retval RepositoryRecord, err error) {
	method := "Repository.get_record"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	selfArg, err := serializeRepositoryRef(fmt.Sprintf("%s(%s)", method, "self"), self)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, selfArg)
	if err != nil {
		return
	}
	retval, err = deserializeRepositoryRecord(method+" -> ", result)
	return
}
