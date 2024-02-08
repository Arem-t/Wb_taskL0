package postgres

import (
	"WB00L0/internal/models"
	"database/sql"
)

func InsertOrder(db *sql.DB, order models.Order) error {

	_, err := db.Exec(`
			INSERT INTO orders
			(order_uid, track_number, entry, locale, internal_signature,
			customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
			VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, order.OrderUid, order.TrackNumber, order.Entry, order.Locale,
		order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId,
		order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
			INSERT INTO deliveries
			(order_uid, name, phone, zip, city,
			address, region, email)
			VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)`, order.OrderUid, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
		order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
			INSERT INTO payments
			(order_uid, transaction, request_id, currency, provider,
			amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, order.OrderUid, order.Payment.Transaction, order.Payment.RequestId,
		order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank,
		order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		_, err = db.Exec(`
			INSERT INTO items
			(order_uid, chrt_id, track_number, price, rid,
			name, sale, size, total_price, nm_id, brand, status)
			VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`, order.OrderUid, item.ChrtId, item.TrackNumber, item.Price, item.Rid,
			item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)
		if err != nil {
			return err
		}
	}

	return nil

}
