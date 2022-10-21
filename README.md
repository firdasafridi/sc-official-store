# Prerequites


### NodeJS & NPM 
personaly i used Node.js v18.6.0
### Truffle

`npm install -g truffle`

You may need to run this if on MacOS or Linux

`sudo npm install -g truffle`

### Ganace

The installation can be found [here](https://trufflesuite.com/ganache/)

### Install  Solidity
Install Solidity [here](https://docs.soliditylang.org/en/v0.8.2/installing-solidity.html)
```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

### Geth
Install Geth [here](https://geth.ethereum.org/docs/install-and-build/installing-geth)
```
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum
```
### Golang
The installation can be found [here](https://go.dev/doc/install)

go version
```bash
> go version
> go version go1.18.4 darwin/amd64
```

# Let's do some coding
## Step #1: Init your project
`truffle init`

## Step #2: Create the contract
Under contracts create OfficialStore.sol
```solidity
// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;

import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

contract OfficialStore is Ownable {

    /*
    * ContractBrand is structure of brand and contract definition.
    * `name` will represent name of toko that assign as principle from Brand
    * `mainBrandAddress` will represent contract brand address
    * `district` will represent location of brand distribution
    * `expiredDate` will represent unit timestamp the brand will active
    * `active` will represent activation brand as principle official store
    */
    struct ContractBrand {
        string name;
        address mainBrandAddress;
        string district;
        uint256 expiredDate;
        bool active;
    }

    // _mainAddressBrand represent wallet address of brand assign
    mapping (address => string) private _mainAddressBrand;

    // _contractBrandList represent list contract brand from the partners
    mapping (address => ContractBrand) private _contractBrandList;

    constructor() {
    }

    /* 
    * setBrandContract will set an address as official store brand name
    * mainBrandAddress is destination the address assign as brand
    * brandName will represent brand name
    */
    function setBrandContract(address mainBrandAddress, string memory brandName) external onlyOwner() {
        _mainAddressBrand[mainBrandAddress] = brandName;
    }

    /* 
    * setContractOS will set an address as principle official store brand
    * contractAddress is destination the address assign as official store
    * contractBrandInfo will keep information of brand info
    */
    function setContractOS(address contractAddress, ContractBrand memory contractBrandInfo) external onlyOwner() {
        _contractBrandList[contractAddress] = contractBrandInfo;
    }

    /* 
    * enableContractOS will set active existing address to activate the brand official store
    * contractAddress is destination the address assign enable as the official store
    */
    function enableContractOS(address contractAddress) external onlyOwner() {
        _contractBrandList[contractAddress].active = true;
    }

    /* 
    * disableContractOS will set inactive existing address to activate the brand official store
    * contractAddress is destination the address to disable as the official store
    */
    function disableContractOS(address contractAddress) external onlyOwner() {
        _contractBrandList[contractAddress].active = false;
    }

    /* 
    * extendContractOS will extend the contract as official store
    * contractAddress is destination the address assign to extend
    * ts is unix timestamp extend the contract address
    */
    function extendContractOS(address contractAddress, uint256 ts) external onlyOwner() {
        _contractBrandList[contractAddress].expiredDate = ts;
    }

    /* 
    * getBrandFromAddress will show information brand
    * will return brand name information
    */
    function getBrandFromAddress(address brandAddress) public view returns (string memory brandName) {
        return _mainAddressBrand[brandAddress];
    }

    /* 
    * getContractBrandFromAddress will show information contract brand and principle
    */
    function getContractBrandFromAddress(address contractAddress) public view returns (ContractBrand memory) {
        return _contractBrandList[contractAddress];
    }

    /* 
    * isContractBrandActive will check contract address still active or not as official store.
    * contractAddress will represent the contract address of principle official store.
    * brandName will contain check what brand assign to the contract.
    * unixTimeNow will contain availability of time that assign.
    */
    function isContractBrandActive(address contractAddress, string memory brandName, uint256 unixTimeNow) public view returns (string memory active) {
        if (!_contractBrandList[contractAddress].active) {
            return "Contract Brand Inactive";
        }

        if (!compareStrings(_mainAddressBrand[_contractBrandList[contractAddress].mainBrandAddress], brandName)) {
            return "Company not official store";
        }

        if (_contractBrandList[contractAddress].expiredDate < unixTimeNow) {
            return "Company Inactive";
        }

        return "ACTIVE";
    }

    // compareStrings from the brand
    function compareStrings(string memory a, string memory b) public pure returns (bool) {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
    }
}
```


## Step #3: Install openzeppelin contracts
openzeppelin is a library modular, reusable, and secure smart contracts, written in Solidity

```npm install @openzeppelin/contracts```

## Step #4: Compile using truffle
Inside *truffle-config.js*
```javascript
module.exports = {
  compilers: {
    solc: {
      version: "0.8.1",   
    }
  },

};
```
Now back to terminal and run:

`truffle compile`

Now, we have build folder with `json` file on it.


## Step #5: Deploying to ganace

Open the ganace and you will find like below. Just click *QUICKSTART* button.
![Alt text](img/1.png?raw=true "Open Ganace")

It should open a screen with 10 wallets.
![Alt text](img/3.png?raw=true "Open Ganace")

Keep note of the port number visible under RPC Server. In this case, its 7545.
![Alt text](img/4.png?raw=true "Open Ganace")

Now, we will have to update the *truffle-config.js* to enable deploying to the local network.
```javascript
module.exports = {
  networks: {
    development: {
      from: "0x0F161646c650001f3E66f94Aa2A47adeB8Ce3e6C", // put the address from your ganace
      host: "127.0.0.1",
      port: 7545,
      network_id: "*"
    }
  },
  compilers: {
    solc: {
      version: "0.8.1",   
    }
  },

};
```


## Step #6: Compile using truffle

Now, we have to develop a migration script for deploying our contract. The following code snippet will enable you to do that  In the migrations folder delete the starter file and create a new file: 2_deploy_contract.js.
```javascript
var OfficialStore = artifacts.require("OfficialStore");
module.exports = function(deployer) {
  deployer.deploy(OfficialStore);
};
```

Once done run `truffle migrate`.

```
Compiling your contracts...
===========================
> Everything is up to date, there is nothing to compile.


Starting migrations...
======================
> Network name:    'development'
> Network id:      5777
> Block gas limit: 6721975 (0x6691b7)


1_initial_migration.js
======================

   Replacing 'Migrations'
   ----------------------
   > transaction hash:    0x489342c1caafa77fd1fa00c4e55322b0d0fac9a2af16abf8b68c094d11b8a9cc
   > Blocks: 0            Seconds: 0
   > contract address:    0xA7Ce8C07168a2D4255163D635484eDE6c7708C9B
   > block number:        13
   > block timestamp:     1660479947
   > account:             0x0F161646c650001f3E66f94Aa2A47adeB8Ce3e6C
   > balance:             99.96005888
   > gas used:            248854 (0x3cc16)
   > gas price:           20 gwei
   > value sent:          0 ETH
   > total cost:          0.00497708 ETH

   > Saving migration to chain.
   > Saving artifacts
   -------------------------------------
   > Total cost:          0.00497708 ETH


2_deploy_contract.js
====================

   Deploying 'OfficialStore'
   -------------------------
   > transaction hash:    0xec8470f0cad5a4d992755806a1719763d98c9ea4c1dafc216192eaab6e4c96d1
   > Blocks: 0            Seconds: 0
   > contract address:    0x5d4f28743E676F6E8fFe7224dB0BA8E8C0219149
   > block number:        15
   > block timestamp:     1660479947
   > account:             0x0F161646c650001f3E66f94Aa2A47adeB8Ce3e6C
   > balance:             99.89943296
   > gas used:            2988783 (0x2d9aef)
   > gas price:           20 gwei
   > value sent:          0 ETH
   > total cost:          0.05977566 ETH

   > Saving migration to chain.
   > Saving artifacts
   -------------------------------------
   > Total cost:          0.05977566 ETH

Summary
=======
> Total deployments:   2
> Final cost:          0.06475274 ETH
```

# Interaction with the smart contract

We can easily interact with the contract in the terminal itself. Run this command in your terminal:
```bash
truffle console

## Let's store an instance of the contract by typing this command in the Truffle console:
let instance = await OfficialStore.deployed()


## Let's also store the addresses in a variable so that we can easily use them:
let accounts = await web3.eth.getAccounts()

##
# - Brand Address: 0x7c3a8Bb6853Fad82Bf21aD775E8bBc567b3564dc
# - Official Store assign to Brand: 0x41b945868EB2e599759DD06A645e5894785c4265

instance.setBrandContract("0x7c3a8Bb6853Fad82Bf21aD775E8bBc567b3564dc", "APPLE")

instance.getBrandFromAddress("0x7c3a8Bb6853Fad82Bf21aD775E8bBc567b3564dc")

instance.setContractOS("0x41b945868EB2e599759DD06A645e5894785c4265",{"name": "APPLE Banjarmasin", "mainBrandAddress": "0x7c3a8Bb6853Fad82Bf21aD775E8bBc567b3564dc", "district": "BANJARMASIN", "expiredDate": 10, "active": true})

instance.getContractBrandFromAddress("0x41b945868EB2e599759DD06A645e5894785c4265");

instance.disableContractOS("0x41b945868EB2e599759DD06A645e5894785c4265");

instance.enableContractOS("0x41b945868EB2e599759DD06A645e5894785c4265");

instance.isContractBrandActive("0x41b945868EB2e599759DD06A645e5894785c4265","APPLE",100)

instance.extendContractOS("0x41b945868EB2e599759DD06A645e5894785c4265",10000000)

instance.isContractBrandActive("0x41b945868EB2e599759DD06A645e5894785c4265","APPLE",1)

```


# Interaction to smartcontract with Golang
## Build
```bash
# Build abi solidity 
solc --optimize --abi ./contracts/OfficialStore.sol -o build

# Build bin solidity
solc --optimize --bin ./contracts/OfficialStore.sol -o build

# Build bin solidity
mkdir api
abigen --abi=./build/OfficialStore.abi --bin=./build/OfficialStore.bin --pkg=api --out=./api/OfficialStore.go
go mod tidy

# Create cmd folder for golang
mkdir cmd
# copy main.go to cmd

go mod tidy

go run cmd/main.go
```
