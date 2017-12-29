package plugins

import (
  "os"
  "path/filepath"
)

type JavaSpring struct {}

func (n JavaSpring) GenerateJSClient(archFileContent []byte) {
  log.Info("Java Spring Generate JS Client Called")

  dir, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }

  baseDir := filepath.Base(dir)

  backendClient := NewBackendClient(nil)
  actions, err := backendClient.GenerateClient("/java/spring/generate-jsclient/" + baseDir, archFileContent)

  if err != nil {
    log.Fatal(err)
  }

  for i:=0; i<len(actions); i++ {
    actions[i].Execute()
  }
}

func (n JavaSpring) InitWithArch(archFileContext []byte) {
  log.Info("Java Spring Init Called")

  dir, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }

  baseDir := filepath.Base(dir)

  backendClient := NewBackendClient(nil)
  actions, err := backendClient.InitWithArch("/java/spring/init/" + baseDir, archFileContext)

  if err != nil {
    log.Fatal(err)
  }

  for i:=0; i<len(actions); i++ {
    actions[i].Execute()
  }
}

func (n JavaSpring) Init() {
  log.Info("Java Spring Init Called")

  dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

  baseDir := filepath.Base(dir)

  backendClient := NewBackendClient(nil)
  actions, err := backendClient.Get("/java/spring/init/" + baseDir)

  if err != nil {
    log.Fatal(err)
  }

  for i:=0; i<len(actions); i++ {
    actions[i].Execute()
  }

}
