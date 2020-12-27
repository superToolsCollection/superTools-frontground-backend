package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"superTools-frontground-backend/configs"
)

/**
* @Author: super
* @Date: 2020-12-26 16:02
* @Description:
**/

type Region struct {
	Id         int64  `json:"id" sql:"id"`
	Code       int64  `json:"code" sql:"code"`
	ParentCode int64  `json:"parent_code" sql:"parent_code"`
	Name       string `json:"name" sql:"name"`
	Level      int    `json:"level" sql:"level"`
	Center     string `json:"center" sql:"center"`
	Polyline   string `json:"polyline" sql:"polyline"`
}

func main() {
	file, err := os.Open("china-region.json")
	if err != nil {
		panic("can not open file")
	}
	defer file.Close()
	regions := []Region{}
	fd, err := ioutil.ReadAll(file)
	err = json.Unmarshal(fd, &regions)
	if err != nil {
		fmt.Println(err)
	}
	dbEngine, err := NewMysqlConn()
	for _, region := range regions {
		err = Insert(dbEngine, region)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//创建mysql 连接
func NewMysqlConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", configs.MysqlURL)
	return
}

func Insert(conn *sql.DB, region Region) error {
	sql := "INSERT `region` SET code=?, parent_code=?, name=?, level=?, center=?, polyline=?"

	stmt, errStmt := conn.Prepare(sql)
	if errStmt != nil {
		return errStmt
	}

	_, errResult := stmt.Exec(region.Code, region.ParentCode, region.Name, region.Level, region.Center, region.Polyline)
	if errResult != nil {
		return errResult
	}
	return nil
}
