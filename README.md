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

#### Contract

* 환자 정보 등록 및 요청
 * 이름
 * 주민번호
 * 성별
 
 
* 환자 정보 수정 및 삭제


* 의료기록 입력
 * 수술기록
 * 진찰기록
 

### UI/UX
![MRSS구조](https://user-images.githubusercontent.com/49246977/64236333-273c7280-cf35-11e9-9fb8-f4c5cc81d94b.png)
* * *

## Let's follw this test

* * *

#### ~2019.09.11 Modifying 'fabcar' to 'MRSSsample'


#### Edit Chaincode_( copy from 'fabcar' to 'MRSSsample' in chaincode )_
```bash
cd ~/fabric-samples/chaincode/
cp -r fabcar MRSSsample
mv fabcar/go/fabcar.go MRSSsample/go/MRSSsample.go
```


##### modify MRSSsample.go
```bash
nano MRSS/go/MRSSsample.go
```
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


##### add function & add createMedical
```go
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
```

#### Edit MRSS_( copy from 'fabcar' to 'MRSSsample' )_
```bash
cd ..
cp -r fabcar MRSS
cd MRSS
```


##### modify startFabric.sh
```bash
nano startFabric.sh
```
* 'fabcar' → 'MRSSsample'


##### change query.js → queryPerson.js
```bash
nano javascript/queryPerson.js
```
* line 41: const result = await contract.evaluateTransaction('queryPerson', 'PERSON12');
* fabcar → MRSSsample
* CAR → PERSON


##### add queryAllPerson.js
```bash
nano javascript/queryAllPerson.js
```
* line 41: const result = await contract.evaluateTransaction('queryAllPersons');
* fabcar → MRSSsample
* CAR → PERSON

##### change invoke.js → personInvoke.js
```bash
nano javascript/personInvoke.js
```
* await contract.submitTransaction('createPerson', 'PERSON12', 'AhnJM', '940409', 'M');
* fabcar → MRSSsample
* Car → Person
* CAR → PERSON


##### add medicalInvoke.js
```bash
nano javascript/medicalInvoke.js
```
* await contract.submitTransaction('createMedical', 'PERSON12', 'O', 'Dimetapp, Alllergy, Tavist');
* fabcar → MRSSsample
* Car → Person
* CAR → PERSON

**RUN**
```bash
cd javascript
npm install // git clone 으로 파일들을 가져오면 .gitignore로 인해 node_modules 설치가 안되어있기 때문에 수행
cd ..
./startFabcar.sh

cd javascript
mkdir wallet // git clone 으로 파일들을 가져오면 .gitignore로 인해 wallet 디렉터리가 없기 때문에 생성
node enrollAdmin.js // admin 계정 생성
node registerUser.js // user1 계정 생성
node queryAllPersons.js // 저장된 정보 출력 확인
node medicalInvoke.js // Key값이 없을 때 정보저장이 안되는 것 확인
node personInvoke.js // 새로운 데이터 저장
node queryPerson.js // 저장된 데이터 출력 확인
node medicalInvoke.js // medical정보 저장되는 것 확인
node queryPerson.js // PERSON의 정보와 medical정보가 출력되는지 확인
node queryAllPersons.js // 이미 저장된 데이터와 새로 저장한 데이터가 출력되는지 확인
```


## Test Complete

* * *
