package main

import (
	"fmt"
	"html"
	"os"

	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yurifrl/poc-websocket/pkg/config"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var _ = pp.Println
var logrusLogger = logrus.New()

func main() {
	// Set the default log level to WarnLevel
	logrusLogger.SetLevel(logrus.DebugLevel)
	logrusLogger.SetOutput(os.Stderr)

	cfg, err := config.New(logrusLogger)
	if err != nil {
		panic(err)
	}

	var rootCmd = &cobra.Command{
		Use:   "poc-websocket",
		Short: "... TODO",
		Long:  `... TODO`,
		// Do not touch.
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			logrusLogger.WithFields(logrus.Fields{
				"config": cfg.ToString(),
				"name":   cmd.Name(),
			}).Info("Starting CLI")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.AddCommand(execCmd(cfg))

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("\n\033[1;31mERROR\033[0m\n\n\033[1;31m%s\033[0m\n\n", html.UnescapeString(err.Error()))
		// Print with stacktrace in debug
		logrusLogger.WithError(err).Debug()
		os.Exit(1)
	}
}

func execCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "exec",
		Short: "Exec ... TODO",
		Long:  `Exec ... TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.Log().Info("Executing Exec")
			client := NewWebSocketClient(cfg)
			return client.Run()
		},
	}
}
