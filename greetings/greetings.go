package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// 返回在消息中嵌入名称的问候语.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init 为函数中使用的变量设置初始值.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
