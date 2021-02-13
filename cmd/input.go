package cmd

import (
	"github.com/massn/ClubTree/pkg/tree"
	"github.com/spf13/cobra"
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
	oldFilename := "clubtree.json"
	newFilename := "new-clubtree.json"
	root, err := tree.ReadJson(oldFilename)
	if err != nil {
		panic(err)
	}

	newUserId := "newuser"
	newNominatorId := "olduser"

	user := tree.NewDummyUser(newUserId, newNominatorId)

	_ = root.AddUser(user)

	if err := root.SaveJSON(newFilename); err != nil {
		panic(err)
	}
}
