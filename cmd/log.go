package cmd

import (
	"fmt"
	"log"

	"github.com/chrillux/go-wyag/git"
	"github.com/chrillux/go-wyag/object"
	"github.com/spf13/cobra"
)

// catfileCmd represents the catfile command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show the wyag log",
	Long:  `Generates a wyag log for all the commits.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wyagLog(args[0])
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

func wyagLog(hash string) {
	repo := git.NewRepo()
	seen := map[string]bool{}
	fmt.Print("digraph wyaglog{")
	logRecurse(repo, hash, seen)
	fmt.Print("}")
}

func logRecurse(repo *git.Repository, hash string, seen map[string]bool) {
	if seen[hash] {
		return
	}
	seen[hash] = true

	commit, err := object.ReadObject(hash)
	if err != nil {
		log.Fatal(err)
	}
	if commit.GetObjType() != "commit" {
		log.Fatalf("object %s is not a commit", hash)
	}

	o := commit.(*object.CommitObject)
	parents := o.GetParents()
	// Base case: the initial commit.
	if len(parents) == 0 {
		return
	}

	for _, p := range parents {
		fmt.Printf("c_%s -> c_%s;", hash, p)
		logRecurse(repo, p, seen)
	}
}
