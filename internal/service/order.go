package service

import (
	"fmt"

	"superTools-frontground-backend/internal/dao"
	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-11-21 11:32
* @Description:
**/
type OrderRequest struct {
	ID int64 `form:"id" binding:"required,gte=1"`
}

type InsertOrderRequest struct {
	ID        int64  `form:"id" binding:"required,gte=1"`
	UserID    string `form:"user_id" binding:"required,min=2,max=4294967295"`
	ProductID int64  `form:"product_id" binding:"required,gte=1"`
	State     int    `form:"state" binding:"required,gte=1"`
}

type UpdateOrderRequest struct {
	ID        int64  `form:"id" binding:"required,gte=1"`
	UserID    string `form:"user_id" binding:"required,min=2,max=4294967295"`
	ProductID int64  `form:"product_id" binding:"required,gte=1"`
	State     int    `form:"state" binding:"required,gte=1"`
}

type DeleteOrderRequest struct {
	ID int64 `form:"id" binding:"required,gte=1"`
}

type GetOrderByUserIDRequest struct {
	UserID string `form:"user_id" binding:"required,min=2,max=4294967295"`
}

type GetOrderListByUserIDRequest struct {
	UserID string `form:"user_id" binding:"required,min=2,max=4294967295"`
}

type Order struct {
	ID        int64  `json:"id"`
	UserID    string `json:"user_id"`
	ProductID int64  `json:"product_id"`
	State     int    `json:"state"`
}

type IOrderService interface {
	GetOrderByID(param *OrderRequest) (*Order, error)
	GetAllOrder() ([]*Order, error)
	GetOrderList(pager *app.Pager) ([]*Order, error)
	GetOrderByUserID(param *GetOrderByUserIDRequest) ([]*Order, error)
	GetOrderListByUserID(param *GetOrderListByUserIDRequest, pager *app.Pager) ([]*Order, error)
	GetAllOrderWithInfo() ([]*dao.OrderRow, error)
	DeleteOrderByID(param *DeleteOrderRequest) bool
	InsertOrder(param *InsertOrderRequest) (int64, error)
	UpdateOrder(param *UpdateOrderRequest) error
}

type OrderService struct {
	orderDao dao.IOrder
}

func NewOrderService(orderDao dao.IOrder) IOrderService {
	return &OrderService{orderDao: orderDao}
}

func (s *OrderService) GetOrderByID(param *OrderRequest) (*Order, error) {
	order, err := s.orderDao.SelectByKey(param.ID)
	if err != nil {
		return nil, err
	}
	return &Order{
		ID:        order.ID,
		UserID:    order.UserID,
		ProductID: order.ProductID,
		State:     order.State,
	}, nil
}

func (s *OrderService) GetAllOrder() ([]*Order, error) {
	orders, err := s.orderDao.SelectAll()
	if err != nil {
		return nil, err
	}
	var orderList []*Order
	for _, order := range orders {
		orderList = append(orderList, &Order{
			ID:        order.ID,
			UserID:    order.UserID,
			ProductID: order.ProductID,
			State:     order.State,
		})
	}
	return orderList, nil
}

func (s *OrderService) GetOrderList(pager *app.Pager) ([]*Order, error) {
	orders, err := s.orderDao.SelectList(pager.Page, pager.PageSize)
	if err != nil {
		return nil, err
	}
	var orderList []*Order
	for _, order := range orders {
		orderList = append(orderList, &Order{
			ID:        order.ID,
			UserID:    order.UserID,
			ProductID: order.ProductID,
			State:     order.State,
		})
	}
	return orderList, nil
}

func (s *OrderService) GetOrderByUserID(param *GetOrderByUserIDRequest) ([]*Order, error) {
	fmt.Println("service user_id", param.UserID)
	orders, err := s.orderDao.SelectByUser(param.UserID)
	if err != nil {
		return nil, err
	}
	var orderList []*Order
	for _, order := range orders {
		orderList = append(orderList, &Order{
			ID:        order.ID,
			UserID:    order.UserID,
			ProductID: order.ProductID,
			State:     order.State,
		})
	}
	return orderList, nil
}

func (s *OrderService) GetOrderListByUserID(param *GetOrderListByUserIDRequest, pager *app.Pager) ([]*Order, error) {
	orders, err := s.orderDao.SelectByUserList(param.UserID, pager.Page, pager.PageSize)
	if err != nil {
		return nil, err
	}
	var orderList []*Order
	for _, order := range orders {
		orderList = append(orderList, &Order{
			ID:        order.ID,
			UserID:    order.UserID,
			ProductID: order.ProductID,
			State:     order.State,
		})
	}
	return orderList, nil
}

func (s *OrderService) GetAllOrderWithInfo() ([]*dao.OrderRow, error) {
	orders, err := s.orderDao.SelectAllWithInfo()
	if err != nil {
		return nil, err
	}
	var orderList []*dao.OrderRow
	for _, order := range orders {
		orderList = append(orderList, &dao.OrderRow{
			OrderID:     order.OrderID,
			UserID:      order.UserID,
			ProductName: order.ProductName,
			OrderState:  order.OrderState,
		})
	}
	return orderList, nil
}

func (s *OrderService) DeleteOrderByID(param *DeleteOrderRequest) bool {
	return s.orderDao.Delete(param.ID)
}

func (s *OrderService) InsertOrder(param *InsertOrderRequest) (int64, error) {
	product := &dao.Order{
		ID:        param.ID,
		UserID:    param.UserID,
		ProductID: param.ProductID,
		State:     param.State,
	}
	return s.orderDao.Insert(product)
}

func (s *OrderService) UpdateOrder(param *UpdateOrderRequest) error {
	product := &dao.Order{
		ID:        param.ID,
		UserID:    param.UserID,
		ProductID: param.ProductID,
		State:     param.State,
	}
	err := s.orderDao.Update(product)
	if err != nil {
		return err
	}
	return nil
}
