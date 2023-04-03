package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of airdroperc20",
	Long:  `All software has versions. This is airdroperc20's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("airdroperc20 v0.1 -- HEAD")
	},
}
