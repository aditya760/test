package plugins

import (
  "os/exec"

  "github.com/op/go-logging"
)

var log = logging.MustGetLogger("codeblue")

type CodebluePlugin interface {
  Init()
  InitWithArch([]byte)
  GenerateJSClient([]byte)
}

func CallCommand(cmd string) error {
  return nil
}

func CallCommands(cmds []*exec.Cmd) error {
  for i:=0; i < len(cmds); i++ {
    err := cmds[i].Run()
    if err != nil {
      log.Critical("Found error: %s", err)
      return err
    }
  }
  return nil
}
