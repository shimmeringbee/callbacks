package callbacks

import (
	"context"
	"reflect"
	"sync"
)

type Adder interface {
	Add(f interface{})
}

type Caller interface {
	Call(ctx context.Context, event interface{}) error
}

type AdderCaller interface {
	Adder
	Caller
}

type Callbacks struct {
	m         *sync.RWMutex
	callbacks map[reflect.Type][]interface{}
}

func Create() *Callbacks {
	return &Callbacks{callbacks: map[reflect.Type][]interface{}{}, m: &sync.RWMutex{}}
}

func (c *Callbacks) Add(f interface{}) {
	funcType := reflect.TypeOf(f)

	if funcType.Kind() != reflect.Func {
		panic("attempted to add callback which was not a function")
	}

	if funcType.NumIn() != 2 || funcType.NumOut() != 1 {
		panic("attempted to add function which did not meet expected interface")
	}

	contextType := funcType.In(0)
	realContextType := reflect.TypeOf((*context.Context)(nil)).Elem()

	if !contextType.Implements(realContextType) {
		panic("attempted to add function which did not meet expected interface (first not context)")
	}

	errorType := funcType.Out(0)
	realErrorType := reflect.TypeOf((*error)(nil)).Elem()

	if !errorType.Implements(realErrorType) {
		panic("attempted to add function which did not meet expected interface (return not error)")
	}

	eventType := funcType.In(1)

	c.m.Lock()
	defer c.m.Unlock()

	_, found := c.callbacks[eventType]

	if !found {
		c.callbacks[eventType] = []interface{}{}
	}

	c.callbacks[eventType] = append(c.callbacks[eventType], f)
}

func (c *Callbacks) Call(ctx context.Context, event interface{}) error {
	eventType := reflect.TypeOf(event)

	c.m.RLock()
	array, found := c.callbacks[eventType]
	c.m.RUnlock()

	if found {
		for _, cb := range array {
			ctxValue := reflect.ValueOf(ctx)
			eventValue := reflect.ValueOf(event)

			in := []reflect.Value{ctxValue, eventValue}

			out := reflect.ValueOf(cb).Call(in)

			if !out[0].IsNil() {
				if err := out[0].Interface().(error); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
