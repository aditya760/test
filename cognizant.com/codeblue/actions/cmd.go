package actions

import "github.com/op/go-logging"

var log = logging.MustGetLogger("codeblue")

type CmdAction struct {
  Cmd string
}

func (a CmdAction) Execute() error {
  agent := NewAgent()
  _, err := agent.RunCmd(a.Cmd)
  return err
}
