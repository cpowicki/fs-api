package server

import "github.com/cpowicki/fs-api/pkg/service"

type ListFilesResponse struct {
	Files []service.FileMetadata `json:"files"`
}

type FileContentResponse struct {
	FilePath string `json:"filePath"`
	Data     string `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
