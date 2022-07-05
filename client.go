package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
	"nats/cache"
	"nats/database"
	"time"
)

func handlerMessage(msg *stan.Msg, ch *cache.Cache, base *database.DataBase) {

}

func main() {

	cach := cache.Init()
	bd, err := database.Init("database/conf.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	bd.Connect()
	defer bd.Disconnect()
	strs := bd.SelectAll()
	cach.LoadCache(&strs)
	connect, err := stan.Connect("test-cluster", "123")
	if err != nil {
		log.Println(err)
		return
	}
	defer connect.Close()
	fu := func(msg *stan.Msg) {
		err := cach.FileInCache(msg.Data)
		if err != nil {
			log.Println(err)
			return
		}
		err = bd.Insert(msg.Data)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Данные успешно добавлены!")
	}
	subscribe, err := connect.Subscribe("1231", fu)
	if err != nil {
		return
	}
	defer subscribe.Close()

	time.Sleep(time.Minute)
}
