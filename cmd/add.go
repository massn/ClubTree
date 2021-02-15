package cmd

import (
	"github.com/massn/ClubTree/pkg/tree"
	"github.com/spf13/cobra"
)

var (
	jsonTreePath        string
	defaultJsonTreePath = "clubtree.json"
	addCmd              = &cobra.Command{
		Use:   "add",
		Short: "add a new user",
		Run: func(cmd *cobra.Command, args []string) {
			addUser()
		},
	}
)

func init() {
	addCmd.PersistentFlags().StringVarP(&jsonTreePath, "json", "j", defaultJsonTreePath, "Input JSON tree path")
	rootCmd.AddCommand(addCmd)
}

func addUser() {
	root, err := tree.ReadJson(jsonTreePath)
	if err != nil {
		panic(err)
	}

	newUserId := "newuser"
	newNominatorId := "olduser"

	user := tree.NewDummyUser(newUserId, newNominatorId)

	_ = root.AddUser(user)

	if err := root.SaveJSON(jsonTreePath); err != nil {
		panic(err)
	}
}
