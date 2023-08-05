package service

import "mime/multipart"

type IFileService interface {
	SaveFile(file multipart.File, name string, creatorId string) error
}
