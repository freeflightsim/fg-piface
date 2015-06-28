

package config


import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	//"errors"
)


// a Flightgear Node
type FgNode struct {
	Id string 		` yaml:"id" `
	Node string 	` yaml:"node" `
	Comment string 	` yaml:"comment" `
	On string 		` yaml:"on" `
	Off string 		` yaml:"off" `
	Toggle bool    ` yaml:"toggle" `
}


// Returns a list of unique input Nodes
func (me *Config) LoadFgNodes() error {
	contents, err_file := ioutil.ReadFile("protocol/flightgear_defs.yaml")
	if err_file != nil {
		fmt.Println("erooor", err_file)
		return err_file
	}
	//fmt.Println("contt", contents)
	//conf := new(Config)
	var nodes[] FgNode
	err_yaml := yaml.Unmarshal(contents, &nodes)
	if err_yaml != nil {
		fmt.Println("err_yaml", err_yaml)
		return err_yaml
	}
	for _, n := range nodes {
		fmt.Println("n=", n)
		me.FgNodes = append(me.FgNodes, n)
	}

	return nil
}
