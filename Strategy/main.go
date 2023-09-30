package main

import "fmt"

// PaymentStrategy defines the strategy interface
type PaymentStrategy interface {
	Pay(amount float64)
}

// CreditCardPayment is a concrete strategy for credit card payments
type CreditCardPayment struct{}

func (cc *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Credit Card\n", amount)
}

// PayPalPayment is a concrete strategy for PayPal payments
type PayPalPayment struct{}

func (pp *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal\n", amount)
}

// PaymentContext is the context that uses the selected payment strategy
type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) ProcessPayment(amount float64) {
	pc.strategy.Pay(amount)
}

func main() {
	creditCard := &CreditCardPayment{}
	payPal := &PayPalPayment{}

	context := NewPaymentContext(creditCard)
	context.ProcessPayment(100.0)

	context.SetStrategy(payPal)
	context.ProcessPayment(50.0)
}
