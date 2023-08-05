package impl

import (
	"github.com/rahularcota/dropbox/dal/models"
	"github.com/rahularcota/dropbox/service"
	"mime/multipart"
)

type FileOrch struct {
	fileMetaDataService service.IFileMetaDataService
	fileService         service.IFileService
}

func NewFileOrch(fileMetaDataService service.IFileMetaDataService, fileService service.IFileService) *FileOrch {
	return &FileOrch{fileMetaDataService: fileMetaDataService, fileService: fileService}
}

func (orch *FileOrch) UploadFile(file multipart.File, name string, creatorId string) error {
	err := orch.fileMetaDataService.AddFileMetaData(name, creatorId)
	if err != nil {
		return err
	}
	return orch.fileService.SaveFile(file, name, creatorId)
}

func (orch *FileOrch) GetFilesInfoForUser(userId string) ([]*models.FileMetaData, error) {
	return orch.fileMetaDataService.GetFilesMetaDataForUser(userId)
}
