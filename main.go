package main

import (
	"github.com/joho/godotenv"
	"ioc-provider/crawler"
	"ioc-provider/db"
	"ioc-provider/helper"
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
		Host:     rbmqHost,
		Port:     rbmqPort,
	}
	clientRB.ConnectRbmq()

	iocHandler := IocHandler{
		IocRepo: repo_impl.NewIocRepo(clientES),
	}
	// time start crawler
	//crawler.GetAllDataSubscribed(iocHandler.IocRepo)
	//crawler.LiveHunting(iocHandler.IocRepo)
	//crawler.MirrorPost(iocHandler.IocRepo)
	crawler.Subscribed(iocHandler.IocRepo)

	// schedule crawler
	schedule(10*time.Second, iocHandler, 1)
}

func schedule(timeSchedule time.Duration, handler IocHandler, crowIlnndex int) {
	ticker := time.NewTicker(timeSchedule)
	func() {
		for {
			switch crowIlnndex {

			case 1:
				<-ticker.C
				//fmt.Println("2")
				//crawler.GetAllDataSubscribed(handler.IocRepo)
				//crawler.LiveHunting(handler.IocRepo)
				//crawler.MirrorPost(handler.IocRepo)
				crawler.Subscribed(handler.IocRepo)
			}
		}
	}()
}