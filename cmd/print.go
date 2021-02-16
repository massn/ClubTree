package cmd

import (
	"github.com/massn/ClubTree/pkg/tree"
	"github.com/spf13/cobra"
)

var (
	printCmd = &cobra.Command{
		Use:   "print",
		Short: "print the club tree",
		Run: func(cmd *cobra.Command, args []string) {
			print()
		},
	}
)

func init() {
	rootCmd.AddCommand(printCmd)
}

func print() {
	root, err := tree.ReadJson(jsonTreePath)
	if err != nil {
		panic(err)
	}
	root.Print()
}
