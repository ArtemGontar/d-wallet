package walletstore

import (
	"fmt"

	dfs "github.com/ArtemGontar/d-wallet/fs"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type Store struct {
	walletsHomePath string
	keystore        *keystore.KeyStore
}

func InitialiseStore(walletsHomePath string) (*Store, error) {
	if err := dfs.EnsureDir(walletsHomePath); err != nil {
		return nil, fmt.Errorf("couldn't ensure directories at %s: %w", walletsHomePath, err)
	}

	return &Store{
		walletsHomePath: walletsHomePath,
		keystore:        keystore.NewKeyStore(walletsHomePath, keystore.StandardScryptN, keystore.StandardScryptP),
	}, nil
}

func (s *Store) WalletExists(address string) bool {
	return false
}

func (s *Store) ListWallets() ([]string, error) {
	return nil, nil
}

func (s *Store) GetWallet(privateKey string, passphrase string) (string, error) {
	return "nil", nil
}

func (s *Store) SaveWallet(passphrase string) (accounts.Account, error) {
	account, err := s.keystore.NewAccount(passphrase)
	if err != nil {
		return accounts.Account{}, err
	}
	return account, nil
}

func (s *Store) ImportWallet(account string, passphrase string, newPassphrase string) (*accounts.Account, error) {
	// f := s.getWallet(account)
	// account, err := s.keystore.Import(jsonKey, passphrase, newPassphrase)
	// if err != nil {
	// 	return accounts.Account{}, err
	// }
	return nil, nil
}

func (s *Store) DeleteWallet(name string) error {
	return nil
}

func (s *Store) getWallet(account string) *accounts.Wallet {
	for _, element := range s.keystore.Wallets() {
		if element.Accounts()[0].Address.Hex() == account {
			return &element
		}
	}

	return nil
}
