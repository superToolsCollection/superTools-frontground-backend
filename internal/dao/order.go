package dao

import (
	"errors"
	"fmt"

	"superTools-frontground-backend/internal/model"
	"superTools-frontground-backend/pkg/app"

	"github.com/jinzhu/gorm"
)

/**
* @Author: super
* @Date: 2020-11-21 10:40
* @Description:
**/

type Order struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
	State     int    `json:"state"`
}

type IOrder interface {
	Insert(order *Order) (int64, error)
	Delete(id string) bool
	Update(order *Order) error
	SelectByKey(id string) (*model.Order, error)
	SelectAll() ([]*model.Order, error)
	SelectByUser(userID string) ([]*model.Order, error)
	SelectAllWithInfo() ([]*OrderRow, error)
	SelectList(page, pageSize int) ([]*model.Order, error)
	SelectByUserList(userID string, page, pageSize int) ([]*model.Order, error)
}

type OrderManager struct {
	table string
	conn  *gorm.DB
}

func NewOrderManager(table string, conn *gorm.DB) IOrder {
	return &OrderManager{table: table, conn: conn}
}

func (m *OrderManager) Insert(order *Order) (int64, error) {
	o := &model.Order{
		ID:        order.ID,
		UserID:    order.UserID,
		ProductID: order.ProductID,
		State:     order.State,
	}
	result := m.conn.Create(o)
	if result.RowsAffected == int64(0) {
		return 0, errors.New("insert error")
	}
	return result.RowsAffected, nil
}

func (m *OrderManager) Delete(id string) bool {
	result := m.conn.Where("id = ?", id).Delete(model.Order{})
	if result.RowsAffected == int64(0) {
		return false
	}
	return true
}

func (m *OrderManager) Update(order *Order) error {
	o := &model.Order{
		ID:        order.ID,
		UserID:    order.UserID,
		ProductID: order.ProductID,
		State:     order.State,
	}
	result := m.conn.Model(o).Where("id=?", o.ID).Updates(o)
	if result.RowsAffected == int64(0) {
		return errors.New("update error")
	}
	return nil
}

func (m *OrderManager) SelectByKey(id string) (*model.Order, error) {
	order := &model.Order{}
	result := m.conn.Where("id = ?", id).Find(order)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return order, nil
}

func (m *OrderManager) SelectAll() ([]*model.Order, error) {
	var orders []*model.Order
	if err := m.conn.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (m *OrderManager) SelectByUser(userID string) ([]*model.Order, error) {
	var orders []*model.Order
	fmt.Println("dao user_id", userID)
	if err := m.conn.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

type OrderRow struct {
	OrderID     int64
	ProductName string
	UserID      int64
	OrderState  int
}

func (m *OrderManager) SelectAllWithInfo() ([]*OrderRow, error) {
	fields := []string{"o.id AS order_id", "p.product_name as product_name"}
	fields = append(fields, []string{"u.id AS user_id", "o.state AS order_state"}...)

	rows, err := m.conn.Select(fields).Table(model.Order{}.TableName() + " AS o").
		Joins("LEFT JOIN `" + model.Product{}.TableName() + "` AS p ON o.product_id = p.id").
		Joins("LEFT JOIN `" + model.User{}.TableName() + "` AS u ON o.user_id = u.id").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderRows []*OrderRow
	for rows.Next() {
		r := &OrderRow{}
		if err := rows.Scan(&r.OrderID, &r.ProductName, &r.UserID, &r.OrderState); err != nil {
			return nil, err
		}
		orderRows = append(orderRows, r)
	}
	return orderRows, nil
}

func (m *OrderManager) SelectList(page, pageSize int) ([]*model.Order, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset < 0 && pageSize < 0 {
		pageOffset = 0
		pageSize = 5
	}
	fields := []string{"id", "product_id", "user_id", "state"}
	rows, err := m.conn.Offset(pageOffset).Limit(pageSize).Select(fields).Table(m.table).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		if err := rows.Scan(&order.ID,
			&order.ProductID,
			&order.UserID,
			&order.State); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (m *OrderManager) SelectByUserList(userID string, page, pageSize int) ([]*model.Order, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset < 0 && pageSize < 0 {
		pageOffset = 0
		pageSize = 5
	}
	fields := []string{"id", "product_id", "user_id", "state"}
	rows, err := m.conn.Offset(pageOffset).Limit(pageSize).Where("user_id=?", userID).Select(fields).Table(m.table).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		if err := rows.Scan(&order.ID,
			&order.ProductID,
			&order.UserID,
			&order.State); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
