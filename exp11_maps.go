package main

import "fmt"

func main()  {
	m1 := make(map[string]int)

	m1["key1"] = 1
	m1["key2"] = 2

	fmt.Println("MAP: ", m1)

	value1 := m1["key1"]
	fmt.Println("MAP - value: ", value1)

	fmt.Println("MAP - LEN: ", len(m1))

	delete(m1, "key2")
	fmt.Println("MAP - delete key2: ", m1)

	_, prs := m1["key2"]
	fmt.Println("MAP - key2 in map: ", prs)

	m2 := map[string]int{"k3": 3, "k4": 4}
	fmt.Println("MAP - create and init new map: ", m2)


}