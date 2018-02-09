// Abstract Factory
//
// The Abstract Factory design pattern is a new layer of grouping to achieve a
// bigger (and more complex) composite object, which is used through its
// interfaces. The idea behind grouping objects in families and grouping
// families is to have big factories that can be interchangeable and can grow
// more easily. In the early stages of development, it is also easier to work
// with factories and abstract factories than to wait until all concrete
// implementations are done to start your code. Also, you won't write an
// Abstract Factory from the beginning unless you know that your object's
// inventory for a particular field is going to be very large and it could be
// easily grouped into families.
//
// Grouping related families of objects is very convenient when your object
// number is growing so much that creating a unique point to get them all seems
// the only way to gain the flexibility of the runtime object creation. The
// following objectives of the Abstract Factory method must be clear to you:
//
// - Provide a new layer of encapsulation for Factory methods that return a
// common interface for all factories
//
// - Group common factories into a super Factory (also called a factory of
// factories)
//
package main
