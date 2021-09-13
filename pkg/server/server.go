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
		resp := ErrorResponse{
			Message: "Method Not Allowed",
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var path = strings.TrimPrefix(r.URL.Path, "/")

	if path == "" {
		s.listRootDir(w, r)
	} else {

		if !s.fsService.CheckFileExists(path) {
			w.WriteHeader(404)
			resp := ErrorResponse{
				Message: "File not found",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		if s.fsService.IsDirectory(path) {
			s.listDir(path, w, r)
		} else {
			s.readFileContents(path, w, r)
		}
	}

}

func (s *FileSystemServer) listRootDir(w http.ResponseWriter, r *http.Request) {
	metdata, _ := s.fsService.ListRootDirContents()

	response := ListFilesResponse{
		Files: metdata,
	}

	json.NewEncoder(w).Encode(response)
}

func (s *FileSystemServer) listDir(dir string, w http.ResponseWriter, r *http.Request) {
	metdata, _ := s.fsService.ListDirContents(dir)

	response := ListFilesResponse{
		Files: metdata,
	}
	json.NewEncoder(w).Encode(response)
}

func (s *FileSystemServer) readFileContents(path string, w http.ResponseWriter, r *http.Request) {
	data, _ := s.fsService.ReadFileContents(path)
	response := FileContentResponse{
		FilePath: path,
		Data:     data,
	}
	json.NewEncoder(w).Encode(response)
}
