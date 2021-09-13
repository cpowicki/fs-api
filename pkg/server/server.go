package server

import (
	"net/http"

	"github.com/cpowicki/fs-api/pkg/config"
	"github.com/cpowicki/fs-api/pkg/service"
)

type HttpServer struct {
	port      int
	fsService service.FileSystemService
}

func (*HttpServer) NewHttpServer(config config.FsApiConfig) HttpServer {
	return HttpServer{
		port: config.ServerPort,
	}
}

func (*HttpServer) Start() {

}

func ListRootDir(w http.ResponseWriter, r *http.Request) {

}
