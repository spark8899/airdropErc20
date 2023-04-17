package cmd

import (
	"github.com/spark8899/airdropErc20/pkg/airdrop"
	"github.com/spf13/cobra"
)

var airdropCmd = &cobra.Command{
	Use:   "airdrop",
	Short: "airdrop token to address.",
	Long:  `airdrop token to address.`,
	Run: func(cmd *cobra.Command, args []string) {
		airdrop.Airdrop()
	},
}
