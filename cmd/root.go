package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	util "github.com/zeten30/metrics-agent/util"
	cldyVersion "github.com/zeten30/metrics-agent/version"
)

// RootCmd is the cobra root command to be executed
// nolint:revive
var RootCmd = &cobra.Command{
	Use:              "metrics-agent [command] [flags]",
	Short:            "Starts the Cloudability Metrics Agent",
	Long:             `Starts the Cloudability Metrics Agent for the configured metrics collectors and polling interval.`,
	Args:             cobra.MinimumNArgs(1),
	TraverseChildren: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return util.SetupLogger()
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	RootCmd.PersistentFlags().String(
		"log_level",
		"INFO",
		"Log level to run the agent at (INFO,WARN,DEBUG)",
	)

	RootCmd.PersistentFlags().String(
		"log_format",
		"PLAIN",
		"Format for log output (JSON,TXT)",
	)

	// set version flag
	RootCmd.Version = cldyVersion.VERSION

	//nolint gosec
	_ = viper.BindPFlag("log_level", RootCmd.PersistentFlags().Lookup("log_level"))
	_ = viper.BindPFlag("log_format", RootCmd.PersistentFlags().Lookup("log_format"))

}

// Execute metrics-agent with arguments
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		log.Fatalln("Unable to execute :", err)
	}
}
