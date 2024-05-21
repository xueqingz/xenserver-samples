package testGoSDK

import (
	"testing"

	"github.com/google/uuid"
	xenapi "github.com/xenserver/xenserver-samples/go/goSDK"
)

func TestSRBase(t *testing.T) {
	var deviceConfig = make(map[string]string)
	var smConfig = make(map[string]string)
	var testSRName = "TestSR: DO NOT USE (created by sr_test.go)"
	var testSRDesc = "Should be automatically deleted"
	var testSRType = "dummy"
	var testSRContent = "contenttype"
	var testSRSize = 100000

	hostRefs, err := xenapi.Host.GetAll(session)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	_, err = xenapi.SR.Create(session, hostRefs[0], deviceConfig, testSRSize, testSRName, testSRDesc, testSRType, testSRContent, true, smConfig)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	taskRef, err := xenapi.SR.AsyncCreate(session, hostRefs[0], deviceConfig, testSRSize, testSRName, testSRDesc, testSRType, testSRContent, true, smConfig)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	err = WaitForTask(taskRef, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	srRefs, err := xenapi.SR.GetByNameLabel(session, testSRName)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	for _, srRef := range srRefs {
		pbdRecords, err := xenapi.PBD.GetAllRecords(session)
		if err != nil {
			t.Log(err)
			t.Fail()
			return
		}
		for pbdRef, pbdRecord := range pbdRecords {
			if pbdRecord.CurrentlyAttached {
				pbdSRRef, err := xenapi.PBD.GetSR(session, pbdRef)
				if err != nil {
					t.Log(err)
					t.Fail()
					return
				}
				if srRef == pbdSRRef {
					err = xenapi.PBD.Unplug(session, pbdRef)
					if err != nil {
						t.Log(err)
						t.Fail()
						return
					}
					err = xenapi.PBD.Destroy(session, pbdRef)
					if err != nil {
						t.Log(err)
					}
				}
			}
		}
		err = xenapi.SR.Forget(session, srRef)
		if err != nil {
			t.Log(err)
			t.Fail()
			return
		}
	}
	_, err = xenapi.SR.Introduce(session, uuid.NewString(), testSRName, testSRDesc, testSRType, testSRContent, true, smConfig)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	taskRef, err = xenapi.SR.AsyncIntroduce(session, uuid.NewString(), testSRName, testSRDesc, testSRType, testSRContent, true, smConfig)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	err = WaitForTask(taskRef, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	srRefs, err = xenapi.SR.GetByNameLabel(session, testSRName)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	err = xenapi.SR.Forget(session, srRefs[0])
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	taskRef, err = xenapi.SR.AsyncForget(session, srRefs[1])
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	err = WaitForTask(taskRef, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
}
