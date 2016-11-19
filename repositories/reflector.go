package repositories

import (
	"fmt"
	"reflect"
)

//Reflector uses golang reflection to get reflection attributes of an object
type Reflector struct {
}

//Fields returns the properties of an object as a map of string and interface
func (r Reflector) Fields(m interface{}) (map[string]interface{}, error) {
	typ := reflect.TypeOf(m)
	val := reflect.ValueOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]interface{})
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		return attrs, fmt.Errorf("%v type can't have attributes inspected\n", typ.Kind())
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			//attrs[p.Name] = p.Type
			attrs[p.Name] = val.Field(i).Interface()
		}
	}

	return attrs, nil
}
