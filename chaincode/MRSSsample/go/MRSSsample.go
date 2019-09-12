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

// Define the Person structure, with 4 properties.  Structure tags are used by encoding/json library
type Person struct {
	Name   string `json:"name"`
	Number  string `json:"number"`
	Sex string `json:"sex"`
}

type Medical struct {
	Operation string `json:"operation"`
	Op_Medication string `json:"op_medication"`
	Diagnosis string `json:"diagnosis"`
	Diag_Medication string `json:"diag_medication"`
}

/*
 * The Init method is called when the Smart Contract "fabPerson" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabPerson"
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

	} else if function == "queryMedical" {
                return s.queryPerson(APIstub, args)
	} else if function == "initLedgerMedical" {
                return s.initLedger(APIstub)
	} else if function == "createMedical" {
                return s.createPerson(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryPerson(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	PersonAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(PersonAsBytes)
}
/*
func (s *SmartContract) queryMedical(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 1")
        }

        PersonAsBytes, _ := APIstub.GetState(args[0])
        return shim.Success(PersonAsBytes)
}

*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	Persons := []Person{
		Person{Name: "Toyota", Number: "111111-1111111", Sex: "blue"},
		Person{Name: "Ford", Number: "222222-2222222", Sex: "red"},
		Person{Name: "Hyundai", Number: "333333-3333333", Sex: "green"},
		Person{Name: "Volkswagen", Number: "444444-4444444", Sex: "yellow"},
		Person{Name: "Tesla", Number: "111112-111111", Sex: "black"},
		Person{Name: "Peugeot", Number: "111112-111112", Sex: "purple"},
		Person{Name: "Chery", Number: "222222-2222223", Sex: "white"},
		Person{Name: "Fiat", Number: "222222-2222244", Sex: "violet"},
		Person{Name: "Tata", Number: "222222-2222555", Sex: "indigo"},
		Person{Name: "Holden", Number: "881212-2111111", Sex: "brown"},
	}

	i := 0
	for i < len(Persons) {
		fmt.Println("i is ", i)
		PersonAsBytes, _ := json.Marshal(Persons[i])
		APIstub.PutState("Person"+strconv.Itoa(i), PersonAsBytes)
		fmt.Println("Added", Persons[i])
		i = i + 1
	}

	return shim.Success(nil)
}
/*
func (s *SmartContract) initLedgerMedical(APIstub shim.ChaincodeStubInterface) sc.Response {
        Medicals := []Medical{
                Medical{Operation: "Cancer surgery", Op_Medication: "morphine", Diagnosis: "cold", Diag_Medication: "Brufen"},
                Medical{Operation: "Cancer surgery", Op_Medication: "morphine", Diagnosis: "cold", Diag_Medication: "Brufen"},
                Medical{Operation: "Cancer surgery", Op_Medication: "morphine", Diagnosis: "cold", Diag_Medication: "Brufen"},
                Medical{Operation: "Cancer surgery", Op_Medication: "morphine", Diagnosis: "cold", Diag_Medication: "Brufen"},
        }

        i := 0
        for i < len(Medicals) {
                fmt.Println("i is ", i)
                PersonAsBytes, _ := json.Marshal(Medicals[i])
                APIstub.PutState("Person"+strconv.Itoa(i), PersonAsBytes)
                fmt.Println("Added", Medicals[i])
                i = i + 1
        }

        return shim.Success(nil)
}
*/
func (s *SmartContract) createPerson(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var Person = Person{Name: args[1], Number: args[2], Sex: args[3]}

	PersonAsBytes, _ := json.Marshal(Person)
	APIstub.PutState(args[0], PersonAsBytes)

	return shim.Success(nil)
}

/*
func (s *SmartContract) createMedical(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

        if len(args) != 4 {
                return shim.Error("Incorrect number of arguments. Expecting 5")
        }

        //var Person = Person{Name: args[1], Number: args[2], Sex: args[3]}
        var Medical = Medical{Operation: args[1], Op_Medication: args[2], Diagnosis: args[3], Diag_Medication: args[4]}
        MedicalAsBytes, _ := json.Marshal(Medical)
        APIstub.PutState(args[0], MedicalAsBytes)

        return shim.Success(nil)
}
*/


func (s *SmartContract) queryAllPersons(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "Person0"
	endKey := "Person999"

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

	PersonAsBytes, _ := APIstub.GetState(args[0])
	Person := Person{}

	json.Unmarshal(PersonAsBytes, &Person)
	Person.Owner = args[1]

	PersonAsBytes, _ = json.Marshal(Person)
	APIstub.PutState(args[0], PersonAsBytes)

	return shim.Success(nil)
}
*/
// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
