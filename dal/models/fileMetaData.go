package models

type FileMetaData struct {
	Id             string `json:"-"`
	Name           string `json:"name"`
	CreatorId      string `json:"-"`
	CreatedAtEpoch int64  `json:"createdAtEpoch"`
}
