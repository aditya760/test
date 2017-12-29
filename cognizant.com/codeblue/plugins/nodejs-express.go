package plugins

import (
  "os"
  "path/filepath"
)

type NodejsExpress struct {}

func (n NodejsExpress) GenerateJSClient(archFileContent []byte) {
  log.Info("NodeJS Express Generate JS Client Called")

  dir, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }

  baseDir := filepath.Base(dir)

  backendClient := NewBackendClient(nil)
  actions, err := backendClient.GenerateClient("/nodejs/express/generate-jsclient/" + baseDir, archFileContent)

  if err != nil {
    log.Fatal(err)
  }

  for i:=0; i<len(actions); i++ {
    actions[i].Execute()
  }
}

func (n NodejsExpress) InitWithArch(archFileContext []byte) {
  log.Info("NodeJS Express Init Called")

  dir, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }

  baseDir := filepath.Base(dir)

  backendClient := NewBackendClient(nil)
  actions, err := backendClient.InitWithArch("/nodejs/express/init/" + baseDir, archFileContext)

  if err != nil {
    log.Fatal(err)
  }

  for i:=0; i<len(actions); i++ {
    actions[i].Execute()
  }
}

func (n NodejsExpress) Init() {
  log.Info("NodeJS Express Init Called")

  dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

  baseDir := filepath.Base(dir)

  backendClient := NewBackendClient(nil)
  actions, err := backendClient.Get("/nodejs/express/init/" + baseDir)

  if err != nil {
    log.Fatal(err)
  }

  for i:=0; i<len(actions); i++ {
    actions[i].Execute()
  }

}
