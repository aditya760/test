package cmd

import (
	"bufio"
	"os"
	"path/filepath"
	"os/user"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("codeblue")

type codebluePlugin interface {

}


var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "codeblue",
	Short: "A quick tool to generate microservices",
	Long: `A quick tool to generate microservices`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
  }

  directory := filepath.Join(usr.HomeDir, ".codeblue")
  if _, err := os.Stat(directory); os.IsNotExist(err) {
    os.Mkdir(directory, 0700)
  }
	logFilePath := filepath.Join(directory, "codeblue.log")
  f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE, 0644)
  if err != nil {
    log.Fatal(err)
  }
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	backend1 := logging.NewLogBackend(w, "", 0644)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.DEBUG, "codeblue")
	logging.SetBackend(backend1Leveled, backend1Formatter)


	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.codeblue.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".codeblue" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".codeblue")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}
