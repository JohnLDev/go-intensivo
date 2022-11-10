package database

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/johnldev/imersao-golang/src/order/entity"
	"github.com/stretchr/testify/suite"

	//sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

// alwasys run before test
func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("create table orders(id varchar(255), tax float, price float, total float)")
	suite.Db = db
}

// always run after test
func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestShouldSaveOrder() {
	order, err := entity.NewOrder("1", 10, 10)
	suite.NoError(err)
	suite.NoError(order.CalculateTotal())
	repo := NewOrderRepoistory(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("select id, tax, price, total from orders where id = ?", order.ID).Scan(&orderResult.ID, &orderResult.Tax, &orderResult.Price, &orderResult.Total)
	suite.NoError(err)
	fmt.Println("OrderResult", orderResult)
	suite.Equal("1", orderResult.ID)
	suite.Equal(10.0, orderResult.Tax)
	suite.Equal(10.0, orderResult.Price)
	suite.Equal(20.0, orderResult.Total)

}
