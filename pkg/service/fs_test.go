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
}

func mockFileSystem() (testFs afero.Fs) {
	testFs = afero.NewMemMapFs()

	testFs.Mkdir("/", 0755)
	testFs.Mkdir("/emptyDirectory", 0755)
	testFs.Mkdir("/directory", 0755)
	afero.WriteFile(testFs, "/data.txt", []byte("this file has contents"), 0777)

	return
}
