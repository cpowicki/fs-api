package main

import (
	"github.com/cpowicki/fs-api/pkg/config"
	"github.com/cpowicki/fs-api/pkg/server"
)

func main() {
	var fsApiConfig = config.ParseCliArgs()
	var server = server.NewFileSystemServer(fsApiConfig)
	server.Start()
}
