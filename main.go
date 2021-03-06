package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"superTools-frontground-backend/pkg/elastic"
	"superTools-frontground-backend/pkg/mq"
	"time"

	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/routers"
	"superTools-frontground-backend/pkg/cache"
	"superTools-frontground-backend/pkg/db"
	"superTools-frontground-backend/pkg/idGenerator"
	"superTools-frontground-backend/pkg/logger"
	"superTools-frontground-backend/pkg/setting"
	"superTools-frontground-backend/pkg/tracer"

	"gopkg.in/natefinch/lumberjack.v2"
)

/**
* @Author: super
* @Date: 2020-08-21 20:37
* @Description:
**/
var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func init() {
	//读取命令行参数
	err := setupFlag()
	if err != nil {
		log.Printf("init.setupFlag err: %v\n", err)
	}
	//初始化配置
	err = setupSetting()
	if err != nil {
		log.Printf("init setupSetting err: %v\n", err)
	}
	//初始化日志
	err = setupLogger()
	if err != nil {
		log.Printf("init setupLogger err: %v\n", err)
	}
	//初始化数据库
	err = setupDBEngine()
	if err != nil {
		log.Printf("init setupDBEngine err: %v\n", err)
	}
	//初始化redis
	err = setupCacheEngine()
	if err != nil {
		log.Printf("init setupCacheEngine err: %v\n", err)
	}
	//初始化RabbitMQ
	err = setupRabbitMQEngine()
	if err != nil {
		log.Printf("init setupRabbitMQEngine err: %v\n", err)
	}
	//初始化elastic
	//err = setupElasticEngine()
	//if err != nil {
	//	log.Printf("init setupElasticEngine err: %v\n", err)
	//}
	//初始化追踪
	err = setupTracer()
	if err != nil {
		log.Printf("init.setupTracer err: %v\n", err)
	}
	//初始化ID生成器
	err = idGenerator.InitSnowflake()
	if err != nil {
		log.Printf("init.snowflak err: %v\n", err)
	}
}

// @title superTools前端后台
// @version 1.0
// @description 超级工具合集
// @Github https://github.com/superToolsCollection
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := pingServer(); err != nil {
			global.Logger.Errorf(context.Background(), "The server has no response, or it might took too long to start up.")
		}
		global.Logger.Info(context.Background(), "The server has been deployed successfully.")
	}()

	global.Logger.Infof(context.Background(), "Start to listening the incoming requests on http address :%s", global.ServerSetting.HttpPort)
	err := s.ListenAndServe()
	if err != nil {
		global.Logger.Fatalf(context.Background(), "start listen server err: %v", err)
	}
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(":" + global.ServerSetting.HttpPort + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		global.Logger.Info(context.Background(), "Waiting for the server, retry in 1 second.")
	}
	return errors.New("cannot connect to the server")
}

func setupSetting() error {
	newSetting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Cache", &global.CacheSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("RabbitMQ", &global.RabbitMQSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Elastic", &global.ElasticSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = db.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupCacheEngine() error {
	var err error
	global.RedisEngine, err = cache.NewRedisEngine(global.CacheSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupRabbitMQEngine() error {
	var err error
	global.RabbitMQEngine, err = mq.NewRabbitMQEngine(global.RabbitMQSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupElasticEngine() error {
	var err error
	global.ElasticEngine, err = elastic.NewElasticEngine(global.ElasticSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	fmt.Println("log file name ", fileName)
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("bedtimeStory", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
