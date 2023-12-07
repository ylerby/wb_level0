package sql

import (
	"gorm.io/gorm"
	"subscriber/api/json"
	models "subscriber/internal/model"
)

type Sql struct {
	DB *gorm.DB
}

type InterfaceSql interface {
	Connect() error
	GetAllRecords() ([]models.Model, bool)
	AddRecord(jsonModel json.ModelJson)
	GetById(id int) (*models.Model, bool)
}
