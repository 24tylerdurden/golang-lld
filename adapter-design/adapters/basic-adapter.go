package adapters

import "fmt"

// Target Interface
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	RefundPayment(transactionId string) error
}

// Adaptee - Existing inCompatible struct
type PayPalService struct{}

func (p *PayPalService) SendMoney(amount float64) {
	fmt.Printf("Sending $%.2f via PayPal\n", amount)
}

func (p *PayPalService) ProcessRefund(txId string) {
	fmt.Printf("Processing Refund for txn : %s \n", txId)
}

// Adaptee - Implements Traget Interface

type PayPalAdapter struct {
	payPalService *PayPalService
}

func NewPayPalAdapter(payPalService *PayPalService) *PayPalAdapter {
	return &PayPalAdapter{
		payPalService: payPalService,
	}
}

func (a *PayPalAdapter) ProcessPayment(amount float64) error {
	a.payPalService.SendMoney(amount)
	return nil
}

func (a *PayPalAdapter) RefundPayment(txId string) error {
	a.payPalService.ProcessRefund(txId)
	return nil
}

// Client Code

type EcommerceApp struct {
	paymentProcessor PaymentProcessor
}

func NewEcommerceApp(processor PaymentProcessor) *EcommerceApp {
	return &EcommerceApp{
		paymentProcessor: processor,
	}
}

func (e *EcommerceApp) Checkout(amount float64) {
	e.paymentProcessor.ProcessPayment(amount)
}

func (e *EcommerceApp) InitiateRefund(txnId string) {
	e.paymentProcessor.RefundPayment(txnId)
}

func TestBascAdapter() {
	payPalService := &PayPalService{}
	payPalAdapter := NewPayPalAdapter(payPalService)

	app := NewEcommerceApp(payPalAdapter)

	app.Checkout(99.99)
	app.InitiateRefund("TX123456")
}

// using adapter desing pattern
// we are delegating the actions here ??
//
