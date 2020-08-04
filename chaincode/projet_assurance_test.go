package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

/*
Testing UUID formats:
_____________________
Administrator UUIDs start with 'A'
Device UUIDs start with 'D'
DeviceClass UUIDs start with 'C'
Organization UUIDs start zith 'O'
PatchFile UUIDs start with 'F'
Patch UUIDs start with 'P'
*/

func displayNewTest(name string) {
	fmt.Println("\n-------\nTest: ", name, "\n-------")
}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	displayNewTest("checkInit")
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.Fail()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	displayNewTest("checkInvoke")
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.Fail()
	}
}

/*
	RUN TEST
*/
func TestExample(t *testing.T) {
	scc := new(ProjetAssurance)
	stub := shim.NewMockStub("ex02", scc)

	checkInit(t, stub, [][]byte{[]byte("init")})

}
