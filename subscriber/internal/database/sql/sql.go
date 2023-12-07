package sql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"subscriber/api/json"
	models "subscriber/internal/model"
	"time"
)

func New() InterfaceSql {
	return &Sql{}
}

func (s *Sql) Connect() error {
	var err error

	dsn := fmt.Sprintf("host=localhost user=postgres password=postgres dbname=NatsDatabase port=5433 sslmode=disable")

	s.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("ошибка при подключении к БД %s", err)
	}
	log.Println("подключена бд")
	err = s.DB.AutoMigrate(&models.Model{}, &models.Delivery{}, &models.Items{}, &models.Payment{})
	if err != nil {
		return fmt.Errorf("ошибка при миграции %s", err)
	}
	log.Println("миграции совершены")
	return nil
}

func (s *Sql) GetAllRecords() ([]models.Model, bool) {
	var modelSlice []models.Model
	s.DB.Preload("Delivery").Preload("Payment").Preload("Items").Find(&modelSlice)
	if modelSlice == nil {
		return nil, true
	}
	return modelSlice, false
}

func (s *Sql) GetById(id int) (*models.Model, bool) {
	var model models.Model
	log.Println("sql id = ", id)
	s.DB.First(&model, id)
	if model.ID == 0 && id != 0 {
		log.Println("value not found")
		return nil, false
	}
	return &model, true
}

func (s *Sql) AddRecord(jsonModel json.ModelJson) {

	delivery := models.Delivery{
		Name:    jsonModel.Delivery.Name,
		Phone:   jsonModel.Delivery.Phone,
		Zip:     jsonModel.Delivery.Zip,
		City:    jsonModel.Delivery.City,
		Address: jsonModel.Delivery.Address,
		Region:  jsonModel.Delivery.Region,
		Email:   jsonModel.Delivery.Email,
	}

	payment := models.Payment{
		Transaction:  jsonModel.Payment.Transaction,
		RequestId:    jsonModel.Payment.RequestId,
		Currency:     jsonModel.Payment.Currency,
		Provider:     jsonModel.Payment.Provider,
		Amount:       jsonModel.Payment.Amount,
		PaymentDt:    jsonModel.Payment.PaymentDt,
		Bank:         jsonModel.Payment.Bank,
		DeliveryCost: jsonModel.Payment.DeliveryCost,
		GoodsTotal:   jsonModel.Payment.GoodsTotal,
		CustomFee:    jsonModel.Payment.CustomFee,
	}

	var items []models.Items
	for _, item := range jsonModel.Items {
		items = append(items, models.Items{
			ChrtId:      item.ChrtId,
			TrackNumber: item.TrackNumber,
			Price:       item.Price,
			Rid:         item.Rid,
			Name:        item.Name,
			Sale:        item.Sale,
			Size:        item.Size,
			TotalPrice:  item.TotalPrice,
			NmId:        item.NmId,
			Brand:       item.Brand,
			Status:      item.Status,
		})
	}

	newModel := models.Model{
		OrderUid:          jsonModel.OrderUid,
		TrackNumber:       jsonModel.TrackNumber,
		Entry:             jsonModel.Entry,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            jsonModel.Locale,
		InternalSignature: jsonModel.InternalSignature,
		CustomerId:        jsonModel.CustomerId,
		DeliveryService:   jsonModel.DeliveryService,
		ShardKey:          jsonModel.Shardkey,
		SmId:              jsonModel.SmId,
		DateCreated:       time.Now(),
		OofShard:          jsonModel.OofShard,
	}

	s.DB.Create(&newModel)
}
