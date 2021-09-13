package service

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/cpowicki/fs-api/pkg/config"
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
}

func NewFileSystemService(fsApiConfig config.FsApiConfig) FileSystemService {
	return FileSystemService{
		root: fsApiConfig.Root,
	}
}

func (service *FileSystemService) ListRootDirContents() (metadata []FileMetadata, err error) {
	return service.ListDirContents("")
}

func (service *FileSystemService) ListDirContents(relativePath string) (metadata []FileMetadata, err error) {
	var path = filepath.Join(service.root, relativePath)

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

	var t string

	if fileInfo.IsDir() {
		t = "Directory"
	} else {
		t = "File"
	}

	return FileMetadata{
		FileName:    fileInfo.Name(),
		Owner:       service.getFileOwner(fileInfo),
		Size:        fmt.Sprintf("%d bytes", fileInfo.Size()),
		Permissions: "",
		Type:        t,
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

func (service *FileSystemService) IsDirectory(path string) bool {
	var fullPath = filepath.Join(service.root, path)
	// TODO handle err
	fileInfo, _ := os.Stat(fullPath)
	return fileInfo.IsDir()

}

func (service *FileSystemService) ReadFileContents(file string) (contents string, err error) {
	var path = filepath.Join(service.root, file)

	bytes, err := ioutil.ReadFile(path)

	contents = string(bytes)

	return
}
