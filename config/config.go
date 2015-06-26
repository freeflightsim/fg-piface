

package config


import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"errors"

)


type Config struct {
	Model string 	` yaml:"model" `
	DInPins []InputPin	` yaml:"digital_inputs" `
	AInPins []InputPin	` yaml:"analog_inputs" `
	DOupPins []OutputPin	` yaml:"digital_outputs" `
	zxOutputs map[int]string
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

	for _, p := range me.OutputDefs {
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
