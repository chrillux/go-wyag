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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		catFile(args[0])
	},
}

var printType bool

func init() {
	rootCmd.AddCommand(catfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
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
