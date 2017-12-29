package actions

type FileStoreAction struct {
  FilePath string
  Content string
}

func (a FileStoreAction) Execute() error {
  agent := NewAgent()
  return agent.StoreFile(a.FilePath, a.Content)
}
