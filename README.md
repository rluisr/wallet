# wallet

### binary use

```
wallet cre
```

### library use

```go
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rluisr/wallet/pkg/wallet"
	"github.com/spf13/cobra"
)

func main() {
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
```

```
paddress    0x258e...b5Df
privatek    0x818a...400a
mnemonic    tea choral speed ... teddy moon brother
```
