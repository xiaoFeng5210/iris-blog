package pointer

import (
	"fmt"
)

func Int10(v int) *int {
	p := v * 10
	fmt.Printf("p: %v\n", p)
	ip := &p
	fmt.Printf("ip: %v\n", ip)
	return ip
}
