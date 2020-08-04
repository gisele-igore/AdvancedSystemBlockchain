package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//AcheteurAssurance declaration of the struct
type AcheteurAssurance struct {
	ObjectType string
	UUID       string
	Nom        string
	Contact    string
	Adresse    string
	PassportID string
	VisaID     string
}

func makeAcheteurAssuranceFromBytes(stub shim.ChaincodeStubInterface, bytes []byte) AcheteurAssurance {
	acheteurAssurance := AcheteurAssurance{}
	err := json.Unmarshal(bytes, &acheteurAssurance)
	panicErr(err)
	return acheteurAssurance
}

func makeBytesFromAcheteurAssurance(stub shim.ChaincodeStubInterface, acheteurAssurance AcheteurAssurance) []byte {
	bytes, err := json.Marshal(acheteurAssurance)
	panicErr(err)
	return bytes
}

//CreateAcheteurAssuranceOnLedger to create an AcheteurAssurance on ledger
func CreateAcheteurAssuranceOnLedger(stub shim.ChaincodeStubInterface, objectType string, uuid string,
	nom string, contact string, adresse string, passportID string, visaID string) []byte {

	acheteurAssurance := AcheteurAssurance{objectType, uuid, nom, contact, adresse, passportID, visaID}
	acheteurAssuranceAsJSONBytes := makeBytesFromAcheteurAssurance(stub, acheteurAssurance)

	uuidIdexKeyAcheteurAssurance := createIndexKey(stub, uuid, "acheteurassurance")

	putEntityInLedger(stub, uuidIdexKeyAcheteurAssurance, acheteurAssuranceAsJSONBytes)
	return acheteurAssuranceAsJSONBytes
}

//CreateAcheteurAssurance Core creation
func (t *AcheteurAssurance) CreateAcheteurAssurance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := args[0]
	nom := args[1]
	contact := args[2]
	adresse := args[3]
	passportID := args[4]
	visaID := args[5]

	uuidIndexKeyAcheteurAssurance := createIndexKey(stub, uuid, "acheteurassurance")
	if checkEntityExist(stub, uuidIndexKeyAcheteurAssurance) == true {
		return entityAlreadyExistMessage(stub, uuid, "acheteurassurance")
	}

	acheteurAssurance := CreateAcheteurAssuranceOnLedger(stub, "acheteurassurance",
		uuid, nom, contact, adresse, passportID, visaID)

	return succeed(stub, "AcheteurAssuranceCreated", acheteurAssurance)
}

//GetAcheteurAssuranceByID method to get an acheteurAssurance by id
func (t *AcheteurAssurance) GetAcheteurAssuranceByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetAcheteurAssuranceByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "acheteurassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "acheteurassurance")
	}
	acheteurAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(acheteurAssuranceAsBytes)
}

//UpdateAcheteurAssuranceByID method to update an acheteurassurance by id
func (t *AcheteurAssurance) UpdateAcheteurAssuranceByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("\n UpdateAcheteteurAssuranceByID - Start")

	uuid := args[0]
	newNom := args[1]
	newContact := args[2]
	newAdresse := args[3]

	uuidIndexKey := createIndexKey(stub, uuid, "acheteurassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "acheteurassurance")
	}
	acheteurAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)
	acheteurassurance := makeAcheteurAssuranceFromBytes(stub, acheteurAssuranceAsBytes)

	acheteurassurance.Nom = newNom
	acheteurassurance.Contact = newContact
	acheteurassurance.Adresse = newAdresse

	acheteurAssuranceAsJSONBytes := makeBytesFromAcheteurAssurance(stub, acheteurassurance)

	putEntityInLedger(stub, uuidIndexKey, acheteurAssuranceAsJSONBytes)
	return succeed(stub, "AcheteurAssuranceUpdatedEvent", acheteurAssuranceAsJSONBytes)

}

//UnregisterAcheteurAssuranceByID method to unregister an acheteurassurance by id
func (t *AcheteurAssurance) UnregisterAcheteurAssuranceByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("\n UnregisterAcheteurAssuranceByID - Start")
	uuid := args[0]

	uuidIndexKey := createIndexKey(stub, uuid, "acheteurassurance")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "acheteurassurance")
	}
	acheteurAssuranceAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	if acheteurAssuranceAsBytes == nil {
		fmt.Println("Impossible to delete non-existent acheteurassurance")
		return entityNotFoundMessage(stub, uuid, "acheteurassurance")
	}

	//delete acheteurassurance
	deleteEntityFromLedger(stub, uuidIndexKey)

	fmt.Println("AcheteurAssurance " + uuid + " was unregistered successfully")
	return succeed(stub, "acheteurAssuranceUnregisteredEvent", []byte("{\"AcheteurAssuranceUUID\":\""+uuid+"\"}"))
}
