// Singleton
//
// The Singleton pattern is easy to remember. As the name implies, it will
// provide you with a single instance of an object, and guarantee that there are
// no duplicates.
//
// As a general guide, we consider using the Singleton pattern when the
// following rule applies:
//
// - Need a single, shared value, of some particular type.
// - Need to restrict object creation of some type to a single unit along the
// entire program.
//
package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
