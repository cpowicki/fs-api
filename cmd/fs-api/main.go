package main

import (
	"github.com/cpowicki/fs-api/pkg/config"
	"github.com/cpowicki/fs-api/pkg/service"
)

func main() {
	var fsApiConfig = config.ParseCliArgs()
	metadata, err := service.ListDirContents(fsApiConfig.Root)

	if err != nil {
		panic(err)
	}

	for _, file := range metadata {
		println(file.FileName)
		println(file.Owner)
	}
}
