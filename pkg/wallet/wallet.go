package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/wallet"
)

type Wallet struct {
	Add ethgo.Address
	Mne string
	Pri *ecdsa.PrivateKey
	Pub *ecdsa.PublicKey
}

func New(con Config) *Wallet {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	var err error

	var see []byte
	{
		see, err = bip39.NewSeedWithErrorChecking(con.Mne, "")
		if err != nil {
			panic(err)
		}
	}

	var mas *hdkeychain.ExtendedKey
	{
		mas, err = hdkeychain.NewMaster(see, &chaincfg.MainNetParams)
		if err != nil {
			panic(err)
		}
	}

	var pri *ecdsa.PrivateKey
	{
		pri, err = wallet.DefaultDerivationPath.Derive(mas)
		if err != nil {
			panic(err)
		}
	}

	return &Wallet{
		Add: pubKeyToAddress(&pri.PublicKey),
		Mne: con.Mne,
		Pri: pri,
		Pub: &pri.PublicKey,
	}
}

func pubKeyToAddress(pub *ecdsa.PublicKey) ethgo.Address {
	var byt []byte
	{
		byt = ethgo.Keccak256(elliptic.Marshal(btcec.S256(), pub.X, pub.Y)[1:])
	}

	var add ethgo.Address
	{
		copy(add[:], byt[12:])
	}

	return add
}
