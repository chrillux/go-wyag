package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chrillux/go-wyag/git"
	"github.com/spf13/cobra"
)

// hashobjectCmd represents the hashobject command
var hashobjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "Hash a file.",
	Long:  `Run man git-hash-object for more information.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hashObject(args)
	},
}

func init() {
	rootCmd.AddCommand(hashobjectCmd)
}

func hashObject(args []string) {
	for _, arg := range args {
		f, err := os.ReadFile(arg)
		if err != nil {
			log.Fatalf("error opening file %s, %v", arg, err)
		}
		gr := git.New()
		h, err := gr.HashObject("blob", f, false)
		if err != nil {
			fmt.Printf("error running init: %v\n", err)
			return
		}
		fmt.Printf("%s\n", *h)
	}
}
