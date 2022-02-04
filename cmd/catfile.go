package cmd

import (
	"fmt"
	"log"

	"github.com/chrillux/go-wyag/object"
	"github.com/spf13/cobra"
)

// catfileCmd represents the catfile command
var catfileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Provide content or type information for repository objects",
	Long:  "Provide content or type information for repository objects",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		catFile(args[0])
	},
}

var printType bool

func init() {
	rootCmd.AddCommand(catfileCmd)
	catfileCmd.Flags().BoolVarP(&printType, "object type", "t", false, "Print the object type.")
}

func catFile(hash string) {
	o, err := object.ReadObject(hash)
	if err != nil {
		log.Fatalf("error running cat-file: %v", err)
	}
	if printType {
		fmt.Println(o.GetObjType())
	} else {
		fmt.Println(o)
	}
}
