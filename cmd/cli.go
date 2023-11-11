package main

import (
	"fmt"
	"strconv"

	"github.com/wingelote/aisprid-alerting/internal/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var conf config.Config

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start alerting application server",
	Run: func(cmd *cobra.Command, args []string) {
		NewServer(conf)
	},
}

var cmdClient = &cobra.Command{
	Use:   "client",
	Short: "Client to interact with alerting application",
	Args:  cobra.ExactArgs(1),
}

var cmdSendClient = &cobra.Command{
	Use:   "send [stack] [resources]",
	Short: "Send resources usage",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		resourcesValue, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			logrus.Error("resources must be a valid number")
			return
		}

		resources := float32(resourcesValue)
		SendMetrics(args[0], resources)
	},
}

var cmdAlertClient = &cobra.Command{
	Use:   "alert [option]",
	Short: "Interact with alerts",
	Args:  cobra.ExactArgs(1),
}

var cmdListAlert = &cobra.Command{
	Use:   "list",
	Short: "List alerts",
	Run: func(cmd *cobra.Command, args []string) {
		GetAlertHistory()
	},
}

func NewCLI(configuration config.Config) {
	conf = configuration
	NewClient(configuration)

	cmdServe.Flags().Int32("port", int32(conf.Server.Port), fmt.Sprintf("specify alternate port (default: %d)", int32(conf.Server.Port)))
	cmdClient.Flags().Int32("port", int32(conf.Server.Port), fmt.Sprintf("specify alternate port (default: %d)", int32(conf.Server.Port)))
	cmdClient.Flags().String("host", conf.Server.Host, fmt.Sprintf("specify alternate host (default: %s)", conf.Server.Host))
	cmdAlertClient.AddCommand(cmdListAlert)
	cmdClient.AddCommand(cmdSendClient, cmdAlertClient)

	var rootCmd = &cobra.Command{Use: "aisprid-alerting"}
	rootCmd.AddCommand(cmdClient, cmdServe)

	rootCmd.Execute()
}
