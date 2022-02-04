package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wyag",
	Short: "wyag is a git client written in golang.",
	Long: `wyag - write yourself a git is a git client written in golang,
				  written by me for learning git internals better.
				  All inspiration comes from this post https://wyag.thb.lt/`,
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
