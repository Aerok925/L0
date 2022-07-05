package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (db *DataBase) Connect() error {
	var err error
	constr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", db.User, db.Pass, db.Dbname, db.Sslmode)
	db.con, err = sql.Open("postgres", constr)
	if err != nil {
		return err
	}
	fmt.Println("connected!")
	return nil

}

func (db *DataBase) Disconnect() {
	db.con.Close()
}
