package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cpowicki/fs-api/pkg/config"
	"github.com/cpowicki/fs-api/pkg/service"
)

type FileSystemServer struct {
	port      int
	fsService service.FileSystemService
}

func NewFileSystemServer(config config.FsApiConfig) FileSystemServer {
	return FileSystemServer{
		port:      config.ServerPort,
		fsService: service.NewFileSystemService(config),
	}
}

func (s *FileSystemServer) Start() {

	http.HandleFunc("/", s.listRootDir)

	http.ListenAndServe(fmt.Sprintf("%s%d", ":", s.port), nil)

}

func (s *FileSystemServer) listRootDir(w http.ResponseWriter, r *http.Request) {
	metdata, _ := s.fsService.ListRootDirContents()
	json.NewEncoder(w).Encode(metdata)
}
