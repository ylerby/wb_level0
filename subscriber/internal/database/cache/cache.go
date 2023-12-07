package cache

import (
	"subscriber/api/json"
	models "subscriber/internal/model"
	"time"
)

func New() InterfaceCache {
	return &Client{}
}

func (c *Client) Connect() {
	c.DB = make(map[int]models.Model)
}

func (c *Client) GetById(id int) (*models.Model, bool) {
	value, ok := c.DB[id]
	if ok {
		return &value, true
	}
	return nil, false
}

func (c *Client) CacheDownloading(modelSlice []models.Model) {
	for index, value := range modelSlice {
		c.DB[index] = value
	}
}

func (c *Client) GetAllRecords() map[int]models.Model {
	return c.DB
}

func (c *Client) GetCacheSize() int {
	return len(c.DB)
}

func (c *Client) AddRecord(jsonModel json.ModelJson) {
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
	c.DB[c.GetCacheSize()-1] = newModel
}
