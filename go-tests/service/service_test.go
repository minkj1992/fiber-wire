package service_test

import (
	"fmt"
	"testing"

	"github.com/minkj1992/go-playground/go-tests/service"
	"github.com/stretchr/testify/mock"
)

type smsServiceMock struct {
	mock.Mock
}

func (m *smsServiceMock) SendChargeNotification(v int) bool {
	fmt.Println("Mocked charge notification is called.")
	args := m.Called(v)
	return args.Bool(0)
}

func TestChargeCustomer(t *testing.T) {
	given := 100
	smsService := new(smsServiceMock)
	smsService.On("SendChargeNotification", given).Return(true) // Return args will be passed to m.Called's return value

	ms := service.MyService{smsService}
	ms.ChargeCustomer(given)

	smsService.AssertExpectations(t)
}
