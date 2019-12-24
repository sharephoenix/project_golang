package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const tagName = "validate"
// 自定义标签 customtag

type User struct {
	Id int `validate:"-"`
	Name string `validate:"presence,min=2,max=32" json:"name"`
	Email string `validate:"email,required"`
}

func main() {
	user := User{
		Id: 1,
		Name: "John Doe",
		Email: "john@example",
	}

	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(user)

	v, _ := json.Marshal(reflect.ValueOf(user).Interface())
	var m map[string]interface{}

	json.Unmarshal(v, &m)
	fmt.Println("map::", m, reflect.TypeOf(reflect.ValueOf(user).Interface()))
	//Get the type and kind of our user variable
	fmt.Println("Type: ", t.Name())
	fmt.Println("Kind: ", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		//Get the field tag value
		tag := field.Tag.Get(tagName)

		tag0 := field.Tag.Get("json")
		//v, _ := json.Marshal(field)
		//var m map[string]interface{}
		//json.Unmarshal(v, &m)
		//fmt.Println(m)

		fmt.Printf("*** %d. %v(%v), tag:'%v', tagJson:'%v' 值：%v\n", i+1, field.Name, field.Type.Name(), tag, tag0, m[field.Name])
	}
}
