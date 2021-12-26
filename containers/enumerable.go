package containers

// EnumerableWithIndex provides function for ordered containers whose value can be fetched by an index.
type EnumerableWithIndex interface {
	// Each calls the given function once for each element, passing the element's index and value.
	Each(func(index int, value interface{}))

	// Map invokes the given function once for each element and returns a
	// container containing the values returned by the given function.
	// TODO: would apperciate help on how to enforce this in containers (don't want to type assert when chaining)
	// Map(func(index int, value interface{}) interface{}) Container

	// Select returns a new container containing all elements for which the given function returns a true value.
	// TODO: need help on how to enforce this in containers (don't want to type assert when chaining)
	// Select(func(index int, value interface{}) bool) bool

	// Any passes each element of the container to the given function and
	// returns true if the function ever returns true for all elements.
	Any(func(index int, value interface{}) bool) bool

	// All passes each element of the container to the given function and
	// returns true if the function returns true for all elements.
	All(func(index int, value interface{}) bool) bool

	// Find passes each element of the container to the given function and returns
	// the first (index, value) for which the function is true or -1,nil otherwise
	// if no element matches the critria.
	Find(func(index int, value interface{}) bool) (int, interface{})
}

// EnumerableWithKey provides functions for ordered containers whose values whose elements are key/value pairs.
type EnumerableWithKey interface {
	Each(func(key interface{}, value interface{}))

	// Map(func(key interface{}, value interface{}) (interface{}, interface{})) Container

	// Select(func(key interface{}, value interface{}) bool) Container

	Any(func(key interface{}, value interface{}) bool) bool

	All(func(key interface{}, value interface{}) bool) bool

	Find(func(key interface{}, value interface{}) bool) (interface{}, interface{})
}
