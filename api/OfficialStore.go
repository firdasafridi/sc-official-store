// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OfficialStoreContractBrand is an auto generated low-level Go binding around an user-defined struct.
type OfficialStoreContractBrand struct {
	Name             string
	MainBrandAddress common.Address
	District         string
	ExpiredDate      *big.Int
	Active           bool
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareStrings\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"disableContractOS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"enableContractOS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"}],\"name\":\"extendContractOS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"brandAddress\",\"type\":\"address\"}],\"name\":\"getBrandFromAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"brandName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getContractBrandFromAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mainBrandAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"district\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiredDate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structOfficialStore.ContractBrand\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"brandName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"unixTimeNow\",\"type\":\"uint256\"}],\"name\":\"isContractBrandActive\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"active\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mainBrandAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"brandName\",\"type\":\"string\"}],\"name\":\"setBrandContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mainBrandAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"district\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiredDate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structOfficialStore.ContractBrand\",\"name\":\"contractBrandInfo\",\"type\":\"tuple\"}],\"name\":\"setContractOS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600d81526020016c4f6666696369616c53746f726560981b8152506040518060400160405280600381526020016213d4d560ea1b815250816000908162000065919062000192565b50600162000074828262000192565b505050620000916200008b6200009760201b60201c565b6200009b565b6200025e565b3390565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200011857607f821691505b6020821081036200013957634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200018d57600081815260208120601f850160051c81016020861015620001685750805b601f850160051c820191505b81811015620001895782815560010162000174565b5050505b505050565b81516001600160401b03811115620001ae57620001ae620000ed565b620001c681620001bf845462000103565b846200013f565b602080601f831160018114620001fe5760008415620001e55750858301515b600019600386901b1c1916600185901b17855562000189565b600085815260208120601f198616915b828110156200022f578886015182559484019460019091019084016200020e565b50858210156200024e5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b611d7d806200026e6000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c80637c0b5664116100de578063c381685011610097578063da30813111610071578063da3081311461033b578063e4d8ce771461034e578063e985e9c514610361578063f2fde38b1461039d57600080fd5b8063c381685014610302578063c87b56dd14610315578063d1527da11461032857600080fd5b80637c0b5664146102905780638da5cb5b146102b057806395d89b41146102c1578063a22cb465146102c9578063b88d4fde146102dc578063bed34bba146102ef57600080fd5b806328e3e3de1161013057806328e3e3de1461021b5780633a05bda41461022e57806342842e0e146102415780636352211e1461025457806370a0823114610267578063715018a61461028857600080fd5b806301ffc9a71461017857806306fdde03146101a0578063081812fc146101b5578063095ea7b3146101e05780630de31795146101f557806323b872dd14610208575b600080fd5b61018b6101863660046114dc565b6103b0565b60405190151581526020015b60405180910390f35b6101a8610402565b6040516101979190611549565b6101c86101c336600461155c565b610494565b6040516001600160a01b039091168152602001610197565b6101f36101ee366004611591565b6104bb565b005b6101f3610203366004611591565b6105d5565b6101f36102163660046115bb565b6105fc565b6101f36102293660046116dc565b61062d565b6101a861023c3660046117ac565b6106c2565b6101f361024f3660046115bb565b610894565b6101c861026236600461155c565b6108af565b61027a610275366004611803565b61090f565b604051908152602001610197565b6101f3610995565b6102a361029e366004611803565b6109a9565b604051610197919061181e565b6006546001600160a01b03166101c8565b6101a8610b63565b6101f36102d7366004611890565b610b72565b6101f36102ea3660046118c3565b610b81565b61018b6102fd36600461193f565b610bb9565b6101a8610310366004611803565b610c12565b6101a861032336600461155c565b610cbe565b6101f3610336366004611803565b610d31565b6101f36103493660046119a3565b610d60565b6101f361035c366004611803565b610d8a565b61018b61036f3660046119e7565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6101f36103ab366004611803565b610db6565b60006001600160e01b031982166380ac58cd60e01b14806103e157506001600160e01b03198216635b5e139f60e01b145b806103fc57506301ffc9a760e01b6001600160e01b03198316145b92915050565b60606000805461041190611a11565b80601f016020809104026020016040519081016040528092919081815260200182805461043d90611a11565b801561048a5780601f1061045f5761010080835404028352916020019161048a565b820191906000526020600020905b81548152906001019060200180831161046d57829003601f168201915b5050505050905090565b600061049f82610e2f565b506000908152600460205260409020546001600160a01b031690565b60006104c6826108af565b9050806001600160a01b0316836001600160a01b0316036105385760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b03821614806105545750610554813361036f565b6105c65760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c0000606482015260840161052f565b6105d08383610e8e565b505050565b6105dd610efc565b6001600160a01b03909116600090815260086020526040902060030155565b6106063382610f56565b6106225760405162461bcd60e51b815260040161052f90611a4b565b6105d0838383610fd5565b610635610efc565b6001600160a01b03821660009081526008602052604090208151829190819061065e9082611ae7565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906106999082611ae7565b50606082015160038201556080909101516004909101805460ff19169115159190911790555050565b6001600160a01b03831660009081526008602052604090206004015460609060ff16610722575060408051808201909152601781527f436f6e7472616374204272616e6420496e616374697665000000000000000000602082015261088d565b6001600160a01b038085166000908152600860209081526040808320600101549093168252600790522080546107e0919061075c90611a11565b80601f016020809104026020016040519081016040528092919081815260200182805461078890611a11565b80156107d55780601f106107aa576101008083540402835291602001916107d5565b820191906000526020600020905b8154815290600101906020018083116107b857829003601f168201915b505050505084610bb9565b61081e575060408051808201909152601a81527f436f6d70616e79206e6f74206f6666696369616c2073746f7265000000000000602082015261088d565b6001600160a01b03841660009081526008602052604090206003015482111561086e575060408051808201909152601081526f436f6d70616e7920496e61637469766560801b602082015261088d565b5060408051808201909152600681526541435449564560d01b60208201525b9392505050565b6105d083838360405180602001604052806000815250610b81565b6000818152600260205260408120546001600160a01b0316806103fc5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161052f565b60006001600160a01b0382166109795760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161052f565b506001600160a01b031660009081526003602052604090205490565b61099d610efc565b6109a76000611171565b565b6109e66040518060a001604052806060815260200160006001600160a01b0316815260200160608152602001600081526020016000151581525090565b6001600160a01b03821660009081526008602052604090819020815160a08101909252805482908290610a1890611a11565b80601f0160208091040260200160405190810160405280929190818152602001828054610a4490611a11565b8015610a915780601f10610a6657610100808354040283529160200191610a91565b820191906000526020600020905b815481529060010190602001808311610a7457829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610ac190611a11565b80601f0160208091040260200160405190810160405280929190818152602001828054610aed90611a11565b8015610b3a5780601f10610b0f57610100808354040283529160200191610b3a565b820191906000526020600020905b815481529060010190602001808311610b1d57829003601f168201915b50505091835250506003820154602082015260049091015460ff16151560409091015292915050565b60606001805461041190611a11565b610b7d3383836111c3565b5050565b610b8b3383610f56565b610ba75760405162461bcd60e51b815260040161052f90611a4b565b610bb384848484611291565b50505050565b600081604051602001610bcc9190611ba7565b6040516020818303038152906040528051906020012083604051602001610bf39190611ba7565b6040516020818303038152906040528051906020012014905092915050565b6001600160a01b0381166000908152600760205260409020805460609190610c3990611a11565b80601f0160208091040260200160405190810160405280929190818152602001828054610c6590611a11565b8015610cb25780601f10610c8757610100808354040283529160200191610cb2565b820191906000526020600020905b815481529060010190602001808311610c9557829003601f168201915b50505050509050919050565b6060610cc982610e2f565b6000610ce060408051602081019091526000815290565b90506000815111610d00576040518060200160405280600081525061088d565b80610d0a846112c4565b604051602001610d1b929190611bc3565b6040516020818303038152906040529392505050565b610d39610efc565b6001600160a01b03166000908152600860205260409020600401805460ff19166001179055565b610d68610efc565b6001600160a01b03821660009081526007602052604090206105d08282611ae7565b610d92610efc565b6001600160a01b03166000908152600860205260409020600401805460ff19169055565b610dbe610efc565b6001600160a01b038116610e235760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161052f565b610e2c81611171565b50565b6000818152600260205260409020546001600160a01b0316610e2c5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161052f565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610ec3826108af565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6006546001600160a01b031633146109a75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161052f565b600080610f62836108af565b9050806001600160a01b0316846001600160a01b03161480610fa957506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b80610fcd5750836001600160a01b0316610fc284610494565b6001600160a01b0316145b949350505050565b826001600160a01b0316610fe8826108af565b6001600160a01b03161461104c5760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b606482015260840161052f565b6001600160a01b0382166110ae5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161052f565b6110b9600082610e8e565b6001600160a01b03831660009081526003602052604081208054600192906110e2908490611c08565b90915550506001600160a01b0382166000908152600360205260408120805460019290611110908490611c1b565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b0316036112245760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161052f565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b61129c848484610fd5565b6112a8848484846113c5565b610bb35760405162461bcd60e51b815260040161052f90611c2e565b6060816000036112eb5750506040805180820190915260018152600360fc1b602082015290565b8160005b811561131557806112ff81611c80565b915061130e9050600a83611caf565b91506112ef565b60008167ffffffffffffffff811115611330576113306115f7565b6040519080825280601f01601f19166020018201604052801561135a576020820181803683370190505b5090505b8415610fcd5761136f600183611c08565b915061137c600a86611cc3565b611387906030611c1b565b60f81b81838151811061139c5761139c611cd7565b60200101906001600160f81b031916908160001a9053506113be600a86611caf565b945061135e565b60006001600160a01b0384163b156114bb57604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290611409903390899088908890600401611ced565b6020604051808303816000875af1925050508015611444575060408051601f3d908101601f1916820190925261144191810190611d2a565b60015b6114a1573d808015611472576040519150601f19603f3d011682016040523d82523d6000602084013e611477565b606091505b5080516000036114995760405162461bcd60e51b815260040161052f90611c2e565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610fcd565b506001949350505050565b6001600160e01b031981168114610e2c57600080fd5b6000602082840312156114ee57600080fd5b813561088d816114c6565b60005b838110156115145781810151838201526020016114fc565b50506000910152565b600081518084526115358160208601602086016114f9565b601f01601f19169290920160200192915050565b60208152600061088d602083018461151d565b60006020828403121561156e57600080fd5b5035919050565b80356001600160a01b038116811461158c57600080fd5b919050565b600080604083850312156115a457600080fd5b6115ad83611575565b946020939093013593505050565b6000806000606084860312156115d057600080fd5b6115d984611575565b92506115e760208501611575565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715611630576116306115f7565b60405290565b600067ffffffffffffffff80841115611651576116516115f7565b604051601f8501601f19908116603f01168101908282118183101715611679576116796115f7565b8160405280935085815286868601111561169257600080fd5b858560208301376000602087830101525050509392505050565b600082601f8301126116bd57600080fd5b61088d83833560208501611636565b8035801515811461158c57600080fd5b600080604083850312156116ef57600080fd5b6116f883611575565b9150602083013567ffffffffffffffff8082111561171557600080fd5b9084019060a0828703121561172957600080fd5b61173161160d565b82358281111561174057600080fd5b61174c888286016116ac565b82525061175b60208401611575565b602082015260408301358281111561177257600080fd5b61177e888286016116ac565b6040830152506060830135606082015261179a608084016116cc565b60808201528093505050509250929050565b6000806000606084860312156117c157600080fd5b6117ca84611575565b9250602084013567ffffffffffffffff8111156117e657600080fd5b6117f2868287016116ac565b925050604084013590509250925092565b60006020828403121561181557600080fd5b61088d82611575565b602081526000825160a0602084015261183a60c084018261151d565b60208501516001600160a01b0316604085810191909152850151848203601f1901606086015290915061186d828261151d565b915050606084015160808401526080840151151560a08401528091505092915050565b600080604083850312156118a357600080fd5b6118ac83611575565b91506118ba602084016116cc565b90509250929050565b600080600080608085870312156118d957600080fd5b6118e285611575565b93506118f060208601611575565b925060408501359150606085013567ffffffffffffffff81111561191357600080fd5b8501601f8101871361192457600080fd5b61193387823560208401611636565b91505092959194509250565b6000806040838503121561195257600080fd5b823567ffffffffffffffff8082111561196a57600080fd5b611976868387016116ac565b9350602085013591508082111561198c57600080fd5b50611999858286016116ac565b9150509250929050565b600080604083850312156119b657600080fd5b6119bf83611575565b9150602083013567ffffffffffffffff8111156119db57600080fd5b611999858286016116ac565b600080604083850312156119fa57600080fd5b611a0383611575565b91506118ba60208401611575565b600181811c90821680611a2557607f821691505b602082108103611a4557634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602e908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526d1c881b9bdc88185c1c1c9bdd995960921b606082015260800190565b601f8211156105d057600081815260208120601f850160051c81016020861015611ac05750805b601f850160051c820191505b81811015611adf57828155600101611acc565b505050505050565b815167ffffffffffffffff811115611b0157611b016115f7565b611b1581611b0f8454611a11565b84611a99565b602080601f831160018114611b4a5760008415611b325750858301515b600019600386901b1c1916600185901b178555611adf565b600085815260208120601f198616915b82811015611b7957888601518255948401946001909101908401611b5a565b5085821015611b975787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008251611bb98184602087016114f9565b9190910192915050565b60008351611bd58184602088016114f9565b835190830190611be98183602088016114f9565b01949350505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156103fc576103fc611bf2565b808201808211156103fc576103fc611bf2565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b600060018201611c9257611c92611bf2565b5060010190565b634e487b7160e01b600052601260045260246000fd5b600082611cbe57611cbe611c99565b500490565b600082611cd257611cd2611c99565b500690565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611d209083018461151d565b9695505050505050565b600060208284031215611d3c57600080fd5b815161088d816114c656fea26469706673582212209220e528b95a79172c9f36b642fe93650339e385174658f150992f22fbf39e7964736f6c63430008100033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Api *ApiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Api *ApiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Api *ApiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, owner)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Api *ApiCaller) CompareStrings(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "compareStrings", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Api *ApiSession) CompareStrings(a string, b string) (bool, error) {
	return _Api.Contract.CompareStrings(&_Api.CallOpts, a, b)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Api *ApiCallerSession) CompareStrings(a string, b string) (bool, error) {
	return _Api.Contract.CompareStrings(&_Api.CallOpts, a, b)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Api *ApiCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Api *ApiSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Api.Contract.GetApproved(&_Api.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Api *ApiCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Api.Contract.GetApproved(&_Api.CallOpts, tokenId)
}

// GetBrandFromAddress is a free data retrieval call binding the contract method 0xc3816850.
//
// Solidity: function getBrandFromAddress(address brandAddress) view returns(string brandName)
func (_Api *ApiCaller) GetBrandFromAddress(opts *bind.CallOpts, brandAddress common.Address) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getBrandFromAddress", brandAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBrandFromAddress is a free data retrieval call binding the contract method 0xc3816850.
//
// Solidity: function getBrandFromAddress(address brandAddress) view returns(string brandName)
func (_Api *ApiSession) GetBrandFromAddress(brandAddress common.Address) (string, error) {
	return _Api.Contract.GetBrandFromAddress(&_Api.CallOpts, brandAddress)
}

// GetBrandFromAddress is a free data retrieval call binding the contract method 0xc3816850.
//
// Solidity: function getBrandFromAddress(address brandAddress) view returns(string brandName)
func (_Api *ApiCallerSession) GetBrandFromAddress(brandAddress common.Address) (string, error) {
	return _Api.Contract.GetBrandFromAddress(&_Api.CallOpts, brandAddress)
}

// GetContractBrandFromAddress is a free data retrieval call binding the contract method 0x7c0b5664.
//
// Solidity: function getContractBrandFromAddress(address contractAddress) view returns((string,address,string,uint256,bool))
func (_Api *ApiCaller) GetContractBrandFromAddress(opts *bind.CallOpts, contractAddress common.Address) (OfficialStoreContractBrand, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getContractBrandFromAddress", contractAddress)

	if err != nil {
		return *new(OfficialStoreContractBrand), err
	}

	out0 := *abi.ConvertType(out[0], new(OfficialStoreContractBrand)).(*OfficialStoreContractBrand)

	return out0, err

}

// GetContractBrandFromAddress is a free data retrieval call binding the contract method 0x7c0b5664.
//
// Solidity: function getContractBrandFromAddress(address contractAddress) view returns((string,address,string,uint256,bool))
func (_Api *ApiSession) GetContractBrandFromAddress(contractAddress common.Address) (OfficialStoreContractBrand, error) {
	return _Api.Contract.GetContractBrandFromAddress(&_Api.CallOpts, contractAddress)
}

// GetContractBrandFromAddress is a free data retrieval call binding the contract method 0x7c0b5664.
//
// Solidity: function getContractBrandFromAddress(address contractAddress) view returns((string,address,string,uint256,bool))
func (_Api *ApiCallerSession) GetContractBrandFromAddress(contractAddress common.Address) (OfficialStoreContractBrand, error) {
	return _Api.Contract.GetContractBrandFromAddress(&_Api.CallOpts, contractAddress)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Api *ApiCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Api *ApiSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Api.Contract.IsApprovedForAll(&_Api.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Api *ApiCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Api.Contract.IsApprovedForAll(&_Api.CallOpts, owner, operator)
}

// IsContractBrandActive is a free data retrieval call binding the contract method 0x3a05bda4.
//
// Solidity: function isContractBrandActive(address contractAddress, string brandName, uint256 unixTimeNow) view returns(string active)
func (_Api *ApiCaller) IsContractBrandActive(opts *bind.CallOpts, contractAddress common.Address, brandName string, unixTimeNow *big.Int) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isContractBrandActive", contractAddress, brandName, unixTimeNow)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IsContractBrandActive is a free data retrieval call binding the contract method 0x3a05bda4.
//
// Solidity: function isContractBrandActive(address contractAddress, string brandName, uint256 unixTimeNow) view returns(string active)
func (_Api *ApiSession) IsContractBrandActive(contractAddress common.Address, brandName string, unixTimeNow *big.Int) (string, error) {
	return _Api.Contract.IsContractBrandActive(&_Api.CallOpts, contractAddress, brandName, unixTimeNow)
}

// IsContractBrandActive is a free data retrieval call binding the contract method 0x3a05bda4.
//
// Solidity: function isContractBrandActive(address contractAddress, string brandName, uint256 unixTimeNow) view returns(string active)
func (_Api *ApiCallerSession) IsContractBrandActive(contractAddress common.Address, brandName string, unixTimeNow *big.Int) (string, error) {
	return _Api.Contract.IsContractBrandActive(&_Api.CallOpts, contractAddress, brandName, unixTimeNow)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCallerSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Api *ApiCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Api *ApiSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Api.Contract.OwnerOf(&_Api.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Api *ApiCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Api.Contract.OwnerOf(&_Api.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Api *ApiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Api *ApiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Api.Contract.SupportsInterface(&_Api.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Api *ApiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Api.Contract.SupportsInterface(&_Api.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCallerSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Api *ApiCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Api *ApiSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Api.Contract.TokenURI(&_Api.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Api *ApiCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Api.Contract.TokenURI(&_Api.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Api *ApiTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Api *ApiSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Api *ApiTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, to, tokenId)
}

// DisableContractOS is a paid mutator transaction binding the contract method 0xe4d8ce77.
//
// Solidity: function disableContractOS(address contractAddress) returns()
func (_Api *ApiTransactor) DisableContractOS(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "disableContractOS", contractAddress)
}

// DisableContractOS is a paid mutator transaction binding the contract method 0xe4d8ce77.
//
// Solidity: function disableContractOS(address contractAddress) returns()
func (_Api *ApiSession) DisableContractOS(contractAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.DisableContractOS(&_Api.TransactOpts, contractAddress)
}

// DisableContractOS is a paid mutator transaction binding the contract method 0xe4d8ce77.
//
// Solidity: function disableContractOS(address contractAddress) returns()
func (_Api *ApiTransactorSession) DisableContractOS(contractAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.DisableContractOS(&_Api.TransactOpts, contractAddress)
}

// EnableContractOS is a paid mutator transaction binding the contract method 0xd1527da1.
//
// Solidity: function enableContractOS(address contractAddress) returns()
func (_Api *ApiTransactor) EnableContractOS(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "enableContractOS", contractAddress)
}

// EnableContractOS is a paid mutator transaction binding the contract method 0xd1527da1.
//
// Solidity: function enableContractOS(address contractAddress) returns()
func (_Api *ApiSession) EnableContractOS(contractAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.EnableContractOS(&_Api.TransactOpts, contractAddress)
}

// EnableContractOS is a paid mutator transaction binding the contract method 0xd1527da1.
//
// Solidity: function enableContractOS(address contractAddress) returns()
func (_Api *ApiTransactorSession) EnableContractOS(contractAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.EnableContractOS(&_Api.TransactOpts, contractAddress)
}

// ExtendContractOS is a paid mutator transaction binding the contract method 0x0de31795.
//
// Solidity: function extendContractOS(address contractAddress, uint256 ts) returns()
func (_Api *ApiTransactor) ExtendContractOS(opts *bind.TransactOpts, contractAddress common.Address, ts *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "extendContractOS", contractAddress, ts)
}

// ExtendContractOS is a paid mutator transaction binding the contract method 0x0de31795.
//
// Solidity: function extendContractOS(address contractAddress, uint256 ts) returns()
func (_Api *ApiSession) ExtendContractOS(contractAddress common.Address, ts *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ExtendContractOS(&_Api.TransactOpts, contractAddress, ts)
}

// ExtendContractOS is a paid mutator transaction binding the contract method 0x0de31795.
//
// Solidity: function extendContractOS(address contractAddress, uint256 ts) returns()
func (_Api *ApiTransactorSession) ExtendContractOS(contractAddress common.Address, ts *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ExtendContractOS(&_Api.TransactOpts, contractAddress, ts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Api *ApiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Api *ApiSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SafeTransferFrom(&_Api.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Api *ApiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SafeTransferFrom(&_Api.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Api *ApiTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Api *ApiSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.SafeTransferFrom0(&_Api.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Api *ApiTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.SafeTransferFrom0(&_Api.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Api *ApiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Api *ApiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Api.Contract.SetApprovalForAll(&_Api.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Api *ApiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Api.Contract.SetApprovalForAll(&_Api.TransactOpts, operator, approved)
}

// SetBrandContract is a paid mutator transaction binding the contract method 0xda308131.
//
// Solidity: function setBrandContract(address mainBrandAddress, string brandName) returns()
func (_Api *ApiTransactor) SetBrandContract(opts *bind.TransactOpts, mainBrandAddress common.Address, brandName string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setBrandContract", mainBrandAddress, brandName)
}

// SetBrandContract is a paid mutator transaction binding the contract method 0xda308131.
//
// Solidity: function setBrandContract(address mainBrandAddress, string brandName) returns()
func (_Api *ApiSession) SetBrandContract(mainBrandAddress common.Address, brandName string) (*types.Transaction, error) {
	return _Api.Contract.SetBrandContract(&_Api.TransactOpts, mainBrandAddress, brandName)
}

// SetBrandContract is a paid mutator transaction binding the contract method 0xda308131.
//
// Solidity: function setBrandContract(address mainBrandAddress, string brandName) returns()
func (_Api *ApiTransactorSession) SetBrandContract(mainBrandAddress common.Address, brandName string) (*types.Transaction, error) {
	return _Api.Contract.SetBrandContract(&_Api.TransactOpts, mainBrandAddress, brandName)
}

// SetContractOS is a paid mutator transaction binding the contract method 0x28e3e3de.
//
// Solidity: function setContractOS(address contractAddress, (string,address,string,uint256,bool) contractBrandInfo) returns()
func (_Api *ApiTransactor) SetContractOS(opts *bind.TransactOpts, contractAddress common.Address, contractBrandInfo OfficialStoreContractBrand) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setContractOS", contractAddress, contractBrandInfo)
}

// SetContractOS is a paid mutator transaction binding the contract method 0x28e3e3de.
//
// Solidity: function setContractOS(address contractAddress, (string,address,string,uint256,bool) contractBrandInfo) returns()
func (_Api *ApiSession) SetContractOS(contractAddress common.Address, contractBrandInfo OfficialStoreContractBrand) (*types.Transaction, error) {
	return _Api.Contract.SetContractOS(&_Api.TransactOpts, contractAddress, contractBrandInfo)
}

// SetContractOS is a paid mutator transaction binding the contract method 0x28e3e3de.
//
// Solidity: function setContractOS(address contractAddress, (string,address,string,uint256,bool) contractBrandInfo) returns()
func (_Api *ApiTransactorSession) SetContractOS(contractAddress common.Address, contractBrandInfo OfficialStoreContractBrand) (*types.Transaction, error) {
	return _Api.Contract.SetContractOS(&_Api.TransactOpts, contractAddress, contractBrandInfo)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Api *ApiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Api *ApiSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Api *ApiTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// ApiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Api contract.
type ApiApprovalIterator struct {
	Event *ApiApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiApproval represents a Approval event raised by the Api contract.
type ApiApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Api *ApiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ApiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiApprovalIterator{contract: _Api.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Api *ApiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ApiApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiApproval)
				if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Api *ApiFilterer) ParseApproval(log types.Log) (*ApiApproval, error) {
	event := new(ApiApproval)
	if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Api contract.
type ApiApprovalForAllIterator struct {
	Event *ApiApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiApprovalForAll represents a ApprovalForAll event raised by the Api contract.
type ApiApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Api *ApiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ApiApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ApiApprovalForAllIterator{contract: _Api.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Api *ApiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ApiApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiApprovalForAll)
				if err := _Api.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Api *ApiFilterer) ParseApprovalForAll(log types.Log) (*ApiApprovalForAll, error) {
	event := new(ApiApprovalForAll)
	if err := _Api.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Api contract.
type ApiOwnershipTransferredIterator struct {
	Event *ApiOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiOwnershipTransferred represents a OwnershipTransferred event raised by the Api contract.
type ApiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ApiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ApiOwnershipTransferredIterator{contract: _Api.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ApiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiOwnershipTransferred)
				if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) ParseOwnershipTransferred(log types.Log) (*ApiOwnershipTransferred, error) {
	event := new(ApiOwnershipTransferred)
	if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Api contract.
type ApiTransferIterator struct {
	Event *ApiTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiTransfer represents a Transfer event raised by the Api contract.
type ApiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Api *ApiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ApiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiTransferIterator{contract: _Api.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Api *ApiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ApiTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiTransfer)
				if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Api *ApiFilterer) ParseTransfer(log types.Log) (*ApiTransfer, error) {
	event := new(ApiTransfer)
	if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
