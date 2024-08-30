package utils

import "strings"

// Errors maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Errors map
// are case-sensitive.
type Errors map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (e Errors) Get(key string) string {
	if e == nil {
		return ""
	}
	vs := e[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (e Errors) Set(key, value string) {
	e[key] = []string{value}
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (e Errors) Add(key, value string) {
	e[key] = append(e[key], value)
}

// Del deletes the values associated with key.
func (e Errors) Del(key string) {
	delete(e, key)
}

// IsNil report whether the errors is empty
func (e Errors) IsNil() bool {
	return len(e) == 0
}

func (e Errors) Error() string {
	errs := []string{}
	for field, errs := range e {
		errs = append(errs, field+":"+strings.Join(errs, ","))
	}
	return strings.Join(errs, ",")
}
