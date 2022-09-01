package wallet

import (
	"strings"

	"github.com/tyler-smith/go-bip39"
)

type Config struct {
	// Mne is the 24 mnemonic word list used to derive the private key of an
	// Ethereum wallet. If Mne is empty, calling wallet.New creates a new
	// Ethereum wallet.
	Mne string
}

func (c Config) Ensure() Config {
	var err error

	if c.Mne == "" {
		var ent []byte
		{
			ent, err = bip39.NewEntropy(256)
			if err != nil {
				panic(err)
			}
		}

		var mne string
		{
			mne, err = bip39.NewMnemonic(ent)
			if err != nil {
				panic(err)
			}
		}

		{
			c.Mne = mne
		}
	}

	return c
}

func (c Config) Verify() {
	if c.Mne == "" {
		panic("Config.Mne must not be empty")
	}
	if len(strings.Split(c.Mne, " ")) != 24 {
		panic("Config.Mne must contain 24 words")
	}
}
