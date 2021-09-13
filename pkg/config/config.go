package config

import (
	"flag"
	"os"
)

// The config for the fs-api
type FsApiConfig struct {
	Root       string // the root directory to expose for browsing
	ServerPort int    // the port to listen on to handle http requests
}

// Parses CLI arguments and returns structured config
func ParseCliArgs() (config FsApiConfig) {

	defaultRoot, set := os.LookupEnv("HOME")
	if !set {
		defaultRoot = "/"
	}

	flag.StringVar(&config.Root, "root", defaultRoot, "the root directory to expose for browsing")
	flag.IntVar(&config.ServerPort, "port", 3030, "the server port")

	flag.Parse()

	return
}
