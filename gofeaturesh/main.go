package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	var m map[string]Vertex //= make (map[string]Vertex)
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Paris"] = Vertex{23.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Paris"])
	var n = map[string]int{"a": 1}
	fmt.Println("n %+v", n)
}
