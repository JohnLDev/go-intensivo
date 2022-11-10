package database

import (
	"database/sql"

	"github.com/johnldev/imersao-golang/src/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepoistory(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("insert into orders(id, tax, price, total) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Tax, order.Price, order.Total)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
