package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	http.HandleFunc("/", s.handleRequest)

	http.ListenAndServe(fmt.Sprintf("%s%d", ":", s.port), nil)

}

func (s *FileSystemServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}

	var path = strings.TrimPrefix(r.URL.Path, "/")

	if path == "" {
		s.listRootDir(w, r)
	} else {
		if s.fsService.IsDirectory(path) {
			s.listDir(path, w, r)
		} else {
			s.readFileContents(path, w, r)
		}
	}

}

func (s *FileSystemServer) listRootDir(w http.ResponseWriter, r *http.Request) {
	metdata, _ := s.fsService.ListRootDirContents()
	json.NewEncoder(w).Encode(metdata)
}

func (s *FileSystemServer) listDir(dir string, w http.ResponseWriter, r *http.Request) {
	metdata, _ := s.fsService.ListDirContents(dir)
	json.NewEncoder(w).Encode(metdata)
}

func (s *FileSystemServer) readFileContents(path string, w http.ResponseWriter, r *http.Request) {
	data, _ := s.fsService.ReadFileContents(path)
	w.Write([]byte(data))
}
