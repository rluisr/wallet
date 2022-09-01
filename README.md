# wallet

```go
var wal *wallet.Wallet
{
	wal = wallet.New(wallet.Config{})
}

{
	fmt.Printf("paddress    %s\n", wal.Add)
	fmt.Printf("privatek    %s\n", hexutil.EncodeBig(wal.Pri.D))
	fmt.Printf("mnemonic    %s\n", wal.Mne)
}
```

```
paddress    0x258e...b5Df
privatek    0x818a...400a
mnemonic    tea choral speed ... teddy moon brother
```
