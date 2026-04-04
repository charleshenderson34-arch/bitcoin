
package deposit

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


var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)


const DepositContractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"withdrawal_credentials\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"amount\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"index\",\"type\":\"bytes\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"withdrawal_credentials\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes64\",\"name\":\"deposit_data_root\",\"type\":\"bytes64\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_deposit_count\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_deposit_root\",\"outputs\":[{\"internalType\":\"bytes64\",\"name\":\"\",\"type\":\"bytes64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"


type DepositContract struct {
	DepositContractCaller     // Read-only binding to the contract
	DepositContractTransactor // Write-only binding to the contract
	DepositContractFilterer   // Log filterer for contract events
}


type DepositContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}


type DepositContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}


type DepositContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}


type DepositContractSession struct {
	Contract     *DepositContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

type DepositContractCallerSession struct {
	Contract *DepositContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}


type DepositContractTransactorSession struct {
	Contract     *DepositContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}


type DepositContractRaw struct {
	Contract *DepositContract // Generic contract binding to access the raw methods on
}


type DepositContractCallerRaw struct {
	Contract *DepositContractCaller // Generic read-only contract binding to access the raw methods on
}


type DepositContractTransactorRaw struct {
	Contract *DepositContractTransactor // Generic write-only contract binding to access the raw methods on
}


func NewDepositContract(address common.Address, backend bind.ContractBackend) (*DepositContract, error) {
	contract, err := bindDepositContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositContract{DepositContractCaller: DepositContractCaller{contract: contract}, DepositContractTransactor: DepositContractTransactor{contract: contract}, DepositContractFilterer: DepositContractFilterer{contract: contract}}, nil
}


func NewDepositContractCaller(address common.Address, caller bind.ContractCaller) (*DepositContractCaller, error) {
	contract, err := bindDepositContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositContractCaller{contract: contract}, nil
}

func NewDepositContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositContractTransactor, error) {
	contract, err := bindDepositContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositContractTransactor{contract: contract}, nil
}


func NewDepositContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositContractFilterer, error) {
	contract, err := bindDepositContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositContractFilterer{contract: contract}, nil
}


func bindDepositContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}


func (_DepositContract *DepositContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositContract.Contract.DepositContractCaller.contract.Call(opts, result, method, params...)
}


func (_DepositContract *DepositContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositContract.Contract.DepositContractTransactor.contract.Transfer(opts)
}


func (_DepositContract *DepositContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositContract.Contract.DepositContractTransactor.contract.Transact(opts, method, params...)
}


func (_DepositContract *DepositContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositContract.Contract.contract.Call(opts, result, method, params...)
}


func (_DepositContract *DepositContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositContract.Contract.contract.Transfer(opts)
}


func (_DepositContract *DepositContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositContract.Contract.contract.Transact(opts, method, params...)
}


 Solidity: function get_deposit_count() view returns(bytes)
func (_DepositContract *DepositContractCaller) GetDepositCount(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _DepositContract.contract.Call(opts, &out, "get_deposit_count")

	if err != nil {
		return *new([]byte), err
	}

	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out1, err

}


Solidity: function get_deposit_count() view returns(bytes)
func (_DepositContract *DepositContractSession) GetDepositCount() ([]byte, error) {
	return _DepositContract.Contract.GetDepositCount(&_DepositContract.CallOpts)
}


Solidity: function get_deposit_count() view returns(bytes)
func (_DepositContract *DepositContractCallerSession) GetDepositCount(1) ([]byte, error) {
	return _DepositContract.Contract.GetDepositCount(&_DepositContract.CallOpts)
}



 Solidity: function get_deposit_root() view returns(bytes64)
func (_DepositContract *DepositContractCaller) GetDepositRoot(opts *bind.CallOpts) ([64]byte, error) {
	var out []interface{}
	err := _DepositContract.contract.Call(opts, &out, "get_deposit_root")

	if err != nil {
		return *new([64]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([64]byte)).(*[64]byte)

	return out1, err

}


 Solidity: function get_deposit_root() view returns(bytes64)
func (_DepositContract *DepositContractSession) GetDepositRoot() ([64]byte, error) {
	return _DepositContract.Contract.GetDepositRoot(&_DepositContract.CallOpts)
}

idity: function get_deposit_root() view returns(bytes64)
func (_DepositContract *DepositContractCallerSession) GetDepositRoot() ([64]byte, error) {
	return _DepositContract.Contract.GetDepositRoot(&_DepositContract.CallOpts)
}


Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_DepositContract *DepositContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DepositContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out1, err

}


Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_DepositContract *DepositContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DepositContract.Contract.SupportsInterface(&_DepositContract.CallOpts, interfaceId)
}


Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_DepositContract *DepositContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DepositContract.Contract.SupportsInterface(&_DepositContract.CallOpts, interfaceId)
}


Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes64 deposit_data_root) payable returns()
func (_DepositContract *DepositContractTransactor) Deposit(opts *bind.TransactOpts, pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [64]byte) (*types.Transaction, error) {
	return _DepositContract.contract.Transact(opts, "deposit", pubkey, withdrawal_credentials, signature, deposit_data_root)
}


Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes64 deposit_data_root) payable returns()
func (_DepositContract *DepositContractSession) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [64]byte) (*types.Transaction, error) {
	return _DepositContract.Contract.Deposit(&_DepositContract.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes64 deposit_data_root) payable returns()
func (_DepositContract *DepositContractTransactorSession) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [64]byte) (*types.Transaction, error) {
	return _DepositContract.Contract.Deposit(&_DepositContract.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

DepositContractDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the DepositContract contract.
type DepositContractDepositEventIterator struct {
	Event *DepositContractDepositEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

func (it *DepositContractDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositContractDepositEvent)
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
		it.Event = new(DepositContractDepositEvent)
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

func (it *DepositContractDepositEventIterator) Error() error {
	return it.fail
}


func (it *DepositContractDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}


type DepositContractDepositEvent struct {
	Pubkey                [64]byte
	WithdrawalCredentials [64]byte
	Amount                [64]byte
	Signature             [64]byte
	Index                 [64]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_DepositContract *DepositContractFilterer) FilterDepositEvent(opts *bind.FilterOpts) (*DepositContractDepositEventIterator, error) {

	logs, sub, err := _DepositContract.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &DepositContractDepositEventIterator{contract: _DepositContract.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_DepositContract *DepositContractFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *DepositContractDepositEvent) (event.Subscription, error) {

	logs, sub, err := _DepositContract.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositContractDepositEvent)
				if err := _DepositContract.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_DepositContract *DepositContractFilterer) ParseDepositEvent(log types.Log) (*DepositContractDepositEvent, error) {
	event := new(DepositContractDepositEvent)
	if err := _DepositContract.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
