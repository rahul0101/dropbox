package repo

import "github.com/rahularcota/dropbox/dal/models"

type IFileMetaDataRepo interface {
	AddFileMetaData(metaData *models.FileMetaData) error
	GetFilesMetaDataForUser(userId string) ([]*models.FileMetaData, error)
}
