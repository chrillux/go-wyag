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
	zread := git.ReadGitFile(hash)
	o, err := git.CatFile(zread)
	if err != nil {
		log.Fatalf("error running cat-file: %v", err)
	}
	if printType {
		fmt.Printf("%s\n", o.GetObjType())
	} else {
		fmt.Printf("%s", o.GetDeserializedData())
	}
}