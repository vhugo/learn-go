// Factory
//
// Its purpose is to abstract the user from the knowledge of the struct he needs
// to achieve for a specific purpose, such as retrieving some value, maybe from
// a web service or a database. The user only needs an interface that provides
// him this value. By delegating this decision to a Factory, this Factory can
// provide an interface that fits the user needs. It also eases the process of
// downgrading or upgrading of the implementation of the underlying type if
// needed.
//
// When using the Factory method design pattern, we gain an extra layer of
// encapsulation so that our program can grow in a controlled environment. With
// the Factory method, we delegate the creation of families of objects to a
// different package or object to abstract us from the knowledge of the pool of
// possible objects we could use.
//
// Objectives:
//
// - Delegating the creation of new instances of structures to a different part
// of the program
//
// - Working at the interface level instead of with concrete implementations
//
// - Grouping families of objects to obtain a family object creator
//
package factory

import "fmt"

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	_ = iota
	Cash
	DebitCard
	CreditCard
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil

	case DebitCard:
		return new(DebitCardPM), nil

	case CreditCard:
		return new(CreditCardPM), nil

	default:
		return nil, fmt.Errorf("A payment method with ID %d must return an error", m)
	}
}

type CashPM struct{}
type DebitCardPM struct{}
type CreditCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

func (c *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using credit card\n", amount)
}
