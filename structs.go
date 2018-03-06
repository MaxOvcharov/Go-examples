package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	fmt.Printf("Struct1: %+v \n", person{"Bob", 25})
	fmt.Printf("Struct2: %+v \n", person{name: "Alice", age: 22})
	fmt.Printf("Struct3: %+v \n", person{name: "Fred"})
	fmt.Printf("Struct4(pointer): %+v \n", &person{name: "Anna", age: 45})

	s := person{name: "Tod", age: 12}
	fmt.Printf("Struct(name): %+v \n", s.name)

	st := &s
	fmt.Printf("Struct(age1)(pointer): %+v \n", st.age)

	st.age = 100
	fmt.Printf("Struct(age2)(pointer): %+v \n", st.age)
}
