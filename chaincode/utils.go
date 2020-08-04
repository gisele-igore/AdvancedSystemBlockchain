package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func panicErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func putEntityInLedger(stub shim.ChaincodeStubInterface, uuid string, payload []byte) {
	/* creatorID := getTransactionCreator(stub) */

	err := stub.PutState(uuid, payload)
	panicErr(err)
	fmt.Println("Put entity " + uuid + " in ledger")
}

func getEntityFromLedger(stub shim.ChaincodeStubInterface, uuid string) []byte {
	entityAsBytes, err := stub.GetState(uuid)
	if err != nil {
		errorMessage := "Failed to get state for Entity with uuid " + uuid
		panicErr(errors.New(errorMessage))
	} else if entityAsBytes == nil {
		fmt.Println("Entitypb " + uuid + " not found. Returningss nil...")
		entityAsBytes = nil
		return entityAsBytes
	}

	fmt.Println("GET Entity with uuid " + uuid + " returned: ")
	fmt.Println(string(entityAsBytes))

	return entityAsBytes
}

func deleteEntityFromLedger(stub shim.ChaincodeStubInterface, uuid string) {
	err := stub.DelState(uuid)
	panicErr(err)
	fmt.Println("DelState done")
}

func succeed(stub shim.ChaincodeStubInterface, eventMessage string, eventPayload []byte) pb.Response {
	stub.SetEvent(eventMessage, eventPayload)
	m := make(map[string]interface{})
	if err := json.Unmarshal(eventPayload, &m); err != nil {
		panic(err)
	}
	test, err1 := json.MarshalIndent(m, "", " ")
	panicErr(err1)

	fmt.Printf("\n m: %q", m)

	return shim.Success(test)
}

func createIndexKey(stub shim.ChaincodeStubInterface, uuid string, objectType string) string {

	indexName := "objectType~uuid"
	uuidIndexKey, err := stub.CreateCompositeKey(indexName, []string{objectType, uuid})
	panicErr(err)
	return uuidIndexKey
}

func checkEntityExist(stub shim.ChaincodeStubInterface, uuid string) bool {
	response, err := stub.GetState(uuid)
	panicErr(err)

	if response == nil {
		return false
	}
	return true
}

//FunctionnalError definition of struct
type FunctionnalError struct {
	errorTag string
	errorMsg string
}

func (e *FunctionnalError) Error() string {
	return e.errorTag
}

//NewFunctionnalError method to define new functionnal error
func NewFunctionnalError(errorTag string, errorMsg string) *FunctionnalError {
	return &FunctionnalError{errorTag: errorTag, errorMsg: errorMsg}
}

func entityAlreadyExistMessage(stub shim.ChaincodeStubInterface, uuid string, objectType string) pb.Response {

	errorMessage := "Entity with uuid " + uuid + " already exists in the ledger."
	errorType := ""
	switch objectType {
	case "compagnieassurance":
		errorType = "compagnieAssuranceAlreadyExist"
	/* case "compagnieAssurance":
	errorType = "compagnieAssuranceAlreadyExist" */
	case "hopital":
		errorType = "hopitalAlreadyExist"
	}

	err := NewFunctionnalError(errorType, errorMessage)
	fmt.Println(errorMessage)
	return shim.Error(err.Error())
}

func entityNotFoundMessage(stub shim.ChaincodeStubInterface, uuid string, objectType string) pb.Response {
	errorMessage := objectType + " with uuid " + uuid + " not found in the ledger."
	errorType := ""
	switch objectType {
	case "compagnieassurance":
		errorType = "noCompagnieAssuranceFound"
		/* case "organizations":
		errorType = "noOrganizationFound" */
	}

	err := NewFunctionnalError(errorType, errorMessage)
	fmt.Println(errorMessage)
	return shim.Error(err.Error())

}
