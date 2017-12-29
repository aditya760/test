package common

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "os/user"
  "bytes"
  "io"
  "path/filepath"
  "github.com/op/go-logging"
)

type CodeBlueConfig struct{
  Url string
}

var log = logging.MustGetLogger("codeblue")

func StoreConfig(config CodeBlueConfig) error {
  usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
  }

  directory := filepath.Join(usr.HomeDir, ".codeblue")
  if _, err := os.Stat(directory); os.IsNotExist(err) {
    os.Mkdir(directory, 0700)
  }

  fileContent, err := json.Marshal(config)
  if err != nil {
    log.Critical("Error found while marshal config: %s", err)
    return err
  }

  configFilePath := filepath.Join(directory, "config.json")
  return ioutil.WriteFile(configFilePath, fileContent, 0600)

}

func DefaultConfig() CodeBlueConfig {
  return CodeBlueConfig{
    Url: "http://localhost:3000",
  }
}

func LoadConfig() (CodeBlueConfig, error) {
  usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
  }

  directory := filepath.Join(usr.HomeDir, ".codeblue")
  if _, err := os.Stat(directory); os.IsNotExist(err) {
    log.Info("Cannot found CodeBlue Configuration file. Using default api target instead")
    return DefaultConfig(), nil
  }

  buf := bytes.NewBuffer(nil)
  configFilePath := filepath.Join(directory, "config.json")
  f, err := os.Open(configFilePath)
  if err != nil {
    return DefaultConfig(), err
  }
  io.Copy(buf, f)
  f.Close()

  result := CodeBlueConfig{}
  err = json.Unmarshal(buf.Bytes(), &result)
  if err != nil {
    return DefaultConfig(), err
  }

  return result, nil
}
