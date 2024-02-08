package master_nats

import (
	"github.com/nats-io/stan.go"
	"log"
)

func ConnectToNats() (stan.Conn, error) {
	sc, err := stan.Connect("test-cluster", "client-123", stan.NatsURL("nats://127.0.0.1:4222"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Успешно подключились к NATS Streaming!")

	return sc, nil
}
