package btcwallet

import (
	"path/filepath"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"

	"github.com/btcsuite/btcwallet/chain"
	"github.com/btcsuite/btcwallet/wallet"

	// This is required to register bdb as a valid walletdb driver. In the
	// init function of the package, it registers itself. The import is used
	// to activate the side effects w/o actually binding the package name to
	// a file-level variable.
	_ "github.com/btcsuite/btcwallet/walletdb/bdb"
)

var (
	// defaultPubPassphrase is the default public wallet passphrase which is
	// used when the user indicates they do not want additional protection
	// provided by having all public data in the wallet encrypted by a
	// passphrase only known to them.
	defaultPubPassphrase = []byte("public")
)

// Config is a struct which houses configuration parameters which modify the
// instance of BtcWallet generated by the New() function.
type Config struct {
	// DataDir is the name of the directory where the wallet's persistent
	// state should be stored.
	DataDir string

	// LogDir is the name of the directory which should be used to store
	// generated log files.
	LogDir string

	// PrivatePass is the private password to the underlying btcwallet
	// instance. Without this, the wallet cannot be decrypted and operated.
	PrivatePass []byte

	// PublicPass is the optional public password to btcwallet. This is
	// optionally used to encrypt public material such as public keys and
	// scripts.
	PublicPass []byte

	// HdSeed is an optional seed to feed into the wallet. If this is
	// unspecified, a new seed will be generated.
	HdSeed []byte

	// Birthday specifies the time at which this wallet was initially
	// created. It is used to bound rescans for used addresses.
	Birthday time.Time

	// RecoveryWindow specifies the address look-ahead for which to scan
	// when restoring a wallet. The recovery window will apply to all
	// default BIP44 derivation paths.
	RecoveryWindow uint32

	// ChainSource is the primary chain interface. This is used to operate
	// the wallet and do things such as rescanning, sending transactions,
	// notifications for received funds, etc.
	ChainSource chain.Interface

	// NetParams is the net parameters for the target chain.
	NetParams *chaincfg.Params

	// CoinType specifies the BIP 44 coin type to be used for derivation.
	CoinType uint32

	// Wallet is an unlocked wallet instance that is set if the
	// UnlockerService has already opened and unlocked the wallet. If this
	// is nil, then a wallet might have just been created or is simply not
	// encrypted at all, in which case it should be attempted to be loaded
	// normally when creating the BtcWallet.
	Wallet *wallet.Wallet
}

// NetworkDir returns the directory name of a network directory to hold wallet
// files.
func NetworkDir(dataDir string, chainParams *chaincfg.Params) string {
	netname := chainParams.Name

	// For now, we must always name the testnet data directory as "testnet"
	// and not "testnet3" or any other version, as the chaincfg testnet3
	// parameters will likely be switched to being named "testnet3" in the
	// future.  This is done to future proof that change, and an upgrade
	// plan to move the testnet3 data directory can be worked out later.
	if chainParams.Net == wire.TestNet3 {
		netname = "testnet"
	}

	return filepath.Join(dataDir, netname)
}
