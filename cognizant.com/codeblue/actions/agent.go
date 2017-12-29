package actions

import (
  "os/exec"
  "strings"
  "io/ioutil"
  "encoding/base64"
  "os"
  "github.com/mattn/go-shellwords"
  "github.com/wsxiaoys/terminal/color"
)

type Agent interface {
  RunCmd(cmd string) ([]byte, error)
  StoreFile(path string, contentBase64 string) error
  DeleteFile(path string) error
  ReadFile(path string) ([]byte, error)
  // StoreFile(content byte[], path string)
}

type AgentImpl struct {}

func NewAgent() Agent {
  return AgentImpl{}
}

func shouldDisplayOutput() bool{
  envVar := os.Getenv("HIDE_OUTPUT")
  return envVar == ""
}

func (a AgentImpl) DeleteFile(path string) error {
  return os.Remove(path)
}

func (a AgentImpl) RunCmd(strCmd string) ([]byte, error) {
  if shouldDisplayOutput() {
    color.Printf("@{b}Running @{g}%s\n", strCmd)
  }

  strCmd = strings.Trim(strCmd, " ")
  if len(strCmd) == 0 {
      return []byte{}, nil
  }

  args, err := shellwords.Parse(strCmd)
  if err != nil {
    return []byte{}, err
  }

  log.Debug("args = %s", args)
  log.Debug("last arg = %s", args[len(args) - 1])

  cmd := exec.Command(args[0], args[1:]...)
  return cmd.Output()
}

func (a AgentImpl) StoreFile(path string, contentBase64 string) error {
  if shouldDisplayOutput() {
    color.Printf("@{m}Saving @{g}%s\n", path)
  }

  data, err := base64.StdEncoding.DecodeString(contentBase64)
  if err != nil {
    return err
  }
  ioutil.WriteFile(path, data, 0644)

  return nil
}

func (a AgentImpl) ReadFile(path string) ([]byte, error) {
    return ioutil.ReadFile(path)
}
