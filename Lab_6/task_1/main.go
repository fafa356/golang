package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Server struct {
	Host       string   `yaml:"host"`
	Port       int      `yaml:"port"`
	Debug      bool     `yaml:"debug"`
	AllowedIPs []string `yaml:"allowed_ips"`
}

func main() {
	v := Server{"localhost", 8080, true, []string{"192.168.1.1", "10.0.0.1"}}

	yamlStr, err := ToYAML(v)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(yamlStr)
}

func ToYAML(v any) (string, error) {
	t := reflect.TypeOf(v)

	if t != nil && t.Kind() != reflect.Struct {
		return "", errors.New("тип даних не struct")
	}

	result := ""

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := reflect.ValueOf(v).Field(i) // значення конкретного поля
		key := field.Tag.Get("yaml")

		switch val.Kind() {
		case reflect.String:
			result += fmt.Sprintf("%s: \"%s\"\n", key, val.String())
		case reflect.Int:
			result += fmt.Sprintf("%s: %d\n", key, val.Int())
		case reflect.Bool:
			result += fmt.Sprintf("%s: %t\n", key, val.Bool())
		case reflect.Slice:
			result += fmt.Sprintf("%s:\n", key)
			for j := 0; j < val.Len(); j++ {
				result += fmt.Sprintf("  - \"%s\"\n", val.Index(j).String())
			}
		}

	}

	return result, nil

}
