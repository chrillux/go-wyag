package cmd

import (
	"fmt"

	"github.com/chrillux/go-wyag/git"
	"github.com/spf13/cobra"
)

// showrefCmd represents the show-ref command
var showrefCmd = &cobra.Command{
	Use:   "show-ref",
	Short: "List references in a local repository",
	Long:  `List references in a local repository, the equivalent to git show-ref.`,
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		wyagShowRef(args)
	},
}

func init() {
	rootCmd.AddCommand(showrefCmd)
}

func wyagShowRef(args []string) {
	repo := git.NewExistingRepo()
	refs := repo.RefList()
	for ref, hash := range refs {
		fmt.Printf("%s %s\n", hash, ref)
	}
}
