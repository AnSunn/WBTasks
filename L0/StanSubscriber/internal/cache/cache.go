package cache

import (
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/model"
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/storage"
	"log"
	"sync"
)

// Cache structure with RWMutex for safe access from several goroutines
type cache struct {
	sync.RWMutex
	Items map[string]model.Order
}

var c cache

// Initialize cache
func InitCache() {
	c = cache{
		Items: make(map[string]model.Order),
	}
}

// Update cache
func UpdateCache(order model.Order) {
	c.Lock()
	defer c.Unlock()
	//Add order to cache
	c.Items[order.OrderUID] = order
}

// Get item from the cache
func GetItemFromCache(orderUID string) (model.Order, bool) {
	c.RLock()
	defer c.RUnlock()
	//Get item from the cache
	item, ok := c.Items[orderUID]
	return item, ok
}

// Load cache from DB
func LoadCacheFromDB() {
	//Get data from all tables except orderitems
	rows, err := storage.GetAllOrders()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//For each row from query scan data to model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(
			&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerID,
			&order.DeliveryService, &order.ShardKey, &order.SMID, &order.DateCreated, &order.OOFShard,
			&order.Delivery.OrderUID, &order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City,
			&order.Delivery.Address, &order.Delivery.Region, &order.Delivery.Email,
			&order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency, &order.Payment.Provider,
			&order.Payment.Amount, &order.Payment.PaymentDT, &order.Payment.Bank, &order.Payment.DeliveryCost,
			&order.Payment.GoodsTotal, &order.Payment.CustomFee,
		); err != nil {
			log.Fatal(err)
		}

		//Several order items can be assigned to one order, that's why we need to add each of then to model order (Items).
		//First of all we need to take date from orderitems where tracknumber = order.tracknumber
		itemRows, err := storage.GetItemsByOrderTrackNumber(order)
		if err != nil {
			log.Fatal(err)
		}
		defer itemRows.Close()
		//Then add each of found item to order.Items
		for itemRows.Next() {
			var item model.OrderItem
			if err := itemRows.Scan(
				&item.TrackNumber, &item.ChrtID, &item.Price, &item.RID, &item.Name, &item.Sale, &item.Size, &item.TotalPrice,
				&item.NMID, &item.Brand, &item.Status,
			); err != nil {
				log.Fatal(err)
			}
			order.Items = append(order.Items, item)
		}
		// Add the order to cache
		UpdateCache(order)
	}

	if err := rows.Err(); err != nil {
		storage.ErrorLogDB.Fatal(err)
	}
}
