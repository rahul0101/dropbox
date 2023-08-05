package impl

import (
	"github.com/rahularcota/dropbox/dal/repo"
	"mime/multipart"
)

type FileService struct {
	fileRepo repo.IFileRepo
}

func NewFileService(fileRepo repo.IFileRepo) *FileService {
	return &FileService{fileRepo: fileRepo}
}

func (svc *FileService) SaveFile(file multipart.File, name string, creatorId string) error {
	return svc.fileRepo.SaveFile(file, name, creatorId)
}
