package models

import (
	"gorm.io/gorm"
	"time"
)

type Delivery struct {
	gorm.Model
	DeliveryID int    `gorm:"primaryKey;autoIncrement"`
	ID         int    `gorm:"type:int"`
	Name       string `gorm:"type:varchar(255)"`
	Phone      string `gorm:"type:varchar(255)"`
	Zip        string `gorm:"type:varchar(255)"`
	City       string `gorm:"type:varchar(255)"`
	Address    string `gorm:"type:varchar(255)"`
	Region     string `gorm:"type:varchar(255)"`
	Email      string `gorm:"type:varchar(255)"`
}

type Payment struct {
	gorm.Model
	PaymentID    int    `gorm:"primaryKey;autoIncrement"`
	ID           int    `gorm:"type:int"`
	Transaction  string `gorm:"type:varchar(255)"`
	RequestId    string `gorm:"type:varchar(255)"`
	Currency     string `gorm:"type:varchar(255)"`
	Provider     string `gorm:"type:varchar(255)"`
	Amount       int    `gorm:"type:int"`
	PaymentDt    int    `gorm:"type:int"`
	Bank         string `gorm:"type:varchar(255)"`
	DeliveryCost int    `gorm:"type:int"`
	GoodsTotal   int    `gorm:"type:int"`
	CustomFee    int    `gorm:"type:int"`
}

type Items struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement"`
	ChrtId      int    `gorm:"type:int"`
	TrackNumber string `gorm:"type:varchar(255)"`
	Price       int    `gorm:"type:int"`
	Rid         string `gorm:"type:varchar(255)"`
	Name        string `gorm:"type:varchar(255)"`
	Sale        int    `gorm:"type:int"`
	Size        string `gorm:"type:varchar(255)"`
	TotalPrice  int    `gorm:"type:int"`
	NmId        int    `gorm:"type:int"`
	Brand       string `gorm:"type:varchar(255)"`
	Status      int    `gorm:"type:int"`
}

type Model struct {
	gorm.Model
	ID                int       `gorm:"primaryKey;autoIncrement"`
	OrderUid          string    `gorm:"type:varchar(255)"`
	TrackNumber       string    `gorm:"type:varchar(255)"`
	Entry             string    `gorm:"type:varchar(255)"`
	Delivery          Delivery  `gorm:"foreignKey:DeliveryID"`
	Payment           Payment   `gorm:"foreignKey:PaymentID"`
	Items             []Items   `gorm:"foreignKey:ID"`
	Locale            string    `gorm:"type:varchar(255)"`
	InternalSignature string    `gorm:"type:varchar(255)"`
	CustomerId        string    `gorm:"type:varchar(255)"`
	DeliveryService   string    `gorm:"type:varchar(255)"`
	ShardKey          string    `gorm:"type:varchar(255)"`
	SmId              int       `gorm:"type:int"`
	DateCreated       time.Time `gorm:"type:timestamp"`
	OofShard          string    `gorm:"type:varchar(255)"`
}
