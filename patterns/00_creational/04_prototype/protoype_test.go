// Acceptance criteria
//
// To achieve what is described in the example, we will use a prototype of
// shirts. Each time we need a new shirt we will take this prototype, clone it
// and work with it. In particular, those are the acceptance criteria for using
// the Prototype pattern design method in this example:
//
// - To have a shirt-cloner object and interface to ask for different types of
// shirts (white, black, and blue at 15.00, 16.00, and 17.00 dollars
// respectively)
//
// - When you ask for a white shirt, a clone of the white shirt must be made, and
// the new instance must be different from the original one
//
// - The SKU of the created object shouldn't affect new object creation
//
// - An info method must give me all the information available on the instance
// fields, including the updated SKU
//
package main

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Error("Item1 cannot be equal to white prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	shirt1.SKU = "abbccc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("Shirt 1 cannot be equal to Shirt 2")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())
	t.Logf("LOG: The memory positions of the shirts are different %p != %p\n\n",
		&shirt1, &shirt2)
}
