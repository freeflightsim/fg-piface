
package fgio

import (
	//"encoding/json"
)


type AddListenerCommand struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
	//Value string ` json:"value" `
}

func NewAddListenerCommand(node string) AddListenerCommand {
	c := AddListenerCommand{Cmd: "addListener", Node: node}
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
