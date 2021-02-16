package cmd

import (
	"github.com/spf13/cobra"
)

var (
	jsonTreePath        string
	defaultJsonTreePath = "clubtree.json"
	rootCmd             = &cobra.Command{
		Use:   "clubtree",
		Short: "make a club tree (social tree graph)",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&jsonTreePath, "json", "j", defaultJsonTreePath, "Input JSON tree path")
}
