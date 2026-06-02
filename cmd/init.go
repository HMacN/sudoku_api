package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sudoku_api/config"
	"sudoku_api/config/config_keys"
	"sudoku_api/services/command_hooks/example"
	"sudoku_api/services/command_hooks/run_server"
	"sudoku_api/services/logging"
	"sudoku_api/services/server"

	"github.com/pkg/errors"
	"github.com/samber/slog-multi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var (
	cfgFile string
	logger  *slog.Logger
	app     *fx.App
)

func init() {
	// Persistent flags available to all commands
	rootCmd.PersistentFlags().BoolP(config_keys.Verbose.String(), "v", false, "verbose output (default value: false)")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, config_keys.ConfigurationFile.String(), "c", "./config.yaml", "config file path and name (default name and location: ./config.yaml)")
	rootCmd.PersistentFlags().StringP(config_keys.LogLevel.String(), "l", "info", "log level")
	rootCmd.PersistentFlags().StringP(config_keys.LogFile.String(), "o", "", "log file path and name (default name and location: ./logs.txt)")

	// Add example subcommand
	rootCmd.AddCommand(example.NewCommand())

	// Run as server subcommand
	rootCmd.AddCommand(run_server.NewCommand())

	err := initialiseConfig(rootCmd)
	if err != nil {
		msg := fmt.Sprintf("failed to initialise config with error: %s", err.Error())
		fmt.Println(msg)
		return
	}

	err = initialiseLogger()
	if err != nil {
		msg := fmt.Sprintf("failed to initialise logger with error: %s", err.Error())
		fmt.Println(msg)
		return
	}

	// Set up Dependency Injection
	loggingService, fxLogger := logging.NewLogger()
	loggerProvider := func() logging.LogWrapper { return loggingService }
	app = fx.New(
		fx.WithLogger(func() fxevent.Logger { return fxLogger }),
		fx.Provide(
			server.NewServer,
			server.NewServeMux,
			example.NewService,
			run_server.NewService,
			loggerProvider,
		),
		//fx.Invoke()		// TODO: Should there be a long-running function here that is shut down elsewhere?  How will that interact with other long-running functions i.e. servers?
	)
	initErr := app.Err()
	if initErr != nil {
		msg := fmt.Sprintf("encountered error during dependency injection: %s", initErr.Error())
		loggingService.LogError(msg)
		return
	}

	// TODO: Context?
	// TODO: Channel here for app shutdown?
	app.Start(context.Background())
	app.Done()
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

	// TODO: Check if all config values have been finalised by this point.  If so then FX should probably get set up here.

	return nil
}

func initialiseLogger() error {

	var handlers []slog.Handler
	isVerbose := config.GetBool(config_keys.Verbose)
	if isVerbose {
		fmt.Println("Log output set to verbose mode")
		handlers = append(handlers, slog.NewTextHandler(os.Stdout, nil))
	}

	logFile := config.GetString(config_keys.LogFile)
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
	switch strings.ToUpper(config.GetString(config_keys.LogLevel)) {
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
