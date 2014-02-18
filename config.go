package main

import (
        "encoding/json"
        "io/ioutil"
)

type Config struct {
        Templates_dir      string
        Template_regexp    string
        Default_output_dir string
}

func MakeConfig(conf *Config, configFile string) {
        // Load in some config
        bytes, err := ioutil.ReadFile(configFile)

        if err != nil {
                panic("Couldn't read config file")
        }

        if err = json.Unmarshal(bytes, conf); err != nil {
                panic("Couldn't unmarshal config file")
        }
}
