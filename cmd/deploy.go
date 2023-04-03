package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spark8899/airdropErc20/pkg/config"
	"github.com/spark8899/airdropErc20/pkg/deploy"
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
			fmt.Println("deploy contract:", config.TOKEN)
			fmt.Println("RPC URL:", config.RPCURL)
			deploy.Deploy(config.TOKEN)
		} else if contact_name == "airdrop" {
			fmt.Println("deploy contract:", config.AIRDROP)
			fmt.Println("RPC URL:", config.RPCURL)
			deploy.Deploy(config.AIRDROP)
		} else {
			fmt.Println("please input contract name in (token|airdrop), you input:", contact_name)
			os.Exit(15)
		}
	},
}
