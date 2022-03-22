package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Load reads the config file provided and unmarshalls to the cfg variable.
// Typically this will be called by slv services on bootup.
func Load(path string, cfg interface{}) {
	configFile, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		panic(err)
	}

}
