package database

import "sync"

func (db *DataBase) Insert(str []byte) error {
	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	_, err := db.con.Exec("insert into test (data) Values ($1)", string(str))
	if err != nil {
		return err
	}

	return nil

}
