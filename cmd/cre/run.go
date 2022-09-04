package cre

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/phoebetron/wallet/pkg/wallet"
	"github.com/spf13/cobra"
)

type run struct {
	flags *flags
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var wal *wallet.Wallet
	{
		wal = wallet.New(wallet.Config{})
	}

	{
		fmt.Printf("paddress    %s\n", wal.Add)
		fmt.Printf("privatek    %s\n", hexutil.EncodeBig(wal.Pri.D))
		fmt.Printf("mnemonic    %s\n", wal.Mne)
	}
}
