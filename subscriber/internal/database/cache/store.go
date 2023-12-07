package cache

import (
	"subscriber/api/json"
	models "subscriber/internal/model"
)

type Client struct {
	DB map[int]models.Model
}

type InterfaceCache interface {
	Connect()
	GetById(id int) (*models.Model, bool)
	GetCacheSize() int
	AddRecord(model json.ModelJson)
	CacheDownloading(modelSlice []models.Model)
	GetAllRecords() map[int]models.Model
}
