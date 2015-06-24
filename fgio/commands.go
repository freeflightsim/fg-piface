
package fgio

import (
	//"encoding/json"
)

type Command struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
	Value string ` json:"value,omitempty" `

}
/*
type CommandVal struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `

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
	return CommandVal{Cmd: "set", Node: node, Value: val}
}
*/
