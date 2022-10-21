// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;

import "../node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

contract OfficialStore is ERC721, Ownable {

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

    event BrandAddress(address indexed _from, string _value);

    event BrandContract(address indexed _from, ContractBrand indexed _contractBrand);


    constructor() ERC721("OfficialStore", "OST") {
    }


    /* 
    * setBrandContract will set an address as official store brand name
    * mainBrandAddress is destination the address assign as brand
    * brandName will represent brand name
    */
    function setBrandContract(address mainBrandAddress, string memory brandName) external onlyOwner() {
       emit BrandAddress(mainBrandAddress, brandName);
       
        _mainAddressBrand[mainBrandAddress] = brandName;
    }

    /* 
    * setContractOS will set an address as principle official store brand
    * contractAddress is destination the address assign as official store
    * contractBrandInfo will keep information of brand info
    */
    function setContractOS(address contractAddress, ContractBrand memory contractBrandInfo) external onlyOwner() {
        
        emit BrandContract(contractAddress, contractBrandInfo);
        
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