package app

import (
	"WB00L0/internal/master_nats"
	"WB00L0/internal/postgres"
	"WB00L0/internal/server"
	"context"
	"log"
	"os/signal"
	"syscall"
)

func StartL0() {

	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()

	db, err := postgres.ConnectToPostgres()
	if err != nil {
		log.Fatal(err.Error())
	}

	sc, err := master_nats.ConnectToNats()
	if err != nil {
		log.Fatal(err.Error())
	}
	server.RestoreCache(db)
	master_nats.ListenAndSubscribe(sc, db)
	server.StartHandler()

	<-ctx.Done()
	db.Close()
	sc.Close()
}
