## MRSS (Medical Record Sharing Service)
**Hyperledger Fabric을 이용한 의료기록 공유 시스템 개발**

### 1. 팀 소개
* 안종민
* 이상아
* 이규환
* 전성원
* 명진우

### 2. 프로젝트 개요
 병원들끼리의 환자의 질병에 대한 데이터 공유를 가능하게 함  
 잘못된 데이터 전달로 인한 환자 상태 악화 및 사망 등을 방지하기 위함(사례: 방광암 환자)

#### Contract

* 환자 정보 입력
* 환자 정보 수정 및 삭제
* 의료 기록 입력
* 선택 환자 정보 출력
* 전체 환자 정보 출력

### UI/UX
![MRSS구조](https://user-images.githubusercontent.com/49246977/64236333-273c7280-cf35-11e9-9fb8-f4c5cc81d94b.png)
* * *


## Let's follow this test

 **2019.09.06 Start Project: Announce Prologue**  
 **2019.09.11 Modifying 'fabcar' to 'MRSSsample'**  
 **2019.09.12 Implemented function to Input or Output medical data**  


### Edit Chaincode_( copy from 'fabcar' to 'MRSSsample' in chaincode )_
```bash
cd ~/fabric-samples/chaincode/
cp -r fabcar MRSSsample
mv fabcar/go/fabcar.go MRSSsample/go/MRSSsample.go
```

* **modify MRSSsample.go**
```bash
nano MRSS/go/MRSSsample.go
```
Car → Person<br>
car → person<br>
Make → Name<br>
Model → Number<br>
Colour → Sex<br>
Owner → DorO<br>
make → name<br>
model → number<br>
colour → sex<br>
owner → doro<br>
CAR → PERSON<br>
add - pm<br>
add - PM<br>

* **Add function: createMedical**
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




### Edit MRSS_( copy from 'fabcar' to 'MRSSsample' )_
```bash
cd ..
cp -r fabcar MRSS
cd MRSS
```

* **modify startFabric.sh**
```bash
nano startFabric.sh
```
'fabcar' → 'MRSSsample'<br>

* **change query.js → queryPerson.js**
```bash
nano javascript/queryPerson.js
```
line 41: const result = await contract.evaluateTransaction('queryPerson', 'PERSON12');<br>
fabcar → MRSSsample<br>
CAR → PERSON<br>

* **add queryAllPerson.js**
```bash
nano javascript/queryAllPerson.js
```
line 41: const result = await contract.evaluateTransaction('queryAllPersons');<br>
fabcar → MRSSsample<br>
CAR → PERSON<br>

* **change invoke.js → personInvoke.js**
```bash
nano javascript/personInvoke.js
```
await contract.submitTransaction('createPerson', 'PERSON12', 'AhnJM', '940409', 'M');<br>
fabcar → MRSSsample<br>
Car → Person<br>
CAR → PERSON<br>

* **add medicalInvoke.js**
```bash
nano javascript/medicalInvoke.js
```
await contract.submitTransaction('createMedical', 'PERSON12', 'O', 'Dimetapp, Alllergy, Tavist');<br>
fabcar → MRSSsample<br>
Car → Person<br>
CAR → PERSON<br>



### RUN
```bash
cd javascript
npm install // git clone 으로 파일들을 가져오면 .gitignore로 인해 node_modules 설치가 안되어있기 때문에 수행
cd ..
./startFabcar.sh

cd javascript
mkdir wallet // git clone 으로 파일들을 가져오면 .gitignore로 인해 wallet 디렉터리가 없기 때문에 생성
node enrollAdmin.js // admin 계정 생성
```
![enrol](https://user-images.githubusercontent.com/49246977/64962469-79d14380-d8d2-11e9-83a6-a5d7b514f93c.PNG)
```bash
node registerUser.js // user1 계정 생성
```
![register](https://user-images.githubusercontent.com/49246977/64962662-c3ba2980-d8d2-11e9-8361-2af06f30da9b.PNG)
```bash
node queryAllPersons.js // 저장된 정보 출력 확인
```
![queryAllPerson](https://user-images.githubusercontent.com/49246977/64962562-9f5e4d00-d8d2-11e9-8581-253b7da50056.PNG)
```bash
node medicalInvoke.js // Key값이 없을 때 정보저장이 안되는 것 확인
```
![medicalInvoke_fail](https://user-images.githubusercontent.com/49246977/64962503-85bd0580-d8d2-11e9-9718-322528e326b6.PNG)
```bash
node personInvoke.js // 새로운 데이터 저장
```
![personInvoke](https://user-images.githubusercontent.com/49246977/64962543-966d7b80-d8d2-11e9-91f5-5636c284404e.PNG)

```bash
node queryPerson.js // 저장된 데이터 출력 확인
```
![queryPerson_no_medical](https://user-images.githubusercontent.com/49246977/64962654-be5cdf00-d8d2-11e9-86d7-2a3e6196141b.PNG)
```bash
node medicalInvoke.js // medical정보 저장되는 것 확인
```
![medicalInvoke_success](https://user-images.githubusercontent.com/49246977/64962526-8eadd700-d8d2-11e9-8d65-f49eab060088.PNG)
```bash
node queryPerson.js // PERSON의 정보와 medical정보가 출력되는지 확인
```
![queryPerson_add_medical](https://user-images.githubusercontent.com/49246977/64962619-b69d3a80-d8d2-11e9-8e57-59c9ba5e077d.PNG)
```bash
node queryAllPersons.js // 이미 저장된 데이터와 새로 저장한 데이터가 출력되는지 확인
```
![queryAllPerson_add_PERSON12](https://user-images.githubusercontent.com/49246977/64962587-a8e7b500-d8d2-11e9-88cc-9d06aa924785.PNG)

* * *

## Test Complete!!!
