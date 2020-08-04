package main

import (
	"encoding/json"
	"fmt"

	//_ "utils.go"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//CompagnieAssurance declaration of the struct
type CompagnieAssurance struct {
	ObjectType string
	UUID       string
	Nom        string
	Contact    string
	Adresse    string
}

func makeCompagnieAssuranceFromBytes(stub shim.ChaincodeStubInterface, bytes []byte) CompagnieAssurance {
	compagnieAssurance := CompagnieAssurance{}
	err := json.Unmarshal(bytes, &compagnieAssurance)
	panicErr(err)
	return compagnieAssurance
}

func makeBytesFromCompagnieAssurance(stub shim.ChaincodeStubInterface, compagnieAssurance CompagnieAssurance) []byte {
	bytes, err := json.Marshal(compagnieAssurance)
	panicErr(err)
	return bytes
}

//CreateCompagnieAssuranceOnLedger to create an CompagnieAssurance on ledger
func CreateCompagnieAssuranceOnLedger(stub shim.ChaincodeStubInterface, objectType string, uuid string,
	nom string, contact string, adresse string) []byte {

	compagnieAssurance := CompagnieAssurance{objectType, uuid, nom, contact, adresse}
	compagnieAssuranceAsJSONBytes := makeBytesFromCompagnieAssurance(stub, compagnieAssurance)

	uuidIndexKeyCompagnieAssurance := createIndexKey(stub, uuid, "compagnieassurance")

	putEntityInLedger(stub, uuidIndexKeyCompagnieAssurance, compagnieAssuranceAsJSONBytes)
	return compagnieAssuranceAsJSONBytes

}

//CreateCompagnieAssurance Core creation
func (t *CompagnieAssurance) CreateCompagnieAssurance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := args[0]
	nom := args[1]
	contact := args[2]
	adresse := args[3]

	uuidIndexKeyCompagnieAssurance := createIndexKey(stub, uuid, "compagnieassurance")
	if checkEntityExist(stub, uuidIndexKeyCompagnieAssurance) == true {
		return entityAlreadyExistMessage(stub, uuid, "compagnieassurance")
	}

	compagnieAssurance := CreateCompagnieAssuranceOnLedger(stub, "compagnieassurance",
		uuid, nom, contact, adresse)
	return succeed(stub, "CompagnieAssuranceCreated", compagnieAssurance)
}

//GetCompagnieAssuranceByID method to get an compagnieAssurance by id
func (t *CompagnieAssurance) GetCompagnieAssuranceByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetCompagnieAssuranceByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "compagnieassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "compagnieassurance")
	}
	compagnieAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(compagnieAssuranceAsBytes)
}

//UpdateCompagnieAssuranceByID method to update an compagnieassurance by id
func (t *CompagnieAssurance) UpdateCompagnieAssuranceByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("\n UpdateCompagnieAssuranceByID - Start")

	uuid := args[0]
	newNom := args[1]
	newContact := args[2]
	newAdresse := args[3]

	uuidIndexKey := createIndexKey(stub, uuid, "compagnieassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "compagnieassurance")
	}
	compagnieAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)
	compagnieassurance := makeCompagnieAssuranceFromBytes(stub, compagnieAssuranceAsBytes)

	compagnieassurance.Nom = newNom
	compagnieassurance.Contact = newContact
	compagnieassurance.Adresse = newAdresse

	compagnieAssuranceAsJSONBytes := makeBytesFromCompagnieAssurance(stub, compagnieassurance)

	putEntityInLedger(stub, uuidIndexKey, compagnieAssuranceAsJSONBytes)
	return succeed(stub, "CompagnieAssuranceUpdatedEvent", compagnieAssuranceAsJSONBytes)

}

//UnregisterCompagnieAssuranceByID method to unregister an compagnieassurance by id
func (t *CompagnieAssurance) UnregisterCompagnieAssuranceByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("\n UnregisterCompagnieAssuranceByID - Start")
	uuid := args[0]

	uuidIndexKey := createIndexKey(stub, uuid, "compagnieassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "compagnieassurance")
	}
	compagnieAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	if compagnieAssuranceAsBytes == nil {
		fmt.Println("Impossible to delete non-existent compagnieassurance")
		return entityNotFoundMessage(stub, uuid, "compagnieassurance")
	}

	//delete compagnieassurance
	deleteEntityFromLedger(stub, uuidIndexKey)

	fmt.Println("CompagnieAssurance " + uuid + " was unregistered successfully")
	return succeed(stub, "compagnieAssuranceUnregisteredEvent", []byte("{\"CompagnieAssuranceUUID\":\""+uuid+"\"}"))
}
