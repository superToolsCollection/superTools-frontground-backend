package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"superTools-frontground-backend/configs"
	"superTools-frontground-backend/pkg/consistentHash"
	"superTools-frontground-backend/pkg/util"
	"superTools-frontground-backend/services/permission/common"
	"sync"
)

/**
* @Author: super
* @Date: 2020-12-11 19:26
* @Description: 用于分布式权限验证
**/

var port = 8081
var hashConsistent *consistentHash.Consistent

type  AccessControl struct {
	sourcesArray map[string]interface{}
	sync.RWMutex
}

var accessControl = &AccessControl{
	sourcesArray:make(map[string]interface{}),
}

//获取用户设置的数据
func (m *AccessControl) GetRecord(uid string)interface{}{
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	data := m.sourcesArray[uid]
	return data
}

//用户设置数据
func (m *AccessControl) SetRecord(uid string){
	m.RWMutex.Lock()
	//todo
	m.sourcesArray[uid] = "hello"
	m.RWMutex.Unlock()
}

func (m *AccessControl) GetDistributedRight(req *http.Request)bool{
	uid, err := req.Cookie("loginUserJson")
	if err != nil{
		return false
	}
	hostRequest, err := hashConsistent.Get(uid.Value)
	if err != nil{
		return false
	}
	//判断是否为本机
	if hostRequest == configs.LocalHost {
		//执行本机数据读取和校验
		return m.GetDataFromMap(uid.Value)
	} else {
		//不是本机充当代理访问数据返回结果
		return GetDataFromRemote(hostRequest, req)
	}
}

//获取本机map，并且处理业务逻辑，返回的结果类型为bool类型
func (m *AccessControl) GetDataFromMap(uid string) (isOk bool) {
	data := m.GetRecord(uid)

	//执行逻辑判断
	if data != nil {
		return true
	}
	return
}

//获取其它节点处理结果
func GetDataFromRemote(host string, request *http.Request) bool {
	//获取Uid
	uidPre, err := request.Cookie("loginUserJson")
	if err != nil {
		return false
	}
	//获取sign
	uidSign, err := request.Cookie("sign")
	if err != nil {
		return false
	}

	//模拟接口访问，
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+host+":"+string(port)+"/check", nil)
	if err != nil {
		return false
	}

	//手动指定，排查多余cookies
	cookieUid := &http.Cookie{Name: "loginUserJson", Value: uidPre.Value, Path: "/"}
	cookieSign := &http.Cookie{Name: "sign", Value: uidSign.Value, Path: "/"}
	//添加cookie到模拟的请求中
	req.AddCookie(cookieUid)
	req.AddCookie(cookieSign)

	//获取返回结果
	response, err := client.Do(req)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false
	}

	//判断状态
	if response.StatusCode == 200 {
		if string(body) == "true" {
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	//负载均衡器设置，采用一致Hash算法
	hashConsistent = consistentHash.NewConsistent()
    //通过一致性Hash算法添加节点
    for _, v := range configs.HOST_ARRAY{
    	hashConsistent.Add(v)
	}
	//创建过滤器
	filter := common.NewFilter()
	//注册拦截器
	filter.RegisterFilterUri("/check", Auth)
	//启动服务
	http.HandleFunc("/check", filter.Handle(Check))
	http.ListenAndServe(":8083", nil)
}

func Auth(rw http.ResponseWriter, r *http.Request) error {
	//添加基于cookie的权限验证
	err := CheckUserInfo(r)
	if err != nil {
		return err
	}
	return nil
}

//身份校验函数
func CheckUserInfo(r *http.Request) error {
	uidCookie, err := r.Cookie("loginUserJson")
	if err != nil {
		return errors.New("用户UID Cookie 获取失败")
	}
	signCookie, err := r.Cookie("sign")
	if err != nil {
		return errors.New("用户加密串 Cookie 获取失败")
	}
	signByte, err := util.DePwdCode(signCookie.Value)
	if err != nil {
		return errors.New("加密串被篡改")
	}
	if !checkInfo(uidCookie.Value, string(signByte)) {
		return errors.New("身份校验失败")
	}
	return nil
}

//自定义加密cookie是否有效
func checkInfo(checkStr string, signStr string) bool {
	if checkStr == signStr {
		return true
	}
	return false
}

//执行正常业务逻辑
func Check(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("执行check")
}
