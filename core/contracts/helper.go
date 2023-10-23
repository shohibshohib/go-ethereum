package contracts

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	// OwnerContractAddr is the address of the owner contract
	OwnerContractAddr = common.HexToAddress(OwnerContractAddrHex)
	// WacContractAddr is the address of the WAC contract
	WacContractAddr = common.HexToAddress(WacContractAddrHex)
	// SidraTokenAddr is the address of the Sidra token contract
	SidraTokenAddr = common.HexToAddress(SidraTokenAddrHex)
	// MainFaucetAddr is the address of the main faucet contract
	MainFaucetAddr = common.HexToAddress(MainFaucetAddrHex)
	// WaqfAddr is the address of the waqf contract
	WaqfAddr = common.HexToAddress(WaqfAddrHex)

	SystemWallets = map[common.Address]bool{
		OwnerContractAddr: true,
		WacContractAddr:   true,
		SidraTokenAddr:    true,
		MainFaucetAddr:    true,
		WaqfAddr:          true,
	}
)

// https://docs.soliditylang.org/en/latest/internals/layout_in_storage.html#mappings-and-dynamic-arrays

func ComputeMappingHash(addr *common.Address, slot *big.Int) common.Hash {
	// Convert the slot number to 32 bytes hash with leading zeros
	p := common.BytesToHash(slot.Bytes())

	// Convert the 20 bytes address to 32 bytes hash with leading zeros
	hK := common.BytesToHash(addr.Bytes())

	// Concatenate the key and slot number and convert it to bytes array of 64 bytes
	concatenated := append(hK.Bytes(), p.Bytes()...)

	// Compute and return the Keccak-256 hash of the concatenated key and slot number
	return crypto.Keccak256Hash(concatenated)
}

func WalletStatus(addr *common.Address, statedb *state.StateDB) *big.Int {
	if IsSystemAddr(addr) {
		// Return true if the address is one of the system wallets.
		return common.Big1
	}
	// Get the state of the WalletAccessControl contract.
	wacState := statedb.GetOrNewStateObject(WacContractAddr)

	// Calculate the keccak256 hash of the key and slot number.
	keyHash := ComputeMappingHash(addr, big.NewInt(3))

	// Get the value of the key from the state.
	value := wacState.GetState(keyHash).Big()

	return value
}

func IsSystemAddr(addr *common.Address) bool {
	if addr == nil {
		return false
	}
	return SystemWallets[*addr]
}

func isWhiteList(value *big.Int) bool {
	if value == nil {
		return false
	}
	return value.Cmp(common.Big1) == 0
}

func isBlackList(value *big.Int) bool {
	if value == nil {
		return true
	}
	return value.Cmp(common.Big0) == 0
}

func isGreyList(value *big.Int) bool {
	if value == nil {
		return false
	}
	return value.Cmp(common.Big2) == 0
}

func IsTransactionAllowed(tx *types.Transaction, addr *common.Address, statedb *state.StateDB) (bool, error) {
	// Get the state of the WalletAccessControl contract.
	if IsSystemAddr(addr) {
		// Return true if the sender is one of the system wallets.
		return true, nil
	}

	// Get the state of the WalletAccessControl contract.
	senderStatus := WalletStatus(addr, statedb)
	receiverStatus := WalletStatus(tx.To(), statedb)

	if isWhiteList(senderStatus) && isWhiteList(receiverStatus) {
		// Return true if both sender and receiver are whitelisted.
		return true, nil
	} else if isGreyList(senderStatus) && IsSystemAddr(tx.To()) {
		// Return true if the sender is greylisted and the receiver is one of the system wallets.
		return true, nil
	} else if isBlackList(senderStatus) {
		// Return false if the sender is blacklisted.
		return false, fmt.Errorf("transaction is not allowed because sender is blacklisted")
	} else if isBlackList(receiverStatus) {
		// Return false if the receiver is blacklisted.
		return false, fmt.Errorf("transaction is not allowed because receiver is blacklisted")
	} else if isGreyList(senderStatus) {
		// Return false if the sender is greylisted and the receiver is not one of the system wallets.
		return false, fmt.Errorf("transaction is not allowed because sender is greylisted")
	} else if isGreyList(receiverStatus) {
		// Return false if the receiver is greylisted.
		return false, fmt.Errorf("transaction is not allowed because receiver is greylisted")
	}

	// Return false if none of the above conditions are met.
	return false, fmt.Errorf("transaction is not allowed")
}

func GetSidraTokenAbi() abi.ABI {
	output, err := abi.JSON(strings.NewReader(SidraTokenAbiString))
	if err != nil {
		// An error here probably means that the abi string is not valid.
		// That should never happen.
		panic(err)
	}
	return output
}

func GetCurrentOwnerAddr(statedb *state.StateDB) common.Address {
	// Get the state of the Owner contract.
	ownerState := statedb.GetOrNewStateObject(OwnerContractAddr)

	// Get the slot number of the owner address.
	slot := common.BigToHash(big.NewInt(0))

	// Get the value of the slot from the state.
	value := ownerState.GetState(slot).Big()

	// Return the address of the owner.
	return common.BigToAddress(value)
}
