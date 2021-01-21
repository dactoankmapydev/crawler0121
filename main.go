package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"ioc-provider/crawler"
	"ioc-provider/db"
	"ioc-provider/repository"
	"ioc-provider/repository/repo_impl"
	"log"
	"os"
	"time"
)

type IocHandler struct {
	IocRepo repository.IocRepo
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("not environment variable")
	}
}

func main() {
	// elastic details
	esHost := os.Getenv("ES_HOST")
	esPort := os.Getenv("ES_PORT")

	// connect elastic
	client := &db.ElasticDB{
		Host: esHost,
		Port: esPort,
	}
	client.NewElasticDB()

	iocHandler := IocHandler{
		IocRepo: repo_impl.NewIocRepo(client),
	}

	// time start crawler
	go scheduleUpdate(60*time.Second, iocHandler)
}

func scheduleUpdate(timeSchedule time.Duration, handler IocHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Crawler data...")
				crawler.Subscribed(handler.IocRepo)
				crawler.LiveHunting(handler.IocRepo)
			}
		}
	}()
}
