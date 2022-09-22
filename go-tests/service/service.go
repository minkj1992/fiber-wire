package service

import "fmt"

type MessageService interface {
	SendChargeNotification(int) bool
}

type SMSService struct{}

type MyService struct {
	MessageService MessageService
}

func (sms SMSService) SendChargeNotification(v int) bool {
	fmt.Println("Sending production charge notification.")
	return true
}

func (ms MyService) ChargeCustomer(v int) error {
	ms.MessageService.SendChargeNotification(v)
	fmt.Printf("Charging the customer for value %d.\n", v)
	return nil
}
