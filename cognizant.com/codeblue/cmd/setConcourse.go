
package cmd

import (
    "os"
    "fmt"
    "github.com/spf13/cobra"
    "os/exec"
    "io/ioutil"
    "path/filepath"
)

// var pipelineName string
// var targetName string
// var paramFileName string
// var flytargetName string
// var flyConcourseUrl string
// var runConcourse string
var Concourse_Username string
var Concourse_Password string
var setConcourseCmd = &cobra.Command{
    Use:   "setconcourse",
    Short: "Local install of Concourse from a docker repository ",
    Long: `Set up Concourse from a Docker repository

      Pre-requistes:
      --------------
         Docker should be installed
    `,
    Run: func(cmd *cobra.Command, args []string) {


    dockerFilePath, _ := filepath.Abs("../src/cognizant.com/codeblue/config/docker-compose.yml")
    cli_commands_FilePath, _ := filepath.Abs("../src/cognizant.com/codeblue/config/concourse_commands.sh")
    //paramFilePath, _ := filepath.Abs("../src/cognizant.com/codeblue/config/params.yml")
    read, err := ioutil.ReadFile(dockerFilePath)
    if err != nil {
    panic(err)
    }

    docker_compose_yml := string(read)

    dir, err := os.Getwd()
    filePath := dir + "/" + "docker-compose.yml"

    err = ioutil.WriteFile(filePath, []byte(docker_compose_yml), 0644)
    if err != nil {
        panic(err)
    }

    out, err := exec.Command("/bin/sh", cli_commands_FilePath).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("output is %s\n", out)

    // keys := exec.Command("mkdir ","-p" ,"keys/web keys/worker")
    // keys.Stdin = os.Stdin
    // keys_out, trigger_err := keys.CombinedOutput()
    // if err != nil {
    //     panic(trigger_err)
    // }
    // fmt.Printf("combined out:\n%s\n", string(trigger_out))

    // tsa_host_key := exec.Command("ssh ","-keygen" ,"-t","rsa","-f","./keys/web/tsa_host_key -N ''")
    // tsa_host_key.Stdin = os.Stdin
    // tsa_host_key_out, tsa_host_key_err := tsa_host_key.CombinedOutput()
    // if err != nil {
    //     panic(tsa_host_key_err)
    // }
    // fmt.Printf("combined out:\n%s\n", string(tsa_host_key_out))

    Con_Url_cmd := exec.Command("export ","CONCOURSE_EXTERNAL_URL=" ,"http://54.169.190.81:8181")
    Con_Url_cmd.Stdin = os.Stdin
    Con_Url_out, Con_Url_err := Con_Url_cmd.CombinedOutput()
    if err != nil {
        panic(Con_Url_err)
    }
    fmt.Printf("combined out external url:\n%s\n", string(Con_Url_out))

    Con_Username_cmd := exec.Command("export ","CONCOURSE_BASIC_AUTH_USERNAME=" ,Concourse_Username )
    Con_Username_cmd.Stdin = os.Stdin
    Con_Username_out, Con_Username_err := Con_Username_cmd.CombinedOutput()
    if err != nil {
        panic(Con_Username_err)
    }
    fmt.Printf("combined out username:\n%s\n", string(Con_Username_out))
    
    Con_Password_cmd := exec.Command("export ","CONCOURSE_BASIC_AUTH_PASSWORD=" ,Concourse_Password)
    Con_Password_cmd.Stdin = os.Stdin
    Con_Password_out, Con_Password_err := Con_Password_cmd.CombinedOutput()
    if err != nil {
        panic(Con_Password_err)
    }
    fmt.Printf("combined out password:\n%s\n", string(Con_Password_out))



    docker_Compose_cmd := exec.Command("docker-compose","up")
    docker_Compose_cmd.Stdin = os.Stdin
    docker_Compose_out, docker_Compose_err := docker_Compose_cmd.CombinedOutput()
    if err != nil {
        panic(docker_Compose_err)
    }
    fmt.Printf("combined out docker compose up:\n%s\n", string(docker_Compose_out))

    },
}

func init() {
    RootCmd.AddCommand(setConcourseCmd)
    setConcourseCmd.Flags().StringVarP(&Concourse_Username, "name", "n", "", "Concourse Username")
    setConcourseCmd.Flags().StringVarP(&Concourse_Password, "target", "p", "", "Concourse Password")
    //setConcourseCmd.Flags().StringVarP(&paramFileName, "paramFileName", "p", "", "Param File Name")
}
