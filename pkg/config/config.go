package config

import (
	"flag"
	"os"
)

type FsApiConfig struct {
	Root       string
	ServerPort int
}

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
