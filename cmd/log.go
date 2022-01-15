/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/chrillux/go-wyag/git"
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

	commit, err := git.ReadObject(hash)
	if err != nil {
		log.Fatal(err)
	}
	if commit.GetObjType() != "commit" {
		log.Fatalf("object %s is not a commit", hash)
	}

	parents := commit.GetParents()
	// Base case: the initial commit.
	if len(parents) == 0 {
		return
	}

	for _, p := range parents {
		fmt.Printf("c_%s -> c_%s;", hash, p)
		logRecurse(repo, p, seen)
	}
}
