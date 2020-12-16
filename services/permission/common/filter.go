package common

import (
	"net/http"
)

/**
* @Author: super
* @Date: 2020-12-11 19:28
* @Description:
**/

type FilterHandle func(rw http.ResponseWriter, req *http.Request) error

type Filter struct {
	//用于存储需要拦截的URI
	filterMap map[string]FilterHandle
}

func NewFilter() *Filter {
	return &Filter{
		filterMap: make(map[string]FilterHandle),
	}
}

//注册拦截器
func (f *Filter) RegisterFilterUri(uri string, handler FilterHandle) {
	f.filterMap[uri] = handler
}

//根据uri获取对应的handler
func (f *Filter) GetFilterHandler(uri string) FilterHandle {
	return f.filterMap[uri]
}

type WebHandle func(rw http.ResponseWriter, req *http.Request)

func (f *Filter) Handle(webHandle WebHandle) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		for path, handle := range f.filterMap {
			if path == r.RequestURI {
				err := handle(rw, r)
				if err != nil {
					//response := app.NewResponse(nil)
					//response.ToErrorResponse(err)
					rw.Write([]byte(err.Error()))
					return
				}
				break
			}
		}
		//处理正常的业务函数
		webHandle(rw, r)
	}
}
