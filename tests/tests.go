package tests

import "fmt"

func testAdd() {
	x := Load{5, 6}
	y := Load{2, 3}
	fmt.Println(x.Add(y))
}
