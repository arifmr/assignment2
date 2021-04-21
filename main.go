package main

import (
	"assignment2/config/postgres"
	"assignment2/http/routes"
	orders "assignment2/repository/postgres"

	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/subosito/gotenv"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if err := gotenv.Load(); err != nil {
		log.Println(err)
	}
}

func main() {
	db := postgres.Connect()
	repo := orders.NewOrderRepo(db)

	errs := make(chan error)

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

		sig := <-sigChan
		log.Printf("Signal notified %v", sig)
		errs <- fmt.Errorf("%v", sig)
	}()

	go func() {
		router := routes.NewRouter(repo)
		if err := router.Run(":" + os.Getenv("PORT")); err != nil {
			errs <- err
		}
	}()

	log.Fatal(<-errs)
}
