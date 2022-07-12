package main

import (
	"fmt"
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
		log.Println(err)
		return
	}
	defer subscribe.Close()

	r := mux.NewRouter()
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
		//http.ServeFile(w, r, "htmlTemlates/index.html")
	})

	http.Handle("/", r)
	http.ListenAndServe(":8181", nil)
}
