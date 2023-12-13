package storage

import (
	"database/sql"
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/model"
	"log"
	"os"
	"time"
)

var (
	Db         *sql.DB
	infoLogDB  = log.New(os.Stdout, "INFO[DB] ", log.Ldate|log.Ltime)
	ErrorLogDB = log.New(os.Stderr, "ERROR[DB] ", log.Ldate|log.Ltime|log.Lshortfile)
)

// Connection to db func
func ConnectToDB(dbinfo string, maxRetries int, IntervalRetry time.Duration) {
	var (
		err     error
		retries int
	)
	//If DB doesn't respond, then there are 'maxRetries' attempts to reconnect every 'IntervalRetry' seconds
	for retries < maxRetries {
		Db, err = sql.Open("postgres", dbinfo)
		if err == nil {
			// It establishes a connection
			if err := Db.Ping(); err == nil {
				infoLogDB.Println("DB connection is successfully created")
				return
			}
		}
		ErrorLogDB.Printf("Error connecting to db (attempt %d): %v", retries+1, err)
		retries++
		time.Sleep(IntervalRetry)
	}
	ErrorLogDB.Fatalf("Unable to connect to db after all attempts")
}

// Function to get data from all tables except orderitems
func GetAllOrders() (*sql.Rows, error) {
	rows, err := Db.Query(
		"SELECT o.orderuid, o.tracknumber, o.entry, o.locale, o.internalsignature, o.customerid, o.deliveryservice, " +
			" o.shardkey, o.smid, o.datecreated, o.oofshard," +
			" d.orderuid, d.name, d.phone, d.zip, d.city, d.address, d.region, d.email," +
			" p.transaction, p.requestid, p.currency, p.provider, p.amount, p.paymentdt, p.bank, p.deliverycost, p.goodstotal, p.customfee" +
			" FROM orders o" +
			" JOIN delivery d ON o.orderuid = d.orderuid" +
			" JOIN payment p ON o.orderuid = p.transaction",
	)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Function to get data from orderitems table by order.Tracknumber
func GetItemsByOrderTrackNumber(order model.Order) (*sql.Rows, error) {
	itemRows, err := Db.Query("SELECT tracknumber, chrtid, price, rid, name, sale, size, totalprice, "+
		"nmid, brand, status FROM orderitems WHERE tracknumber = $1", order.TrackNumber)
	if err != nil {
		return nil, err
	}
	return itemRows, nil
}

// function to insert the order to DB
func InsertOrder(order model.Order) error {
	//The beginnig of transaction - it is required as if we face any problems inserting data to one of the table,
	//then we need to rollback. Else - Commit
	tx, err := Db.Begin()
	if err != nil {
		return err
	}

	// Delayed call of RollBack in case of error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Insert data to orders table
	_, err = tx.Exec("INSERT INTO orders (orderuid, tracknumber, entry, locale, internalsignature, customerid, deliveryservice, shardkey, smid, datecreated, oofshard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID, order.DateCreated, order.OOFShard)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert data to delivery table
	_, err = tx.Exec("INSERT INTO delivery (orderuid, name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert data to payment table
	_, err = tx.Exec("INSERT INTO payment (transaction, requestid, currency, provider, amount, paymentdt, bank, deliverycost, goodstotal, customfee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		order.OrderUID, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert data to orderitems table (the loop as several order items can be assigned to one order)
	for _, item := range order.Items {
		_, err := tx.Exec("INSERT INTO orderitems (tracknumber, chrtid, price, rid, name, sale, size, totalprice, nmid, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
			item.TrackNumber, item.ChrtID, item.Price, item.RID, item.Name, item.Sale, item.Size, item.TotalPrice, item.NMID, item.Brand, item.Status)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
