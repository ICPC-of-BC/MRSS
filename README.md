## MRSS (Medical Record Sharing Service)
**Hyperledger Fabric을 이용한 의료기록 공유 시스템 개발**

* * *

### 1. 팀 소개
* 안종민
* 이상아
* 이규환
* 전성원
* 명진우

### 2. 프로젝트 개요
 - 병원들끼리의 

**Contract**
```
● 환자 정보 등록 및 요청
  ○ 이름
  ○ 주민번호
  ○ 성별
● 환자 정보 수정 및 삭제
● 의료기록 입력
  ○ 수술기록
  ○ 진찰기록
```

### UI/UX


![MRSS구조](https://user-images.githubusercontent.com/49246977/64236333-273c7280-cf35-11e9-9fb8-f4c5cc81d94b.png)

* * *

## PROGRESS

* * *

### ~2019.09.11 Modifying'fabcar'file to 'MRSSsample'

**Edit Chaincode**
```
cd ~/fabric-samples/chaincode/
cp -r fabcar MRSSsample
mv fabcar/go/fabcar.go MRSSsample/go/MRSSsample.go
```
> rename **'fabcar'** to **'MRSSsample'**

**Change MRSSsample.go**
```
nano MRSS/go/MRSSsample.go
```
> change inside file's function and variables name
like this:
* Car → Person
* car → person
* Make → Name
* Model → Number
* Colour → Sex
* Owner → DorO
* add - PM
* make → name
* model → number
* colour → sex
* owner → doro
* add - pm
* CAR → PERSON

> add function
**add createMedical**

func (s *SmartContract) createMedical( APIstub shim.ChaincodeStubInterface, args []string ) sc.Response {
        tmpAsBytes, _:= APIstub.GetState(args[0])

        if tmpAsBytes == nil {
                return shim.Error( "No data" )
        }

        if len( args ) != 3 {
                return shim.Error( "Incorrect number of arguments. Expection 3" )
        }

        personAsBytes, _:= APIstub.GetState( args[0] )
        person := Person{}

        json.Unmarshal( personAsBytes, &person )
        person.DorO = args[1]
        person.PM = args[2]

        personAsBytes, _ = json.Marshal( person )
        APIstub.PutState( args[0], personAsBytes )

        return shim.Success( nil )
}


**Edit MRSS**
```
cd ..
cp -r fabcar MRSS
cd MRSS
```
> modify file's name (fabcar → MRSS), command: 'cp -r'

**change startFabric.sh**
```
nano startFabric.sh
```
> change whole word 'fabcar' to 'MRSS'(fabcar → MRSS)

**change query.js → queryPerson.js**
```
nano javascript/queryPerson.js
```
> changed:
- line 41: const result = await contract.evaluateTransaction('queryPerson', 'USER5');
- fabcar → MRSSsample
- CAR → PERSON

**add queryAllPerson.js**
```
nano javascript/queryAllPerson.js
```
> changed:
- line 41: const result = await contract.evaluateTransaction('queryAllPersons');
- fabcar → MRSSsample
- CAR → PERSON

**change invoke.js → personInvoke.js**
```
nano javascript/personInvoke.js
```
> modified:
- await contract.submitTransaction('createPerson', 'PERSON12', 'AhnJM', '940409', 'M');
- fabcar → MRSSsample
- Car → Person
- CAR → PERSON

**add medicalInvoke.js**
```
nano javascript/medicalInvoke.js
```
> modified:
- await contract.submitTransaction('createMedical', 'PERSON12', 'O', 'Dimetapp, Alllergy, Tavist');
- fabcar → MRSSsample
- Car → Person
- CAR → PERSON

**RUN**
```
cd javascript
npm install
cd ..
./startFabcar.sh

cd javascript
node enrollAdmin.js
node registerUser.js
node query.js
node invoke.js
```
> **!!!CAUTION!!!**
You Must Follow This Order

* * *
