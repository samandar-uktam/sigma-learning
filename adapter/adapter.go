package adapter

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// Here legacy payment system has a different method than PaymentProcessor interface method
type LegacyPaymentSystem struct{}

func (l *LegacyPaymentSystem) MakeTransaction(value float64) string {
	return fmt.Sprintf("Processed legacy payment if $%.2f", value)
}

// So, we need to make an adapter.
type LegacyAdapter struct {
	legacy *LegacyPaymentSystem
}

func (l *LegacyAdapter) ProcessPayment(amount float64) string {
	return l.legacy.MakeTransaction(amount)
}

func Checkout(p PaymentProcessor, total float64) {
	fmt.Println(p.ProcessPayment(total))
}

func main() {
	legacy := &LegacyPaymentSystem{}
	adapter := &LegacyAdapter{legacy: legacy}

	Checkout(adapter, 99.99)
}
