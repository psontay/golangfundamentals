package manager

import "fmt"

type CommandPromptManager struct {
}

func NewCommandPromptManager() CommandPromptManager {
	return CommandPromptManager{}
}

func (commandPromptManager *CommandPromptManager) ReadJSON(data interface{}) error {
	fmt.Println(data)
	return nil
}

func (commandPromptManager *CommandPromptManager) WriteJSON(data interface{}) error {
	fmt.Println(data)
	return nil
}
