package config

import (
	"gopkg.in/ini.v1"
	"strings"
)

type Config struct {
	ConfigReader *ini.File // config reader
}

// load app config
func (c *Config) Load() error {
	if c.ConfigReader == nil {
		conf, err := ini.Load("config/app.ini")
		if err != nil {
			return err
		}
		runMode := conf.Section("").Key("run_mode").String()
		if len(runMode) > 0 {
			env := "config/env/"+runMode+".ini"
			conf.Append(env)
		}
		c.ConfigReader = conf
	}
	return nil
}

// loading the config for page
func (c *Config) LoadPageConfig(page string) (*ini.File, error)  {
	conf, err := ini.Load("config/panges/"+page+"/app.ini")
	if err != nil {
		return nil,err
	}
	return conf, nil
}

func (c *Config) GetReader() *ini.File {
	if c.ConfigReader == nil {
		return nil
	}
	return c.ConfigReader
}
// split string  "Section.Key"
func (c *Config) getSplitKey(keys string) *ini.Key {
	var section, key string
 	arr := strings.Split(keys, ".")
 	if len(arr) == 1 {
		section = ""
		key = arr[0]
	}
	if len(arr) == 2 {
		section = arr[0]
		key = arr[1]
	}
	s := c.ConfigReader.Section(section).Key(key)
	return s
}

// get string
func (c *Config) GetString(keys string) string {
	s := c.getSplitKey(keys)
	if s == nil {
		return ""
	}
	return s.String()
}

// get int32
func (c *Config) GetInt32(keys string) int32 {
	s := c.getSplitKey(keys)
	if s == nil {
		return 0
	}
	vInt, _ := s.Int()
	return int32(vInt)
}

// get uint32
func (c *Config) GetUint32(keys string) uint32 {
	s := c.getSplitKey(keys)
	if s == nil {
		return 0
	}
	vInt, _ := s.Uint()
	return uint32(vInt)
}

// get Int64
func (c *Config) GetInt64(keys string) int64 {
	s := c.getSplitKey(keys)
	if s == nil {
		return 0
	}
	vInt, _ := s.Int64()
	return vInt
}

// get Uint64
func (c *Config) GetUint64(keys string) uint64 {
	s := c.getSplitKey(keys)
	if s == nil {
		return 0
	}
	vInt, _ := s.Uint64()
	return vInt
}

// get float32
func (c *Config) GetFloat32(keys string) float32 {
	s := c.getSplitKey(keys)
	if s == nil {
		return 0
	}
	vInt, _ := s.Float64()
	return float32(vInt)
}

// get float64
func (c *Config) GetFloat64(keys string) float64 {
	s := c.getSplitKey(keys)
	if s == nil {
		return 0
	}
	vInt, _ := s.Float64()
	return vInt
}

