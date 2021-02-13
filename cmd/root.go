package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "clubtree",
		Short: "make a social tree graph.",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
