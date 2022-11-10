package useCases

import (
	"github.com/johnldev/imersao-golang/src/order/entity"
)

type CalculateOrderInputDTO struct {
	ID    string
	Tax   float64
	Price float64
}

type CalculateOrderOutputDTO struct {
	ID    string
	Tax   float64
	Price float64
	Total float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface //ACOPLAMENTO
}

func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{OrderRepository: orderRepository}
}

func (c *CalculateFinalPriceUseCase) Execute(input CalculateOrderInputDTO) (*CalculateOrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Tax, input.Price)
	if err != nil {
		return nil, err
	}

	err = order.CalculateTotal()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &CalculateOrderOutputDTO{
		ID:    order.ID,
		Tax:   order.Tax,
		Price: order.Price,
		Total: order.Total,
	}, nil
}
