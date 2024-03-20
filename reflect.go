package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readFile(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	t := reflect.TypeOf(value)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()

			result = data != ""

			if !result {
				return result
			}
		}
	}

	return result
}

func main() {
	readFile(Sample{"Rizki"})
	readFile(Person{"Rizki", "", ""})

	person := Person{
		Name:    "Rizki",
		Address: "Dimana yah",
		Email:   "test@gmail.com",
	}
	fmt.Println(IsValid(person))

	person.Email = ""
	fmt.Println(IsValid(person))
}
