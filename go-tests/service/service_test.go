package service_test

import (
	"testing"

	"github.com/minkj1992/go-playground/go-tests/service"
	"github.com/minkj1992/go-playground/go-tests/service/mocks"
)

func TestChargeCustomer(t *testing.T) {
	given := 100
	smsService := new(mocks.MessageService)
	smsService.On("SendChargeNotification", given).Return(true) // Return args will be passed to m.Called's return value

	ms := service.MyService{smsService}
	ms.ChargeCustomer(given)

	smsService.AssertExpectations(t)
}
