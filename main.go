package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main () {
	// 设置预定义Logger的属性，包括
    // 日志条目前缀和禁用打印的标志
    //  时间、源文件和行号.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
		message, err := greetings.Hello("")
		if err != nil {
			log.Fatal(err)
		}
    fmt.Println(message)
}
