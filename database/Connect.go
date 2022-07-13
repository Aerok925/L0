package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// Connect - подключение к БД
// return value - в случае успеха nil
// в случае провала error
func (db *DataBase) Connect() (err error) {
	constr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", db.User, db.Pass, db.Dbname, db.Sslmode)
	db.con, err = sql.Open("postgres", constr)
	_, err = db.con.Query("select (data) from test")
	if err != nil {
		return errors.New("Not connect")
	}
	log.Println("Connected!")
	return nil

}

// Disconnect - Отключение от БД
func (db *DataBase) Disconnect() {
	db.con.Close()
}
