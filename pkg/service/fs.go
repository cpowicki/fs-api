package service

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/cpowicki/fs-api/pkg/config"
	"github.com/spf13/afero"
)

type FileMetadata struct {
	FileName    string `json:"fileName"`
	Owner       string `json:"owner"`
	Size        string `json:"size"`
	Permissions string `json:"permissions"`
	Type        string `json:"type"`
}

type FileSystemService struct {
	root string
	fs   afero.Fs
}

func NewFileSystemService(fsApiConfig config.FsApiConfig) FileSystemService {
	return FileSystemService{
		root: fsApiConfig.Root,
		fs:   afero.NewOsFs(),
	}
}

func (service *FileSystemService) ListRootDirContents() (metadata []FileMetadata, err error) {
	return service.ListDirContents("")
}

func (s *FileSystemService) CheckFileExists(file string) bool {
	var fullPath = filepath.Join(s.root, file)
	_, err := s.fs.Stat(fullPath)
	return !errors.Is(err, os.ErrNotExist)
}

func (s *FileSystemService) ListDirContents(relativePath string) (metadata []FileMetadata, err error) {
	var path = filepath.Join(s.root, relativePath)

	fileinfo, err := afero.ReadDir(s.fs, path)
	if err != nil {
		return
	}

	metadata = make([]FileMetadata, len(fileinfo))
	for i, info := range fileinfo {
		metadata[i] = s.buildFileMetadata(path, info)
	}

	return
}

func (s *FileSystemService) buildFileMetadata(dir string, entry fs.FileInfo) FileMetadata {
	var fullPath = filepath.Join(dir, entry.Name())

	fileInfo, err := s.fs.Stat(fullPath)
	if err != nil {
		panic(err)
	}

	var t string
	if fileInfo.IsDir() {
		t = "Directory"
	} else {
		t = "File"
	}

	return FileMetadata{
		FileName:    fileInfo.Name(),
		Owner:       s.getFileOwner(fileInfo),
		Size:        s.getFileSize(fileInfo),
		Permissions: s.getFilePermissions(fileInfo),
		Type:        t,
	}
}

func (service *FileSystemService) getFileSize(fileInfo fs.FileInfo) string {
	// TODO make human readable
	return fmt.Sprintf("%d bytes", fileInfo.Size())
}

func (service *FileSystemService) getFilePermissions(fileInfo fs.FileInfo) string {
	return fmt.Sprintf("%o", fileInfo.Mode().Perm())
}

func (service *FileSystemService) getFileOwner(fileInfo fs.FileInfo) (username string) {

	if sysInfo, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
		user_id := strconv.FormatUint(uint64(sysInfo.Uid), 10)
		usr, err := user.LookupId(user_id)

		if err != nil {
			// TODO replace with logger
			fmt.Println("error resolving user for requested file", fileInfo.Name(), err)
			username = "Unknown"
		} else {
			username = usr.Name
		}
	} else {
		username = "Unknown"
	}

	return
}

func (s *FileSystemService) IsDirectory(path string) bool {
	var fullPath = filepath.Join(s.root, path)
	// TODO handle err
	fileInfo, _ := s.fs.Stat(fullPath)
	return fileInfo.IsDir()

}

func (s *FileSystemService) ReadFileContents(file string) (contents string, err error) {
	var path = filepath.Join(s.root, file)

	bytes, err := afero.ReadFile(s.fs, path)

	contents = string(bytes)

	return
}
