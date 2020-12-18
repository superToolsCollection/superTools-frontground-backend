package mq

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/pkg/setting"
	"sync"
)

/**
* @Author: super
* @Date: 2020-08-13 08:44
* @Description: todo：修改rabbitMQ的实现方式，采用连接池的形式
**/

type RabbitMQ struct {
	conn    *amqp.Connection
	Channel *amqp.Channel

	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
	sync.Mutex
}

// 创建RabbitMQ实例
func NewRabbitMQ(queuqName string,
	exchange string, key string,
	rabbitMQSetting *setting.RabbitMQSettingS) (*RabbitMQ, error) {

	rabbitMQUrl := "amqp://" + rabbitMQSetting.UserName + ":" + rabbitMQSetting.Password + "@" + rabbitMQSetting.Host + "/"
	rabbitmq := &RabbitMQ{
		QueueName: queuqName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     rabbitMQUrl,
	}
	//创建rabbitmq连接
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil{
		return nil, err
	}
	rabbitmq.Channel, err = rabbitmq.conn.Channel()
	if err != nil{
		return nil, err
	}
	return rabbitmq,  nil
}

// 断开channel和connection的连接释放资源
func (r *RabbitMQ) Destory() {
	r.Channel.Close()
	r.conn.Close()
}

//自定义错误处理函数
func (r *RabbitMQ) FailOnErr(err error, message string) {
	if err != nil {
		global.Logger.Error(context.Background(), err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}
