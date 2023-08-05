package repo

import "mime/multipart"

type IFileRepo interface {
	SaveFile(file multipart.File, name string, creatorId string) error
}
