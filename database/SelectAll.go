package database

import "fmt"

func (db *DataBase) SelectAll() []string {
	var retValue []string
	query, err := db.con.Query("select (data) from test")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer query.Close()
	for query.Next() {
		tempstr := ""
		query.Scan(&tempstr)
		retValue = append(retValue, tempstr)
	}
	return retValue
}
