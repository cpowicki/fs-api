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

	flag.StringVar(&config.Root, "root", os.Getenv("HOME"), "the root directory to expose for browsing")
	flag.IntVar(&config.ServerPort, "port", 3030, "the server port")

	flag.Parse()

	return
}
