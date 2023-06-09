package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spark8899/airdropErc20/pkg/config"
	"github.com/spark8899/airdropErc20/pkg/contract"
	dp "github.com/spark8899/airdropErc20/pkg/deploy"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy contract (token|airdrop)",
	Long:  `deploy contract (token|airdrop) on chain.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contact_name := strings.TrimSpace(args[0])
		config, err := config.LoadConfig()
		if err != nil {
			log.Fatal("cannot load config:", err)
		}

		if contact_name == "token" {
			fmt.Println("Deploy contract:", config.TOKEN)
			fmt.Println("RPC URL:", config.RPCURL)
			dp.Deploy(config.TOKEN)
		} else if contact_name == "airdrop" {
			fmt.Println("Deploy contract:", config.AIRDROP)
			fmt.Println("RPC URL:", config.RPCURL)
			dp.Deploy(config.AIRDROP)
		} else {
			fmt.Println("Please input contract name in (token|airdrop), you input:", contact_name)
			os.Exit(15)
		}
	},
}

var getAddrBalanceCmd = &cobra.Command{
	Use:   "getAddrBalance",
	Short: "get eth addr balance.",
	Long:  "get eth addr balance.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addrStr := strings.TrimSpace(args[0])
		if addrStr == "" {
			fmt.Println("Enter addrStr is empty")
			os.Exit(5)
		}
		contract.AddrBalance(addrStr)
	},
}

var getTokenBalanceCmd = &cobra.Command{
	Use:   "getTokenBalance",
	Short: "get token addr balance.",
	Long:  "get token addr balance.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addrStr := strings.TrimSpace(args[0])
		if addrStr == "" {
			fmt.Println("Enter addrStr is empty")
			os.Exit(5)
		}
		contract.TokenBalance(addrStr)
	},
}
