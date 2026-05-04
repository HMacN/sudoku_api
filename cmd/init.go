package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sudoku_api/services/command_hooks/example"

	"github.com/pkg/errors"
	"github.com/samber/slog-multi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

const (
	flagVerbose    string = "verbose"
	flagConfigFile string = "config"
	flagLogLevel   string = "log_level"
	flagLogFile    string = "log_file"
)

var (
	cfgFile string
	logger  *slog.Logger
	app     *fx.App
)

func init() {
	// Persistent flags available to all commands
	rootCmd.PersistentFlags().BoolP(flagVerbose, "v", false, "verbose output (default value: false)")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, flagConfigFile, "c", "./config.yaml", "config file path and name (default name and location: ./config.yaml)")
	rootCmd.PersistentFlags().StringP(flagLogLevel, "l", "info", "log level")
	rootCmd.PersistentFlags().StringP(flagLogFile, "o", "", "log file path and name (default name and location: ./logs.txt)")

	// Add subcommands
	// TODO: Add subcommand for running as a server, and add flags for port number, &c.
	rootCmd.AddCommand(example.NewCommand())
}

func initialiseConfig(cmd *cobra.Command) error {
	viper.SetEnvPrefix("SUDOKU")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for a config file with the name "config.yaml".
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		// It's okay if the config file doesn't exist.
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return errors.Wrap(err, "failed to read config file")
		}
	}

	// Bind Cobra flags to Viper - makes the flag values available through Viper.
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return errors.Wrap(err, "failed to bind flags")
	}

	fmt.Println("Configuration initialised. Using config file:", viper.ConfigFileUsed())
	return nil
}

func initialiseLogger() error {

	var handlers []slog.Handler
	isVerbose := viper.IsSet(flagVerbose)
	if isVerbose {
		fmt.Println("Log output set to verbose mode")
		handlers = append(handlers, slog.NewTextHandler(os.Stdout, nil))
	}

	logFile := viper.GetString(flagLogFile)
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			return errors.Wrap(err, "Could not open log file")
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				return
			}
		}(f)

		handlers = append(handlers, slog.NewJSONHandler(f, nil))
	}

	logLevel := slog.LevelInfo
	switch strings.ToUpper(viper.GetString(flagLogLevel)) {
	case "DEBUG":
		logLevel = slog.LevelDebug
		break
	case "INFO":
		logLevel = slog.LevelInfo
		break
	case "WARN":
		logLevel = slog.LevelWarn
		break
	case "ERROR":
		logLevel = slog.LevelError
		break
	default:
		logLevel = slog.LevelInfo
		break
	}

	logger = slog.New(slogmulti.Fanout(handlers...))
	slog.SetLogLoggerLevel(logLevel)
	slog.SetDefault(logger)
	slog.Info(fmt.Sprintf("Logger initialised (verbose=%t, level=%s, logfile=%s)", isVerbose, logLevel, logFile))

	return nil
}
