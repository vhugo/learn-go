// Acceptance criteria
//
// Using the previous description, the requirements for the acceptance criteria
// are the following:
//
// - To have a common method for every payment method called Pay
//
// - To be able to delegate the creation of payments methods to the Factory
//
// - To be able to add more payment methods to the library by just adding it to
// the factory method
//
package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Error("A payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(22.30)

	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestGetPaymentMethodCreditCard(t *testing.T) {
	payment, err := GetPaymentMethod(CreditCard)

	if err != nil {
		t.Error("A payment method of type 'CreditCard' must exist")
	}

	msg := payment.Pay(22.30)

	if !strings.Contains(msg, "paid using credit card") {
		t.Error("The credit card payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExisting(t *testing.T) {
	_, err := GetPaymentMethod(20)

	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
