package model

/**
* @Author: super
* @Date: 2020-12-17 19:04
* @Description: 用于高并发秒杀消息传递
**/

type Message struct {
	ProductId string
	UserID    string
}

func NewMessage(productId string, userId string) *Message {
	return &Message{
		ProductId: productId,
		UserID:    userId,
	}
}
