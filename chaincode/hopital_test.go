package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewHopital(t *testing.T, stub *shim.MockStub, uuid string,
	nom string, contact string, adresse string) {
	displayNewTest("Create Hopital Test When Hopital does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateHopital"),
		[]byte(uuid), []byte(nom), []byte(contact), []byte(adresse)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func checkGetExistingHopital(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Get Existing Hopital " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetHopitalByID"), []byte(uuid)})
	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}

	org := Hopital{}
	_ = json.Unmarshal(response.Payload, &org)

	if org.UUID != uuid {
		t.Fail()
	}
}

func TestCreateHopital(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewHopital(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
	checkGetExistingCompagnieAssurance(t, stub, "O0")
	checkCreateNewHopital(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
}

func TestGetHopitalByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewHopital(t, stub, "O0", "SCHAIN", "Scorechain", "ISP")
	checkGetExistingHopital(t, stub, "O0")
}
