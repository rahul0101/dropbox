package impl

import (
	"github.com/google/uuid"
	"github.com/rahularcota/dropbox/dal/models"
	"github.com/rahularcota/dropbox/dal/repo"
)

type FileMetaDataService struct {
	fileMetaDataRepo repo.IFileMetaDataRepo
}

func NewFileMetaDataService(fileMetaDataRepo repo.IFileMetaDataRepo) *FileMetaDataService {
	return &FileMetaDataService{fileMetaDataRepo: fileMetaDataRepo}
}

func (svc *FileMetaDataService) AddFileMetaData(name string, creatorId string) error {
	return svc.fileMetaDataRepo.AddFileMetaData(&models.FileMetaData{
		Id:        uuid.New().String(),
		Name:      name,
		CreatorId: creatorId,
	})
}

func (svc *FileMetaDataService) GetFilesMetaDataForUser(userId string) ([]*models.FileMetaData, error) {
	return svc.fileMetaDataRepo.GetFilesMetaDataForUser(userId)
}
