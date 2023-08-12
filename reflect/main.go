package main

import (
	"fmt"
	"reflect"
)

type monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s monster) Print() {
	fmt.Println("----start-----")
	fmt.Println(s)
	fmt.Println("-----end-------")
}

func (s monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func testStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()

	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("Field %d : 值为=%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag=%v", i, tagVal)
		}
	}

	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())

}

func main() {

	/*
		str := "tom"
		fs := reflect.ValueOf(&str)
		fs.Elem().SetString("jack")
		fmt.Printf("%v\n", str)
	*/
	a := monster{
		Name:  "小明",
		Age:   400,
		Score: 30.9,
		Sex:   "男",
	}
	testStruct(a)
}
