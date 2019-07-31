package core

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitMysql () *sql.DB {
	mysqlConf, err := Conf.Map("mysql")
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlConf["user"].(string), mysqlConf["passwd"].(string), mysqlConf["host"].(string), mysqlConf["port"].(int), mysqlConf["db"].(string))
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	return db
}