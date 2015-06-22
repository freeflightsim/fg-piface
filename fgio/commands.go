
package fgio

import (
	//"encoding/json"
)


//{"command":"get","node":"/instrumentation/comm/station-name"}
//{"command":"get","node":"/instrumentation/comm[1]/frequencies/selected-mhz"}
//{"command":"addListener","node":"/instrumentation/comm/station-name"}




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


