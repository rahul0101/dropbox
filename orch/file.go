package orch

import (
	"github.com/rahularcota/dropbox/dal/models"
	"mime/multipart"
)

type IFileOrch interface {
	UploadFile(file multipart.File, name string, creatorId string) error
	GetFilesInfoForUser(userId string) ([]*models.FileMetaData, error)
}
