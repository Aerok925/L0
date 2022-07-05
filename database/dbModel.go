package database

import (
	"database/sql"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type DataBase struct {
	fileName string
	User     string `yaml:"user"`
	Pass     string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
	con      *sql.DB
}

func Init(name string) (*DataBase, error) {
	db := &DataBase{
		fileName: name,
	}
	yamlFile, err := ioutil.ReadFile(db.fileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
