package cache

import (
	"subscriber/api/json"
	models "subscriber/internal/model"
	"time"
)

func New() InterfaceCache {
	return &Cache{}
}

func (c *Cache) Connect() {
	c.DB = make(map[int]models.Model)
}

func (c *Cache) GetCacheSize() int {
	return len(c.DB)
}

func (c *Cache) GetById(id int) (*models.Model, bool) {
	value, ok := c.DB[id]
	if ok {
		return &value, true
	}
	return nil, false
}

func (c *Cache) AddRecord(jsonModel json.ModelJson) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delivery := models.Delivery{
		DeliveryID: c.GetCacheSize() - 1,
		Name:       jsonModel.Delivery.Name,
		Phone:      jsonModel.Delivery.Phone,
		Zip:        jsonModel.Delivery.Zip,
		City:       jsonModel.Delivery.City,
		Address:    jsonModel.Delivery.Address,
		Region:     jsonModel.Delivery.Region,
		Email:      jsonModel.Delivery.Email,
	}

	payment := models.Payment{
		PaymentID:    c.GetCacheSize() - 1,
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
			ID:          c.GetCacheSize() - 1,
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
		ID:                c.GetCacheSize() - 1,
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
	c.DB[c.GetCacheSize()-1] = newModel
}

func (c *Cache) CacheDownloading(modelSlice []models.Model) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for index, value := range modelSlice {
		c.DB[index+1] = value
	}
}

func (c *Cache) GetAllRecords() map[int]models.Model {
	return c.DB
}
