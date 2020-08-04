package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//Hopital declaration of structure
type Hopital struct {
	ObjectType string
	UUID       string
	Nom        string
	Contact    string
	Adresse    string
}

func makeHopitalFromBytes(stub shim.ChaincodeStubInterface, bytes []byte) Hopital {
	hopital := Hopital{}
	err := json.Unmarshal(bytes, &hopital)
	panicErr(err)
	return hopital
}

func makeBytesFromHopital(stub shim.ChaincodeStubInterface, hopital Hopital) []byte {
	bytes, err := json.Marshal(hopital)
	panicErr(err)
	return bytes
}

//CreateHopitalOnLedger to create an Hopital on ledger
func CreateHopitalOnLedger(stub shim.ChaincodeStubInterface, objectType string, uuid string,
	nom string, contact string, adresse string) []byte {

	hopital := Hopital{objectType, uuid, nom, contact, adresse}
	hopitalAsJSONBytes := makeBytesFromHopital(stub, hopital)

	uuidIndexKeyHopital := createIndexKey(stub, uuid, "hopital")

	putEntityInLedger(stub, uuidIndexKeyHopital, hopitalAsJSONBytes)
	return hopitalAsJSONBytes
}

//CreateHopital Core creation
func (t *Hopital) CreateHopital(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := args[0]
	nom := args[1]
	contact := args[2]
	adresse := args[3]

	uuidIndexKeyHopital := createIndexKey(stub, uuid, "hopital")
	if checkEntityExist(stub, uuidIndexKeyHopital) == true {
		return entityAlreadyExistMessage(stub, uuid, "hopital")
	}

	hopital := CreateHopitalOnLedger(stub, "Hopital", uuidIndexKeyHopital, nom, contact, adresse)
	return succeed(stub, "HopitalCreated", hopital)
}

//GetHopitalByID method to get an hopital by id
func (t *Hopital) GetHopitalByID(stub shim.ChaincodeStubInterface, args string) pb.Response {
	fmt.Println("\n GetHopitalByID - Start", args)

	uuid := args

	uuidIndexKey := createIndexKey(stub, uuid, "hopital")
	if checkEntityExist(stub, uuidIndexKey) == false {
		return entityNotFoundMessage(stub, uuid, "hopital")
	}
	hopitalAsBytes := getEntityFromLedger(stub, uuidIndexKey)

	return shim.Success(hopitalAsBytes)
}
