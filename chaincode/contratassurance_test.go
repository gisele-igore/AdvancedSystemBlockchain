package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewContratAssurance(t *testing.T, stub *shim.MockStub, uuid string, iDCompagnieAssurance string,
	codeAcheteurAssurance string, dateDebut string, dateFin string, contratAssurancePDF string, signatureAcheteur string, signatureCompagnie string) {

	displayNewTest("Create ContratAssurance Test When ContratAssurance does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateContratAssurance"),
		[]byte(uuid), []byte(iDCompagnieAssurance), []byte(codeAcheteurAssurance), []byte(dateDebut), []byte(dateFin), []byte(contratAssurancePDF), []byte(signatureAcheteur), []byte(signatureCompagnie)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func checkGetExistingContratAssurance(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Get Existing ContratAssurance " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetContratAssuranceByID"), []byte(uuid)})
	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}

	ass := ContratAssurance{}
	_ = json.Unmarshal(response.Payload, &ass)

	if ass.UUID != uuid {
		t.Fail()
	}
}

func checkUpdateContratAssurance(t *testing.T, stub *shim.MockStub, uuid string, newDateDebut string, newDateFin string, newContratAssurancePDF string) {
	displayNewTest("checkUpdateContratAssurance")
	res := stub.MockInvoke("1", [][]byte{[]byte("UpdateContratAssuranceByID"), []byte(uuid), []byte(newDateDebut), []byte(newDateFin), []byte(newContratAssurancePDF)})
	if res.Status != shim.OK {
		fmt.Println("UpdadeAsset", uuid, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("UpdadeAsset", uuid, "failed to get value")
		t.FailNow()
	}
}

func checkUnregisterContratAssurance(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Unregister existing Contrat Assurance " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("UnregisterContratAssuranceByID"), []byte(uuid)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func TestCreateContratAssurance(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	/*checkCreateNewContratAssurance(t, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")
	checkGetExistingContratAssurance(t, stub, "O0")*/
	checkCreateNewContratAssurance(t, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")
}

func TestGetContratAssuranceByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewContratAssurance(t, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")
	checkGetExistingContratAssurance(t, stub, "O0")
}

func TestUpdateContratAssuranceByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewContratAssurance(t, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")

	checkUpdateContratAssurance(t, stub, "O0", "01022020", "01072020", "yyyyyyyyyy")
	checkGetExistingContratAssurance(t, stub, "O0")
}

func TestCreateContratAssuranceAlreadyExist(t *testing.T) {
	displayNewTest("Create ContratAssurance Test When ContratAssurance already exist")
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewContratAssurance(t, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")
	checkCreateNewContratAssurance(t, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")
}

func TestUnregisterContratAssurance(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewContratAssurance(t, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")
	checkUnregisterContratAssurance(t, stub, "O0")
}
