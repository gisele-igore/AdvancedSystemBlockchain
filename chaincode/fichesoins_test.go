package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkCreateNewFicheSoins(t *testing.T, stub *shim.MockStub, uuid string, iDContrat string,
	iDCompagnieAssurance string, iDHopital string, codeAcheteurAssurance string, dateDebut string, dateFin string,
	ficheSoinsPDF string, signatureAcheteur string, signatureCompagnie string, signatureHopital string) {
	displayNewTest("Create FicheSoins Test When FicheSoins does not exist")

	response := stub.MockInvoke("1", [][]byte{[]byte("CreateFicheSoins"),
		[]byte(uuid), []byte(iDContrat), []byte(iDCompagnieAssurance), []byte(iDHopital), []byte(codeAcheteurAssurance), []byte(dateDebut),
		[]byte(dateFin), []byte(ficheSoinsPDF), []byte(signatureAcheteur), []byte(signatureCompagnie), []byte(signatureHopital)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func checkGetExistingFicheSoins(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Get Existing FicheSoins " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("GetFicheSoinsByID"), []byte(uuid)})
	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}

	org := FicheSoins{}
	_ = json.Unmarshal(response.Payload, &org)

	if org.UUID != uuid {
		t.Fail()
	}
}

func checkUpdateFicheSoins(t *testing.T, stub *shim.MockStub, uuid string, newDateDebut string, newDateFin string, newFicheSoinsPDF string) {
	displayNewTest("checkUpdateFicheSoins")
	res := stub.MockInvoke("1", [][]byte{[]byte("UpdateFicheSoinsByID"), []byte(uuid), []byte(newDateDebut), []byte(newDateFin), []byte(newFicheSoinsPDF)})
	if res.Status != shim.OK {
		fmt.Println("UpdadeAsset", uuid, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("UpdadeAsset", uuid, "failed to get value")
		t.FailNow()
	}
}

func checkUnregisterFicheSoins(t *testing.T, stub *shim.MockStub, uuid string) {
	displayNewTest("Unregister existing Compagnie Assurance " + uuid + " From Ledger Test")

	response := stub.MockInvoke("1", [][]byte{[]byte("UnregisterFicheSoinsByID"), []byte(uuid)})

	if response.Status != shim.OK || response.Payload == nil {
		t.Fail()
	}
}

func TestCreateFicheSoins(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	//checkCreateNewFicheSoins(t, stub, "O0", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00")
	//checkGetExistingFicheSoins(t, stub, "O0")
	checkCreateNewFicheSoins(t, stub, "O0", "00", "00", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00", "00")
}

func TestGetFicheSoinsByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewFicheSoins(t, stub, "O0", "00", "00", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00", "00")
	checkGetExistingFicheSoins(t, stub, "O0")
}

func TestUpdateFicheSoinsByKey(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewFicheSoins(t, stub, "O0", "00", "00", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00", "00")

	checkUpdateFicheSoins(t, stub, "O0", "01022020", "01072020", "yyyyyyyyyy")
	checkGetExistingFicheSoins(t, stub, "O0")
}

func TestCreateFicheSoinsAlreadyExist(t *testing.T) {
	displayNewTest("Create FicheSoins Test When FicheSoins already exist")
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewFicheSoins(t, stub, "O0", "00", "00", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00", "00")
	checkCreateNewFicheSoins(t, stub, "O0", "00", "00", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00", "00")
}

func TestUnregisterFicheSoins(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkCreateNewFicheSoins(t, stub, "O0", "00", "00", "00", "00", "01012020", "01062020", "xxxxxx", "00", "00", "00")
	checkUnregisterFicheSoins(t, stub, "O0")
}
