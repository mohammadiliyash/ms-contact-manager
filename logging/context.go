package logging

import (
	"context"
	"fmt"
)

var contextDict map[string]contextEntry

type contextKey string

var contextListKey = contextKey("contextListKey")

type contextEntry struct {
	active bool
	label  string
}

// fieldsFromContext reads through the submitted context for fields stored
// to it by this package. It returns a list of label->value pairs that can
// be added to any logger instance.
func fieldsFromContext(c context.Context) (map[string]interface{}, error) {
	rtn := map[string]interface{}{}

	// retrieve the list of field names (labels) that we've saved in this context
	ok := false
	var fieldSlice []contextKey
	fieldListField := c.Value(contextListKey)
	if fieldListField == nil {
		//spew.Dump("Fieldslice nil")
		// no fields, nothing to do
		return rtn, nil
	}
	if fieldSlice, ok = fieldListField.([]contextKey); !ok || len(fieldSlice) < 1 {
		//spew.Dump("Fieldslice empty")
		// no fields, nothing to do
		return rtn, nil
	}
	//spew.Dump("Fieldslice: ", fieldSlice)

	// loop over the list of fields and populate the return map
	// with label->value pairs
	for _, v := range fieldSlice {
		label := string(v)
		cv := c.Value(v)
		//spew.Dump(fmt.Sprintf("adding %v = %v to logger", label, cv))
		rtn[label] = cv
	}
	//spew.Dump("Returning: ", rtn)
	return rtn, nil
}

// AddToContext adds a value to the context
// Any future logging using this context will include this field
// with the given label
func AddToContext(c context.Context, label string, value interface{}) context.Context {
	var rtn context.Context
	if c == nil {
		return nil
	}

	// store the value in the context with a private key
	key := contextKey(label)
	rtn = context.WithValue(c, key, value)

	// store the key in the list of keys also stored in the context
	var fieldSlice []contextKey
	fieldListField := c.Value(contextListKey)
	found, ok := false, false
	if fieldListField == nil {
		//spew.Dump("Fieldslice nil")
		fieldSlice = []contextKey{}
	} else if fieldSlice, ok = fieldListField.([]contextKey); !ok {
		//spew.Dump("Fieldslice not casted")
		fieldSlice = []contextKey{}
	}

	//spew.Dump("Fieldslice: ", fieldSlice)
	for i := 0; i < len(fieldSlice) && !found; i++ {
		if fieldSlice[i] == key {
			found = true
		}
	}
	if !found {
		fieldSlice = append(fieldSlice, key)
	}
	rtn = context.WithValue(rtn, contextListKey, fieldSlice)
	return rtn
}

// HeadersFromContext accepts a map[string]string with context keys as keys
// and header names as values. It returns a map[string]string that can be used
// to set headers on new requests or reaponses
func HeadersFromContext(c context.Context, headerMap map[string]string) map[string][]string {
	rtn := map[string][]string{}
	for k, v := range headerMap {
		if k == "" || v == "" {
			continue
		}
		ckey := contextKey(k)
		cval := c.Value(ckey)
		rtn[v] = []string{fmt.Sprintf("%v", cval)}
	}

	return rtn
}
