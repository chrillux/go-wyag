package cmd

import (
	"fmt"
	"log"

	"github.com/chrillux/go-wyag/object"
	"github.com/spf13/cobra"
)

// catfileCmd represents the catfile command
var lstreeCmd = &cobra.Command{
	Use:   "ls-tree",
	Short: "List the contents of a tree object",
	Long:  `Lists the contents of a given tree object, the equivalent to git ls-tree.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lstree(args[0])
	},
}

func init() {
	rootCmd.AddCommand(lstreeCmd)
}

func lstree(hash string) {
	obj, err := object.ReadObject(hash)
	if err != nil {
		log.Fatal(err)
	}

	if obj.GetObjType() != "tree" {
		log.Fatalf("object %s is not a tree object", hash)
	}

	o := obj.Obj().(*object.TreeObject)
	for _, item := range o.Items() {
		i, err := object.ReadObject(item.SHA())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\t%s\t%s\n", item.Mode(), i.GetObjType(), item.SHA(), item.Path())
	}
}
