package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewCompagnieAssurance(t *testing.T, stub *shim.MockStub, uuid string,
	nom string, contact string, adresse string) {
	displayNewTest("Create CompagnieAssurance Test When CompagnieAssurance does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateCompagnieAssurance"), []byte(uuid), []byte(nom), []byte(contact), []byte(adresse)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func checkGetExistingCompagnieAssurance(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Get Existing CompagnieAssurance " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetCompagnieAssuranceByID"), []byte(uuid)})
	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}

	org := CompagnieAssurance{}
	_ = json.Unmarshal(response.Payload, &org)

	if org.UUID != uuid {
		t.Fail()
	}
}

func checkUpdateCompagnieAssurance(t *testing.T, stub *shim.MockStub, uuid string, newNom string, newContact string, newAdresse string) {
	displayNewTest("checkUpdateCompagnieAssurance")
	res := stub.MockInvoke("1", [][]byte{[]byte("UpdateCompagnieAssuranceByID"), []byte(uuid), []byte(newNom), []byte(newContact), []byte(newAdresse)})
	if res.Status != shim.OK {
		fmt.Println("UpdadeOrganization", uuid, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("UpdadeOrganization", uuid, "failed to get value")
		t.FailNow()
	}
}

func checkUnregisterCompagnieAssurance(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Unregister existing Compagnie Assurance " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("UnregisterCompagnieAssuranceByID"), []byte(uuid)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func TestCreateCompagnieAssurance(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	/*checkCreateNewCompagnieAssurance(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
	checkGetExistingCompagnieAssurance(t, stub, "O0")*/
	checkCreateNewCompagnieAssurance(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
}

func TestGetCompagnieAssuranceByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewCompagnieAssurance(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
	checkGetExistingCompagnieAssurance(t, stub, "O0")
}

func TestUpdateCompagnieAssuranceByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewCompagnieAssurance(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")

	checkUpdateCompagnieAssurance(t, stub, "O0", "Scorechain Luxembourg", "Scorechain", "ISP")
	checkGetExistingCompagnieAssurance(t, stub, "O0")
}

func TestCreateCompagnieAssuranceAlreadyExist(t *testing.T) {
	displayNewTest("Create CompagnieAssurance Test When CompagnieAssurance already exist")
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewCompagnieAssurance(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
}

func TestUnregisterCompagnieAssurance(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewCompagnieAssurance(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
	checkUnregisterCompagnieAssurance(t, stub, "O0")
}
