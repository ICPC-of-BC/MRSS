/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the person structure, with 4 properties.  Structure tags are used by encoding/json library
type Person struct {
	Name   string `json:"name"`
	Number  string `json:"number"`
	Sex string `json:"sex"`
}

type Medical struct {
	DorO string `json:"doro"`
	PM string `json:"pm"`
}

/*
 * The Init method is called when the Smart Contract "MRSSsample" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "MRSSsample"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryPerson" {
		return s.queryPerson(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createPerson" {
		return s.createPerson(APIstub, args)
	} else if function == "queryAllPersons" {
		return s.queryAllPersons(APIstub)
//	} else if function == "changePersonOwner" {
//		return s.changePersonOwner(APIstub, args)
	} else if function == "createMedical" {
		return s.createMedical(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryPerson(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	personAsBytes, _ := APIstub.GetState(args[0])
//	medicalAsBytes, _ := APIstub.GetState( args[0] )
//	fmt.Printf( "%s\n%s\n", personAsBytes, medicalAsBytes )
	fmt.Printf( "%s\n", personAsBytes )

	return shim.Success( personAsBytes )
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	persons := []Person{
		Person{Name: "Toyota", Number: "Prius", Sex: "blue"},
		Person{Name: "Ford", Number: "Mustang", Sex: "red"},
		Person{Name: "Hyundai", Number: "Tucson", Sex: "green"},
	}

//	medicals := []Medical {
//		Medical{ DorO: "D", PM: "Advil, Anacin, Bayer, Bufferin"},
//		Medical{ DorO: "D", PM: "Halls, Robitussin, Sucrets, Vicks"},
//		Medical{ DorO: "O", PM: "GasX, Maalox, Rolaid, Tums"},
//	}

	i := 0
	for i < len(persons) {
		fmt.Println("i is ", i)
		personAsBytes, _ := json.Marshal(persons[i])
//		medicalAsBytes, _ := json.Marshal( medicals[i] )
		APIstub.PutState("PERSON"+strconv.Itoa(i), personAsBytes)
//		APIstub.PutState( "PERSON"+strconv.Itoa( i ), medicalAsBytes )
		fmt.Println("Added", persons[i])
//		fmt.Println( "Added", medicals[i] )
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createPerson(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var person = Person{Name: args[1], Number: args[2], Sex: args[3]}
//	var medical = Medical{}

	personAsBytes, _ := json.Marshal(person)
	APIstub.PutState(args[0], personAsBytes)
//	medicalAsBytes, _ := json.Marshal( medical )
//	APIstub.PutState( args[0], medicalAsBytes )

	return shim.Success(nil)
}

func (s *SmartContract) queryAllPersons(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "PERSON0"
	endKey := "PERSON999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllPersons:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
/*
func (s *SmartContract) changePersonOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	personAsBytes, _ := APIstub.GetState(args[0])
	person := Person{}

	json.Unmarshal(personAsBytes, &person)
	person.Owner = args[1]

	personAsBytes, _ = json.Marshal(person)
	APIstub.PutState(args[0], personAsBytes)

	return shim.Success(nil)
}
*/

func (s *SmartContract) createMedical( APIstub shim.ChaincodeStubInterface, args []string ) sc.Response {

	if len( args ) != 3 {
		return shim.Error( "Incorrect number of arguments. Expection 3" )
	}

	var medical = Medical{ DorO: args[1], PM: args[2] }
	medicalAsBytes, _ := json.Marshal( medical )
	APIstub.PutState(args[0], medicalAsBytes)

	return shim.Success( nil )
}


// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
