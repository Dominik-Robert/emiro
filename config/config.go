package config

import "github.com/spf13/viper"

// CfgStruct is the structure for complicated fields
var CfgStruct struct {
	Remote struct {
		PermanentRemote string   `yaml:"permanentRemote"`
		AllowIPs        []string `yaml:"allowIPs"`
		AllowRange      []struct {
			From string `yaml:"from"`
			Till string `yaml:"till"`
		}
	} `yaml:"remote"`
}

// Config is the global config of this program to access it in other packages
var Config *viper.Viper
