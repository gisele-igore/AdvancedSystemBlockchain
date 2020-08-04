package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewAcheteurAssurance(t *testing.T, stub *shim.MockStub, uuid string,
	nom string, contact string, adresse string, passportID string, visaID string) {

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateAcheteurAssurance"),
		[]byte(uuid), []byte(nom), []byte(contact), []byte(adresse), []byte(passportID), []byte(visaID)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func checkGetExistingAcheteurAssurance(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Get Existing AcheteurAssurance " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetAcheteurAssuranceByID"), []byte(uuid)})
	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}

	org := AcheteurAssurance{}
	_ = json.Unmarshal(response.Payload, &org)

	if org.UUID != uuid {
		t.Fail()
	}
}

func checkUpdateAcheteurAssurance(t *testing.T, stub *shim.MockStub, uuid string, name string, prettyName string, typeOfOrganization string) {
	displayNewTest("checkUpdateAcheteurAssurance")
	res := stub.MockInvoke("1", [][]byte{[]byte("UpdateAcheteurAssuranceByID"), []byte(uuid), []byte(name), []byte(prettyName), []byte(typeOfOrganization)})
	if res.Status != shim.OK {
		fmt.Println("UpdadeOrganization", uuid, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("UpdadeOrganization", uuid, "failed to get value")
		t.FailNow()
	}
}

func checkUnregisterAcheteurAssurance(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Unregister existing Acheteur Assurance " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("UnregisterAcheteurAssuranceByID"), []byte(uuid)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func TestCreateAcheteurAssurance(t *testing.T) {
	displayNewTest("Create AcheteurAssurance Test When AcheteurAssurance does not exist")
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewAcheteurAssurance(t, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
}

func TestGetAcheteurAssuranceByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewAcheteurAssurance(t, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
	checkGetExistingAcheteurAssurance(t, stub, "O0")
}

func TestUpdateAcheteurAssuranceByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewAcheteurAssurance(t, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")

	checkUpdateAcheteurAssurance(t, stub, "O0", "DEKPO", "vietnam", "ISP")
	checkGetExistingAcheteurAssurance(t, stub, "O0")
}

func TestCreateAcheteurAssuranceAlreadyExist(t *testing.T) {
	displayNewTest("Create AcheteurAssurance Test When AcheteurAssurance already exist")
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewAcheteurAssurance(t, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
	checkCreateNewAcheteurAssurance(t, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
}

func TestUnregisterAcheteurAssurance(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewAcheteurAssurance(t, stub, "O0", "DEKPE", "Hanoi", "KTX", "0OK12", "DH23")
	checkUnregisterAcheteurAssurance(t, stub, "O0")
}
