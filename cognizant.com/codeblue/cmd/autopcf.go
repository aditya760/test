
package cmd

import (
    "os"
    "fmt"
    "github.com/spf13/cobra"
    "os/exec"
    "io/ioutil"
    "path/filepath"
    "cognizant.com/codeblue/lib"
)

var pipelineName string
var targetName string
var paramFileName string
var flytargetName string
var flyConcourseUrl string
var runConcourse string

var autoPcfCmd = &cobra.Command{
    Use:   "autopcf",
    Short: "Deploy PCF instance on specified target",
    Long: `Deploy PCF instance on specified target (AWS, AZURE)

      Pre-requistes:
      --------------
        1) AWS Credentials - (ACCESS KEY and SECRET KEY)
        2) Domain Name and Hosted Zone ID
        3) AWS Certificate ARN (AMAZON RESOURE NAME)

    `,
    Run: func(cmd *cobra.Command, args []string) {

    var paramFilePath string

    fmt.Printf("Enter Flyway Concurl Url\n")
    fmt.Scanln(&flyConcourseUrl)

    fmt.Printf("Enter Flyway target\n")
    fmt.Scanln(&flytargetName)

    fmt.Printf("Enter Flyway target Login Credentials for deployment\n")
    pipeline_login_cmd := exec.Command("fly","-t" + flytargetName,"login", "-c"+flyConcourseUrl)
    pipeline_login_cmd.Stdin = os.Stdin

    pipeline_login_out, err := pipeline_login_cmd.CombinedOutput()

    if err != nil {
      panic(err)
    }
    fmt.Printf("combined out:\n%s\n", string(pipeline_login_out))

    if paramFileName == ""     {
       lib.UpdateParams()
       dir, err := os.Getwd()
       if err != nil {
         panic(err)
       }
       paramFilePath = dir + "/" + "params.yml"
       fmt.Printf("\n****************************************")
       fmt.Printf("\n********  Generate Param file **********")
       fmt.Printf("\n****************************************")
       read, err := ioutil.ReadFile(paramFilePath)
       if err != nil {
         panic(err)
       }
       fmt.Printf(string(read))
       fmt.Printf("\n****************************************")
       fmt.Printf("\n******* Generated file is available in\n")
       fmt.Printf(paramFilePath)
    } else {
       paramFilePath = paramFileName
    }

    fmt.Printf("\nAre you want to run the fly pipeline now or later :::  y/N\n")
    fmt.Scanln(&runConcourse)

    if (runConcourse != "y") {
      os.Exit(1)
    }
    // command to start the pcf pipeline using fly command
    fmt.Printf("Apply Configuration changes for setting up pipeline (y/N) :\n")

    pipelineFilePath, _ := filepath.Abs("../src/cognizant.com/codeblue/config/pipeline.yml")
    //paramFilePath, _ := filepath.Abs("../src/cognizant.com/codeblue/config/params.yml")

    pipeline_cmd := exec.Command("fly","-t" + flytargetName,"set-pipeline","-p", pipelineName ,"-c" + pipelineFilePath,"-l" + paramFilePath)
    pipeline_cmd.Stdin = os.Stdin
    out, err := pipeline_cmd.CombinedOutput()

    if err != nil {

        panic(err)
    }
    fmt.Printf("combined out:\n%s\n", string(out))

    // command to unpause the pipeline
    un_pause_cmd := exec.Command("fly","-t" + flytargetName,"unpause-pipeline", "-p", pipelineName)
    un_pause_cmd.Stdin = os.Stdin
    un_pause_out, un_pause_err := un_pause_cmd.CombinedOutput()
    if err != nil {
        panic(un_pause_err)
    }
    fmt.Printf("combined out:\n%s\n", string(un_pause_out))

    trigger_cmd := exec.Command("fly","-t" + flytargetName,"trigger-job", "-j", pipelineName + "/bootstrap-terraform-state")
    trigger_cmd.Stdin = os.Stdin
    trigger_out, trigger_err := trigger_cmd.CombinedOutput()
    if err != nil {
        panic(trigger_err)
    }
    fmt.Printf("combined out:\n%s\n", string(trigger_out))

    trigger_ci_cmd := exec.Command("fly","-t" + flytargetName,"trigger-job", "-j", pipelineName + "/create-infrastructure")
    trigger_ci_cmd.Stdin = os.Stdin
    trigger_ci_out, trigger_ci_err := trigger_ci_cmd.CombinedOutput()
    if err != nil {
        panic(trigger_ci_err)
    }
    fmt.Printf("combined out:\n%s\n", string(trigger_ci_out))

    },
}

func init() {
    RootCmd.AddCommand(autoPcfCmd)
    autoPcfCmd.Flags().StringVarP(&pipelineName, "name", "n", "", "Pipeline Name")
    autoPcfCmd.Flags().StringVarP(&targetName, "target", "t", "", "AWS, Azure, GCP")
    autoPcfCmd.Flags().StringVarP(&paramFileName, "paramFileName", "p", "", "Param File Name")
}
