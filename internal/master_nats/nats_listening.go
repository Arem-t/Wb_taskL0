package master_nats

import (
	"WB00L0/internal/models"
	"WB00L0/internal/postgres"
	"WB00L0/internal/server"
	"database/sql"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

func ListenAndSubscribe(sc stan.Conn, db *sql.DB) {
	log.Println("Прослушивание началось...")

	_, err := sc.Subscribe("your_channel", func(msg *stan.Msg) {
		var order models.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Println(err)
			return
		}

		server.AddCache(order.OrderUid, order)

		err = postgres.InsertOrder(db, order)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Сообщение успешно доставлено")
	})
	if err != nil {
		log.Fatal(err)
	}
}
