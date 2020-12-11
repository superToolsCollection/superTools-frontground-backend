package common

import "net/http"

/**
* @Author: super
* @Date: 2020-12-11 19:28
* @Description:
**/

type FilterHandle func(rw http.ResponseWriter, req *http.Request) error

