package main

import (
	"github.com/joho/godotenv"
	"ioc-provider/db"
	"ioc-provider/helper"
	"ioc-provider/repository"
	"log"
	"os"
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
	clientES := &db.ElasticDB{
		Host: esHost,
		Port: esPort,
	}
	clientES.NewElasticDB()

	rbmqHost := os.Getenv("RBMQ_HOST")
	rbmqPort := os.Getenv("RBMQ_PORT")
	rbmqUserName := os.Getenv("RBMQ_USER_NAME")
	rbmqPassword := os.Getenv("RBMQ_PASSWORD")

	clientRB := &helper.Rbmq{
		UserName: rbmqUserName,
		Password: rbmqPassword,
		Host: rbmqHost,
		Port: rbmqPort,
	}
	clientRB.ConnectRbmq()

	/*iocHandler := IocHandler{
		IocRepo: repo_impl.NewIocRepo(clientES),
	}

	// time start crawler
	go scheduleUpdate(60*time.Second, iocHandler)*/
}

/*func scheduleUpdate(timeSchedule time.Duration, handler IocHandler) {
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
}*/
