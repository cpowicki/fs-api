package service

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/cpowicki/fs-api/pkg/config"
)

type FileMetadata struct {
	FileName    string
	Owner       string
	Size        string
	Permissions string
}

type FileSystemService struct {
	root string
}

func NewFileSystemService(fsApiConfig config.FsApiConfig) FileSystemService {
	return FileSystemService{
		root: fsApiConfig.Root,
	}
}

func (service *FileSystemService) ListRootDirContents() (metadata []FileMetadata, err error) {
	return service.ListDirContents(service.root)
}

func (service *FileSystemService) ListDirContents(path string) (metadata []FileMetadata, err error) {

	fileinfo, err := os.ReadDir(path)
	if err != nil {
		return
	}

	metadata = make([]FileMetadata, len(fileinfo))
	for i, info := range fileinfo {
		metadata[i] = service.buildFileMetadata(path, info)
	}

	return
}

func (service *FileSystemService) buildFileMetadata(dir string, entry fs.DirEntry) FileMetadata {

	var fullPath = filepath.Join(dir, entry.Name())

	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		panic(err)
	}

	return FileMetadata{
		FileName:    fileInfo.Name(),
		Owner:       service.getFileOwner(fileInfo),
		Size:        "",
		Permissions: "",
	}
}

func (service *FileSystemService) getFileOwner(fileInfo fs.FileInfo) (username string) {

	sysInfo := fileInfo.Sys().(*syscall.Stat_t)
	user_id := strconv.FormatUint(uint64(sysInfo.Uid), 10)
	usr, err := user.LookupId(user_id)

	if err != nil {
		// TODO replace with logger
		fmt.Println("error resolving user for requested file", fileInfo.Name(), err)
		username = "Unknown"
	} else {
		username = usr.Name
	}

	return
}
