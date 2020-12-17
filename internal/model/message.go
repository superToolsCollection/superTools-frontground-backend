package model

/**
* @Author: super
* @Date: 2020-12-17 19:04
* @Description: 用于高并发秒杀消息传递
**/

type Message struct {
	ProductId int64
	UserID    string
}

func NewMessage(productId int64, userId string) *Message {
	return &Message{
		ProductId: productId,
		UserID:    userId,
	}
}
