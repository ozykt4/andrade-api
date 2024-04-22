package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ozykt4/andrade-api/config"
	"github.com/ozykt4/andrade-api/internal/api"
	"github.com/spf13/cobra"
)

func main() {
	config.LoadConfig()

	var host string
	var port string
	var module string

	rootCmd := &cobra.Command{
		Use:   "Use --bind or -b for specify host and --port or -p for specify port",
		Short: "Start jzap-api application",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Start Andrade API application ", module, " on ", host, ":", port)
			if module == "andrade" {
				log.Fatal(api.Run(host, port))
			} else {
				log.Fatal("module not found")
			}
		},
	}

	rootCmd.PersistentFlags().StringVarP(&host, "bind", "b", "0.0.0.0", "Use --bind or -b for specify host application -b 0.0.0.0")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", config.GetConfig().Port, "Use --port or -p for specify port -p 5001")
	rootCmd.PersistentFlags().StringVarP(&module, "module", "m", config.GetConfig().ApplicationName, "Use --module or -m for specify module -m api, minizap, worker, campaign-summary")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
