package cmd

import (
	"github.com/mhshajib/oasis_boilerplate/pkg/cmd/migration"
	"github.com/mhshajib/oasis_boilerplate/pkg/cmd/seeder"
	"os"

	"github.com/mhshajib/oasis_boilerplate/pkg/config"
	"github.com/mhshajib/oasis_boilerplate/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// cfgFile store the configuration file name
	cfgFile                 string
	verbose, prettyPrintLog bool

	// rootCmd is the root command of backup service
	rootCmd = &cobra.Command{
		Use:   "projectName",
		Short: "projectName service provide deals & offer releated core functionalities",
		Long:  `projectName service provide deals & offer releated core functionalities`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(migration.MigrationCmd)
	rootCmd.AddCommand(seeder.SeederCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")
	rootCmd.PersistentFlags().BoolVarP(&prettyPrintLog, "pretty", "p", false, "pretty print verbose/log")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// set the value to viper config
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func initConfig() {
	log.Info("Loading configurations")
	if err := config.Init(cfgFile); err != nil {
		log.Warn("Failed to load configuration")
		log.Fatal(err)
	}
	log.Info("Configurations loaded successfully!")

	// Log as JSON instead of the default ASCII formatter.
	log.SetLogFormatter(&logrus.JSONFormatter{
		PrettyPrint: prettyPrintLog,
	})

	// by defualt only log the warning severity or above.
	log.SetLogLevel(logrus.WarnLevel)
	if config.HttpApp().Verbose { // if verbose set true show trace level
		log.SetLogLevel(logrus.TraceLevel)
	}
	if verbose { // if -v flag pass override previous value
		log.SetLogLevel(logrus.TraceLevel)
	}
}
