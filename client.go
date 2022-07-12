package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"html/template"
	"log"
	"nats/cache"
	"nats/database"
	"net/http"
)

func main() {
	// инициализация кэша
	cach := cache.Init()
	// инициализация сонфика БД
	bd, err := database.Init("database/conf.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	// подключение к БД
	err = bd.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer bd.Disconnect()
	// Получение данных из БД
	strs := bd.SelectAll()
	// Загрузка данных в кэш
	cach.LoadCache(&strs)
	// Подключение к nats-streaming
	connect, err := stan.Connect("test-cluster", "123")
	if err != nil {
		log.Println(err)
		return
	}
	defer connect.Close()
	// Анонимная функция обработки полученных данных из nats-streaming
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
	// Подписка/объявление канала nats-streaming
	subscribe, err := connect.Subscribe("1231", fu)
	if err != nil {
		log.Println(err)
		return
	}
	defer subscribe.Close()

	r := mux.NewRouter()
	// хендлеры сервера
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmlTemlates/FoundUid.html")
	})
	r.HandleFunc("/Uid", func(w http.ResponseWriter, r *http.Request) {

		uid := r.FormValue("uid")

		forUid, err := cach.FoundForUid(uid)
		tmpl, _ := template.ParseFiles("htmlTemlates/Uid.html")
		if err != nil {
			tmpl.Execute(w, forUid)
			return
		}
		tmpl.Execute(w, forUid)
	})

	http.Handle("/", r)
	http.ListenAndServe(":8181", nil)
}
