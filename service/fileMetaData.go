package service

import "github.com/rahularcota/dropbox/dal/models"

type IFileMetaDataService interface {
	AddFileMetaData(name string, creatorId string) error
	GetFilesMetaDataForUser(userId string) ([]*models.FileMetaData, error)
}
