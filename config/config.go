

package config


import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"errors"

)


type Config struct {
	Model string 	` yaml:"model" `
	Inputs []InputPin	` yaml:"inputs" `
	Outputs []OutputPin	` yaml:"outputs" `
}

func Load(file_path string) (*Config, error) {



	contents, err_file := ioutil.ReadFile(file_path)
	if err_file != nil {
		return nil, err_file
	}

	conf := new(Config)
	err_yaml := yaml.Unmarshal(contents, &conf)
	if err_yaml != nil {
		return nil, err_yaml
	}
	/*
	conf.LedMap = make(map[string]int)
	for _, led := range conf.Leds {
		conf.LedMap[led.Node] = led.Index
	}
	*/
	err := conf.Validate()
	if err != nil {
		return conf, err
	}

	return conf, nil
}


func (me *Config) Validate() error {

	exists := make(map[int]bool)
	mess := ""

	for _, p := range me.Outputs {
		if p.Pin > 7 {
			mess +=  "OutPin " + p.Node + " has index > 7\n"
		}
		_, found := exists[p.Pin]
		if found {
			mess +=  "OutPin " + p.Node + " has duplicate index\n"
		}
		exists[p.Pin] = true
	}
	if mess == "" {
		return nil
	}
	return errors.New(mess)

}

// Returns a list of unique output Nodes
func (me *Config) GetOutputNodes() []string {

	nodes := make(map[string]bool)
	for _, p := range me.Outputs {
		_, found := nodes[p.Node]
		if found == false {
			nodes[p.Node] = true
		}
	}
	lst := make([]string, 0)
	for n, _ := range nodes {
		lst = append(lst, n)
	}
	return lst

}

// Returns a list of unique input Nodes
func (me *Config) GetInputNodes() []string {

	nodes := make(map[string]bool)
	for _, p := range me.Inputs {
		_, found := nodes[p.Node]
		if found == false {
			nodes[p.Node] = true
		}
	}
	lst := make([]string, 0)
	for n, _ := range nodes {
		lst = append(lst, n)
	}
	return lst

}
