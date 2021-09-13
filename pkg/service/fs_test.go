package service

import (
	"testing"

	"github.com/spf13/afero"
)

func TestListDirContents(t *testing.T) {

	var fsService = FileSystemService{
		root: "/",
		fs:   mockFileSystem(),
	}

	metadata, err := fsService.ListRootDirContents()

	if err != nil {
		t.Fatalf(`%v`, err)
	}

	if len(metadata) != 3 {
		t.Fatalf(`expecting len 3, got %d`, len(metadata))
	}

	if metadata[0].FileName != "data.txt" {
		t.Fatalf(`expecting 'data.txt', got %s`, metadata[0].FileName)
	}

	if metadata[0].Owner != "Unknown" {
		t.Fatalf(`expecting 'Unknown', got %s`, metadata[0].Owner)
	}

	if metadata[0].Type != "File" {
		t.Fatalf(`expecting 'File', got %s`, metadata[0].Type)
	}

	if metadata[0].Permissions != "777" {
		t.Fatalf(`expecting '777', got %s`, metadata[0].Permissions)
	}

	if metadata[1].FileName != "directory" {
		t.Fatalf(`expecting 'directory', got %s`, metadata[0].FileName)
	}

	if metadata[1].Type != "Directory" {
		t.Fatalf(`expecting 'Directory', got %s`, metadata[0].Type)
	}
}

func TestReadFileContents(t *testing.T) {
	var fsService = FileSystemService{
		root: "/",
		fs:   mockFileSystem(),
	}

	contents, _ := fsService.ReadFileContents("data.txt")

	if contents != "this file has contents" {
		t.Fatalf(`expecting 'this file has contents', got '%s'`, contents)
	}
}

func mockFileSystem() (testFs afero.Fs) {
	testFs = afero.NewMemMapFs()

	testFs.Mkdir("/", 0755)
	testFs.Mkdir("/emptyDirectory", 0755)
	testFs.Mkdir("/directory", 0755)
	afero.WriteFile(testFs, "/data.txt", []byte("this file has contents"), 0777)

	return
}
