package impl

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/rahularcota/dropbox/dal/models"
	"time"
)

type PGFileMetaDataRepo struct {
	db *gorm.DB
}

func NewPGFileMetaDataRepo(db *gorm.DB) *PGFileMetaDataRepo {
	return &PGFileMetaDataRepo{db: db}
}

func (repo *PGFileMetaDataRepo) AddFileMetaData(metaData *models.FileMetaData) error {
	metaData.CreatedAtEpoch = time.Now().UnixMilli()
	var data models.FileMetaData
	err := repo.db.Where("creator_id=? and name=?", metaData.CreatorId, metaData.Name).First(&data).Error
	if gorm.IsRecordNotFoundError(err) {
		return repo.db.Create(metaData).Error
	}
	return fmt.Errorf("file already exists, please use a different name")
}

func (repo *PGFileMetaDataRepo) GetFilesMetaDataForUser(userId string) ([]*models.FileMetaData, error) {
	var metaDatum []*models.FileMetaData
	err := repo.db.Where("creator_id = ?", userId).Find(&metaDatum).Error
	if err != nil {
		return nil, err
	}
	return metaDatum, nil
}
