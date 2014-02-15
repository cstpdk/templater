package main

import (
        "testing"
)

func TestNew(t *testing.T) {
        var conf Config

        MakeConfig(&conf, "test_data/config.json")

        if conf.Templates_dir != "templates" {
                t.Fatalf("Wrong templates dir read\n It was %s", conf.Templates_dir)
        }
}
