package database

func (db *DataBase) Insert(str []byte) error {
	_, err := db.con.Exec("insert into test (data) Values ($1)", string(str))
	if err != nil {
		return err
	}

	return nil

}
