package cre

import (
	"github.com/spf13/cobra"
)

const (
	use = "cre"
	sho = "Creater a new Web3 wallet."
	lon = `Creater a new Web3 wallet.

    wallet cre
`
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var f *flags
	{
		f = &flags{}
	}

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{flags: f}).run,
		}
	}

	{
		f.Create(c)
	}

	return c, nil
}
