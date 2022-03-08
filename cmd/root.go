package cmd

import (
	"crabada-bot/pkg/constant"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crabada-bot",
	Short: "Crabada cli bot",
	Long:  "Crabada cli bot",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	zerolog.TimestampFieldName = constant.TimeLogField
	zerolog.TimeFieldFormat = constant.TimeLogFieldFormat
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Logger()
}
