
package fgio

import (
	//"encoding/json"
)


//{"command":"get","node":"/instrumentation/comm/station-name"}
//{"command":"get","node":"/instrumentation/comm[1]/frequencies/selected-mhz"}
//{"command":"addListener","node":"/instrumentation/comm/station-name"}

type Command struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
}
type CommandVal struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
	Value string ` json:"value" `
}


func AddListenerCmd(node string) Command {
	return Command{Cmd: "addListener", Node: node}
}

func RemoveListenerCmd(node string) Command {
	return Command{Cmd: "removeListener", Node: node}
}


func GetCmd(node string) Command {
	return Command{Cmd: "get", Node: node}
}

func SetCmd(node string, val string) CommandVal {
	return CommandVal{Cmd: "get", Node: node, Value: val}
}

/*
type DEADRemoveListenerCommand struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
	//Value string ` json:"value" `
}
type deadGetCommand struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
}

type DEADAddListenerCommand struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
	//Value string ` json:"value" `
}
*/
