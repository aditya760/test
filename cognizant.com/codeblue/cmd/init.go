package cmd

import (
	"cognizant.com/codeblue/plugins"
	"cognizant.com/codeblue/actions"
	"github.com/spf13/cobra"
	"fmt"
	"strings"
	"errors"
)

var archFile string
var language string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new microservice",
	Long: `The init command initializes a new microservice`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("init called")
		log.Debug("archFile = " + archFile)
		log.Debug("language = " + language)
		p, err := selectPlugin();
		if err != nil {
			log.Fatal(err)
		}

		if archFile != "" {
			agent := actions.AgentImpl{}
			archFileContext, err := agent.ReadFile(archFile)
			if err != nil {
				log.Critical("Found error: %s", err)
			}
			p.InitWithArch(archFileContext)
		} else {
			p.Init();
		}
	},
}

func selectPlugin() (plugins.CodebluePlugin, error) {
	language = strings.ToLower(language)
	fmt.Println("language = " + language)
	if language == "java" {
		return plugins.JavaSpring{}, nil
	} else if language == "nodejs" {
		return plugins.NodejsExpress{}, nil
	} else {
		errorMsg := "Unknown Language: " + language
		fmt.Println(errorMsg)
		return plugins.JavaSpring{}, errors.New(errorMsg)
	}

}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&archFile, "arch", "a", "", "Initialize microservice with an architecture layout file")
	initCmd.Flags().StringVarP(&language, "language", "l", "java", "Choose a language to initialize the microservice with. Default with Java")
}
