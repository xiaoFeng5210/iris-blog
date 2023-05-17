package main

import (
	"fmt"
	"log"
  "encoding/csv"
	"os"
)

func main () {
	// 首先我们要打开source文件夹下的csv文件
	f, err := os.Open("source/pwd.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 创建一个新的csv读取器
	r := csv.NewReader(f)
  // 读取csv文件， 这将返回一个切片，其中每个元素是一个字符串切片， 表示csv文件的一行
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(records)

	for i, record := range records {
		fmt.Printf("Record %d: %v\n", i, record)
		// 把每个record切片里面的字符串拼接在一起
		newStr := record[0] + record[1]
		fmt.Println(newStr)
	}
}
