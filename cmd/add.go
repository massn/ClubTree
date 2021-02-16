package cmd

import (
	"github.com/massn/ClubTree/pkg/tree"
	"github.com/spf13/cobra"
	"os"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a new user",
		Run: func(cmd *cobra.Command, args []string) {
			addUser()
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}

func addUser() {
	_, err := os.Stat(jsonTreePath)
	var root *tree.User
	if err == nil {
		root, err = tree.ReadJson(jsonTreePath)
		if err != nil {
			panic(err)
		}
	} else {
		root = tree.NewEmptyRoot()
	}

	user := tree.NewUser()
	_ = root.AddUser(user)

	if err := root.SaveJSON(jsonTreePath); err != nil {
		panic(err)
	}
}
