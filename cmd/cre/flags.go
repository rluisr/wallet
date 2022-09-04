package cre

import (
	"github.com/spf13/cobra"
)

type flags struct {
	// Dst string
}

func (f *flags) Create(cmd *cobra.Command) {
	// cmd.Flags().StringVarP(&f.Dst, "dst", "d", "", "The destination exchange to execute trades on, e.g. dydx.")
}

func (f *flags) Verify() {
	// if f.Dst == "" {
	// 	panic("-d/--dst must not be empty")
	// }
}
