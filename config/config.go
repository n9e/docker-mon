package config

import (
	"fmt"
	"io/ioutil"

	"github.com/toolkits/file"
	"gopkg.in/yaml.v2"
)

type LoggerSection struct {
	Dir       string `yaml:"dir"`
	Level     string `yaml:"level"`
	KeepHours uint   `yaml:"keepHours"`
}

type ConfYaml struct {
	CadvisorUrl    string        `yaml:"cadvisorUrl"`
	DockerSockFile string        `yaml:"dockerSockFile"`
	Identify       string        `yaml:"identify"`
	Step           int           `yaml:"step"`
	Logger         LoggerSection `yaml:"logger"`
}

var (
	G *ConfYaml
)

func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuratio file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("configuration file %s is nonexistent", cfg)
	}

	bs, err := ioutil.ReadFile(cfg)
	if err != nil {
		return fmt.Errorf("read configuration file %s fail %s", cfg, err.Error())
	}

	var c ConfYaml
	err = yaml.Unmarshal(bs, &c)
	if err != nil {
		return fmt.Errorf("parse configuration file %s fail %s", cfg, err.Error())
	}

	G = &c
	if G.Identify == "" {
		G.Identify = "hostname"
	}

	if G.CadvisorUrl == "" {
		G.CadvisorUrl = "http://127.0.0.1:4194/"
	}

	if G.DockerSockFile == "" {
		G.DockerSockFile = "/var/run/docker.sock"
	}

	if G.Step == 0 {
		G.Step = 10
	}

	return nil
}
