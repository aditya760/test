// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"cognizant.com/codeblue/actions"
	"github.com/spf13/cobra"
)

// generateJsclientCmd represents the generateJsclient command
var generateJsclientCmd = &cobra.Command{
	Use:   "generate-jsclient",
	Short: "Generate Javascript Client",
	Long: `Generate Javascript Client according to Architecture file

For example:
codeblue generate-jsclient -a arch.yml
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("generate js client called")
		log.Debug("archFile = " + archFile)
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
			p.GenerateJSClient(archFileContext)
		} else {
			log.Critical("Architecture File Required!")
		}
	},
}

func init() {
	RootCmd.AddCommand(generateJsclientCmd)
	generateJsclientCmd.Flags().StringVarP(&archFile, "arch", "a", "", "An architecture layout file")
}
