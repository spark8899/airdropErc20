package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "airdroperc20",
	Short: "a erc20 airdrop tools",
	Long:  "airdrop erc20 tool.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(getAddrBalanceCmd)
	rootCmd.AddCommand(getTokenBalanceCmd)
	rootCmd.AddCommand(airdropCmd)
}
