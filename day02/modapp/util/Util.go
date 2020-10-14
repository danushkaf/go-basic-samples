package util

import (
	"flag"

	"github.com/magiconair/properties"
)

var props *properties.Properties
var CleanDB *bool

func init() {

	configFile := flag.String("configFile", "application.properties", "Configuration File")
	CleanDB = flag.Bool("cleanDB", false, "Specify to recreate the DB")
	flag.Parse()
	props = properties.MustLoadFile(*configFile, properties.UTF8)
}

func GetProperties() *properties.Properties {
	return props
}
