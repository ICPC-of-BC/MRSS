## MRSS( Medical Record Sharing Service )
**Hyperledger Fabric을 이용한 의료기록 공유 시스템**

### 팀 소개
* 안종민
* 이상아
* 이규환
* 전성원
* 명진우


### 프로젝트 개요
 현재 우리나라의 병원들끼리는 정보공유가 되지 않고있다. 따라서 환자는 병원을 옮길 때마다 자신이 직접 이전 병원에서
 수술, 진료 기록을 떼어 옮기려는 병원에 내야 하는 어려움이 있는데 기록을 받으려면 어느정도의 돈도 지불해야한다.
 만약 수술, 진료 기록을 해당 병원에 알려주지 않으면 그 병원에서는 이 환자가 어떠한 이력을 가지고 있는지, 어떤 수술을 받았었는지 등을
 알 수가 없다. 그로인해 실제 환자가 사망하는 사례도 적지 않으며 병이 악화된 사례는 더욱 많다.


### Contract
* 환자 정보 입력
* 환자 정보 수정 및 삭제
* 의료 기록 입력
* 선택 환자 정보 출력
* 전체 환자 정보 출력


### 사용하는 util
* docker v17.06.2 이상
* npm 6.12.0
* python 2.7.x
* node v12.13.0


### UI/UX
![그림01](https://user-images.githubusercontent.com/49246977/65002812-779dd200-d930-11e9-9a25-ff8e6e071313.png)


* * *


## Let's follow this test !!
 **2019.09.06 Start Project: Announce Prologue**  
 **2019.09.11 Modifying 'fabcar' to 'MRSSsample'**  
 **2019.09.12 Implemented function to Input or Output medical data**  


### 1. Edit Chaincode( copy from 'fabcar' to 'MRSSsample' in chaincode )
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


### 2. Edit MRSS( copy from 'fabcar' to 'MRSSsample' )
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


### 3. RUN
```bash
git clone https://github.com/ICPC-of-BC/MRSS.git
cd MRSS

curl -sSL https://goo.gl/6wtTN5 | bash -s // 현재 이미지가 없는 셋팅이기 때문에 docker image 들을 다운받아야 한다.

cd MRSS
cd javascript
npm install // git clone 으로 파일들을 가져오면 .gitignore로 인해 node_modules 설치가 안되어있기 때문에 수행
cd ..
./startFabcar.sh

cd javascript
mkdir wallet // git clone 으로 파일들을 가져오면 wallet디렉터리 안에 데이터가 있을 수 있는데 모두 지워 빈 상태에서 실습을 진행해야한다.
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
```bash
node delPerson.js // 새로 저장된 데이터 지우기
```
![delPerson](https://user-images.githubusercontent.com/49246977/65012546-aa58c200-d952-11e9-806a-6522740da804.PNG)
```bash
node queryPerson.js // 데이터 삭제됐는지 확인
```
![queryPerson_del_PERSON12](https://user-images.githubusercontent.com/49246977/65012557-b6dd1a80-d952-11e9-879a-7f70a6163e4f.PNG)
```bash
node queryAllPersons.js // 전체 데이터에서 삭제 됐는지 확인
```
![queryAllPerson_del_PERSON12](https://user-images.githubusercontent.com/49246977/65012567-c3617300-d952-11e9-98bf-81b5826bdfdf.PNG)

## Test Complete!!!

* * *

## After Plans
본래에 병원 데이터를 관리하고 있는 데이터관리업체와 협업하여 org로 등록할 병원의 수를 늘리고 현직에 종사하는 병원관계자와 피드백을 받아서 편리성을 증진할 시스템을 구축 및 의료기관간의 정보 공유를 확장할 것을 목표로 하고 있다.  

