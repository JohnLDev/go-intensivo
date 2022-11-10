package useCases

import (
	"database/sql"
	"testing"

	"github.com/johnldev/imersao-golang/src/order/entity"
	"github.com/johnldev/imersao-golang/src/order/infra/database"
	"github.com/stretchr/testify/suite"

	//sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type CalculateFinalPriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository entity.OrderRepositoryInterface
	Db              *sql.DB
}

// alwasys run before test
func (suite *CalculateFinalPriceUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("create table orders(id varchar(255), tax float, price float, total float)")
	suite.OrderRepository = database.NewOrderRepoistory(db)
	suite.Db = db
}

// always run after test
func (suite *CalculateFinalPriceUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculateFinalPriceUseCaseTestSuite))
}

func (suite *CalculateFinalPriceUseCaseTestSuite) TestShouldCalculateFinalPrice() {
	order, err := entity.NewOrder("1", 10, 10)

	suite.NoError(err)
	suite.NoError(order.CalculateTotal())

	input := CalculateOrderInputDTO{
		ID:    order.ID,
		Tax:   order.Tax,
		Price: order.Price,
	}

	useCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)

	result, err := useCase.Execute(input)
	suite.NoError(err)
	suite.Equal(order.Total, result.Total)
	suite.Equal(order.Price, result.Price)
	suite.Equal(order.ID, result.ID)
	suite.Equal(order.Tax, result.Tax)

}
