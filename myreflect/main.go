package main

import (
	"fmt"
	"reflect"
)

// new
type myInt int
type Student struct {
	Name  string `json:"name" form:"username"`
	Age   int    `json:"age" form:"userage"`
	Score int    `json:"score" form:"userscore"`
}

func (s *Student) GetInfo() string {
	var str = fmt.Sprintf("name: %v; age: %v; score: %v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s *Student) Print() {
	fmt.Println("this is print functions")
}

// print fields

func PrintStudentFields(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("this is not s struct")
		return
	}

	//1. using field
	f0 := t.Elem().Field(0)
	fmt.Printf("Type: %v, Kind: %v, Name: %v, Tag1: %v, Tag2: %v\n", f0, f0.Type, f0.Name, f0.Tag.Get("json"), f0.Tag.Get("form"))

	//2.  get FiledbyName
	f1, ok := t.Elem().FieldByName("Age")
	if ok {
		fmt.Printf("Type: %v, Kind: %v, Name: %v, Tag1: %v, Tag2: %v\n", f1, f1.Type, f1.Name, f1.Tag.Get("json"), f1.Tag.Get("form"))
	}

	//3. from Typeof to get numfiled
	fCnt := t.Elem().NumField()
	fmt.Println("struct has ", fCnt, "fields")

	//4. To get value from struct from Value of
	v0 := v.Elem().FieldByName("Score")
	fmt.Printf("Type: %v, Kind: %v\n", v0, v.Elem().Type())

	for i := 0; i < fCnt; i++ {
		f0 = t.Elem().Field(i)
		fmt.Printf("Type: %v, Kind: %v, Name: %v, Tag1: %v, Tag2: %v\n", f0, f0.Type, f0.Name, f0.Tag.Get("json"), f0.Tag.Get("form"))

	}

}

func PrintStudentFn(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("this is not s struct")
		return
	}

	//1. Method(0)
	method0 := t.Method(0) // not the order in the file. It's according to the ASCII
	fmt.Println(method0.Name, method0.Type)

	//2. MethodByName
	m1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(m1.Name, m1.Type)
	}

	// 3. exec method from value
	v.MethodByName("Print").Call(nil)
	str := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(str)

	// 3.1 pass params to call
	var params []reflect.Value
	params = append(params, reflect.ValueOf("ting"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(90))
	v.MethodByName("SetInfo").Call(params)
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)

	// 4. get NumMethod
	fmt.Println(t.NumMethod(), v.NumMethod())

}

func ChangeStudentFn(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Ptr {
		fmt.Println("this is not a ptr")
		return
	} else if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("this is not a struct ptr")
		return
	}

	// update struct filed
	name := v.Elem().FieldByName("Name")
	name.SetString("li")
	age := v.Elem().FieldByName("Age")
	age.SetInt(22)
}

func structTest() {
	stu1 := Student{
		Name:  "xiang",
		Age:   30,
		Score: 100,
	}

	//PrintStudentFields(&stu1)
	// PrintStudentFn(&stu1)
	ChangeStudentFn(&stu1)
	fmt.Println(stu1)
}

func main() {
	//reflectTypeOfTest()
	//reflectValueOfTest()
	structTest()
}

func reflectTypeOf(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("Type: %v, Kind: %v, Name: %v\n", v, v.Kind(), v.Name())
}

func reflectValueOf(x interface{}) {

	// user reflect to get the value the original

	// vi := v.Int() + 12
	// fmt.Println("the sum is :", vi)

	// set

	// v, _ := x.(*int64)
	// *v = 120

	v := reflect.ValueOf(x)
	fmt.Printf("Value: %v, Kind: %v\n", v, v.Kind())
	if v.Elem().Kind() == reflect.Int64 { // if x is ptr then add v.Elem().Kind()
		v.Elem().SetInt(120)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("updated")
	}
}

func reflectTypeOfTest() {
	a := 10
	b := 23.4
	c := true
	d := "hello"

	reflectTypeOf(a)
	reflectTypeOf(b)
	reflectTypeOf(c)
	reflectTypeOf(d)

	var e myInt = 34
	var f = Student{
		Name: "xiang",
		Age:  20,
	}
	reflectTypeOf(e)
	reflectTypeOf(f)

	var h = 25        // pointer
	reflectTypeOf(&h) // *int

	var i = [3]int{1, 2, 3} // arrage
	reflectTypeOf(i)

	var j = []int{11, 22, 33} // slice
	reflectTypeOf(j)

}

func reflectValueOfTest() {
	var a int64 = 13
	reflectValueOf(&a)
	fmt.Println("value of a is: ", a)

	var b string = "13"
	reflectValueOf(&b)
	fmt.Println("value of a is: ", b)
}
