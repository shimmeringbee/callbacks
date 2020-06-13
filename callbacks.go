package callbacks

import (
	"context"
	"reflect"
)

type Callbacks struct {
	callbacks map[reflect.Type][]interface{}
}

func Create() *Callbacks {
	return &Callbacks{callbacks: map[reflect.Type][]interface{}{}}
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

	_, found := c.callbacks[eventType]

	if !found {
		c.callbacks[eventType] = []interface{}{}
	}

	c.callbacks[eventType] = append(c.callbacks[eventType], f)
}

func (c *Callbacks) Call(ctx context.Context, event interface{}) error {
	eventType := reflect.TypeOf(event)

	array, found := c.callbacks[eventType]

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
