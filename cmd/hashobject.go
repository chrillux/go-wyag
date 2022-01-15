package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/chrillux/go-wyag/git"
	"github.com/chrillux/go-wyag/object"
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
		gr := git.NewRepo()
		f, err := os.ReadFile(arg)
		if err != nil {
			log.Fatalf("error opening file %s, %v", arg, err)
		}

		o := object.NewObject(gr, object.NewBlobObject(gr, bytes.NewReader(f)), "blob")
		hash, err := o.WriteObject(false)
		if err != nil {
			fmt.Printf("error running hash-object: %v\n", err)
			return
		}
		fmt.Println(*hash)
	}
}
