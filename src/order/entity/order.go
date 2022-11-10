package entity

import "errors"

type Order struct {
	ID    string
	Tax   float64
	Price float64
	Total float64
}

func NewOrder(id string, tax, price float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Tax:   tax,
		Price: price,
		Total: 0,
	}

	err := order.isValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) isValid() error {
	if o.ID == "" {
		return errors.New("order id is required")
	}

	if o.Tax <= 0 {
		return errors.New("order tax is required")
	}

	if o.Price <= 0 {
		return errors.New("order price is required")
	}

	return nil
}

func (o *Order) CalculateTotal() error {
	o.Total = o.Price + o.Tax
	err := o.isValid()

	if err != nil {
		return err
	}
	return nil
}
