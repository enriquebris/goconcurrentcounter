package goconcurrentcounter

// function to be executed by concurrentInt
type concurrentIntFunc func(currentValue int, previousValue int)

type Int interface {
	initialize(value int)
	executeTriggerFunctions(currentValue int, previousValue int)
	// GetValue returns the current value of the counter
	GetValue() int
	// Update updates the counter with the given value
	Update(upd int)
	// GetTriggerOnValue returns the trigger function for the given value and name
	GetTriggerOnValue(value int, name string) concurrentIntFunc
	// SetTriggerOnValue sets a trigger function to be executed when the counter reaches the given value
	SetTriggerOnValue(value int, name string, fn concurrentIntFunc)
	// UnsetTriggerOnValue removes the trigger function for the given value and name
	UnsetTriggerOnValue(value int, name string)
	// UnsetTriggersOnValue removes all trigger functions for the given value
	UnsetTriggersOnValue(value int)
	// EnqueueToRunAfterCurrentTriggerFunctions enqueues the given function to be executed after the triggered functions. This function is executed only once.
	EnqueueToRunAfterCurrentTriggerFunctions(fn concurrentIntFunc)
}
