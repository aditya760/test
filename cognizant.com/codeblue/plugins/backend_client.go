package plugins

import (
  "net/http"
  "net/url"
  "io"
  "io/ioutil"
  "encoding/json"
  "encoding/base64"
  "cognizant.com/codeblue/actions"
  "cognizant.com/codeblue/common"
  "bytes"
)

type InitWithArchBody struct {
  Content string
}

type BackendClient interface {
  HandleShake(url string) (bool, error)
  Get(url string) ([]Action, error)
  InitWithArch(url string, content []byte) ([]Action, error)
  GenerateClient(url string, content []byte) ([]Action, error)
}

type HttpClient interface {
    Do(req *http.Request) (*http.Response, error)
    Get(url string) (resp *http.Response, err error)
    Head(url string) (resp *http.Response, err error)
    Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
    PostForm(url string, data url.Values) (resp *http.Response, err error)
}

type BackendClientImpl struct {
  client HttpClient
}

type HandShake struct {
  Name string
  Version string
}

func (c BackendClientImpl) HandleShake(url string) (bool, error) {
  log.Debugf("Establishing Handshake with %s", url);
  res, err := c.client.Get(url)
  if err != nil {
		log.Critical("Got error: %s", err)
    return false, err
	}
	body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
		log.Critical("Got error: %s", err)
    return false, err
	}

  handShake := &HandShake{}
  err = json.Unmarshal(body, handShake)
  if err != nil {
		log.Critical("Got error: %s", err)
    return false, err
	}

  if(handShake.Name == "codeblue" && handShake.Version == "0.0.0") {
    return true, nil
  } else {
    return false, nil
  }
}

func (c BackendClientImpl) GenerateClient(url string, content []byte) ([]Action, error) {
  encodedContent := base64.StdEncoding.EncodeToString(content)

  log.Debug("encodedContent = " + encodedContent)

  body := InitWithArchBody{
    Content: encodedContent,
  }

  jsonBody, err := json.Marshal(body)
  if err != nil {
		log.Fatal(err)
	}

  log.Debug("jsonBody = " + string(jsonBody))
  config, err := common.LoadConfig()
  if err != nil {
    log.Fatal(err)
  }

  url = config.Url + url
  res, err := c.client.Post(url, "Application/JSON", bytes.NewReader(jsonBody))
  if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
		log.Fatal(err)
	}

  return parse(robots)
}

func (c BackendClientImpl) InitWithArch(url string, content []byte) ([]Action, error) {
  encodedContent := base64.StdEncoding.EncodeToString(content)

  log.Debug("encodedContent = " + encodedContent)

  body := InitWithArchBody{
    Content: encodedContent,
  }

  jsonBody, err := json.Marshal(body)
  if err != nil {
		log.Fatal(err)
	}

  log.Debug("jsonBody = " + string(jsonBody))
  config, err := common.LoadConfig()
  if err != nil {
    log.Fatal(err)
  }

  url = config.Url + url
  res, err := c.client.Post(url, "Application/JSON", bytes.NewReader(jsonBody))
  if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
		log.Fatal(err)
	}

  return parse(robots)
}


func (c BackendClientImpl) Get(url string) ([]Action, error) {
  config, err := common.LoadConfig()
  if err != nil {
    log.Fatal(err)
  }

  url = config.Url + url
  res, err := c.client.Get(url)
  if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
		log.Fatal(err)
	}

  return parse(robots)
}

func parse(strScript []byte) ([]Action, error){
  var data []map[string]interface{}

  if err := json.Unmarshal(strScript, &data); err != nil {
    panic(err)
  }

  var result []Action

  for i:=0; i<len(data); i++ {
    actionType := data[i]["type"]
    actionContent := data[i]["content"].(map[string]interface{})

    if actionType == "storeFile" {
      newAction := actions.FileStoreAction{
        FilePath: actionContent["file-path"].(string),
        Content: actionContent["content"].(string),
      }

      result = append(result, newAction)
    } else if actionType == "cmd" {
      newAction :=  actions.CmdAction{
        Cmd: actionContent["cmd"].(string),
      }
      result = append(result, newAction)
    }
  }

  return result, nil
}

func NewBackendClient(c HttpClient) BackendClient {
  if c != nil {
    return BackendClientImpl{
      client: c,
    }
  } else {
    return BackendClientImpl{
      client: &http.Client{},
    }
  }
}


type Action interface {
  Execute() error
}
