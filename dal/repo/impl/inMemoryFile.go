package impl

import (
	"fmt"
	"github.com/rahularcota/dropbox/utils"
	"io"
	"mime/multipart"
	"os"
)

type InMemoryFileRepo struct{}

func NewInMemoryFileRepo() *InMemoryFileRepo {
	return &InMemoryFileRepo{}
}

func (repo *InMemoryFileRepo) SaveFile(file multipart.File, name string, creatorId string) error {
	fPath := os.ExpandEnv(fmt.Sprintf("./%s/%s", creatorId, name))
	utils.EnsureBaseDir(fPath)
	f, err := os.Create(fPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = f.Write(fileBytes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
