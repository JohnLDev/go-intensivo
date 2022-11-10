package useCases

import "github.com/johnldev/imersao-golang/src/order/entity"

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) GetTotalUseCase {
	return GetTotalUseCase{
		OrderRepository: orderRepository,
	}
}

func (useCase GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	count, err := useCase.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{
		Total: count,
	}, nil
}
