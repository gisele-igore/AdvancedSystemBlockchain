package main

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//ProjetAssurance principal chaincode class
type ProjetAssurance struct {
	compagnieAssurance CompagnieAssurance
	acheteurAssurance  AcheteurAssurance
	contratAssurance   ContratAssurance
	hopital            Hopital
	ficheSoins         FicheSoins
}

//Init function to Initiate the chaincode
func (t *ProjetAssurance) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("Init")
	return shim.Success([]byte("Init success"))
}

//Invoke function to invoke the chaincode
func (t *ProjetAssurance) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	defer func() {
		if r := recover(); r != nil {
			functionnalError, ok := r.(FunctionnalError)
			if ok {
				shim.Error(functionnalError.errorTag)
			}
			err, ok := r.(error)
			if ok {
				shim.Error(fmt.Sprintf("%v", err))
			} else {
				shim.Error("unknownError")
			}
		}
	}()

	fc, args := stub.GetFunctionAndParameters()

	switch {
	// COMPAGNIEASSURANCE
	case strings.Compare(fc, "CreateCompagnieAssurance") == 0:
		return t.compagnieAssurance.CreateCompagnieAssurance(stub, args)

	case strings.Compare(fc, "GetCompagnieAssuranceByID") == 0:
		return t.compagnieAssurance.GetCompagnieAssuranceByID(stub, args[0])

	case strings.Compare(fc, "UpdateCompagnieAssuranceByID") == 0:
		return t.compagnieAssurance.UpdateCompagnieAssuranceByID(stub, args)

	case strings.Compare(fc, "UnregisterCompagnieAssuranceByID") == 0:
		return t.compagnieAssurance.UnregisterCompagnieAssuranceByID(stub, args)

	// ACHETEURASSURANCE
	case strings.Compare(fc, "CreateAcheteurAssurance") == 0:
		return t.acheteurAssurance.CreateAcheteurAssurance(stub, args)

	case strings.Compare(fc, "GetAcheteurAssuranceByID") == 0:
		return t.acheteurAssurance.GetAcheteurAssuranceByID(stub, args[0])

	case strings.Compare(fc, "UpdateAcheteurAssuranceByID") == 0:
		return t.acheteurAssurance.UpdateAcheteurAssuranceByID(stub, args)

	case strings.Compare(fc, "UnregisterAcheteurAssuranceByID") == 0:
		return t.acheteurAssurance.UnregisterAcheteurAssuranceByID(stub, args)

	// CONTRATASSURANCE
	case strings.Compare(fc, "CreateContratAssurance") == 0:
		return t.contratAssurance.CreateContratAssurance(stub, args)

	case strings.Compare(fc, "GetContratAssuranceByID") == 0:
		return t.contratAssurance.GetContratAssuranceByID(stub, args[0])

	case strings.Compare(fc, "UpdateContratAssuranceByID") == 0:
		return t.contratAssurance.UpdateContratAssuranceByID(stub, args)

	case strings.Compare(fc, "UnregisterContratAssuranceByID") == 0:
		return t.contratAssurance.UnregisterContratAssuranceByID(stub, args)

	// HOPITAL
	case strings.Compare(fc, "CreateHopital") == 0:
		return t.hopital.CreateHopital(stub, args)

	case strings.Compare(fc, "GetHopitalByID") == 0:
		return t.hopital.GetHopitalByID(stub, args[0])

	// FICHESOINS
	case strings.Compare(fc, "CreateFicheSoins") == 0:
		return t.ficheSoins.CreateFicheSoins(stub, args)

	case strings.Compare(fc, "GetFicheSoinsByID") == 0:
		return t.ficheSoins.GetFicheSoinsByID(stub, args[0])

	case strings.Compare(fc, "UpdateFicheSoinsbyID") == 0:
		return t.ficheSoins.GetFicheSoinsByID(stub, args[0])

	case strings.Compare(fc, "UnregisterFicheSoinsByID") == 0:
		return t.ficheSoins.UnregisterFicheSoinsByID(stub, args)

	default:
		return shim.Error("Called function is not defined in the chaincode ")
	}

}

func main() {
	err := shim.Start(new(ProjetAssurance))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
