package cmd

import (
	"github.com/rluisr/wallet/cmd/com"
	"github.com/rluisr/wallet/cmd/cre"
	"github.com/rluisr/wallet/cmd/ver"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

var (
	use = "wallet"
	sho = "Manage wallet keys."
	lon = "Manage wallet keys."
)

func New() (*cobra.Command, error) {
	var err error

	// --------------------------------------------------------------------- //

	var cmdCom *cobra.Command
	{
		c := com.Config{}

		cmdCom, err = com.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cmdCre *cobra.Command
	{
		c := cre.Config{}

		cmdCre, err = cre.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cmdVer *cobra.Command
	{
		c := ver.Config{}

		cmdVer, err = ver.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// --------------------------------------------------------------------- //

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
			// We slience errors because we do not want to see spf13/cobra printing.
			// The errors returned by the commands will be propagated to the main.go
			// anyway, where we have custom error printing for the command line
			// tool.
			SilenceErrors: true,
			SilenceUsage:  true,
		}
	}

	{
		c.SetHelpCommand(&cobra.Command{Hidden: true})
	}

	{
		c.AddCommand(cmdCom)
		c.AddCommand(cmdCre)
		c.AddCommand(cmdVer)
	}

	return c, nil
}
