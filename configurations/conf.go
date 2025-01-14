package configurations

import (
	"errors"
	"fmt"
)

// conf file struct to set path, file & extension
type confFile struct {
	path string
	file string
	ext  string
}

// define new configuraion interface
type Configuration interface {
	// Read Configuration file
	Read() (interface{}, error)
}

// Type of congfiguration file
type ConfType int

// enum const for type of configuration files
const (
	CONF_TYPE_COLLECTOR ConfType = iota

	CONF_TYPE_IP2LOCATION
)

// return filename related to requested configuration
func (ct ConfType) String() string {
	return [...]string{
		"collector",
		"ip2location",
	}[ct]
}

// create new configuration
func New(ct ConfType, path string) (Configuration, error) {

	// define default configuration file path, name, ext
	cf := confFile{
		path: "/opt/nfcollector/etc/",
		file: "collector",
		ext:  "yml",
	}

	if path != "" {
		cf.path = path
	}

	switch ct {
	case CONF_TYPE_COLLECTOR:
		return Configuration(
			&Collector{
				confFile: cf,
			},
		), nil
	case CONF_TYPE_IP2LOCATION:
		return Configuration(
			&IP2Location{
				confFile: confFile{
					path: cf.path,
					file: "ip2location",
					ext:  "yml",
				},
			},
		), nil
	default:
		return Configuration(&Collector{}), errors.New(fmt.Sprintf("No valid configuration type found"))
	}
}
