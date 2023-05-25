package pointer

import (
	"fmt"
)

// func Int10(v int) *int {
// 	p := v * 10
// 	fmt.Printf("p: %v\n", p)
// 	ip := &p
// 	fmt.Printf("ip: %v\n", ip)
// 	return ip
// }

func Run() {
	a := [3]int{1, 2, 3}
	b := a[1:]
	b[0] = 0

	c := &a
	(*c)[0] = -1
	d := (*c)[2:]
	d[0] = -2

	fmt.Println(a)
}
