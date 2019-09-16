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

### ~2019.09.11 Modifying'fabcar'file to 'usermanager'

**Edit Chaincode**
```
cd ~/fabric-samples/chaincode/
cp -r fabcar usermanager
mv usermanager/go/fabcar.go usermanager/go/usermanager.go
```
> rename **'fabcar'** to **'usermanager'**

**Change usermanager.go**
```
nano usermanager/go/usermanager.go
```
> change inside file's function and variables name
like this:
* Car → Person
* car → person
* Make → Firstname
* Model → Lastname
* Colour → Email
* Owner → Phone
* make → firstname
* model → lastname
* colour → email
* owner → phone
* CAR → PERSON

**Edit usermanager**
```
cd ..
cp -r fabcar usermanager
cd usermanager
```
> modify file's name (fabcar → usermanager), command: 'cp -r'

**change startFabric.sh**
```
nano startFabric.sh
```
> change whole word 'fabcar' to 'usermanager'(fabcar → usermanager)

**change query.js**
```
nano javascript/query.js
```
> changed:
- line 41: const result = await contract.evaluateTransaction('queryUser', 'USER5'); #remark
- line 41: const result = await contract.evaluateTransaction('queryAllUsers'); #add
- fabcar → User
- CAR → PERSON

**change invoke.js**
```
nano javascript/invoke.js
```
> modified:
- await contract.submitTransaction('createUser', 'USER5', 'Ahn', 'jongmin', 'ggg@ggg.com', '010-7777-7777'); #code
- fabcar → usermanager
- Car → Person
- CAR → PERSON
- Owner → Phone

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
