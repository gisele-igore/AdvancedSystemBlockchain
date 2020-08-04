package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//ContratAssurance declaration of the struct
type ContratAssurance struct {
	ObjectType            string
	UUID                  string
	IDCompagnieAssurance  CompagnieAssurance
	CodeAcheteurAssurance AcheteurAssurance
	DateDebut             string
	DateFin               string
	ContratAssurancePDF   string
	SignatureAcheteur     string
	SignatureCompagnie    string
}

func makeContratAssuranceFromBytes(stub shim.ChaincodeStubInterface, bytes []byte) ContratAssurance {
	contratAssurance := ContratAssurance{}
	err := json.Unmarshal(bytes, &contratAssurance)
	panicErr(err)
	return contratAssurance
}

func makeBytesFromContratAssurance(stub shim.ChaincodeStubInterface, contratAssurance ContratAssurance) []byte {
	bytes, err := json.Marshal(contratAssurance)
	panicErr(err)
	return bytes
}

//CreateContratAssuranceOnLedger to create an CompagnieAssurance on ledger
func CreateContratAssuranceOnLedger(stub shim.ChaincodeStubInterface, objectType string, uuid string,
	iDCompagnieAssurance CompagnieAssurance, codeAcheteurAssurance AcheteurAssurance, dateDebut string,
	dateFin string, contratAssurancePDF string, signatureAcheteur string, signatureCompagnie string) []byte {

	contratAssurance := ContratAssurance{objectType, uuid, iDCompagnieAssurance, codeAcheteurAssurance,
		dateDebut, dateFin, contratAssurancePDF, signatureAcheteur, signatureCompagnie}
	contratAssuranceAsJSONBytes := makeBytesFromContratAssurance(stub, contratAssurance)

	uuidIndexKeyContratAssurance := createIndexKey(stub, uuid, "contratassurance")

	putEntityInLedger(stub, uuidIndexKeyContratAssurance, contratAssuranceAsJSONBytes)
	return contratAssuranceAsJSONBytes
}

//CreateContratAssurance Core creation
func (t *ContratAssurance) CreateContratAssurance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := args[0]
	iDCompagnieAssurance := args[1]
	codeAcheteurAssurance := args[2]
	dateDebut := args[3]
	dateFin := args[4]
	contratAssurancePDF := args[5]
	signatureAcheteur := args[6]
	signatureCompagnie := args[7]

	uuidIndexKeyContratAssurance := createIndexKey(stub, uuid, "contratassurance")
	if checkEntityExist(stub, uuidIndexKeyContratAssurance) == true {
		return entityAlreadyExistMessage(stub, uuid, "contratassurance")
	}

	uuidIndexKeyCompagnieAssurance := createIndexKey(stub, iDCompagnieAssurance, "compagnieassurance")
	if checkEntityExist(stub, uuidIndexKeyCompagnieAssurance) == true {
		return entityAlreadyExistMessage(stub, iDCompagnieAssurance, "compagnieassurance")
	}
	compagnieAssurance := getEntityFromLedger(stub, uuidIndexKeyCompagnieAssurance)
	compagnieAssuranceAsJSONBytes := makeCompagnieAssuranceFromBytes(stub, compagnieAssurance)

	uuidIndexKeyAcheteurAssurance := createIndexKey(stub, codeAcheteurAssurance, "acheteurassurance")
	if checkEntityExist(stub, uuidIndexKeyAcheteurAssurance) == true {
		return entityAlreadyExistMessage(stub, codeAcheteurAssurance, "acheteurassurance")
	}
	acheteurAssurance := getEntityFromLedger(stub, uuidIndexKeyAcheteurAssurance)
	acheteurAssuranceAsJSONBytes := makeAcheteurAssuranceFromBytes(stub, acheteurAssurance)

	contratAssurance := CreateContratAssuranceOnLedger(stub, "contratassurance",
		uuid, compagnieAssuranceAsJSONBytes, acheteurAssuranceAsJSONBytes, dateDebut, dateFin,
		contratAssurancePDF, signatureAcheteur, signatureCompagnie)

	return succeed(stub, "ContratAssuranceCreated", contratAssurance)
}

//GetContratAssuranceByID method to get an contratAssurance by id
func (t *ContratAssurance) GetContratAssuranceByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetContratAssuranceByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "contratassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "contratassurance")
	}
	contratAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(contratAssuranceAsBytes)
}

//UpdateContratAssuranceByID method to update an contratassurance by id
func (t *ContratAssurance) UpdateContratAssuranceByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("\n UpdateContratAssuranceByID - Start")

	uuid := args[0]
	newDateDebut := args[1]
	newDateFin := args[2]
	newContratAssurancePDF := args[3]

	uuidIndexKey := createIndexKey(stub, uuid, "contratassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "contratassurance")
	}
	contratAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)
	contratassurance := makeContratAssuranceFromBytes(stub, contratAssuranceAsBytes)

	contratassurance.DateDebut = newDateDebut
	contratassurance.DateFin = newDateFin
	contratassurance.ContratAssurancePDF = newContratAssurancePDF

	contratAssuranceAsJSONBytes := makeBytesFromContratAssurance(stub, contratassurance)

	putEntityInLedger(stub, uuidIndexKey, contratAssuranceAsJSONBytes)
	return succeed(stub, "ContratAssuranceUpdatedEvent", contratAssuranceAsJSONBytes)

}

//UnregisterContratAssuranceByID method to unregister an contratassurance by id
func (t *ContratAssurance) UnregisterContratAssuranceByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("\n UnregisterContratAssuranceByID - Start")
	uuid := args[0]

	uuidIndexKey := createIndexKey(stub, uuid, "contratassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "contratassurance")
	}
	contratAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	if contratAssuranceAsBytes == nil {
		fmt.Println("Impossible to delete non-existent contratassurance")
		return entityNotFoundMessage(stub, uuid, "contratassurance")
	}

	//delete contratsassurance
	deleteEntityFromLedger(stub, uuidIndexKey)

	fmt.Println("ContratAssurance " + uuid + " was unregistered successfully")
	return succeed(stub, "contratAssuranceUnregisteredEvent", []byte("{\"ContratAssuranceUUID\":\""+uuid+"\"}"))
}
