package main

import (
	"context"
	"net/http"
	"superTools-frontground-backend/global"
	"sync"
)

/**
* @Author: super
* @Date: 2020-12-17 18:13
* @Description:
**/

var sum int64 = 0

//预存的商品数量
var productSum int64 = 10000

//互斥锁
var mutex sync.Mutex

func main() {
	http.HandleFunc("/getOne", GetProduct)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		global.Logger.Fatal(context.Background(), err)
	}
}

func GetProduct(w http.ResponseWriter, req *http.Request) {
	if GetOneProduct() {
		w.Write([]byte("true"))
	}
	w.Write([]byte("false"))
}

func GetOneProduct() bool {
	mutex.Lock()
	defer mutex.Unlock()
	if sum < productSum {
		sum += 1
		return true
	}
	return false
}
