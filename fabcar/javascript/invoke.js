/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { FileSystemWallet, Gateway } = require('fabric-network'); //fabric-network라는 SDK를 사용할것이다. 하는 내용이 여기에! 
const path = require('path');
console.log(path);

const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');

async function main() {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('user1');	//user1의 프라이빗 키를 가져와서
        if (!userExists) {			//user1이 없으면 if문 코드, error를 밷어낸다. 있으면 if문 다음 코드로 넘어가!
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });
	//user1이 접속한다

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');
	// mychannel이란 정보를 가져온다.

        // Get the contract from the network.
        const contract = network.getContract('fabcar');
	// mychannel이란 네트워크로 부터 fabcar라는 컨트랙트를 호출해온다.

        // Submit the specified transaction.
        // createCar transaction - requires 5 argument, ex: ('createCar', 'CAR12', 'Honda', 'Accord', 'Black', 'Tom')
        // changeCarOwner transaction - requires 2 args , ex: ('changeCarOwner', 'CAR10', 'Dave')
        await contract.submitTransaction('changeCarOwner', 'CAR12','SANGA');
        console.log('Transaction has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

main();
