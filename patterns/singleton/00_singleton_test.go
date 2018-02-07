// Requirements and acceptance criteria
//
// There are some requirements and acceptance criteria to write the described
// single counter..
//
// They are as follows:
//
// 1 - When no counter has been created before, a new one is created with the
// value 0
//
// 2 - If a counter has already been created, return this instance that holds the
// actual count
//
// 3 - If we call the method AddOne, the count must be incremented by 1 We have a
// scenario with three tests to check in our unit tests.
//
package singleton

import "testing"

func TestGetInstance(t *testing.T) {

	counter1 := GetInstance()
	if counter1 == nil {
		// Test of acceptance criteria 1 failed
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 "+
			"but it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		// Test 2 failed
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current "+
			"count must be 2 but was %d\n", currentCount)
	}
}
