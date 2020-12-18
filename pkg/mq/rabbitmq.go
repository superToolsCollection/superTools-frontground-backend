package mq

/**
* @Author: super
* @Date: 2020-11-18 11:50
* @Description: rabbitMQ连接池
**/

func SendMessage(message string, rabbitMQEngine * RabbitMQ){
	rabbitMQEngine.PublishSimple(message)
}