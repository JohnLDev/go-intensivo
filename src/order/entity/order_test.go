package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorWhenIDIsEmpty(t *testing.T) {
	order, err := NewOrder("", 0, 0)
	assert.Nil(t, order)
	assert.Equal(t, "order id is required", err.Error())
}

func Test_ShouldReturnErrorWhenTaxIsEmpty(t *testing.T) {
	order, err := NewOrder("asdas", 0, 0)
	fmt.Println(err)
	assert.Nil(t, order)
	assert.Equal(t, "order tax is required", err.Error())
}

func TestShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("ID", 2, 10)
	assert.Nil(t, err)
	order.CalculateTotal()
	assert.Equal(t, 12.0, order.Total)
}
