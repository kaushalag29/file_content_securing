// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_owner_list\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"}],\"name\":\"change_owner_power\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"add_key\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"get_owner_power\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_ipfs_hash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_key\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"}],\"name\":\"add_owner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"update_hash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"update_key\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x608060405234801561001057600080fd5b50604051611c59380380611c598339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061009581600461009b60201b60201c565b506103fa565b81600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561015c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220616c72656164792065786973747300000000000000000000000081525060200191505060405180910390fd5b60038060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661021c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156102d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b60016000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505082600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b611850806104096000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638d065180116100715780638d065180146102ce578063d131f1f514610326578063e072b5f6146103a9578063eac738131461042c578063fc65e5a71461047a578063fcb6c39a14610535576100a9565b8063025e7c27146100ae57806302d05d3f1461011c57806383ced62714610166578063866309fa146101c55780638c5cd3e114610213575b600080fd5b6100da600480360360208110156100c457600080fd5b81019080803590602001909291905050506105f0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61012461062c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61016e610651565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101b1578082015181840152602081019050610196565b505050509050019250505060405180910390f35b610211600480360360408110156101db57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610856565b005b6102cc6004803603602081101561022957600080fd5b810190808035906020019064010000000081111561024657600080fd5b82018360208201111561025857600080fd5b8035906020019184600183028401116401000000008311171561027a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610b7d565b005b610310600480360360208110156102e457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b97565b6040518082815260200191505060405180910390f35b61032e610d57565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561036e578082015181840152602081019050610353565b50505050905090810190601f16801561039b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103b1610f70565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103f15780820151818401526020810190506103d6565b50505050905090810190601f16801561041e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104786004803603604081101561044257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611189565b005b6105336004803603602081101561049057600080fd5b81019080803590602001906401000000008111156104ad57600080fd5b8201836020820111156104bf57600080fd5b803590602001918460018302840111640100000000831117156104e157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506114e8565b005b6105ee6004803603602081101561054b57600080fd5b810190808035906020019064010000000081111561056857600080fd5b82018360208201111561057a57600080fd5b8035906020019184600183028401116401000000008311171561059c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611679565b005b600181815481106105fd57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610714576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156107c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b600180548060200260200160405190810160405280929190818152602001828054801561084b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610801575b505050505091505090565b816000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156108fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806117fb6021913960400191505060405180910390fd5b82600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166109bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b60038060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610a7d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610b32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b83600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b8060049080519060200190610b93929190611755565b5050565b60006001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c5a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610d0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054915050919050565b60606001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610e1a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610ecf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b60058054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f655780601f10610f3a57610100808354040283529160200191610f65565b820191906000526020600020905b815481529060010190602001808311610f4857829003601f168201915b505050505091505090565b60606001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611033576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156110e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b60048054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561117e5780601f106111535761010080835404028352916020019161117e565b820191906000526020600020905b81548152906001019060200180831161116157829003601f168201915b505050505091505090565b81600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561124a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220616c72656164792065786973747300000000000000000000000081525060200191505060405180910390fd5b60038060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661130a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156113bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b60016000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505082600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b6002600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166115a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561165e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b8160059080519060200190611674929190611755565b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461173b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6f6e6c792063726561746f722063616e2061636365737300000000000000000081525060200191505060405180910390fd5b8060049080519060200190611751929190611755565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061179657805160ff19168380011785556117c4565b828001600101855582156117c4579182015b828111156117c35782518255916020019190600101906117a8565b5b5090506117d191906117d5565b5090565b6117f791905b808211156117f35760008160009055506001016117db565b5090565b9056fe63616e6e6f7420706572666f726d206f7065726174696f6e206f6e206f776e6572a265627a7a72315820ab28d181a207ec5bc29ccadc0ced72a62819a0a67d266ef5d579d692f3e0f63c64736f6c634300050b0032"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _creator common.Address) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend, _creator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Store *StoreCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Store *StoreSession) Creator() (common.Address, error) {
	return _Store.Contract.Creator(&_Store.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Store *StoreCallerSession) Creator() (common.Address, error) {
	return _Store.Contract.Creator(&_Store.CallOpts)
}

// GetIpfsHash is a free data retrieval call binding the contract method 0xd131f1f5.
//
// Solidity: function get_ipfs_hash() constant returns(string)
func (_Store *StoreCaller) GetIpfsHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "get_ipfs_hash")
	return *ret0, err
}

// GetIpfsHash is a free data retrieval call binding the contract method 0xd131f1f5.
//
// Solidity: function get_ipfs_hash() constant returns(string)
func (_Store *StoreSession) GetIpfsHash() (string, error) {
	return _Store.Contract.GetIpfsHash(&_Store.CallOpts)
}

// GetIpfsHash is a free data retrieval call binding the contract method 0xd131f1f5.
//
// Solidity: function get_ipfs_hash() constant returns(string)
func (_Store *StoreCallerSession) GetIpfsHash() (string, error) {
	return _Store.Contract.GetIpfsHash(&_Store.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0xe072b5f6.
//
// Solidity: function get_key() constant returns(string)
func (_Store *StoreCaller) GetKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "get_key")
	return *ret0, err
}

// GetKey is a free data retrieval call binding the contract method 0xe072b5f6.
//
// Solidity: function get_key() constant returns(string)
func (_Store *StoreSession) GetKey() (string, error) {
	return _Store.Contract.GetKey(&_Store.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0xe072b5f6.
//
// Solidity: function get_key() constant returns(string)
func (_Store *StoreCallerSession) GetKey() (string, error) {
	return _Store.Contract.GetKey(&_Store.CallOpts)
}

// GetOwnerList is a free data retrieval call binding the contract method 0x83ced627.
//
// Solidity: function get_owner_list() constant returns(address[])
func (_Store *StoreCaller) GetOwnerList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "get_owner_list")
	return *ret0, err
}

// GetOwnerList is a free data retrieval call binding the contract method 0x83ced627.
//
// Solidity: function get_owner_list() constant returns(address[])
func (_Store *StoreSession) GetOwnerList() ([]common.Address, error) {
	return _Store.Contract.GetOwnerList(&_Store.CallOpts)
}

// GetOwnerList is a free data retrieval call binding the contract method 0x83ced627.
//
// Solidity: function get_owner_list() constant returns(address[])
func (_Store *StoreCallerSession) GetOwnerList() ([]common.Address, error) {
	return _Store.Contract.GetOwnerList(&_Store.CallOpts)
}

// GetOwnerPower is a free data retrieval call binding the contract method 0x8d065180.
//
// Solidity: function get_owner_power(address _owner) constant returns(uint256)
func (_Store *StoreCaller) GetOwnerPower(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "get_owner_power", _owner)
	return *ret0, err
}

// GetOwnerPower is a free data retrieval call binding the contract method 0x8d065180.
//
// Solidity: function get_owner_power(address _owner) constant returns(uint256)
func (_Store *StoreSession) GetOwnerPower(_owner common.Address) (*big.Int, error) {
	return _Store.Contract.GetOwnerPower(&_Store.CallOpts, _owner)
}

// GetOwnerPower is a free data retrieval call binding the contract method 0x8d065180.
//
// Solidity: function get_owner_power(address _owner) constant returns(uint256)
func (_Store *StoreCallerSession) GetOwnerPower(_owner common.Address) (*big.Int, error) {
	return _Store.Contract.GetOwnerPower(&_Store.CallOpts, _owner)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_Store *StoreCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_Store *StoreSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Store.Contract.Owners(&_Store.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_Store *StoreCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Store.Contract.Owners(&_Store.CallOpts, arg0)
}

// AddKey is a paid mutator transaction binding the contract method 0x8c5cd3e1.
//
// Solidity: function add_key(string _key) returns()
func (_Store *StoreTransactor) AddKey(opts *bind.TransactOpts, _key string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "add_key", _key)
}

// AddKey is a paid mutator transaction binding the contract method 0x8c5cd3e1.
//
// Solidity: function add_key(string _key) returns()
func (_Store *StoreSession) AddKey(_key string) (*types.Transaction, error) {
	return _Store.Contract.AddKey(&_Store.TransactOpts, _key)
}

// AddKey is a paid mutator transaction binding the contract method 0x8c5cd3e1.
//
// Solidity: function add_key(string _key) returns()
func (_Store *StoreTransactorSession) AddKey(_key string) (*types.Transaction, error) {
	return _Store.Contract.AddKey(&_Store.TransactOpts, _key)
}

// AddOwner is a paid mutator transaction binding the contract method 0xeac73813.
//
// Solidity: function add_owner(address _owner, uint256 _power) returns()
func (_Store *StoreTransactor) AddOwner(opts *bind.TransactOpts, _owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "add_owner", _owner, _power)
}

// AddOwner is a paid mutator transaction binding the contract method 0xeac73813.
//
// Solidity: function add_owner(address _owner, uint256 _power) returns()
func (_Store *StoreSession) AddOwner(_owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _Store.Contract.AddOwner(&_Store.TransactOpts, _owner, _power)
}

// AddOwner is a paid mutator transaction binding the contract method 0xeac73813.
//
// Solidity: function add_owner(address _owner, uint256 _power) returns()
func (_Store *StoreTransactorSession) AddOwner(_owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _Store.Contract.AddOwner(&_Store.TransactOpts, _owner, _power)
}

// ChangeOwnerPower is a paid mutator transaction binding the contract method 0x866309fa.
//
// Solidity: function change_owner_power(address _owner, uint256 _power) returns()
func (_Store *StoreTransactor) ChangeOwnerPower(opts *bind.TransactOpts, _owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "change_owner_power", _owner, _power)
}

// ChangeOwnerPower is a paid mutator transaction binding the contract method 0x866309fa.
//
// Solidity: function change_owner_power(address _owner, uint256 _power) returns()
func (_Store *StoreSession) ChangeOwnerPower(_owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangeOwnerPower(&_Store.TransactOpts, _owner, _power)
}

// ChangeOwnerPower is a paid mutator transaction binding the contract method 0x866309fa.
//
// Solidity: function change_owner_power(address _owner, uint256 _power) returns()
func (_Store *StoreTransactorSession) ChangeOwnerPower(_owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangeOwnerPower(&_Store.TransactOpts, _owner, _power)
}

// UpdateHash is a paid mutator transaction binding the contract method 0xfc65e5a7.
//
// Solidity: function update_hash(string _hash) returns()
func (_Store *StoreTransactor) UpdateHash(opts *bind.TransactOpts, _hash string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "update_hash", _hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0xfc65e5a7.
//
// Solidity: function update_hash(string _hash) returns()
func (_Store *StoreSession) UpdateHash(_hash string) (*types.Transaction, error) {
	return _Store.Contract.UpdateHash(&_Store.TransactOpts, _hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0xfc65e5a7.
//
// Solidity: function update_hash(string _hash) returns()
func (_Store *StoreTransactorSession) UpdateHash(_hash string) (*types.Transaction, error) {
	return _Store.Contract.UpdateHash(&_Store.TransactOpts, _hash)
}

// UpdateKey is a paid mutator transaction binding the contract method 0xfcb6c39a.
//
// Solidity: function update_key(string _key) returns()
func (_Store *StoreTransactor) UpdateKey(opts *bind.TransactOpts, _key string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "update_key", _key)
}

// UpdateKey is a paid mutator transaction binding the contract method 0xfcb6c39a.
//
// Solidity: function update_key(string _key) returns()
func (_Store *StoreSession) UpdateKey(_key string) (*types.Transaction, error) {
	return _Store.Contract.UpdateKey(&_Store.TransactOpts, _key)
}

// UpdateKey is a paid mutator transaction binding the contract method 0xfcb6c39a.
//
// Solidity: function update_key(string _key) returns()
func (_Store *StoreTransactorSession) UpdateKey(_key string) (*types.Transaction, error) {
	return _Store.Contract.UpdateKey(&_Store.TransactOpts, _key)
}
