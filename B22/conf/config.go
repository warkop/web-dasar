package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var shared *Config

type Config struct {
	Server struct {
		Port         int           `json:"post"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
	} `json:"server"`

	Log struct {
		Verbose bool `json:"verbose"`
	} `json:"log"`
}

func init() {
	if shared != nil {
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	bts, err := ioutil.ReadFile(filepath.Join(basePath, "conf", "conf.json"))
	if err != nil {
		panic(err)
	}

	shared = new(Config)
	err = json.Unmarshal(bts, &shared)
	if err != nil {
		panic(err)
	}
}

//Configuration is
func Configuration() Config {
	return *shared
}
