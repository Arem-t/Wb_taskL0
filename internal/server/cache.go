package server

import (
	"WB00L0/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

var Cache = map[string]models.Order{}

func RestoreCache(db *sql.DB) {
	rdb := db
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {

		var order models.Order
		err := rows.Scan(
			&order.OrderUid,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerId,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmId,
			&order.DateCreated,
			&order.OofShard)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
		}

		rowsD, err := rdb.Query("SELECT * FROM deliveries WHERE order_uid = ($1)", order.OrderUid)
		if err != nil {
			log.Println(err)
		}
		for rowsD.Next() {
			err := rowsD.Scan(
				&order.OrderUid,
				&order.Delivery.Name,
				&order.Delivery.Phone,
				&order.Delivery.Zip,
				&order.Delivery.City,
				&order.Delivery.Address,
				&order.Delivery.Region,
				&order.Delivery.Email)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
			}
		}

		rowsP, err := rdb.Query("SELECT * FROM payments WHERE order_uid = ($1)", order.OrderUid)
		if err != nil {
			log.Println(err)
		}
		for rowsP.Next() {
			err := rowsP.Scan(
				&order.OrderUid,
				&order.Payment.Transaction,
				&order.Payment.RequestId,
				&order.Payment.Currency,
				&order.Payment.Provider,
				&order.Payment.Amount,
				&order.Payment.PaymentDt,
				&order.Payment.Bank,
				&order.Payment.DeliveryCost,
				&order.Payment.GoodsTotal,
				&order.Payment.CustomFee)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
			}
		}

		rowsI, err := rdb.Query("SELECT * FROM items WHERE order_uid = ($1)", order.OrderUid)
		if err != nil {
			log.Println(err)
		}
		var items []models.Item
		for rowsI.Next() {
			var item models.Item
			if err := rowsI.Scan(
				&order.OrderUid,
				&item.ChrtId,
				&item.TrackNumber,
				&item.Price,
				&item.Rid,
				&item.Name,
				&item.Sale,
				&item.Size,
				&item.TotalPrice,
				&item.NmId,
				&item.Brand,
				&item.Status,
			); err != nil {
				log.Fatal(err)
			}
			items = append(items, item)
		}
		order.Items = items

		defer rowsI.Close()
		defer rowsD.Close()
		defer rowsP.Close()
		if err := rowsI.Err(); err != nil {
			log.Fatal(err)
		}
		AddCache(order.OrderUid, order)

	}

}

func AddCache(id string, order models.Order) {
	Cache[id] = order
}
func GetCache(id string) []byte {
	if _, ok := Cache[id]; ok {
		jsonData, err := json.MarshalIndent(Cache[id], "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		return jsonData
	} else {
		return nil
	}

}
