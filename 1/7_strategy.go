package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardStrategy struct {
	cardNumber string
	cvv        string
	expiration string
}

func (c *CreditCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Оплата по кредитной карте %s в размере %.2f", c.cardNumber, amount)
}

type PayPalStrategy struct {
	email    string
	password string
}

func (p *PayPalStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Оплата через PayPal для %s в размере %.2f", p.email, amount)
}

type BankTransferStrategy struct {
	accountNumber string
	bankName      string
}

func (b *BankTransferStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Оплата через банковский перевод на счет %s в банке %s в размере %.2f", b.accountNumber, b.bankName, amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) string {
	return p.strategy.Pay(amount)
}

// func main() {
// 	creditCardStrategy := &CreditCardStrategy{
// 		cardNumber: "1234-5678-9012-3456",
// 		cvv:        "123",
// 		expiration: "12/2025",
// 	}

// 	paypalStrategy := &PayPalStrategy{
// 		email:    "example@example.com",
// 		password: "password",
// 	}

// 	bankTransferStrategy := &BankTransferStrategy{
// 		accountNumber: "1234567890",
// 		bankName:      "Example Bank",
// 	}

// 	paymentContext := &PaymentContext{}

// 	paymentContext.SetStrategy(creditCardStrategy)
// 	fmt.Println(paymentContext.Pay(100.0)) // оплата по кредитной карте 1234-5678-9012-3456 в размере 100.00

// 	paymentContext.SetStrategy(paypalStrategy)
// 	fmt.Println(paymentContext.Pay(200.0)) // оплата через PayPal для example@example.com в размере 200.00

// 	paymentContext.SetStrategy(bankTransferStrategy)
// 	fmt.Println(paymentContext.Pay(300.0)) // оплата через банковский перевод на счет 1234567890 в банке Example Bank в размере 300.00
// }
