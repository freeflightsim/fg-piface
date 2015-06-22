
package fgio

import (
	//"encoding/json"
)


type AddListenerCommand struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
	//Value string ` json:"value" `
}

func NewAddListenerCmd(node string) AddListenerCommand {
	c := AddListenerCommand{Cmd: "addListener", Node: node}
	return c
}

type RemoveListenerCommand struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
	//Value string ` json:"value" `
}

func NewRemoveListenerCmd(node string) RemoveListenerCommand {
	c := RemoveListenerCommand{Cmd: "removeListener", Node: node}
	return c
}


type Command struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
	Value string ` json:"value" `
}


type Frame struct {
	Node string ` json:"path" `
	Name string ` json:"name" `
	Type string ` json:"type" `
	Index int ` json:"index" `
	Value interface{} ` json:"value" `

}
