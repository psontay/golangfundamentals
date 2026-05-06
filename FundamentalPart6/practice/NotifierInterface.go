package practice

import (
	"fmt"
)

type Notifier interface {
	Send(msg string) error
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Send(msg string) error {
	fmt.Println("sending email to :", e.Email, " with msg:", msg)
	return nil
}

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Send(msg string) error {
	fmt.Println("sending sms to :", s.Phone, " with msg:", msg)
	return nil
}

func notify(n Notifier, msg string) {
	err := n.Send(msg)
	if err != nil {
		fmt.Println("error:", err)
	}
}
