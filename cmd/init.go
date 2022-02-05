package cmd

import (
	"fmt"

	"github.com/chrillux/go-wyag/git"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <directory>",
	Short: "Init a new git repo.",
	Long:  `Init a new git repo. If no argument is given, init will be done in current directory.`,
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		wyagInit(args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func wyagInit(args []string) {
	var path string
	if len(args) == 0 {
		path = "."
	} else {
		path = args[0]
	}
	gr := git.NewRepo(path)
	err := gr.Init(path, true)
	if err != nil {
		fmt.Printf("error running init: %v\n", err)
		return
	}
	fmt.Printf("created repo %s\n", path)
}
