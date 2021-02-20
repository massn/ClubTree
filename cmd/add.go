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
	var ct *tree.ClubTree
	if err == nil {
		ct, err = tree.ReadJson(jsonTreePath)
		if err != nil {
			panic(err)
		}
	} else {
		ct = tree.NewClubTree("New ClubTree")
	}

	user := tree.NewUser()
	_ = ct.AddUser(user)

	if err := ct.Tree.SaveJSON(jsonTreePath); err != nil {
		panic(err)
	}
}
